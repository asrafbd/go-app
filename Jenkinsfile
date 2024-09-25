pipeline {
agent any

environment {

DOCKERHUB_CREDENTIALS=credentials('dockerhub')

}

stages {

stage("Downloading code") {

steps {

echo "code cloning"

git credentialsId: 'github-jenkins-key', url: 'git@github.com:asrafbd/go-app.git'

}

}

stage("Build Docker image") {

steps {

echo "Docker image building"

sh 'docker build -t newgoapp:v1.0.$BUILD_ID .'

sh 'docker tag newgoapp:v1.0.$BUILD_ID asrafbd/newgoapp:v1.0'


}

}

stage('Docker Hub Login') {

steps {

sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'

}

}

stage('Push Image') {

steps {

sh "docker push asrafbd/newgoapp:v1.0"

}

}

stage('Cleaning up') { 

steps { 

sh "docker rmi newgoapp:v1.0.$BUILD_ID" 

sh "docker rmi asrafbd/newgoapp:v1.0"

sh 'docker logout'

}

} 

stage('Deploy to Kubernetes') {

steps {
echo "copying yaml file to k8s server"




sshagent(['ansible-k8s']) {

sh 'scp -o StrictHostKeyChecking=no newgoapp-deployment.yaml asraf@192.168.0.120:/home/asraf'
sh 'scp -o StrictHostKeyChecking=no newgoapp-svc.yaml asraf@192.168.0.120:/home/asraf'


script {

sh 'ssh -o StrictHostKeyChecking=no asraf@192.168.0.120 kubectl apply -f newgoapp-deployment.yaml'
sh 'ssh -o StrictHostKeyChecking=no asraf@192.168.0.120 kubectl apply -f newgoapp-svc.yaml'

sh 'ssh -o StrictHostKeyChecking=no asraf@192.168.56.102 kubectl rollout restart deployment newgoapp-v1'

} 

}

}
}
} 
}

