variable "PROJECT_ID" {
  type = string
}
variable "topic_name" {
  type = string
}
variable "subscription_name" {
  type = string
}
variable "ack_deadline_seconds" {
  type    = number
  default = 20
}
variable "push_endpoint" {
  type = string
}
