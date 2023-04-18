module "scraper_secret_env" {
  source = "../../modules/secret_manager"
  secrets = merge(
    jsondecode(file("credentials/database.json")),
    jsondecode(file("credentials/api.json")), {
      ENV                              = local.ENV
      COMPRESSED_LOG_BUCKET_NAME       = module.tmp_tenhou_log_bucket.name
      COMPRESSED_LOG_TRAIL_BUCKET_NAME = module.tmp_tenhou_log_trail_bucket.name
      SCRAPER_PORT                     = module.scraper_app_container.port
      PROJECT_ID                       = module.project.id
      PROJECT_ID_NUMBER                = module.project.number
  })
}
