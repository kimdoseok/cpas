package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

type (
	Message struct {
		UserID string `json:"userid" `
		Key    string `json:"key" `
		Body   string `json:"body" `
	}
)

const (
	charset string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	routing string = "json.messages"
	userid  string = "doseok"
)

var (
	key string
)

func getRandStr(charset string, charlen int) string {
	chars := []rune(charset)
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < charlen; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func main() {
	//Connection()

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

	fmt.Println(">>>doseok:kim7795004")
	conn, err := amqp.DialTLS("amqps://doseok:kim7795004@localhost:5671/", tlsConf)
	//conn, err := amqp.DialTLS("amqps://localhost:5671/", tlsConf)
	if err != nil {
		fmt.Println("Error3:", err)
	}
	fmt.Println("Connection:", conn)
	// If you used credentials in docker, we did not!
	// return amqp.DialTLS("amqps://user:pass@localhost:5671/", tlsConf)

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

	msg := Message{UserID: userid, Key: getRandStr(charset, 64), Body: getRandStr(charset, 256)}
	body, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error5:", err)
	}

	err = ch.Publish(
		"imgs_topic", // exchange
		routing,      // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		fmt.Println("Error6:", err)
	}
	log.Printf(" [x] Sent %s", body)
}
