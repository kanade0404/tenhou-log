locals {
  accounts = [
    {
      name         = "scraper-invoker",
      display_name = "scraper invoker"
    }
  ]
  run_invokers = [
    module.service_account["scraper-invoker"].email
  ]
  sa_token_creators = [
    "service-${module.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
  ]
}
module "service_account" {
  for_each     = { for v in local.accounts : v.name => v }
  source       = "../../modules/service_account"
  PROJECT_ID   = var.PROJECT_ID
  display_name = each.value["display_name"]
  id           = each.value["name"]
  depends_on   = [module.enabled_services.services]
}
module "run_invoker_iam_role" {
  for_each   = toset(local.run_invokers)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = "roles/run.invoker"
  depends_on = [module.enabled_services.services]
}
module "sa_token_creator_iam_role" {
  for_each   = toset(local.sa_token_creators)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = "roles/iam.serviceAccountTokenCreator"
  depends_on = [module.enabled_services.services]
}
locals {
  artifact_registry_writer = [
    module.workload-identity-service-account.email
  ]
}
module "artifact_registry_writer" {
  for_each   = toset(local.artifact_registry_writer)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = "roles/artifactregistry.writer"
}
