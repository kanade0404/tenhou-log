module "artifact_registry_web" {
  source     = "../../modules/artifact_registry"
  PROJECT_ID = var.PROJECT_ID
  id         = "web"
  location   = var.location
  depends_on = [module.enabled_services.services]
}
