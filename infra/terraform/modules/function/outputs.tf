output "uri" {
  value = lookup(google_cloudfunctions2_function.function.service_config[0], "uri")
}
