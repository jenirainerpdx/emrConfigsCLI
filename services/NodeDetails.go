package services

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type NodeDetail struct {
	VCPU   int64
	Memory int64
}

type nodeTypesMap map[string]NodeDetail

// this is not a singleton.  It can be updated and is not immutable.
var instance nodeTypesMap
var once sync.Once

func getNodeTypesMap() nodeTypesMap {
	once.Do(func() {
		instance = buildNodeTypesMap()
	})

	return instance
}

func buildNodeTypesMap() nodeTypesMap {
	nodeMap := make(map[string]NodeDetail)
	data := GetInstanceTypesJson()
	var instanceTypes InstanceTypeBody

	marshallErr := json.Unmarshal(data, &instanceTypes)
	if marshallErr != nil {
		fmt.Println("error unmarshalling file into instance types:  ", marshallErr)
	}

	for _, instanceType := range instanceTypes.InstanceTypes {

		node := NodeDetail{
			VCPU:   instanceType.VCpuInfo.DefaultVCpus,
			Memory: instanceType.MemoryInfo.SizeInMiB,
		}

		nodeMap[instanceType.InstanceType] = node
	}

	return nodeMap
}

func GetNodeTypeDetails(nodeType string, updateMap bool) NodeDetail {
	nodeTypeLC := strings.ToLower(nodeType)
	var lookupMap map[string]NodeDetail
	if updateMap {
		lookupMap = RefreshNodeDetailsFromAws()
	} else {
		lookupMap = getNodeTypesMap()
	}
	details, ok := lookupMap[nodeTypeLC]
	if !ok {
		handleError(details, lookupMap)
	}
	return details
}

func handleError(detail NodeDetail, lookupMap map[string]NodeDetail) {
	if detail.Memory == 0 {
		println("Invalid instance type was entered.  Available instance types include:  ")
		var names []string
		for name := range lookupMap {
			names = append(names, name)
		}

		sort.Strings(names)

		fmt.Println("Valid values:", names)

	}
}
