module "lb-http-serverless_neg" {
  source     = "../../modules/serverless_neg-lb-http"
  PROJECT_ID = var.PROJECT_ID
  region     = var.location
  name       = "serverless-neg"
  run_name   = module.web_app_container.name
  domains    = ["touka-ryuumonbuchi.com"]
  depends_on = [module.enabled_services.services]
}
