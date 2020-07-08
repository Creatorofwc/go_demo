pipeline {
    agent any
    tools {
           go 'go1.14'
          }
     stages {         
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go get -t -d  -v ./... && go build -v ./...'
                 }
               }

        stage('Test') {
            steps {     
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                  }
                } 
} 
}
