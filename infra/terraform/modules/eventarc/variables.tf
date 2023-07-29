variable "PROJECT_ID" {
  type = string
}
variable "location" {
  type = string
}
variable "events" {
  type = list(object({
    name = string
    matching_criterion = list(object({
      attribute = string
      value     = string
    }))
    destinations = list(object({
      path    = string
      service = string
    }))
    service_account = string
  }))
}
