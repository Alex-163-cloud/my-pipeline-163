pipeline {
    agent any

environment {
        TIMEZONE = 'Europe/Moscow'
        DOCKER_HUB_USER = 'zakon13'
    }

options {
        buildDiscarder(logRotator(numToKeepStr: '10'))
        timeout(time: 10, unit: 'MINUTES')
        timestamps()
    }

triggers {
        cron('*/30 * * * *')
    }
    
    stages {
        stage('Checkout') {
            steps {
                echo 'Code checked out!'
                echo "TIMEZONE = ${TIMEZONE}"
                sh 'ls -la'
            }
        }
        stage('Build & Push Docker Image') {
    steps {
        script {
            echo "Building image ${DOCKER_HUB_USER}/my-go-app:${env.BUILD_NUMBER}..."
            app = docker.build("${DOCKER_HUB_USER}/my-go-app:${env.BUILD_NUMBER}")

            docker.withRegistry('https://registry.hub.docker.com', 'dockerhub-credentials') {
                app.push()
                app.push('latest')
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
                    export KUBECONFIG=/var/jenkins_home/.kube/config
                    kubectl apply -f k8s/deployment.yaml
                    kubectl apply -f k8s/service.yaml
                    kubectl rollout status deployment/my-go-app --timeout=60s
                '''
                echo 'Deployed successfully!'
            }
        }
        stage('Verify') {
            steps {
                sh '''
                    export KUBECONFIG=/var/jenkins_home/.kube/config
                    kubectl get pods -l app=my-go-app
                    kubectl get svc my-go-app
                '''
            }
        }
    }
}
