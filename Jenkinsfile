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
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
                sh 'go get -u github.com/golang/dep/cmd/dep'
                sh './dep init'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build .'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                   }
            }
        }
    }
     post {
        always {
            archiveArtifacts artifacts: 'build/*', fingerprint: true 
        }
    }
}
