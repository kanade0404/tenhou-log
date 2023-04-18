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
variable "ingress" {
  type    = string
  default = "all"
  validation {
    error_message = "Ingress must be one of: [\"all\", \"internal\", \"internal-and-cloud-load-balancing\"]."
    condition     = contains(["all", "internal", "internal-and-cloud-load-balancing"], var.ingress)
  }
}

variable "port" {
  type = number
}

variable "secrets" {
  type = list(object({
    name    = string
    version = optional(string)
  }))
  default = []
}
variable "iam_members" {
  type = list(object({
    name = string
    role = string
  }))
  default = []
}
variable "service_account_name" {
  type = string
}
variable "mapping_domains" {
  type    = list(string)
  default = []
}
