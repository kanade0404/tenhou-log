module "firestore" {
  source     = "../../modules/firestore"
  PROJECT_ID = var.PROJECT_ID
  location   = local.location
  depends_on = [module.enabled_services.services]
}
