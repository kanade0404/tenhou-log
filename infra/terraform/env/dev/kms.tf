module "kms" {
  source   = "../../modules/kms"
  ENV      = local.ENV
  location = local.location
}
