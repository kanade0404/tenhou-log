module "log_converter" {
  source      = "../../modules/function"
  entry_point = "Convert"
  location    = var.location
  name        = "log-converter"
  runtime     = "go120"
  storage_source = {
    bucket = "gcf-v2-sources-457582600725-asia-northeast1"
    object = "log-converter/function-source.zip"
  }
  PROJECT_ID = var.PROJECT_ID
  secrets = [
    {
      name = "PROJECT_ID"
    }
  ]
  service_account_email = module.service_account["log-converter-runner"].email
}
module "log_decompressor" {
  source      = "../../modules/function"
  entry_point = "Decompress"
  location    = var.location
  name        = "log-decompressor"
  runtime     = "go120"
  storage_source = {
    bucket = "gcf-v2-sources-457582600725-asia-northeast1"
    object = "log-decompressor/function-source.zip"
  }
  PROJECT_ID = var.PROJECT_ID
  secrets = [
    {
      name = "PROJECT_ID"
    }
  ]
  service_account_email = module.service_account["log-decompressor-runner"].email
}
module "log_downloader" {
  source      = "../../modules/function"
  entry_point = "Download"
  location    = var.location
  name        = "log-downloader"
  runtime     = "go120"
  storage_source = {
    bucket = "gcf-v2-sources-457582600725-asia-northeast1"
    object = "log-downloader/function-source.zip"
  }
  PROJECT_ID = var.PROJECT_ID
  secrets = [
    {
      name = "PROJECT_ID"
    }
  ]
  service_account_email = module.service_account["log-downloader-runner"].email
}
