module "web_app_container" {
  source     = "../../modules/run"
  PROJECT_ID = var.PROJECT_ID
  location   = local.location
  name       = "web"
  image_name = "${module.artifact_registry_web.image}/next"
  ingress    = "all"
  port       = 3000
  depends_on = [module.enabled_services.services]
  iam_members = [
    {
      name = "allUsers"
      role = "roles/run.invoker"
    }
  ]
  service_account_name = module.service_account["web-runner"].email
}
module "scraper_app_container" {
  source     = "../../modules/run"
  PROJECT_ID = var.PROJECT_ID
  image_name = "${module.artifact_registry_backend.image}/scraper"
  location   = local.location
  ingress    = "internal"
  name       = "scraper"
  port       = 8088
  secrets = [
    {
      name = "PROJECT_ID_NUMBER"
    },
    {
      name = "COMPRESSED_LOG_TRAIL_BUCKET_NAME"
    },
    {
      name = "PROJECT_ID"
    }
  ]
  service_account_name = module.service_account["scraper-runner"].email
  depends_on           = [module.enabled_services.services]
}
module "api_app_container" {
  source     = "../../modules/run"
  PROJECT_ID = var.PROJECT_ID
  location   = local.location
  name       = "api"
  image_name = "${module.artifact_registry_backend.image}/api"
  ingress    = "internal"
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
    },
    {
      name = "PROJECT_ID"
    }
  ]
  depends_on           = [module.enabled_services.services]
  service_account_name = module.service_account["api-runner"].email
}
