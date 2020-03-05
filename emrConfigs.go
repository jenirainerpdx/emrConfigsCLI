package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	nodeCountPtr := flag.Int("nodes", 0, "Count of worker nodes.")
	instanceTypePtr := flag.String("instance", "r5.8xlarge", "Instance type to use for worker node clusters.")
	flag.Parse()

	if *nodeCountPtr == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("value of nodeCountPtr is: %d\nvalue of instanceTypePtr: %s\n",
		*nodeCountPtr, *instanceTypePtr)
}
