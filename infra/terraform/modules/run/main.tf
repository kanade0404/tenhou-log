resource "google_cloud_run_service" "run" {
  location = var.location
  name     = var.name
  project  = var.PROJECT_ID
  template {
    spec {
      service_account_name = var.service_account_name
      containers {
        image = var.image_name
        ports {
          container_port = var.port
          name           = "http1"
        }
        resources {
          limits = {
            "cpu"    = "1000m"
            "memory" = "512Mi"
          }
        }
        dynamic "env" {
          for_each = { for v in var.secrets : v.name => v }
          content {
            name = env.value.name
            value_from {
              secret_key_ref {
                key  = data.google_secret_manager_secret_version.version[env.value.name].version
                name = data.google_secret_manager_secret_version.version[env.value.name].secret
              }
            }
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "1000"
        "run.googleapis.com/client-name"   = "terraform"
        "client.knative.dev/user-image"    = var.image_name
      }
    }
  }
  traffic {
    percent         = var.traffic_percentile
    latest_revision = true
  }
  autogenerate_revision_name = true
  metadata {
    annotations = {
      "run.googleapis.com/ingress"        = var.ingress
      "run.googleapis.com/ingress-status" = var.ingress
    }
  }
  lifecycle {
    ignore_changes = [
      metadata[0].annotations["run.googleapis.com/operation-id"],
      template[0].metadata[0].annotations["client.knative.dev/user-image"],
      template[0].metadata[0].annotations["run.googleapis.com/client-name"],
      template[0].metadata[0].annotations["run.googleapis.com/client-version"],
      template[0].metadata[0].annotations["run.googleapis.com/sandbox"],
      metadata[0].annotations["serving.knative.dev/creator"],
      metadata[0].annotations["serving.knative.dev/lastModifier"],
      metadata[0].annotations["run.googleapis.com/ingress-status"],
      metadata[0].annotations["run.googleapis.com/launch-stage"],
      metadata[0].labels["cloud.googleapis.com/location"],
    ]
  }
}
