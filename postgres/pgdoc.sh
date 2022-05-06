#docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
#docker-compose up --build

#SELECT 'CREATE DATABASE offiworks' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'offiworks')
#psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'offiworks'" | grep -q 1 | psql -U postgres -c "CREATE DATABASE offiworks"
#eval $(minikube docker-env)

#docker build .
#docker tag postgres:alpine localhost:5000/postgres
#docker push localhost:5000/postgres

microk8s enable registry:size=50Gi

docker build . -t localhost:32000/postgres:registry
docker tag <ImageID> localhost:32000/postgres:registry

docker push localhost:32000/postgres:registry

/etc/docker/daemon.json and add:
{
  "insecure-registries" : ["localhost:32000"]
}
sudo systemctl restart docker

    spec:
      containers:
      - name: postgres
        image: localhost:32000/postgres:registry
        ports:
        - containerPort: 80