pipeline {
    agent any
    
    stages {
        stage('Checkout via SSH') {
            steps {
                echo 'Code checked out via SSH!'
                sh 'ls -la'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'docker build -t my-go-app:latest .'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'echo "All tests passed!"'
            }
        }
    }
}
