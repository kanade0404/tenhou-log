resource "google_storage_bucket_object" "object" {
  bucket = var.bucket_name
  name   = var.object_name
  source = var.object_source
}
