package services

import (
	"github.com/jenirainerpdx/emrConfigsCLI/domain"
	"strconv"
)

const CoreCount = 5

func GenerateCustomConfigs(workerCount int, nodeType string) []domain.Configuration {
	nodeDetails := domain.GetNodeTypeDetails(nodeType)
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

func GetSparkDefaultConfig(coreCount int, memory int, executorInstances int, parallelism int) domain.Configuration {
	memoryString := (strconv.Itoa(memory)) + "M"
	return domain.Configuration{
		Classification: "spark-defaults",
		Properties: domain.SparkDefaults{
			NetworkTimeout:                "800s",
			ExecutorHeartbeatInterval:     "60s",
			DynamicAllocation:             "false",
			DriverMemory:                  memoryString,
			ExecutorMemory:                memoryString,
			ExecutorCores:                 strconv.Itoa(coreCount),
			Instances:                     strconv.Itoa(executorInstances),
			MemoryFraction:                "0.80",
			StorageFraction:               "0.30",
			ExecutorExtraJavaOptions:      "-XX:+UseG1GC -XX:+UnlockDiagnosticVMOptions -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintTenuringDistribution -XX:OnOutOfMemoryError='kill -9 %p'",
			DriverExtraJavaOptions:        "-XX:+UseG1GC -XX:+UnlockDiagnosticVMOptions -XX:+PrintGCDetails -XX:+PrintGCDateStamps -XX:+PrintTenuringDistribution -XX:OnOutOfMemoryError='kill -9 %p'",
			YarnReporterThreadMaxFailures: "5",
			StorageLevel:                  "MEMORY_AND_DISK_SER",
			RddCompress:                   "true",
			ShuffleCompress:               "true",
			SpillCompress:                 "true",
			DefaultParallelism:            strconv.Itoa(parallelism),
			KyroBufferMax:                 "1028m",
		},
	}
}

func GetExecutorsPerInstance(vcpu int) int {
	return (vcpu - 1) / CoreCount
}

func GetTotalExecutorMemory(memory int, executorsPerInstance int) int {
	return (memory - 1) / executorsPerInstance
}

func GetSparkMemory(totalExecutorMemory int) int {
	return int(0.9 * float64(totalExecutorMemory) * 1000)
}