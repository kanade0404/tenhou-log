resource "google_cloud_run_service" "run" {
  location = var.location
  name     = var.name
  project  = var.PROJECT_ID
  template {
    spec {
      containers {
        image = var.image_name
        ports {
          container_port = 3000
          name           = "http1"
        }
        resources {
          limits = {
            "cpu"    = "1000m"
            "memory" = "512Mi"
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "1000"
        #        "run.googleapis.com/cloudsql-instances" = var.sql_instance_name
        "run.googleapis.com/client-name" = "terraform"
        "client.knative.dev/user-image"  = var.image_name
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
