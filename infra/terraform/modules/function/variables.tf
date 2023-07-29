variable "PROJECT_ID" {
  type = string
}
variable "name" {
  type = string
}
variable "location" {
  type = string
}
variable "runtime" {
  type = string
}
variable "entry_point" {
  type = string
}
variable "storage_source" {
  type = object({
    bucket = string
    object = string
  })
}
variable "max_instance_count" {
  type    = number
  default = 100
}
variable "secrets" {
  type = list(object({
    name    = string
    version = optional(string, "1")
  }))
  default = []
}
variable "service_account_email" {
  type = string
}
