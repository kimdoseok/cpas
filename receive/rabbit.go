package main

import (
	"encoding/json"
	"fmt"
	"log"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/streadway/amqp"
)


type (
	Message struct {
		UserID string `json:"userid" `
		Key    string `json:"key" `
		Body   string `json:"body" `
	}
)


func main() {


	caCert, err := ioutil.ReadFile("../cert/ca_certificate.pem")
	if err != nil {
		fmt.Println("Error1:", err)
	}
 
	cert, err := tls.LoadX509KeyPair("../cert/client_certificate.pem", "../cert/client_key.pem")
	if err != nil {
		fmt.Println("Error2:", err)
	}
 
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)
 
	tlsConf := &tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost", // Optional
	}
 
	conn, err := amqp.DialTLS("amqps://doseok:kim7795004@localhost:5671/", tlsConf)
	if err != nil {
		fmt.Println("Error3:", err)
	}
	fmt.Println("Connection:", conn)

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error4:", err)
	}
	
	
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"imgs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		fmt.Println("Error5:", err)
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		fmt.Println("Error6:", err)
	}

	routing := "*.messages"

	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "imgs_topic", routing)
	err = ch.QueueBind(
		q.Name,       // queue name
		routing,      // routing key
		"imgs_topic", // exchange
		false,
		nil)

	if err != nil {
		fmt.Println("Error7:", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	if err != nil {
		fmt.Println("Error8:", err)
	}

	forever := make(chan bool)

	go func() {
		var m Message
		for d := range msgs {
			err = json.Unmarshal(d.Body, &m)

			//log.Printf(" [x] %s", d.Body)
			log.Printf(">%s", m.UserID)
			log.Printf(">>%s", m.Key)
			log.Printf(">>>%s", m.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}