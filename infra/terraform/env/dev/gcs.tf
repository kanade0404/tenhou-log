module "tmp_tenhou_log_bucket" {
  source       = "../../modules/gcs"
  ENV          = var.ENV
  location     = var.location
  bucket_name  = "tmp-tenhou-log"
  service_name = var.service_name
  depends_on   = [module.enabled_services.services]
  random_id    = module.random_id.id
}
module "tenhou_log_bucket" {
  source       = "../../modules/gcs"
  ENV          = var.ENV
  location     = var.location
  bucket_name  = "tenhou-log"
  service_name = var.service_name
  depends_on   = [module.enabled_services.services]
  random_id    = module.random_id.id
}
