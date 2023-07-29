resource "google_workflows_workflow" "workflow" {
  name            = var.name
  region          = var.region
  description     = var.description
  service_account = var.service_account_id
  source_contents = templatefile("${path.module}/workflow.yaml", {
    source_content_vars = var.source_content_vars
    project_id          = var.PROJECT_ID
  })
}
