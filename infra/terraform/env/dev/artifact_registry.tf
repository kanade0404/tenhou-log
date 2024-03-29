module "artifact_registry_web" {
  source     = "../../modules/artifact_registry"
  PROJECT_ID = var.PROJECT_ID
  id         = "web"
  location   = local.location
  depends_on = [module.enabled_services.services]
}

module "artifact_registry_backend" {
  source     = "../../modules/artifact_registry"
  PROJECT_ID = var.PROJECT_ID
  id         = "backend"
  location   = local.location
  depends_on = [module.enabled_services.services]
}
