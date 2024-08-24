package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	ys "main/ystruct"

	"gopkg.in/yaml.v2"
)

func check(argErr error) {
	if argErr != nil {
		log.Fatalf("Error: %v", argErr)
	}
}

func SEND_REST_DATA(argAddr string, argYamlData interface{}) (*http.Response, string) {

	yamlData, err := yaml.Marshal(&argYamlData)
	check(err)

	resp, err := http.Post(argAddr, "application/x-yaml", bytes.NewBuffer(yamlData))
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return resp, string(body)
}

func ReadYaml(argPath string) {
	// Open the YAML file
	file, err := os.Open(argPath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file content
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Unmarshal the YAML data into the struct
	var config ys.Example
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v\n", err)
		return
	}

	// Print the content of the struct
	fmt.Printf("Server: %s\n", config.Server)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Database Username: %s\n", config.Database.Username)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
}

func MakeYamlFile(argData interface{}, argPath string) {

	// Write the YAML data to a file
	file, err := os.Create(argPath)
	if err != nil {
		fmt.Printf("Error while creating file: %v\n", err)
		return
	}
	defer file.Close()

	// YAML로 직렬화 (serialize)하고 파일에 저장
	encoder := yaml.NewEncoder(file)
	// encoder.SetIndent(2) // YAML 파일의 가독성을 위해 인덴트를 설정합니다.
	err = encoder.Encode(argData)
	if err != nil {
		log.Fatalf("Error encoding YAML to file: %v", err)
	}

	fmt.Println("YAML file created successfully.")
}
