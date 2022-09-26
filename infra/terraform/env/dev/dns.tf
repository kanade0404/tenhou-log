locals {
  name   = "touka-ryuumonbuchi"
  domain = "${local.name}.com"
}
module "dns" {
  source        = "../../modules/dns/managed_zone"
  PROJECT_ID    = var.PROJECT_ID
  name          = local.name
  dns_name      = local.domain
  is_visible    = true
  state         = "on"
  non_existence = "nsec"
  depends_on    = [module.enabled_services.services]
}

module "soa_record" {
  source        = "../../modules/dns/record_set"
  PROJECT_ID    = var.PROJECT_ID
  dns_zone_name = "touka-ryuumonbuchi"
  name          = module.dns.dns_name
  record_type   = "SOA"
  rrdatas = [
    "ns-cloud-d1.googledomains.com. cloud-dns-hostmaster.google.com. 1 21600 3600 259200 300"
  ]
  ttl = 21600
}
module "ns_record" {
  source        = "../../modules/dns/record_set"
  PROJECT_ID    = var.PROJECT_ID
  name          = module.dns.dns_name
  dns_zone_name = "touka-ryuumonbuchi"
  record_type   = "NS"
  rrdatas = [
    "ns-cloud-d1.googledomains.com.",
    "ns-cloud-d2.googledomains.com.",
    "ns-cloud-d3.googledomains.com.",
    "ns-cloud-d4.googledomains.com."
  ]
  ttl = 21600
}
module "a_record" {
  source        = "../../modules/dns/record_set"
  PROJECT_ID    = var.PROJECT_ID
  name          = "www.${module.dns.dns_name}"
  dns_zone_name = "touka-ryuumonbuchi"
  record_type   = "A"
  rrdatas       = [module.lb-http-serverless_neg.global_ipv4_address]
  ttl           = 21600
}

module "a__record" {
  source        = "../../modules/dns/record_set"
  PROJECT_ID    = var.PROJECT_ID
  name          = module.dns.dns_name
  dns_zone_name = "touka-ryuumonbuchi"
  record_type   = "A"
  rrdatas       = [module.lb-http-serverless_neg.global_ipv4_address]
  ttl           = 21600
}
