package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	ys "main/ystruct"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

func ReqResourceAllocInfo(argAddr, argInputPath string) ys.RespResource {
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
			reqYaml.Request.Containers = append(reqYaml.Request.Containers, tmpContainer)
		}
	}

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

	var ackBody ys.RespResource

	ack, body := SEND_REST_DATA(argAddr, reqYaml)
	if ack.StatusCode == http.StatusOK {
		// fmt.Println("=== Request successful ===")
		err = yaml.Unmarshal([]byte(body), &ackBody)
		check(err)

	} else {
		fmt.Printf("[ReqResource] Request failed with status: %s\n", ack.Status)
	}

	return ackBody
}

func MadeFinalWorkloadYAML(argBody ys.RespResource, argInputPath, argOutputPath string) map[string]interface{} {

	// YAML 파일 읽기
	yamlFile, err := ioutil.ReadFile(argInputPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// YAML 데이터를 저장할 변수
	var data map[string]interface{}

	// YAML 데이터 언마샬링
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML data: %v", err)
	}

	// templates 섹션에서 모든 container의 image 값을 출력하고 조건에 따라 새로운 키를 추가
	spec, ok := data["spec"].(map[interface{}]interface{})
	if ok {
		templates, ok := spec["templates"].([]interface{})
		if ok {
			for _, template := range templates {
				templateMap, ok := template.(map[interface{}]interface{})
				if ok {
					for _, val := range argBody.Response.Container {
						if templateMap["name"] == val.Name {
							templateMap["nodeSelector"] = ys.NodeSelect{
								Node: val.Node,
							}
						}
					}
				}
			}
		}
	}

	return data
}
