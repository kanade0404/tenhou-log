resource "google_compute_router" "router" {
  name    = "private-router"
  network = var.network_id
}
resource "google_compute_router_nat" "router_nat" {
  name                               = "private-router-nat"
  nat_ip_allocate_option             = "AUTO_ONLY"
  router                             = google_compute_router.router.name
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}
