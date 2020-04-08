package services

import (
	"github.com/jenirainerpdx/emrConfigsCLI/domain"
	"strconv"
)

const CoreCount = 5

func GenerateCustomConfigs(workerCount int, nodeType string, updateAwsInstanceTypes bool) []domain.Configuration {
	nodeDetails := GetNodeTypeDetails(nodeType, updateAwsInstanceTypes)
	if nodeDetails.Memory == 0 {
		println("Request failed; invalid instance type.  " + nodeType)
		return nil
	}
	executorsPerInstance := GetExecutorsPerInstance(nodeDetails.VCPU)
	totalExecutorMemory := GetTotalExecutorMemory(nodeDetails.Memory, executorsPerInstance)
	// sparkYarnExecutorMemoryOverhead := (0.1 * float64(totalExecutorMemory))
	sparkMemory := GetSparkMemory(totalExecutorMemory)
	executorInstances := (executorsPerInstance * workerCount) - 1
	defaultParallelism := executorInstances * 2

	return []domain.Configuration{
		GetSparkDefaultConfig(CoreCount, sparkMemory, executorInstances, defaultParallelism),
	}
}

func GetSparkDefaultConfig(coreCount int, memory int64, executorInstances int, parallelism int) domain.Configuration {
	memoryString := (strconv.Itoa(int(memory))) + "M"
	return domain.Configuration{
		Classification: "spark-defaults",
		Properties: domain.SparkDefaults{
			NetworkTimeout:                "800s",
			ExecutorHeartbeatInterval:     "60s",
			DynamicAllocation:             "false",
			DriverMemory:                  memoryString,
			ExecutorMemory:                memoryString,
			ExecutorCores:                 strconv.Itoa(coreCount),
			Instances:                     strconv.Itoa(int(executorInstances)),
			MemoryFraction:                "0.80",
			StorageFraction:               "0.30",
			ExecutorExtraJavaOptions:      "-XX:+UseG1GC -XX:+UnlockDiagnosticVMOptions -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintTenuringDistribution -XX:OnOutOfMemoryError='kill -9 %p'",
			DriverExtraJavaOptions:        "-XX:+UseG1GC -XX:+UnlockDiagnosticVMOptions -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintTenuringDistribution -XX:OnOutOfMemoryError='kill -9 %p'",
			YarnReporterThreadMaxFailures: "5",
			StorageLevel:                  "MEMORY_AND_DISK_SER",
			RddCompress:                   "true",
			ShuffleCompress:               "true",
			SpillCompress:                 "true",
			DefaultParallelism:            strconv.Itoa(int(parallelism)),
			KyroBufferMax:                 "1028m",
		},
	}
}

func GetExecutorsPerInstance(vcpu int64) int {
	return int((vcpu - 1) / CoreCount)
}

func GetTotalExecutorMemory(memory int64, executorsPerInstance int) int64 {
	return (memory - 1) / int64(executorsPerInstance)
}

func GetSparkMemory(totalExecutorMemory int64) int64 {
	return int64(0.9 * float64(totalExecutorMemory))
}
