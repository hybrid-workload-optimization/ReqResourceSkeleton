package main

import (
	"encoding/base64"
	"fmt"
	ys "main/ystruct"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

func MadeFinalWorkloadYAML(argAddr, argInputPath, argOutputPath string) {
	var err error
	data, err := os.ReadFile(argInputPath)
	check(err)

	var workflow ys.Workflow
	err = yaml.Unmarshal(data, &workflow)
	check(err)

	// yaml file encoding by base64
	encodedData := base64.StdEncoding.EncodeToString(data)

	//////////////////////////////////////////////////////////////
	// made resource request yaml file (send to kware)
	reqYaml := ys.ReqResource{}

	uuid := "dmkim"
	currentTime := time.Now()
	nowTime := currentTime.Format("2006-01-02 15:04:05")

	reqYaml.Version = "0.12"
	reqYaml.Request.Name = workflow.Metadata.GenerateName
	reqYaml.Request.ID = uuid
	reqYaml.Request.Date = nowTime

	size := len(workflow.Spec.Templates)
	containers := make([]ys.Container, size)

	for idx, value := range workflow.Spec.Templates {

		if value.Container == nil {
			// fmt.Println("NIL: " + value.Name)
			// DAG 처리 및 구조 정의 필요
			containers[idx].Name = value.Name
		} else {
			// fmt.Println("NOT NIL: " + value.Name)
			containers[idx].Name = value.Name
			containers[idx].Resources.Requests.CPU = value.Container.Resources.Requests.CPU
			containers[idx].Resources.Requests.CPU = value.Container.Resources.Requests.CPU
			containers[idx].Resources.Requests.GPU = value.Container.Resources.Requests.GPU
			containers[idx].Resources.Requests.Memory = value.Container.Resources.Requests.Memory
			containers[idx].Resources.Requests.EphemeralStorage = value.Container.Resources.Requests.EphemeralStorage

			containers[idx].Resources.Limits.CPU = value.Container.Resources.Limits.CPU
			containers[idx].Resources.Limits.GPU = value.Container.Resources.Limits.GPU
			containers[idx].Resources.Limits.Memory = value.Container.Resources.Limits.Memory
			containers[idx].Resources.Limits.EphemeralStorage = value.Container.Resources.Limits.EphemeralStorage
		}
	}
	reqYaml.Request.Containers = containers

	reqYaml.Request.Attribute.WorkloadType = "ML"
	reqYaml.Request.Attribute.IsCronJob = true
	reqYaml.Request.Attribute.DevOpsType = "DEV"
	reqYaml.Request.Attribute.GPUDriverVersion = 12.34
	reqYaml.Request.Attribute.CudaVersion = 342.12
	reqYaml.Request.Attribute.MaxReplicas = 5
	reqYaml.Request.Attribute.IsNetworking = false
	reqYaml.Request.Attribute.TotalSize = 875
	reqYaml.Request.Attribute.PredictedExecutionTime = 599
	reqYaml.Request.Attribute.UserID = uuid

	reqYaml.Yaml = encodedData
	//////////////////////////////////////////////////////////////

	ack, body := SEND_REST_DATA(argAddr, reqYaml)
	if ack.StatusCode == http.StatusOK {
		fmt.Println("=== Request successful ===")

		var ackBody ys.RespResource
		err = yaml.Unmarshal([]byte(body), &ackBody)
		check(err)

		workflow.Spec.NodeSelector.Node = ackBody.Response.Result.Node

		MakeYamlFile(workflow, argOutputPath)

		// send final yaml data to kware
		ack, _ := SEND_REST_DATA(BASE_URL+FINAL_YAML_PATH, workflow)
		if ack.StatusCode == http.StatusOK {
			fmt.Println("=== Done ===")
		} else {
			fmt.Printf("[FinalYAML] Request failed with status: %s\n", ack.Status)
		}
	} else {
		fmt.Printf("[ReqResource] Request failed with status: %s\n", ack.Status)
	}
}
