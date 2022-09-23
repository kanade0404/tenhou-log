resource "google_app_engine_application" "firestore" {
  project       = var.PROJECT_ID
  location_id   = var.location
  database_type = "CLOUD_FIRESTORE"
}
