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

	// Container 키의 개수 확인
	containerCount := 0
	for _, template := range workflow.Spec.Templates {
		if template.Container != nil {
			containerCount++
		}
	}

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

	containers := make([]ys.Container, containerCount)

	for _, value := range workflow.Spec.Templates {

		if value.Container == nil {
			// fmt.Println("NIL: " + value.Name)

			continue
		} else {
			// fmt.Println("NOT NIL: " + value.Name)

			tmpContainer := ys.Container{
				Name: value.Name,
				Resources: ys.Resources{
					Requests: ys.ResourceDetails{
						CPU:              value.Container.Resources.Requests.CPU,
						GPU:              value.Container.Resources.Requests.GPU,
						Memory:           value.Container.Resources.Requests.Memory,
						EphemeralStorage: value.Container.Resources.Requests.EphemeralStorage,
					},
					Limits: ys.ResourceDetails{
						CPU:              value.Container.Resources.Limits.CPU,
						GPU:              value.Container.Resources.Limits.GPU,
						Memory:           value.Container.Resources.Limits.Memory,
						EphemeralStorage: value.Container.Resources.Limits.EphemeralStorage,
					},
				},
			}

			containers = append(containers, tmpContainer)
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

	reqYaml.Request.Attribute.Yaml = encodedData
	//////////////////////////////////////////////////////////////

	ack, body := SEND_REST_DATA(argAddr, reqYaml)
	if ack.StatusCode == http.StatusOK {
		fmt.Println("=== Request successful ===")

		var ackBody ys.RespResource
		err = yaml.Unmarshal([]byte(body), &ackBody)
		check(err)

		// log.Println(body)
		for _, val := range ackBody.Response.Container {
			for idx, _ := range workflow.Spec.Templates {
				if workflow.Spec.Templates[idx].Name == val.Name {
					workflow.Spec.Templates[idx].NodeSelector.Node = val.Node
				}
			}
		}

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
