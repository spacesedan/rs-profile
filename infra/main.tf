provider "aws" {
  region = var.aws-region
}

terraform {
  cloud {
    organization = "rarityshark"

    workspaces {
      name = "rs-profile-api"
    }
  }
}

data "tfe_outputs" "platform" {
  organization = var.org-name
  workspace    = var.ws-name
}