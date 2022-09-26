module "scraper_invoker_pubsub" {
  source     = "../../modules/pubsub/push"
  PROJECT_ID = var.PROJECT_ID
  // 後で変える
  push_endpoint     = module.web_app_container.url
  subscription_name = "scraper-invoker-subscription"
  topic_name        = "scraper-invoker"
  depends_on        = [module.enabled_services.services]
}
