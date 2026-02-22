pipeline {
    // 'agent any' tells Jenkins to run this on any available executor
    agent any 
    
    environment {
        // We use localhost:5000 because Jenkins is using your host's Docker daemon
        REGISTRY = "localhost:5000"
        IMAGE_NAME = "sample-app"
        // BUILD_NUMBER is a built-in Jenkins variable (1, 2, 3, etc.)
        // This ensures every build creates a uniquely versioned, immutable artifact
        IMAGE_TAG = "v1.0.${env.BUILD_NUMBER}" 
    }

    stages {
        stage('Checkout Code') {
            steps {
                // This pulls the code from the Git repo configured in the Jenkins job
                checkout scm 
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    echo "🔨 Building Image: ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
                    
                    // Build the multi-stage Dockerfile
                    sh "docker build -t ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} ."
                    
                    // Also tag it as 'latest' for convenience
                    sh "docker tag ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} ${REGISTRY}/${IMAGE_NAME}:latest"
                }
            }
        }
        
        stage('Push to Registry') {
            steps {
                script {
                    echo "🚀 Pushing Artifacts to Local Registry..."
                    
                    // Push the uniquely tagged image and the latest tag
                    sh "docker push ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
                    sh "docker push ${REGISTRY}/${IMAGE_NAME}:latest"
                }
            }
        }
    }
    
    post {
        always {
            // Clean up the Jenkins workspace so we don't run out of disk space
            cleanWs()
            // Optional: Clean up local docker images on the host to save space
            sh "docker rmi ${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG} || true"
            sh "docker rmi ${REGISTRY}/${IMAGE_NAME}:latest || true"
        }
    }
}
