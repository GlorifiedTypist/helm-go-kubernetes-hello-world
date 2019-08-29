node {
   def mvnHome
   stage('Preparation') { // for display purposes
      // Get some code from a GitHub repository
      git 'https://github.com/GlorifiedTypist/helm-go-kubernetes-hello-world.git'
         
   }
   stage('Build') {
       // build
       sh label: 'build', script: 'docker build . -t hello-world'
      
   }
   stage('Deploy') {

   }
}