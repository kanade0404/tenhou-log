resource "google_pubsub_topic" "topic" {
  name    = var.topic_name
  project = var.PROJECT_ID
}
resource "google_pubsub_subscription" "subscription" {
  for_each             = { for sub in var.subscriptions : sub.name => sub }
  name                 = each.key
  topic                = google_pubsub_topic.topic.name
  ack_deadline_seconds = each.value.ack_deadline_seconds
  dynamic "retry_policy" {
    for_each = each.value.retry_policy == null ? [] : [1]
    content {
      maximum_backoff = each.value.retry_policy.maximum_backoff
      minimum_backoff = each.value.retry_policy.minimum_backoff
    }
  }
  dynamic "dead_letter_policy" {
    for_each = each.value.dead_letter_policy == null ? [] : [1]
    content {
      dead_letter_topic     = each.value.dead_letter_policy.dead_letter_topic_id
      max_delivery_attempts = each.value.dead_letter_policy.max_deliver_attempts
    }
  }
}
