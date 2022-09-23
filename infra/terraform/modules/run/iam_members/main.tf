resource "google_cloud_run_service_iam_member" "member" {
  member   = var.member_name
  role     = var.role
  service  = var.run_name
  location = var.location
  project  = var.PROJECT_ID
}
