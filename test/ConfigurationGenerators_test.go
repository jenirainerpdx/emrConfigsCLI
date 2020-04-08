package test

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
	vcpu int64
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
	memory int64
	executorsPerInstance int
	expected int64
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
	totalExecMemory int64
	expected int64
}{
	{1, 0},
	{15, 13},
	{23, 20},
	{85, 76},

}

func TestGetSparkMemory(t *testing.T) {
	for _, testCase := range sparkMemTests {
		actual := services.GetSparkMemory(testCase.totalExecMemory)
		if actual != testCase.expected {
			t.Errorf("GetSparkMemory for totalExecutorMemory of %d:  expected %d, actual %d",
				testCase.totalExecMemory, testCase.expected, actual)
		}
	}
}