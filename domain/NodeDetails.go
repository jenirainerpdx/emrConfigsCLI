package domain

import "strings"

type NodeDetails struct {
	VCPU   int
	Memory int
}

func GetNodeTypeDetails(nodeType string) NodeDetails {
	switch strings.ToLower(nodeType) {
	case "r5.2xlarge":
		return NodeDetails{
			VCPU:   8,
			Memory: 64,
		}
	case "r5.4xlarge":
		return NodeDetails{
			VCPU:   16,
			Memory: 128,
		}
	case "r5.8xlarge":
		return NodeDetails{
			VCPU:   32,
			Memory: 256,
		}
	case "r5.12xlarge":
		return NodeDetails{
			VCPU:   48,
			Memory: 384,
		}
	case "r5.16xlarge":
		return NodeDetails{
			VCPU:   64,
			Memory: 512,
		}
	case "c5.2xlarge":
		return NodeDetails{
			VCPU:   8,
			Memory: 16,
		}
	case "c5.4xlarge":
		return NodeDetails{
			VCPU:   16,
			Memory: 32,
		}
	case "c5.9xlarge":
		return NodeDetails{
			VCPU:   36,
			Memory: 72,
		}
	default:
		{
			println("Node type was not supported:  " + nodeType + "  Defaulting to r5.8xlarge")
			return NodeDetails{
				VCPU: 32,
				Memory:256,
			}
		}
	}
}

