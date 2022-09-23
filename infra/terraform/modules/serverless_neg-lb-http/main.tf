resource "google_compute_region_network_endpoint_group" "serverless_neg" {
  provider              = google-beta
  name                  = var.name
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  project               = var.PROJECT_ID
  cloud_run {
    service = var.run_name
  }
}
resource "google_compute_global_address" "global_ipv4_address" {
  name = "${var.name}-address"
}
resource "google_compute_managed_ssl_certificate" "ssl_cert" {
  provider = google-beta
  name     = "${var.name}-cert"
  managed {
    domains = var.domains
  }
}
resource "google_compute_backend_service" "backend" {
  name        = "${var.name}-backend"
  protocol    = "HTTP"
  port_name   = "http"
  timeout_sec = 30
  backend {
    group = google_compute_region_network_endpoint_group.serverless_neg.id
  }
}
resource "google_compute_url_map" "url_map" {
  name            = "${var.name}-urlmap"
  default_service = google_compute_backend_service.backend.id
}
resource "google_compute_target_https_proxy" "https_proxy" {
  name = "${var.name}-https-proxy"

  url_map = google_compute_url_map.url_map.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.ssl_cert.id
  ]
}
resource "google_compute_global_forwarding_rule" "lb" {
  name       = "${var.name}-lb"
  target     = google_compute_target_https_proxy.https_proxy.id
  port_range = "443"
  ip_address = google_compute_global_address.global_ipv4_address.address
}
resource "google_compute_url_map" "https_redirect" {
  name = "${var.name}-https-redirect"
  default_url_redirect {
    strip_query            = false
    https_redirect         = true
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
  }
}
resource "google_compute_target_http_proxy" "https_redirect" {
  name    = "${var.name}-http-proxy"
  url_map = google_compute_url_map.https_redirect.id
}
resource "google_compute_global_forwarding_rule" "https_redirect" {
  name       = "${var.name}-lb-http"
  target     = google_compute_target_http_proxy.https_redirect.id
  port_range = "80"
  ip_address = google_compute_global_address.global_ipv4_address.address
}
