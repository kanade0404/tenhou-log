module "firestore" {
  source     = "../../modules/firestore"
  PROJECT_ID = var.PROJECT_ID
  location   = var.location
  depends_on = [module.enabled_services.services]
}
