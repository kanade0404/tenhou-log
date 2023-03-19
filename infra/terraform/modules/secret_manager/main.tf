resource "google_secret_manager_secret" "secret" {
  for_each  = var.secrets
  secret_id = each.key
  replication {
    automatic = true
  }
}
resource "google_secret_manager_secret_version" "version" {
  for_each    = var.secrets
  secret      = google_secret_manager_secret.secret[each.key].id
  secret_data = each.value
}
