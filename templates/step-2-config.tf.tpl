provider "aws" {
  version = "~> 2.0"
  region  = "{{.region}}"
}

terraform {
  required_version = ">= 0.12"
  backend "s3" {
    bucket  = "{{.tf_state_bucket}}"
    key     = "{{.step_2_s3_key}}"
    region  = "{{.region}}"
    encrypt = true
  }
}
