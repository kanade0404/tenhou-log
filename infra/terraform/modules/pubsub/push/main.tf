resource "google_pubsub_topic" "topic" {
  name    = var.topic_name
  project = var.PROJECT_ID
}
resource "google_pubsub_subscription" "subscription" {
  name                 = var.subscription_name
  topic                = google_pubsub_topic.topic.name
  ack_deadline_seconds = var.ack_deadline_seconds
  push_config {
    push_endpoint = var.push_endpoint
  }
}
