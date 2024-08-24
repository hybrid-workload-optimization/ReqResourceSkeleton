package main

var BASE_URL = "http://10.0.2.210:9001"
var RESOURCE_PATH = "/resource"
var FINAL_YAML_PATH = "/final"

func main() {

	restURL := BASE_URL + RESOURCE_PATH
	oriPath := "./input/sample.yaml"
	resultPath := "./output/result.yaml"

	MadeFinalWorkloadYAML(restURL, oriPath, resultPath)

}
