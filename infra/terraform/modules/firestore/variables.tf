variable "PROJECT_ID" {
  type = string
}
variable "location" {
  type = string
}
variable "depends_on_services" {
  type    = list(any)
  default = []
}
