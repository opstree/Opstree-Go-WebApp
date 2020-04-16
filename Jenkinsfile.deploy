node("master") {
    stage("Checking out codebase") {
        checkout scm
        config = readProperties file: 'Configuration'
    }
    lintHelmChart(
        helmChartDir: "${config.helm_chart_dir}"
    )
    applyHelmChart(
        helmChartDir: "${config.helm_chart_dir}",
        applicationName: "${config.application_name}"
    )
    testHelmChart(
        helmChartDir: "${config.helm_chart_dir}",
        applicationName: "${config.application_name}"
    )
    notificationStage(
        status: "good",
        environment: "dev",
        message: "New version app is successfully deployed"
    )
}

def lintHelmChart(Map stepParams) {
    try {
        stage("Linting helm chart for application") {
            dir("${stepParams.helmChartDir}") {
                sh "helm lint ./"
            }
        }
    } catch (Exception e) {
        echo "There is an error while linting helm chart. Please check the logs!!!!"
        echo e.toString()
        throw e
    }
}

def applyHelmChart(Map stepParams) {
    try {
        stage("Setting up the GoWebApp application") {
            dir("${stepParams.helmChartDir}") {
                sh "helm upgrade ${stepParams.applicationName} ./ -f values.yaml --namespace go-webapp --install"
            }
        }
    } catch (Exception e) {
        echo "There is an error while setting up application. Please check the logs!!!!"
        echo e.toString()
        throw e
    }
}

def testHelmChart(Map stepParams) {
    try {
        stage("Testing GoWebApp Application") {
            dir("${stepParams.helmChartDir}") {
                sh "helm test ${stepParams.applicationName} --namespace go-webapp"
            }
        }
    } catch (Exception e) {
        echo "There is an error while testing application. Please check the logs!!!!"
        echo e.toString()
        throw e
    }
}

def notificationStage(Map stepParams) {
    stage("Sending notification") {
        slackSend channel: 'build-status', color: "${stepParams.status}", message: "ENVIRONMENT:- ${stepParams.environment}\n BUILD_ID:- ${env.BUILD_ID}\n JOB_NAME:- ${env.JOB_NAME}\n Message:- ${stepParams.message}\n BUILD_URL:- ${env.BUILD_URL}"
    }
}
