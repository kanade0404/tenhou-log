variable "GOOGLE_CREDENTIALS" {}
variable "PROJECT_ID" {
  type = string
}
variable "ENV" {
  type = string
}
variable "location" {
  type    = string
  default = "asia-northeast1"
}
variable "default_zone" {
  type    = string
  default = "asia-northeast1-a"
}
variable "service_name" {
  type    = string
  default = "mj-log"
}
