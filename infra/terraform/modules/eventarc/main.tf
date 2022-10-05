variable "PROJECT_ID" {
  type = string
}
variable "location" {
  type = string
}
variable "name" {
  type = string
}
resource "google_eventarc_trigger" "event_trigger" {
  location = var.location
  project  = var.PROJECT_ID
  name     = var.name
  destination {}
  matching_criteria {
    attribute = ""
    value     = ""
  }
}
