output "name" {
  value = google_cloud_run_service.run.name
}
output "url" {
  value = google_cloud_run_service.run.status[0].url
}
