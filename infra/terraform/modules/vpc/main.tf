resource "google_compute_global_address" "private_ip_address" {
  name          = "vpc-private-ip-address"
  address_type  = "INTERNAL"
  network       = google_compute_network.private_network.id
  prefix_length = 16
  purpose       = "VPC_PEERING"
}

resource "google_service_networking_connection" "private_vpc_connection" {
  network = google_compute_network.private_network.id
  reserved_peering_ranges = [
    "vpc-private-ip-address"
  ]
  service = "servicenetworking.googleapis.com"
}
resource "google_compute_subnetwork" "vpc" {
  ip_cidr_range = "10.0.0.0/16"
  name          = "private"
  network       = google_compute_network.private_network.id
}
resource "google_compute_network" "private_network" {
  name                     = "private"
  enable_ula_internal_ipv6 = false
  auto_create_subnetworks  = false
}
