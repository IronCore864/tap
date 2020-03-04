# Universal Template Tool

Naming: unitet means UNIversal TEmplate Tool.

The naming convention here is my way of paying tribute to redis (REmote DIctionary Server).

It's easy to pronounce, say, than, "utt" (Universal Template Tool), or its original name "tftpl" (TerraForm TemPLate tool, which is why it was created in the first place).

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
$ ./unitet -h
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
./unitet -inputFile ./test/input/test.yaml -outputDir . -templateDir ./test/templates/
```

## Build Docker

```
docker build -t ironcore864/unitet:latest .
docker push ironcore864/unitet:latest
```

## Run in Docker

```
docker pull ironcore864/unitet:latest
docker run -it ironcore864/unitet:latest
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
				sh '/app/tftpl -inputFile ./${USER_INPUT_FILE} -outputDir ./${OUTPUT_DIR} -templateDir ./${TEMPLATE_DIR}'
			}
		}
		stage("XXX") {
			// ...
		}
	}
}
```
