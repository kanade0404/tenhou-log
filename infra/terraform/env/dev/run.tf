module "web_app_container" {
  source            = "../../modules/run"
  PROJECT_ID        = var.PROJECT_ID
  location          = var.location
  name              = "web"
  image_name        = "${module.artifact_registry_web.image}/next"
  sql_instance_name = ""
  ingress           = "internal-and-cloud-load-balancing"
  depends_on        = [module.enabled_services.services]
}
module "web_app_container_public_iam_member" {
  source      = "../../modules/run/iam_members"
  PROJECT_ID  = var.PROJECT_ID
  location    = var.location
  member_name = "allUsers"
  role        = "roles/run.invoker"
  run_name    = module.web_app_container.name
}
