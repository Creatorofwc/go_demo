pipeline {
    agent any
    tools {
           go 'go1.14'
          }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
     stages {         
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go get -u golang.org/x/lint/golint'
                sh 'go get -t -d -v && go build -v -o App'
                 }
               }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet'
                    echo 'Running linting'
                    sh 'golint'
                  }
                } 
        
        post {
        always {
            archiveArtifacts artifacts: '*', fingerprint: true
               }
    } 
}
}
