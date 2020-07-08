pipeline {
    agent any
    tools {
        go 'go1.14'
    }
          stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go get -t -d  -v ./... && go build -v ./...'
            }
        }

        stage('Test') {
            steps {     
                    echo 'Running linting'
                    sh 'go test -v ./...'
                   }
                }
            
    
     post {
        always {
            archiveArtifacts artifacts: '*', fingerprint: true 
        }
    }
}
