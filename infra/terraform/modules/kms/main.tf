resource "google_kms_key_ring" "keyring" {
  location = var.location
  name     = "keyring-${var.ENV}"
}
resource "google_kms_crypto_key" "key" {
  for_each = toset(local.keys)
  key_ring = google_kms_key_ring.keyring.id
  name     = each.value
  lifecycle {
    prevent_destroy = true
  }
}

