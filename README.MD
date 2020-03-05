# emrConfigs Utility

## Building the tool
- install Go
- go build emrConfigs.go

## Running the tool
- build the tool
- from command line, call ./emrConfigs
    - running with no flags will result in a printout of the flags and how to use them.
    - tool gets run with --nodes and --instance flags.
    - The --nodes flag is an integer type and it represents the count of worker nodes that you want to have in your cluster.
    - The --instance flag describes the type of ec2 instances to be used for the worker nodes. 
        - valid types are limited, in current version:  r5.2xlarge - r5.12xlarge and a few of the c type instances.
    - Output will be an ugly single line of json - but you can copy and paste this into the configuration window of the EMR console.
    
### Sample of running the tool: 

`./emrConfigs -nodes=20 -instance=r5.2xlarge`

        