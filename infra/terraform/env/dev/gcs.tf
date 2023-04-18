module "tmp_tenhou_log_bucket" {
  source       = "../../modules/gcs"
  ENV          = local.ENV
  location     = local.location
  bucket_name  = "tmp-tenhou-log"
  service_name = var.service_name
  depends_on   = [module.enabled_services.services]
  random_id    = module.random_id.id
}
module "tmp_tenhou_log_trail_bucket" {
  source       = "../../modules/gcs"
  ENV          = local.ENV
  bucket_name  = "tmp-tenhou-log-trail"
  location     = local.location
  random_id    = module.random_id.id
  service_name = var.service_name
}
module "tenhou_log_bucket" {
  source       = "../../modules/gcs"
  ENV          = local.ENV
  location     = local.location
  bucket_name  = "tenhou-log"
  service_name = var.service_name
  depends_on   = [module.enabled_services.services]
  random_id    = module.random_id.id
}
module "tenhou_log_trail_bucket" {
  source       = "../../modules/gcs"
  ENV          = local.ENV
  location     = local.location
  bucket_name  = "tenhou-log-trail"
  service_name = var.service_name
  depends_on   = [module.enabled_services.services]
  random_id    = module.random_id.id
}


module "function-source" {
  source       = "../../modules/gcs"
  ENV          = local.ENV
  bucket_name  = "function-source-${local.ENV}"
  location     = local.location
  random_id    = module.random_id.id
  service_name = var.service_name
}
