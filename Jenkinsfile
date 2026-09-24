pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                echo 'Code checked out!'
            }
        }
        stage('Use GitLab Token') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'gitlab-token',
                    usernameVariable: 'GITLAB_USER',
                    passwordVariable: 'GITLAB_TOKEN'
                )]) {
                    sh '''
                        echo "GitLab User: $GITLAB_USER"
                        echo "Token is available: ${#GITLAB_TOKEN} chars"
                    '''
                }
            }
        }
    }
}
