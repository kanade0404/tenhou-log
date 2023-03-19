variable "ENV" {
  type = string
}
variable "location" {
  type = string
}

locals {
  keys = ["terraform", "scraper"]
}
