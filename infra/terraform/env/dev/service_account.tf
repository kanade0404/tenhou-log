locals {
  accounts = [
    {
      name         = "scraper-runner",
      display_name = "scraper runner"
    },
    {
      name         = "web-runner",
      display_name = "web app runner"
    },
    {
      name         = "api-runner",
      display_name = "api app runner"
    },
  ]
}
module "service_account" {
  for_each     = { for v in local.accounts : v.name => v }
  source       = "../../modules/service_account"
  PROJECT_ID   = var.PROJECT_ID
  display_name = each.value.display_name
  id           = each.value.name
  depends_on   = [module.enabled_services.services]
}
locals {
  run_invokers = [
    module.service_account["scraper-runner"].email,
    module.service_account["web-runner"].email,
    module.service_account["api-runner"].email
  ]
}
module "run_invoker_iam_role" {
  for_each   = toset(local.run_invokers)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = "roles/run.invoker"
  depends_on = [module.enabled_services.services]
}
locals {
  sa_token_creators = [
    "service-${module.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"

  ]
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
  secret_manager_accessors = [
    module.service_account["scraper-runner"].email,
    module.service_account["api-runner"].email
  ]
}
module "secret_manager_accessor" {
  for_each   = toset(local.secret_manager_accessors)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = "roles/secretmanager.secretAccessor"
  depends_on = [module.enabled_services.services]
}
locals {
  storage_admins = [
    module.service_account["scraper-runner"].email
  ]
}
module "storage_admin" {
  for_each   = toset(local.storage_admins)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = "roles/storage.admin"
}

module "github_actions_role" {
  source = "../../modules/iam/custom_role"
  id     = "actions.runner"
  permissions = [
    "iam.serviceAccounts.actAs",
    "run.services.create",
    "run.services.update",
    "run.services.get",
    "artifactregistry.repositories.uploadArtifacts"
  ]
  title = "GitHub Actions Role"
}
locals {
  github_actions = [
    module.workload-identity-service-account.email
  ]
}
module "github_actions" {
  for_each   = toset(local.github_actions)
  source     = "../../modules/iam"
  PROJECT_ID = var.PROJECT_ID
  name       = each.value
  role       = module.github_actions_role.id
}
