output "id" {
  value = google_artifact_registry_repository.repository.id
}
output "name" {
  value = google_artifact_registry_repository.repository.name
}
output "image" {
  value = "${var.location}-docker.pkg.dev/${var.PROJECT_ID}/${var.id}"
}
