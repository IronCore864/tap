# Terraform Template

Naming: tftpl means "TerraForm TemPLating"

## Purpose

Templating all files under a given template directory, using input file which is in YAML format.

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
  -inputFile string
        input file path and name, example: ./input/eks.yaml (default "./input/eks.yaml")
  -outputDir string
        output directory, example: ./output (default ".")
  -templateDir string
        directory containing all templates to be rendered, example: ./templates (default "./templates")
```

## Run

```
./tftpl -inputFile ./input/eks.yaml -outputDir . -templateDir ./templates/
```

## Build Docker

```
docker build -t ironcore864/tftpl:latest .
docker push ironcore864/tftpl:latest
```

## Todo

- UT
