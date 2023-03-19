module "web_app_container" {
  source     = "../../modules/run"
  PROJECT_ID = var.PROJECT_ID
  location   = local.location
  name       = "web"
  image_name = "${module.artifact_registry_web.image}/next"
  ingress    = "internal-and-cloud-load-balancing"
  port       = 3000
  depends_on = [module.enabled_services.services]
  iam_members = [
    {
      name = "allUsers"
      role = "roles/run.invoker"
    }
  ]
  service_account_name = module.service_account["web-invoker"].email
}
module "scraper_app_container" {
  source               = "../../modules/run"
  PROJECT_ID           = var.PROJECT_ID
  image_name           = "${module.artifact_registry_backend.image}/scraper"
  location             = local.location
  name                 = "scraper"
  port                 = 8080
  service_account_name = module.service_account["scraper-invoker"].email
}
module "api_app_container" {
  source     = "../../modules/run"
  PROJECT_ID = var.PROJECT_ID
  location   = local.location
  name       = "api"
  image_name = "${module.artifact_registry_backend.image}/api"
  ingress    = "internal-and-cloud-load-balancing"
  port       = 8080
  secrets = [
    {
      name = "HASURA_GRAPHQL_METADATA_DATABASE_URL"
    },
    {
      name = "HASURA_GRAPHQL_ADMIN_SECRET"
    },
    {
      name = "HASURA_GRAPHQL_ENABLE_CONSOLE"
    },
    {
      name = "PG_DATABASE_URL"
    },
    {
      name = "HASURA_GRAPHQL_DATABASE_URL"
    }
  ]
  depends_on = [module.enabled_services.services]
  iam_members = [
    {
      name = "allUsers"
      role = "roles/run.invoker"
    }
  ]
  service_account_name = module.service_account["api-invoker"].email
}
