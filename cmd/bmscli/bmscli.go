package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/hawkingrei/bms/model"
	"github.com/shirou/gopsutil/cpu"
)

func flagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("cli", flag.ExitOnError)
	flagSet.String("f", "", "flag")
	return flagSet
}

func main() {
	flagSet := flagSet()
	flagSet.Parse(os.Args[1:])
	flagString := flagSet.Lookup("f").Value.String()
	// Read data from JSON file
	data, err := os.ReadFile("result.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	uploadOrigin(data, flagString)
}

func uploadOrigin(data []byte, f string) {
	var benchOutput BenchOutput
	json.Unmarshal(data, &benchOutput)
	if f != "" {
		benchOutput.TaskName = f
	}
	benchOutput.Arch = runtime.GOARCH
	row, err := json.Marshal(&benchOutput)
	if err != nil {
		log.Fatalf("Failed to Marshal json: %v", err)
	}

	url := "https://gobench.hawkingrei.com/saveData"
	token := "TPhXfzNEpD6DNQn8CJ6z"
	// Create new http request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(row))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	// Create a Client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	// Close the response body
	defer resp.Body.Close()

	// Log the status
	log.Println("Response Status:", resp.Status)
}

func uploadBms(data []byte, f string) {
	var benchOutput model.BenchOutput
	json.Unmarshal(data, &benchOutput)
	if f != "" {
		benchOutput.Flag = f
	}
	benchOutput.Arch = runtime.GOARCH
	benchOutput.CompilerVersion = runtime.Version()
	cpuinfo, err := cpu.Info()
	if err != nil {
		log.Fatalf("Failed to get cpu info: %v", err)

	}
	benchOutput.CPUInfo = cpuinfo[0].ModelName
	fmt.Println(benchOutput)
}
