resource "google_cloudfunctions2_function" "function" {
  name     = var.name
  location = var.location
  build_config {
    runtime     = var.runtime
    entry_point = var.entry_point
    source {
      storage_source {
        bucket = var.storage_source.bucket
        object = var.storage_source.object
      }
    }
  }
  service_config {
    service_account_email = var.service_account_email
    max_instance_count    = var.max_instance_count
    available_memory      = "256M"
    timeout_seconds       = 1800
    ingress_settings      = "ALLOW_INTERNAL_ONLY"
    dynamic "secret_environment_variables" {
      for_each = { for secret in var.secrets : secret.name => secret }
      content {
        key        = secret_environment_variables.value.name
        project_id = var.PROJECT_ID
        secret     = data.google_secret_manager_secret_version.version[secret_environment_variables.value.name].secret
        version    = secret_environment_variables.value.version
      }
    }
  }
  lifecycle {
    ignore_changes = [
      labels["deployment-tool"]
    ]
  }
}
data "google_secret_manager_secret_version" "version" {
  for_each = { for secret in var.secrets : secret.name => secret }
  secret   = each.value.name
  version  = each.value.version
}
