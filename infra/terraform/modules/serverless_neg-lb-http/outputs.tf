output "id" {
  value = google_compute_region_network_endpoint_group.serverless_neg.id
}
output "load_balancer_ip" {
  value = google_compute_global_address.global_ipv4_address.address
}
output "global_ipv4_address" {
  value = google_compute_global_address.global_ipv4_address.address
}
