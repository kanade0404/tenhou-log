variable "PROJECT_ID" {
  type = string
}
variable "location" {
  type = string
}
variable "name" {
  type = string
}
variable "image_name" {
  type = string
}
variable "traffic_percentile" {
  type    = number
  default = 100
}
variable "sql_instance_name" {
  type = string
}
variable "ingress" {
  type    = string
  default = "all"
  validation {
    error_message = "Ingress must be one of: [\"all\", \"internal\", \"internal-and-cloud-load-balancing\"]."
    condition     = contains(["all", "internal", "internal-and-cloud-load-balancing"], var.ingress)
  }
}
