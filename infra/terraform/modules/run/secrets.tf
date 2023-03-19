data "google_secret_manager_secret_version" "version" {
  for_each = { for v in var.secrets : v.name => v }
  secret   = each.value.name
  version  = each.value.version == null ? "latest" : each.value.version
}
