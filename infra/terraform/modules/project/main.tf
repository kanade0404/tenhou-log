resource "google_project" "project" {
  name        = var.PROJECT_ID
  project_id  = var.PROJECT_ID
  skip_delete = false
}
