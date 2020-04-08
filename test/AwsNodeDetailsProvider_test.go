package test

import (
	"github.com/jenirainerpdx/emrConfigsCLI/services"
	"testing"
)

func TestRefreshNodeDetails(t *testing.T) {

	actual := services.RefreshNodeDetailsFromAws()

	expectedR54xlarge := services.NodeDetail{
		VCPU:   16,
		Memory: 131072,
	}

	actualR54 := actual["r5.4xlarge"]

	if actualR54 != expectedR54xlarge {
		t.Errorf("Expected: %v  Got: %v", expectedR54xlarge, actualR54)
	}
}

