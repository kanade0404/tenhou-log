resource "google_artifact_registry_repository" "repository" {
  format        = "DOCKER"
  repository_id = var.id
  location      = var.location
}
