resource "google_cloud_run_service_iam_member" "member" {
  for_each = { for v in var.iam_members : v.name => v }
  member   = each.value.name
  role     = each.value.role
  service  = google_cloud_run_service.run.name
  location = var.location
  project  = var.PROJECT_ID
}
