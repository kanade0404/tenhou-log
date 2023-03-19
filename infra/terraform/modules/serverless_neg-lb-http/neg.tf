resource "google_compute_region_network_endpoint_group" "serverless_neg" {
  for_each              = toset(var.runs)
  provider              = google-beta
  name                  = "${each.value}-${substr(uuid(), 0, 6)}"
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  project               = var.PROJECT_ID
  cloud_run {
    service = each.value
  }
  lifecycle {
    create_before_destroy = true
    ignore_changes = [
      name
    ]
  }
}
resource "google_compute_backend_service" "backend" {
  for_each    = toset(var.runs)
  name        = "${each.value}-backend"
  protocol    = "HTTP"
  port_name   = "http"
  timeout_sec = 30
  backend {
    group = google_compute_region_network_endpoint_group.serverless_neg[each.value].id
  }
  iap {
    oauth2_client_id     = var.oauth.client_id
    oauth2_client_secret = var.oauth.client_secret
  }
}
resource "google_compute_url_map" "url_map" {
  name            = "${var.name}-urlmap-${substr(uuid(), 0, 6)}"
  default_service = google_compute_backend_service.backend[var.default_run_name].id
  host_rule {
    hosts        = ["api.touka-ryuumonbuchi.com"]
    path_matcher = "console"
  }
  path_matcher {
    name            = "console"
    default_service = google_compute_backend_service.backend[var.api.name].id
  }
  test {
    host    = "touka-ryuumonbuchi.com"
    path    = "/"
    service = google_compute_backend_service.backend[var.default_run_name].id
  }
  test {
    host    = "api.touka-ryuumonbuchi.com"
    path    = "/"
    service = google_compute_backend_service.backend[var.api.name].id
  }
  lifecycle {
    ignore_changes = [
      name
    ]
  }
}

resource "google_iap_web_backend_service_iam_member" "member" {
  for_each            = { for v in var.iap_members : "${v.name}-${v.run_name}" => v }
  member              = each.value.name
  role                = each.value.role
  web_backend_service = google_compute_backend_service.backend[each.value.run_name].name
}
