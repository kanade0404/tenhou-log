module "router" {
  source     = "../../modules/router"
  network_id = module.vpc.network_id
}
