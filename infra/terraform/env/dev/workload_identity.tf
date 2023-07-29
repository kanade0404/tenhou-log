module "workload-identity-service-account" {
  source       = "../../modules/service_account"
  PROJECT_ID   = var.PROJECT_ID
  id           = "github-actions"
  display_name = "github-actions"
  depends_on   = [module.enabled_services.services]
}
module "github-actions-workload-identity" {
  source                = "../../modules/workload_identifty"
  PROJECT_ID            = var.PROJECT_ID
  pool_id               = "github-actions-pool"
  pool_display_name     = "github-actions-pool"
  provider_id           = "github-actions-provider"
  provider_display_name = "github-actions-provider"
  github_repository     = "kanade0404/tenhou-log"
  service_account_id    = module.workload-identity-service-account.name
  depends_on            = [module.enabled_services.services]
}
