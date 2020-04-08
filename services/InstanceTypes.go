package services

type InstanceTypeBody struct {
	InstanceTypes []InstanceType
}

type InstanceType struct {
	InstanceType string
	VCpuInfo VcpuInfo
	MemoryInfo MemoryInfo
}

type VcpuInfo struct {
	DefaultVCpus int64
}

type MemoryInfo struct {
	SizeInMiB int64
}

func GetInstanceTypesJson() []byte {
	return []byte(`
{
  "InstanceTypes": [
    {
      "InstanceType": "m5d.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m5ad.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "c5d.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 98304
      }
    },
    {
      "InstanceType": "m5a.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "m2.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 70041
      }
    },
    {
      "InstanceType": "p3.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      }
    },
    {
      "InstanceType": "m5d.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "c1.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 7168
      }
    },
    {
      "InstanceType": "m5.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "r3.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "unsupported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "r5a.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "m5ad.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "d2.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      }
    },
    {
      "InstanceType": "a1.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 4096
      }
    },
    {
      "InstanceType": "m5.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "m5.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "t3a.small",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 2048
      }
    },
    {
      "InstanceType": "r5ad.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "m5.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "r5dn.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "c3.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 61440
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "unsupported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "m4.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "m5a.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "c5d.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 4096
      }
    },
    {
      "InstanceType": "m5dn.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "r5.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "m5dn.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "m5dn.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m5d.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "t1.micro",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ]
      },
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 627
      }
    },
    {
      "InstanceType": "d2.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 36
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "m5dn.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "m5ad.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "g2.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 15360
      },
      "GpuInfo": {
        "Gpus": [
          {
            "Name": "K520",
            "Manufacturer": "NVIDIA",
            "Count": 1,
            "MemoryInfo": {
              "SizeInMiB": 4096
            }
          }
        ],
        "TotalGpuMemoryInMiB": 4096
      }
    },
    {
      "InstanceType": "c4.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 30720
      }
    },
    {
      "InstanceType": "r5a.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "t3.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "c5.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "f1.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      }
    },
    {
      "InstanceType": "r5.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "g3.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 499712
      }
    },
    {
      "InstanceType": "t3a.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "r5n.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "x1e.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 999424
      }
    },
    {
      "InstanceType": "t3.medium",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 4096
      }
    },
    {
      "InstanceType": "c4.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 3840
      }
    },
    {
      "InstanceType": "c5d.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "c5.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "t3.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m5ad.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "r5d.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "m5.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "g3.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "m5dn.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "c5d.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r4.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 499712
      }
    },
    {
      "InstanceType": "t3a.medium",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 4096
      }
    },
    {
      "InstanceType": "z1d.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "m5ad.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "c3.large",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ],
        "SustainedClockSpeedInGhz": 2.8
      },
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 3840
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "unsupported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "r5n.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "r5ad.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          12,
          18,
          24,
          36,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "h1.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "r3.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 15360
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "unsupported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "m5a.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "inf1.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m5d.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "r5ad.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "c1.medium",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ]
      },
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 1740
      }
    },
    {
      "InstanceType": "c5n.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 21504
      }
    },
    {
      "InstanceType": "m4.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "c4.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 36
      },
      "MemoryInfo": {
        "SizeInMiB": 61440
      }
    },
    {
      "InstanceType": "x1.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 999424
      }
    },
    {
      "InstanceType": "g4dn.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r5a.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "x1e.32xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 128
      },
      "MemoryInfo": {
        "SizeInMiB": 3997696
      }
    },
    {
      "InstanceType": "p2.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 499712
      }
    },
    {
      "InstanceType": "m5d.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "t2.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "t3.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "a1.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m2.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 35020
      }
    },
    {
      "InstanceType": "m5n.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r5a.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "a1.medium",
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 2048
      }
    },
    {
      "InstanceType": "i2.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "m5n.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "m5n.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "x1e.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "r5.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24,
          26,
          28,
          30,
          32,
          34,
          36,
          38,
          40,
          42,
          44,
          46,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "i3en.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "t2.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "c5.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 98304
      }
    },
    {
      "InstanceType": "i3en.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "c5n.9xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 36
      },
      "MemoryInfo": {
        "SizeInMiB": 98304
      }
    },
    {
      "InstanceType": "c5n.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 5376
      }
    },
    {
      "InstanceType": "c5n.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 43008
      }
    },
    {
      "InstanceType": "i3en.3xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 12,
        "ValidCores": [
          2,
          4,
          6
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 98304
      }
    },
    {
      "InstanceType": "r5a.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "z1d.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "i2.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "t2.medium",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ],
        "SustainedClockSpeedInGhz": 2.3
      },
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 4096
      }
    },
    {
      "InstanceType": "c5d.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "r5ad.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "c5.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r3.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 31232
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "h1.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "r5n.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m4.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "r5d.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "r3.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "r4.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 31232
      }
    },
    {
      "InstanceType": "m5n.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "inf1.6xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 24,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 49152
      }
    },
    {
      "InstanceType": "c5.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 4096
      }
    },
    {
      "InstanceType": "z1d.6xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 24,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r5ad.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "a1.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "inf1.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "m5dn.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "r5.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "r5.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "c3.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 7680
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "r5n.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "m5.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "x1e.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 1998848
      }
    },
    {
      "InstanceType": "a1.metal",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "r5a.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "m5.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m4.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "z1d.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "z1d.metal",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "r5.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "m5d.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24,
          26,
          28,
          30,
          32,
          34,
          36,
          38,
          40,
          42,
          44,
          46,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "t3a.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m5a.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "r5dn.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "t2.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "c3.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 30720
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "h1.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "i2.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 31232
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "m5d.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r5dn.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "c4.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 15360
      }
    },
    {
      "InstanceType": "c5.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r5d.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "r5d.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m5dn.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "r5dn.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "m5n.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "c5n.metal",
      "VCpuInfo": {
        "DefaultVCpus": 72
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "r5d.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "c5d.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "r4.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "c4.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 7680
      }
    },
    {
      "InstanceType": "m3.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 15360
      }
    },
    {
      "InstanceType": "r5dn.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "r4.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      }
    },
    {
      "InstanceType": "i3.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "c5n.18xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 72
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "c5d.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "t3.nano",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 512
      }
    },
    {
      "InstanceType": "g4dn.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "c5d.18xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 72,
        "ValidCores": [
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24,
          26,
          28,
          30,
          32,
          34,
          36
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 147456
      }
    },
    {
      "InstanceType": "i3.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 31232
      }
    },
    {
      "InstanceType": "m4.10xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 40,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 163840
      }
    },
    {
      "InstanceType": "g3.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      }
    },
    {
      "InstanceType": "r5n.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "r5ad.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m4.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "r5.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "r5d.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "r5.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "r5n.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "c5.9xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 36
      },
      "MemoryInfo": {
        "SizeInMiB": 73728
      }
    },
    {
      "InstanceType": "t2.micro",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ],
        "SustainedClockSpeedInGhz": 2.5
      },
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 1024
      }
    },
    {
      "InstanceType": "m3.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 30720
      }
    },
    {
      "InstanceType": "m5n.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    },
    {
      "InstanceType": "m5d.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "h1.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "r5.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "r5d.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24,
          26,
          28,
          30,
          32,
          34,
          36,
          38,
          40,
          42,
          44,
          46,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "r5d.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "i3.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 499712
      }
    },
    {
      "InstanceType": "r5ad.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "i3.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 15616
      }
    },
    {
      "InstanceType": "r3.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "c5.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "z1d.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "m5.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "m5.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "m1.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 15360
      }
    },
    {
      "InstanceType": "m5ad.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "t3.micro",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 1024
      }
    },
    {
      "InstanceType": "a1.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16,
        "ValidCores": [
          16
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m5ad.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          12,
          18,
          24,
          36,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "c3.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 15360
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "supported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "m5n.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "i2.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "unsupported",
        "EncryptionSupport": "supported"
      }
    },
    {
      "InstanceType": "m5dn.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "r5d.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "c5d.9xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 36
      },
      "MemoryInfo": {
        "SizeInMiB": 73728
      }
    },
    {
      "InstanceType": "g4dn.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "f1.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 999424
      }
    },
    {
      "InstanceType": "r5a.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "r5n.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "inf1.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "i3en.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "f1.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "g3s.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 31232
      }
    },
    {
      "InstanceType": "p2.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      }
    },
    {
      "InstanceType": "d2.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      }
    },
    {
      "InstanceType": "m1.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 7680
      }
    },
    {
      "InstanceType": "r5dn.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "r5a.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "r5ad.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "i3en.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "i3.metal",
      "VCpuInfo": {
        "DefaultVCpus": 72
      },
      "MemoryInfo": {
        "SizeInMiB": 524288
      }
    },
    {
      "InstanceType": "t3.small",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 2048
      }
    },
    {
      "InstanceType": "r5n.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "p3dn.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          4,
          6,
          8,
          10,
          12,
          14,
          16,
          18,
          20,
          22,
          24,
          26,
          28,
          30,
          32,
          34,
          36,
          38,
          40,
          42,
          44,
          46,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "r4.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      }
    },
    {
      "InstanceType": "m3.medium",
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 3840
      }
    },
    {
      "InstanceType": "m5a.24xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 96,
        "ValidCores": [
          12,
          18,
          24,
          36,
          48
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 393216
      }
    },
    {
      "InstanceType": "x1.32xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 128,
        "ValidCores": [
          4,
          8,
          12,
          16,
          20,
          24,
          28,
          32,
          36,
          40,
          44,
          48,
          52,
          56,
          60,
          64
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 1998848
      }
    },
    {
      "InstanceType": "m5d.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "r4.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 15616
      }
    },
    {
      "InstanceType": "c5.18xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 72
      },
      "MemoryInfo": {
        "SizeInMiB": 147456
      }
    },
    {
      "InstanceType": "g4dn.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "r5dn.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 65536
      }
    },
    {
      "InstanceType": "cc2.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 61952
      },
      "EbsInfo": {
        "EbsOptimizedSupport": "unsupported",
        "EncryptionSupport": "unsupported"
      }
    },
    {
      "InstanceType": "i3en.6xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 24,
        "ValidCores": [
          2,
          4,
          6,
          8,
          10,
          12
        ]
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "i3en.metal",
      "VCpuInfo": {
        "DefaultVCpus": 96
      },
      "MemoryInfo": {
        "SizeInMiB": 786432
      }
    },
    {
      "InstanceType": "m1.medium",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ]
      },
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 3788
      }
    },
    {
      "InstanceType": "g4dn.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m5a.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "t2.nano",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ],
        "SustainedClockSpeedInGhz": 2.4
      },
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 512
      }
    },
    {
      "InstanceType": "i3en.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "m1.small",
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 1740
      }
    },
    {
      "InstanceType": "m5a.12xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 48
      },
      "MemoryInfo": {
        "SizeInMiB": 196608
      }
    },
    {
      "InstanceType": "x1e.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 499712
      }
    },
    {
      "InstanceType": "d2.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 31232
      }
    },
    {
      "InstanceType": "p3.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 499712
      }
    },
    {
      "InstanceType": "r5dn.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "t3a.nano",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 512
      }
    },
    {
      "InstanceType": "t3a.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 16384
      }
    },
    {
      "InstanceType": "t2.small",
      "ProcessorInfo": {
        "SupportedArchitectures": [
          "i386",
          "x86_64"
        ],
        "SustainedClockSpeedInGhz": 2.5
      },
      "VCpuInfo": {
        "DefaultVCpus": 1
      },
      "MemoryInfo": {
        "SizeInMiB": 2048
      }
    },
    {
      "InstanceType": "t3a.micro",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 1024
      }
    },
    {
      "InstanceType": "i3.4xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 16
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      }
    },
    {
      "InstanceType": "m5n.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 32768
      }
    },
    {
      "InstanceType": "p2.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 749568
      }
    },
    {
      "InstanceType": "m5ad.16xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 64
      },
      "MemoryInfo": {
        "SizeInMiB": 262144
      }
    },
    {
      "InstanceType": "i3.2xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 8
      },
      "MemoryInfo": {
        "SizeInMiB": 62464
      }
    },
    {
      "InstanceType": "z1d.3xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 12
      },
      "MemoryInfo": {
        "SizeInMiB": 98304
      }
    },
    {
      "InstanceType": "m2.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 17510
      }
    },
    {
      "InstanceType": "c5n.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 10752
      }
    },
    {
      "InstanceType": "g2.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 61440
      }
    },
    {
      "InstanceType": "p3.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 249856
      }
    },
    {
      "InstanceType": "x1e.xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 4
      },
      "MemoryInfo": {
        "SizeInMiB": 124928
      }
    },
    {
      "InstanceType": "m3.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 7680
      }
    },
    {
      "InstanceType": "g4dn.8xlarge",
      "VCpuInfo": {
        "DefaultVCpus": 32
      },
      "MemoryInfo": {
        "SizeInMiB": 131072
      }
    },
    {
      "InstanceType": "m5a.large",
      "VCpuInfo": {
        "DefaultVCpus": 2
      },
      "MemoryInfo": {
        "SizeInMiB": 8192
      }
    }
  ]
}
`)
}
