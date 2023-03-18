variable "PROJECT_ID" {
  type = string
}
variable "region" {
  type = string
}
variable "name" {
  type = string
}
variable "default_run_name" {
  type = string
}
variable "runs" {
  type = list(string)
}
variable "api" {
  type = object({
    name = string
    host_rule = object({
      hosts        = list(string)
      path_matcher = string
    })
  })
}
variable "domains" {
  type = list(string)
}
variable "oauth" {
  type = object({
    client_id     = string
    client_secret = string
  })
}
variable "iap_members" {
  type = list(object({
    name     = string
    role     = string
    run_name = string
  }))
}
