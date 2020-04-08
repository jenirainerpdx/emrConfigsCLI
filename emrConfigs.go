package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jenirainerpdx/emrConfigsCLI/services"
	"os"
)

func main() {
	nodeCountPtr := flag.Int("nodes", 0, "Count of worker nodes.")
	instanceTypePtr := flag.String("instance", "r5.8xlarge", "Instance type to use for worker node clusters.")
	updateFromAWS := flag.Bool("update", false, "Set this flag to true if you want to refresh the default set of Instance data with a call to AWS cli")
	flag.Parse()

	if *nodeCountPtr == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	configs := services.GenerateCustomConfigs(*nodeCountPtr, *instanceTypePtr, *updateFromAWS)

	configJson, err := json.Marshal(configs)

	if err != nil {
		fmt.Println("An error occurred trying to create json:  ", err)
	}

	fmt.Printf("%s\n", string(configJson))
}
