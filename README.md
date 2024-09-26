# Solution

## 1. Provision a cloud storage bucket using Infrastructure as Code (IaC).
$ ansible-playbook s3-bucket-create.yml --syntax-check

$ ansible-playbook s3-bucket-create.yml
## 3. Containerize the Go application.

$ cd sre_airports_api

$ vi Dockerfile

FROM golang:1.21.0

#Set the Current Working Directory inside the container

WORKDIR /app

#Copy go mod file

COPY go.mod ./

#Download all dependencies.

RUN go mod download

#Copy the source code into the container

COPY . .

#Build the Go app

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-app

#Expose port 8080 to the outside world

EXPOSE 8080

#Command to run the executable

CMD ["/go-docker-app"]

# Build docker image
$ docker build -t airports_api .

$ docker tag airports_api:latest asrafbd/airports_api:v1.0

$ docker login

$ docker push asrafbd/airports_api:v1.0

# 4. Prepare a deployment and service resource to deploy in Kubernetes.

$ vi newgoapp-deployment.yaml

$ kubectl apply -f newgoapp-deployment.yaml

deployment.apps/newgoapp-v1 created

$ vi newgoapp-svc.yaml

$ kubectl apply -f newgoapp-svc.yaml

service/newgoapp-v1 created

## 5. Use API gateway Create routing rules to send 20% of traffic to the `/airports_v2` endpoint.

kubectl apply -f newgoapp-ingress-canary.yaml

## 6. CI/CD Pipeline with Jenkinsfile

Need to create a Pipeline project on a Jenkins server and point to the Jenkinsfile in a GitHub repository.
