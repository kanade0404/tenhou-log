output "name" {
  value = google_cloud_run_service.run.name
}
output "url" {
  value = google_cloud_run_service.run.status[0].url
}
output "port" {
  value = google_cloud_run_service.run.template[0].spec[0].containers[0].ports[0].container_port
}
