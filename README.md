# Template Tool in Golang for All Purpose

Naming: tap means Template All Product.

It's a tool written in Golang, which takes yaml (or json) as input, and renders templates in Golang syntax.

Can be used for all kinds of purposes, hence the name, TAP.

## Purpose

Using input values (defined in a single yaml file) and templates (all golang tempaltes under a given directory), render every template and write the output to files.

It was originally designed to generate config.tf and terraform.tfvars for usage of terraform, but it can be for sure used for other purposes.

## Example Input and Templates

Input example see `test/input/test.yaml`.

Template examples see `test/templates/test.tfvars.tpl`.

Both are relatively simple examples, but golang template full feature are supported. For example, to render a yaml list into an array with square brackets and comma separated items:

```
arr = [{{ range  $i, $e := .list }}{{ if $i }}, {{ end }}"{{ $e }}"{{ end }}]
```

## Dependencies

- Golang 1.13.

Tested with golang `1.13.4 darwin/amd64`, as well as `1.14` in CircleCI.

Other dependencies can be fetched via go module command `go get`, see below.

## Build

```
go get ./...
go build
```

## Test

```
go test ./...
```

## Usage

```
$ ./tap -h
Usage of ./tap:
  -inputFile string
        input file path and name, example: ./input/eks.yaml (default "./input/eks.yaml")
  -outputDir string
        output directory, example: ./output (default ".")
  -template string
        if given a file, the file of the template to be rendered; if given a directory, all templates under the directory will be rendered (default "./templates")
```

## Run

```
./tap -inputFile ./test/input/test.yaml -outputDir . -template ./test/templates/
# template can be either a directory (in which case, all templates under it will be rendered), or a file
./tap -inputFile ./test/input/test.yaml -outputDir . -template ./test/templates/a.tpl
```

## Build Docker

```
docker build -t ironcore864/tap:latest .
docker push ironcore864/tap:latest
```

## Run in Docker

```
docker pull ironcore864/tap:latest
docker run -it ironcore864/tap:latest
```

## Use with Jenkins Pipeline with Kubernetes Plugin in a Multi-Container Pipeline

```
podTemplate(label: 'pipeline', cloud: 'kubernetes', containers: [
  containerTemplate(name: 'tftpl', 
    image: 'docker.io/ironcore864/tftpl:latest', 
    alwaysPullImage: true,
    ttyEnabled: true,
    args: '',
    command: 'tail -f /dev/null',
    resourceRequestCpu: '100m',
    resourceLimitCpu: '200m',
    resourceRequestMemory: '200Mi',
    resourceLimitMemory: '400Mi'),
  containerTemplate(name: 'jnlp',
    image: 'docker.io/jenkinsci/jnlp-slave:alpine',
    command: '',
    args: '${computer.jnlpmac} ${computer.name}',
    resourceRequestCpu: '50m',
    resourceLimitCpu: '600m',
    resourceRequestMemory: '100Mi',
    resourceLimitMemory: '500Mi')
  ]) {
  node('pipeline') {
    stage("Render Templates") {
      container('tftpl') {
        unstash 'userinput'
        sh '/app/tftpl -inputFile ./${USER_INPUT_FILE} -outputDir ./${OUTPUT_DIR} -template ./${TEMPLATE_DIR}'
      }
    }
    stage("XXX") {
      // ...
    }
  }
}
```
