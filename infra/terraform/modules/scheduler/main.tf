resource "google_cloud_scheduler_job" "job" {
  name        = var.name
  description = var.description
  time_zone   = "Asia/Tokyo"
  schedule    = var.schedule
  http_target {
    http_method = "POST"
    uri         = var.http_target_uri
    oidc_token {
      service_account_email = var.service_account_email
    }
  }
}

