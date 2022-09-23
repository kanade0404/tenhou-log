resource "google_storage_bucket" "bucket" {
  location                    = var.location
  name                        = "${var.bucket_name}-${var.ENV}-${var.random_id}"
  uniform_bucket_level_access = true
  versioning {
    enabled = true
  }
  force_destroy = var.force_destroy
  labels = {
    service_name = var.service_name
  }
  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      num_newer_versions = 5
    }
  }
  timeouts {}
}
