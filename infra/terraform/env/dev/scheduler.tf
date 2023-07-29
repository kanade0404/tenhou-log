module "scheduler" {
  source                = "../../modules/scheduler"
  description           = "牌譜ログを取得してdecodeするworkflowをキックします。"
  schedule              = "0 */12 * * *"
  http_target_uri       = "https://workflowexecutions.googleapis.com/v1/${module.workflow.id}/executions"
  name                  = "trigger-workflow"
  service_account_email = module.service_account["scheduler-runner"].email
}
