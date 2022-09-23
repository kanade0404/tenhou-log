variable "PROJECT_ID" {
  type = string
}
variable "name" {
  type = string
}
variable "dns_name" {
  type = string
}
variable "is_visible" {
  type    = bool
  default = false
}
variable "state" {
  type = string
  validation {
    error_message = "state must be equal \"on\" or \"off\" or \"transfer\"."
    condition     = contains(["on", "off", "transfer", null], var.state)
  }
}
variable "non_existence" {
  type = string
  validation {
    error_message = "non_existence must be equal \"nsec\" or \"nsec3\"."
    condition     = contains(["nsec", "nsec3"], var.non_existence)
  }
}
