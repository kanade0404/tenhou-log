variable "name" {
  type = string
}
variable "region" {
  type = string
}
variable "description" {
  type = string
}
variable "service_account_id" {
  type = string
}
variable "PROJECT_ID" {
  type = string
}
variable "source_content_vars" {
  type = object({
    scraper_url          = string
    log_converter_url    = string
    log_decompressor_url = string
    log_downloader_url   = string
  })
}
