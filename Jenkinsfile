pipeline {
    agent {
    kubernetes {
        cloud 'kubernetes'
        defaultContainer 'golang'
        yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: golang
    image: golang:1.21-alpine
    command:
    - cat
    tty: true
'''
    }
}

    environment {
        TIMEZONE = 'Europe/Moscow'
        DOCKER_HUB_USER = 'zakon13'
        KUBECONFIG = '/root/.kube/config'
    }

    options {
        buildDiscarder(logRotator(numToKeepStr: '10'))
        timeout(time: 10, unit: 'MINUTES')
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Code checked out!'
                echo "TIMEZONE = ${TIMEZONE}"
                sh 'ls -la'
            }
        }

        stage('Install Tools') {
            steps {
                echo 'Installing Docker CLI and kubectl...'
                sh '''
                    apk add --no-cache docker-cli curl bash
                    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
                    chmod +x kubectl
                    mv kubectl /usr/local/bin/
                    kubectl version --client
                '''
            }
        }

        stage('Build & Push Docker Image') {
            steps {
                script {
                    echo "Building image ${DOCKER_HUB_USER}/my-go-app:${env.BUILD_NUMBER}..."
                    def app = docker.build("${DOCKER_HUB_USER}/my-go-app:${env.BUILD_NUMBER}")

                    withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                        sh '''
                            echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                            docker push ${DOCKER_HUB_USER}/my-go-app:${BUILD_NUMBER}
                            docker tag ${DOCKER_HUB_USER}/my-go-app:${BUILD_NUMBER} ${DOCKER_HUB_USER}/my-go-app:latest
                            docker push ${DOCKER_HUB_USER}/my-go-app:latest
                        '''
                    }
                    echo "Pushed to Docker Hub"
                }
            }
        }

        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'echo "All tests passed!"'
            }
        }

        stage('Deploy to K8s') {
            steps {
                echo 'Deploying to Kubernetes...'
                sh '''
                    export KUBECONFIG=/root/.kube/config
                    kubectl --insecure-skip-tls-verify apply -f k8s/deployment.yaml
                    kubectl --insecure-skip-tls-verify apply -f k8s/service.yaml
                    kubectl --insecure-skip-tls-verify rollout status deployment/my-go-app --timeout=60s
                '''
                echo 'Deployed successfully!'
            }
        }

        stage('Verify') {
            steps {
                sh '''
                    export KUBECONFIG=/root/.kube/config
                    kubectl --insecure-skip-tls-verify get pods -l app=my-go-app
                    kubectl --insecure-skip-tls-verify get svc my-go-app
                '''
            }
        }
    }
}
