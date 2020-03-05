package services_test

import (
	"github.com/jenirainerpdx/emrConfigsCLI/domain"
	"github.com/jenirainerpdx/emrConfigsCLI/services"
	"testing"

)

func TestConfigurationGenerator(t *testing.T) {
	actual := services.GetSparkDefaultConfig(5, 300, 99, 40)
	actualProps, ok := actual.Properties.(domain.SparkDefaults)
	if !ok  {
		t.Errorf("Properties type did not cast successfully to SparkDefaults.")
	}

	if actualProps.DefaultParallelism != "40" {
		t.Errorf("DefaultParallelism expected to be 40, but was: %s", actualProps.DefaultParallelism)
	}

	if actualProps.DriverMemory != "300M" {
		t.Errorf("DriverMemory expected to be 300M, but was: %s", actualProps.DriverMemory)
	}

	if actualProps.ExecutorMemory != "300M" {
		t.Errorf("ExecutorMemory expected to be 300M, but was: %s", actualProps.ExecutorMemory)
	}


	if actualProps.ExecutorCores != "5" {
		t.Errorf("ExecutorCores expected to be 5, but was: %s", actualProps.ExecutorCores)
	}

	if actualProps.Instances != "99" {
		t.Errorf("Instances expected to be 99, but was: %s", actualProps.Instances)
	}
}

var execPerInstanceTests = []struct{
	vcpu int
	expected int
} {
	{8, 1},
	{16, 3},
	{32, 6},
	{48, 9},
	{64, 12},
}


func TestGetExecutorsPerInstance(t *testing.T) {
	for _, tt := range execPerInstanceTests {
		actual := services.GetExecutorsPerInstance(tt.vcpu)
		if actual != tt.expected {
			t.Errorf("GetExecutorsPerInstance for %d: expected %d, actual: %d", tt.vcpu, tt.expected, actual)
		}
	}
}

var totalMemoryTests = []struct{
	memory int
	executorsPerInstance int
	expected int
}{
	{16, 1, 15},
	{72, 3, 23},
	{512, 6, 85},
	{384, 9, 42},
	{16, 12, 1},
}

func TestGetTotalExecutorMemory(t *testing.T) {
	for _, tt := range totalMemoryTests {
		actual := services.GetTotalExecutorMemory(tt.memory, tt.executorsPerInstance)
		if actual != tt.expected {
			t.Errorf("GetTotalExecutorMemory for mem= %d, executors per instance = %d:  expected %d, actual %d",
				tt.memory, tt.executorsPerInstance, tt.expected, actual)
		}
	}
}

var sparkMemTests = []struct{
	totalExecMemory int
	expected int
}{
	{1, 900},
	{15, 13500},
	{23, 20700},
	{85, 76500},

}

func TestGetSparkMemory(t *testing.T) {
	for _, tt := range sparkMemTests {
		actual := services.GetSparkMemory(tt.totalExecMemory)
		if actual != tt.expected {
			t.Errorf("GetSparkMemory for totalExecutorMemory of %d:  expected %d, actual %d",
				tt.totalExecMemory, tt.expected, actual)
		}
	}
}