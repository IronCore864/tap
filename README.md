# Terraform Template

Naming: tftpl means "TerraForm TemPLating"

## Purpose

Templating `config.tf` and `terraform.tfvars`.

## Example Input and Templates
Input example see `input/eks.yaml.example`, all are mandatory fields.

Templates see `templates/*`.

## Build

Tested with `go version go1.13.4 darwin/amd64`

```
go get ./...
```

## Usage

```
Usage of ./tftpl:
  -configTemplate string
        config.tf template path and name (default "./templates/config.tf.tpl")
  -inputFile string
        input file path and name, example: ./input/eks.yaml (default "eks.yaml")
  -outputDir string
        output directory, example: ./output (default ".")
  -tfvarsTemplate string
        terraform.tfvars template path and name (default "./templates/terraform.tfvars.tpl")
```

## Run

```
./tftpl -inputFile ./input/eks.yaml -outputDir ./output
```

## Build Docker

```
docker build .
```

## Todo

- UT
