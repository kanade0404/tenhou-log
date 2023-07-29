#module "gcs_obj_put_trigger" {
#  source     = "../../modules/eventarc"
#  PROJECT_ID = var.PROJECT_ID
#  events = [
#    {
#      name = "putted_compressed_log",
#      matching_criterion = [
#        {
#          attribute = "type"
#          value     = "google.cloud.storage.object.v1.finalized"
#        },
#        {
#          attribute = "bucket"
#          value     = module.tmp_tenhou_log_trail_bucket.name
#        }
#      ]
#      destinations = [
#        {
#          path    = ""
#          service = ""
#        }
#      ]
#    }
#  ]
#  location        = local.location
#  service_account = "${module.project.number}-compute@developer.gserviceaccount.com"
#}
