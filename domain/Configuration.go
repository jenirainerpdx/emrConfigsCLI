package domain


type Configuration struct {
	Classification string
	Properties     interface{}
}

type SparkDefaults struct {
	NetworkTimeout                string `json:"spark.network.timeout"`
	ExecutorHeartbeatInterval     string `json:"spark.executor.heartbeatInterval"`
	DynamicAllocation             string `json:"spark.dynamicAllocation.enabled"`
	DriverMemory                  string `json:"spark.driver.memory"`
	ExecutorMemory                string `json:"spark.executor.memory"`
	ExecutorCores                 string `json:"spark.executor.cores"`
	Instances                     string `json:"spark.executor.instances"`
	MemoryFraction                string `json:"spark.memory.fraction"`
	StorageFraction               string `json:"spark.memory.storageFraction"`
	ExecutorExtraJavaOptions      string `json:"spark.executor.extraJavaOptions"`
	DriverExtraJavaOptions        string `json:"spark.driver.extraJavaOptions"`
	YarnReporterThreadMaxFailures string `json:"spark.yarn.scheduler.reporterThread.maxFailures"`
	StorageLevel                  string `json:"spark.storage.level"`
	RddCompress                   string `json:"spark.rdd.compress"`
	ShuffleCompress               string `json:"spark.shuffle.compress"`
	SpillCompress                 string `json:"spark.shuffle.spill.compress"`
	DefaultParallelism            string `json:"spark.default.parallelism"`
	KyroBufferMax                 string `json:"spark.kyroserializer.buffer.max"`
}
