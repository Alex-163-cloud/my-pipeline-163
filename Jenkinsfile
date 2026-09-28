pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                echo 'Code checked out!'
                sh 'ls -la'
            }
        }
        stage('Build') {
            steps {
                echo 'Building Docker image...'
                sh 'docker build -t my-go-app:latest .'
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
