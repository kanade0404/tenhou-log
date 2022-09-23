resource "google_project_service" "service" {
  for_each           = toset(local.services)
  project            = var.PROJECT_ID
  service            = each.value
  disable_on_destroy = false
}
