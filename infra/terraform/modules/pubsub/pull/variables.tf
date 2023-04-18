variable "PROJECT_ID" {
  type = string
}
variable "topic_name" {
  type = string
}
variable "subscriptions" {
  type = list(object({
    name                 = string
    ack_deadline_seconds = optional(number, 10)
    retry_policy = optional(object({
      minimum_backoff = optional(string, "10s")
      maximum_backoff = optional(string, "600s")
    }))
    dead_letter_policy = optional(object({
      dead_letter_topic_id = string
      max_deliver_attempts = optional(number, 5)
    }))
  }))
}
