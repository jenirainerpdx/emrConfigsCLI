package domain_test

import (
	"github.com/jenirainerpdx/emrConfigsCLI/domain"
	"testing"
)

var nodeDetailsTests = []struct {
	instanceType string
	expected domain.NodeDetails
}{
	{"R5.2xlarge", domain.NodeDetails{
		VCPU:   8,
		Memory: 64,
	}},
	{"r5.8Xlarge", domain.NodeDetails{
		VCPU:   32,
		Memory: 256,
	}},
	{"R5.16XLARGE", domain.NodeDetails{
		VCPU:   64,
		Memory: 512,
	}},
	{"c5.2xlarge", domain.NodeDetails{
		VCPU:   8,
		Memory: 16,
	}},
	{"c5.9xlarge", domain.NodeDetails{
		VCPU:   36,
		Memory: 72,
	}},
	{"invalid type", domain.NodeDetails{
		VCPU:   32,
		Memory: 256,
	}},
}

func TestGetNodeDetails(t *testing.T) {
	for _, tt := range nodeDetailsTests {
		actual := domain.GetNodeTypeDetails(tt.instanceType)
		if actual != tt.expected {
			t.Errorf("GetNodeDetails for %s:  expected %+v, actual %+v", tt.instanceType, tt.expected, actual)
		}
	}
}

