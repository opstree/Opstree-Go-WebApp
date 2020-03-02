@Library('golang-ci@master')

def codeUtils = new org.opstree.golang.golangCIPipeline()

properties([[$class: 'JiraProjectProperty'], gitLabConnection(''), [$class: 'RebuildSettings', autoRebuild: false, rebuildDisabled: false], pipelineTriggers([githubPush()])])

node{
  codeUtils.call()
}
