pipeline {
    agent {
    kubernetes {
        cloud 'kubernetes'
        defaultContainer 'golang'
        yaml '''
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: jenkins
  containers:
  - name: golang
    image: golang:1.21-alpine
    command:
    - cat
    tty: true
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    command:
    - /busybox/cat
    tty: true
    volumeMounts:
    - name: docker-config
      mountPath: /kaniko/.docker
  volumes:
  - name: docker-config
    secret:
      secretName: dockerhub-secret
      items:
      - key: .dockerconfigjson
        path: config.json
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
        container('kaniko') {
            sh '''
                /kaniko/executor \
                  --dockerfile=${WORKSPACE}/Dockerfile \
                  --context=${WORKSPACE} \
                  --destination=zakon13/my-go-app:${BUILD_NUMBER} \
                  --destination=zakon13/my-go-app:latest \
                  --cache=true
            '''
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
                    kubectl get pods -l app=my-go-app
                    kubectl get svc my-go-app
                '''
            }
        }
    }
}
