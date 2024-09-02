package main

import (
	"fmt"
	"net/http"
)

var BASE_URL = "http://10.0.2.210:9001"
var RESOURCE_PATH = "/resource"
var FINAL_YAML_PATH = "/final"

func main() {

	restURL := BASE_URL + RESOURCE_PATH

	fileName := "ml-comparison-pipeline-1.0-before"
	// fileName := "sample"

	oriPath := "./input/" + fileName + ".yaml"
	resultPath := "./output/" + fileName + "-after.yaml"

	// first. request information by node allocator
	ackBody := ReqResourceAllocInfo(restURL, oriPath)

	// second. made final workload yaml by response data
	finalYaml := MadeFinalWorkloadYAML(ackBody, oriPath, resultPath)

	// third. make workload yaml file
	MakeYamlFile(finalYaml, resultPath)

	// forth. send deploied yaml data to node allocator (kware).
	ack, _ := SEND_REST_DATA(BASE_URL+FINAL_YAML_PATH, finalYaml)
	if ack.StatusCode == http.StatusOK {
		fmt.Println("=== Done ===")
	} else {
		fmt.Printf("[FinalYAML] Request failed with status: %s\n", ack.Status)
	}
}
