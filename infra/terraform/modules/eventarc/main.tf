resource "google_eventarc_trigger" "trigger" {
  for_each = { for event in var.events : event.name => event }
  location = var.location
  project  = var.PROJECT_ID
  name     = each.key
  destination {
    dynamic "cloud_run_service" {
      for_each = { for destination in each.value.destinations : destination.service => destination }
      content {
        service = cloud_run_service.value.service
        region  = var.location
        path    = cloud_run_service.value.path
      }
    }
  }
  dynamic "matching_criteria" {
    for_each = { for criteria in each.value.matching_criterion : criteria.value => criteria }
    content {
      attribute = matching_criteria.value.attribute
      value     = matching_criteria.value.value
    }
  }
  service_account = each.value.service_account
}
