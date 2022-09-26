terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.37.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.4.3"
    }
  }
}
provider "google" {
  credentials = var.GOOGLE_CREDENTIALS
  project     = var.PROJECT_ID
  region      = var.location
  zone        = var.default_zone
}

provider "google-beta" {
  credentials = var.GOOGLE_CREDENTIALS
  project     = var.PROJECT_ID
  region      = var.location
  zone        = var.default_zone
}

module "project" {
  source     = "../../modules/project"
  PROJECT_ID = var.PROJECT_ID
}

module "enabled_services" {
  source     = "../../modules/services"
  PROJECT_ID = var.PROJECT_ID
}

module "random_id" {
  source = "../../modules/random_id"
}
