module "lb-http-serverless_neg" {
  source           = "../../modules/serverless_neg-lb-http"
  PROJECT_ID       = var.PROJECT_ID
  region           = local.location
  name             = "serverless-neg"
  default_run_name = module.web_app_container.name
  domains          = ["touka-ryuumonbuchi.com", "api.touka-ryuumonbuchi.com"]
  oauth = {
    client_id     = var.OAUTH_CLIENT_ID
    client_secret = var.OAUTH_CLIENT_SECRET
  }
  runs = [module.web_app_container.name, module.api_app_container.name]
  api = {
    name = module.api_app_container.name
    host_rule = {
      hosts        = ["api.touka-ryuumonbuchi.com"]
      path_matcher = "admin"
    }
  }
  iap_members = [
    {
      name     = local.owner
      role     = "roles/iap.httpsResourceAccessor"
      run_name = module.web_app_container.name
    },
    {
      name     = local.owner
      role     = "roles/iap.httpsResourceAccessor"
      run_name = module.api_app_container.name
    }
  ]
  depends_on = [module.enabled_services.services]
}
