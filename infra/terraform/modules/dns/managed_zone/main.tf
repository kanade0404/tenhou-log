resource "google_dns_managed_zone" "dns_zone" {
  dns_name   = "${var.dns_name}."
  name       = var.name
  project    = var.PROJECT_ID
  visibility = var.is_visible ? "public" : "private"
  dnssec_config {
    state         = var.state
    non_existence = var.non_existence
  }
}
