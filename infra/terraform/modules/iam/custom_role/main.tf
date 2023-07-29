resource "google_project_iam_custom_role" "custom_role" {
  permissions = var.permissions
  role_id     = var.id
  title       = var.title
}
