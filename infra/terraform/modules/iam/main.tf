resource "google_project_iam_member" "member" {
  member  = "serviceAccount:${var.name}"
  project = var.PROJECT_ID
  role    = var.role
}
