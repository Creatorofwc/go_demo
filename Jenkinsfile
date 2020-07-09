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
	}
	     stage('Archive') {
		     steps {
		        archiveArtifacts artifacts: 'App, config.json, index.html, scripts', fingerprint: true, followSymlinks: false
		     }
	     }
	
	      stage('SSH transfer') {
        steps([$class: 'BapSshPromotionPublisherPlugin']) {
            sshPublisher(
                continueOnError: false, failOnError: true,
                publishers: [
                    sshPublisherDesc(
                        configName: "hello_world",
                        verbose: true,
                        transfers: [
                            sshTransfer(execCommand: "/bin/rm -rf /app/*"),
                            sshTransfer(sourceFiles: "App index.html config.json",),
		            sshTransfer(execCommand: "chmod 755 /app/App")	
                        ]
                    )
                ]
            )
        }
    }
      //  stage('Deploy') {
      //       steps([$class: 'AWSCodeDeployPublisher', applicationName: '', awsAccessKey: '', awsSecretKey: '', credentials: 'awsAccessKey', deploymentGroupAppspec: false, deploymentGroupName: '', deploymentMethod: 'deploy', excludes: '', iamRoleArn: '', includes: '**', proxyHost: '', proxyPort: 0, region: 'ap-northeast-1', s3bucket: '', s3prefix: '', subdirectory: '', versionFileName: '', waitForCompletion: false])
       //      }
       }
}
