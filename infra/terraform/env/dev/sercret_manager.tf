module "scraper_secret_env" {
  source = "../../modules/secret_manager"
  secrets = merge(
    jsondecode(file("credentials/scraper.json")),
    jsondecode(file("credentials/database.json")),
    jsondecode(file("credentials/api.json")),
    jsondecode(file("credentials/common.json")), {
      ENV = local.ENV
  })
}
