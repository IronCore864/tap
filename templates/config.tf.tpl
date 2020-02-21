provider "aws" {
  version = "~> 2.0"
  region  = "{{.Region}}"
}

terraform {
  required_version = ">= 0.12"
  backend "s3" {
    bucket  = "{{.TFStateBucketName}}"
    key     = "{{.TFStateFileName}}"
    region  = "{{.Region}}"
    encrypt = true
  }
}
