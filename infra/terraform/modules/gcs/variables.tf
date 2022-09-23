variable "ENV" {
  type = string
}
variable "location" {
  type = string
}
variable "random_id" {
  type = string
}
variable "bucket_name" {
  type = string
}
variable "service_name" {
  type = string
}
variable "force_destroy" {
  type    = bool
  default = false
}
