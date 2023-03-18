output "dns_name" {
  value = google_dns_managed_zone.dns_zone.dns_name
}
output "name_servers" {
  value = google_dns_managed_zone.dns_zone.name_servers
}
