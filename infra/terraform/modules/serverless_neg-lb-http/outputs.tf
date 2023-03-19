output "load_balancer_ip" {
  value = google_compute_global_address.global_ipv4_address.address
}
