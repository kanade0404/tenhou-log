resource "google_dns_record_set" "dns_zone" {
  project      = var.PROJECT_ID
  managed_zone = var.dns_zone_name
  name         = var.name
  type         = var.record_type
  rrdatas      = var.rrdatas
  ttl          = var.ttl
}
