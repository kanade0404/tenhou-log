resource "google_dataflow_job" "job" {
  name              = var.name
  temp_gcs_location = ""
  template_gcs_path = ""
}
