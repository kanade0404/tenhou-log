locals {
  ENV          = "dev"
  location     = "asia-northeast1"
  default_zone = "asia-northeast1-a"
  service_name = "mj-log"
}
variable "GOOGLE_CREDENTIALS" {}
variable "PROJECT_ID" {
  type = string
}
variable "location" {
  type    = string
  default = "asia-northeast1"
}
variable "service_name" {
  type    = string
  default = "mj-log"
}
variable "OAUTH_CLIENT_ID" {
  type = string
}
variable "OAUTH_CLIENT_SECRET" {
  type = string
}
locals {
  owner = "user:melty0404@gmail.com"
}
