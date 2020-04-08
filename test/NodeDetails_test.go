package test

import (
	"fmt"
	"github.com/jenirainerpdx/emrConfigsCLI/services"
	"testing"
)

var nodeDetailsTests = []struct {
	instanceType string
	expected     services.NodeDetail
}{
	{"R5.2xlarge", services.NodeDetail{
		VCPU:   8,
		Memory: 65536,
	}},
	{"r5.8Xlarge", services.NodeDetail{
		VCPU:   32,
		Memory: 262144,
	}},
	{"R5.16XLARGE", services.NodeDetail{
		VCPU:   64,
		Memory: 524288,
	}},
	{"c5.2xlarge", services.NodeDetail{
		VCPU:   8,
		Memory: 16384,
	}},
	{"c5.9xlarge", services.NodeDetail{
		VCPU:   36,
		Memory: 73728,
	}},
	{"invalid type", services.NodeDetail{
		VCPU:   0,
		Memory: 0,
	}},
}

func TestGetNodeDetails(t *testing.T) {
	for _, testCase := range nodeDetailsTests {
		actual := services.GetNodeTypeDetails(testCase.instanceType, false)
		if actual != testCase.expected {
			t.Errorf("GetNodeDetails for %s:  expected %+v, actual %+v\n", testCase.instanceType, testCase.expected, actual)
		}
	}
}

func TestGetNodeDetailsWithRefresh(t *testing.T) {
	for _, testCase := range nodeDetailsTests {
		actual := services.GetNodeTypeDetails(testCase.instanceType, true)
		if actual != testCase.expected {
			t.Errorf("GetNodeDetails while calling refresh for %s:  expected %+v, actual %+v", testCase.instanceType, testCase.expected, actual)
		} else {
			fmt.Printf("Pass: %s:  actual: %+v", testCase.instanceType, actual)
		}
	}
}

