module "workflow" {
  source             = "../../modules/workflows"
  PROJECT_ID         = var.PROJECT_ID
  description        = "牌譜ログを取得してdecodeします。"
  name               = "log-transformer"
  region             = var.location
  service_account_id = module.service_account["workflow-runner"].id
  source_content_vars = {
    scraper_url          = module.scraper_app_container.url
    log_converter_url    = module.log_converter.uri
    log_decompressor_url = module.log_decompressor.uri
    log_downloader_url   = module.log_downloader.uri
  }
}
