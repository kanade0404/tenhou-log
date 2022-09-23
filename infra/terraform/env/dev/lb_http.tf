#module "lb_http" {
#  source            = "../../modules/lb-http"
#  PROJECT_ID        = var.PROJECT_ID
#  name              = "lb_http"
#  serverless_neg_id = module.gclb_backend.id
#}
