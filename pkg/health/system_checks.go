package health

import (
	"fmt"
	"log"
	"strings"
)

type PoolsDefault struct {
	Name  string `json:"name"`
	Nodes []struct {
		SystemStats struct {
			CPUUtilizationRate int `json:"cpu_utilization_rate"`
			CPUStolenRate      int `json:"cpu_stolen_rate"`
			SwapTotal          int `json:"swap_total"`
			SwapUsed           int `json:"swap_used"`
			MemTotal           int `json:"mem_total"`
			MemFree            int `json:"mem_free"`
			MemLimit           int `json:"mem_limit"`
			CPUCoresAvailable  int `json:"cpu_cores_available"`
			Allocstall         int `json:"allocstall"`
		} `json:"systemStats"`
		InterestingStats struct {
		} `json:"interestingStats"`
		Uptime               string `json:"uptime"`
		MemoryTotal          int    `json:"memoryTotal"`
		MemoryFree           int    `json:"memoryFree"`
		McdMemoryReserved    int    `json:"mcdMemoryReserved"`
		McdMemoryAllocated   int    `json:"mcdMemoryAllocated"`
		CouchAPIBase         string `json:"couchApiBase"`
		ClusterMembership    string `json:"clusterMembership"`
		RecoveryType         string `json:"recoveryType"`
		Status               string `json:"status"`
		OtpNode              string `json:"otpNode"`
		ThisNode             bool   `json:"thisNode"`
		Hostname             string `json:"hostname"`
		NodeUUID             string `json:"nodeUUID"`
		ClusterCompatibility int    `json:"clusterCompatibility"`
		Version              string `json:"version"`
		Os                   string `json:"os"`
		CPUCount             int    `json:"cpuCount"`
		Ports                struct {
			Direct  int `json:"direct"`
			DistTCP int `json:"distTCP"`
			DistTLS int `json:"distTLS"`
		} `json:"ports"`
		Services           []string `json:"services"`
		NodeEncryption     bool     `json:"nodeEncryption"`
		ConfiguredHostname string   `json:"configuredHostname"`
		AddressFamily      string   `json:"addressFamily"`
		ExternalListeners  []struct {
			Afamily        string `json:"afamily"`
			NodeEncryption bool   `json:"nodeEncryption"`
		} `json:"externalListeners"`
	} `json:"nodes"`
	Buckets struct {
		URI                       string `json:"uri"`
		TerseBucketsBase          string `json:"terseBucketsBase"`
		TerseStreamingBucketsBase string `json:"terseStreamingBucketsBase"`
	} `json:"buckets"`
	RemoteClusters struct {
		URI         string `json:"uri"`
		ValidateURI string `json:"validateURI"`
	} `json:"remoteClusters"`
	Alerts           []interface{} `json:"alerts"`
	AlertsSilenceURL string        `json:"alertsSilenceURL"`
	Controllers      struct {
		AddNode struct {
			URI string `json:"uri"`
		} `json:"addNode"`
		Rebalance struct {
			URI string `json:"uri"`
		} `json:"rebalance"`
		FailOver struct {
			URI string `json:"uri"`
		} `json:"failOver"`
		StartGracefulFailover struct {
			URI string `json:"uri"`
		} `json:"startGracefulFailover"`
		ReAddNode struct {
			URI string `json:"uri"`
		} `json:"reAddNode"`
		ReFailOver struct {
			URI string `json:"uri"`
		} `json:"reFailOver"`
		EjectNode struct {
			URI string `json:"uri"`
		} `json:"ejectNode"`
		SetRecoveryType struct {
			URI string `json:"uri"`
		} `json:"setRecoveryType"`
		SetAutoCompaction struct {
			URI         string `json:"uri"`
			ValidateURI string `json:"validateURI"`
		} `json:"setAutoCompaction"`
		ClusterLogsCollection struct {
			StartURI  string `json:"startURI"`
			CancelURI string `json:"cancelURI"`
		} `json:"clusterLogsCollection"`
		Replication struct {
			CreateURI   string `json:"createURI"`
			ValidateURI string `json:"validateURI"`
		} `json:"replication"`
	} `json:"controllers"`
	RebalanceStatus        string `json:"rebalanceStatus"`
	RebalanceProgressURI   string `json:"rebalanceProgressUri"`
	StopRebalanceURI       string `json:"stopRebalanceUri"`
	NodeStatusesURI        string `json:"nodeStatusesUri"`
	MaxBucketCount         int    `json:"maxBucketCount"`
	AutoCompactionSettings struct {
		ParallelDBAndViewCompaction    bool `json:"parallelDBAndViewCompaction"`
		DatabaseFragmentationThreshold struct {
			Percentage int    `json:"percentage"`
			Size       string `json:"size"`
		} `json:"databaseFragmentationThreshold"`
		ViewFragmentationThreshold struct {
			Percentage int    `json:"percentage"`
			Size       string `json:"size"`
		} `json:"viewFragmentationThreshold"`
		IndexCompactionMode     string `json:"indexCompactionMode"`
		IndexCircularCompaction struct {
			DaysOfWeek string `json:"daysOfWeek"`
			Interval   struct {
				FromHour     int  `json:"fromHour"`
				ToHour       int  `json:"toHour"`
				FromMinute   int  `json:"fromMinute"`
				ToMinute     int  `json:"toMinute"`
				AbortOutside bool `json:"abortOutside"`
			} `json:"interval"`
		} `json:"indexCircularCompaction"`
		IndexFragmentationThreshold struct {
			Percentage int `json:"percentage"`
		} `json:"indexFragmentationThreshold"`
	} `json:"autoCompactionSettings"`
	Tasks struct {
		URI string `json:"uri"`
	} `json:"tasks"`
	Counters struct {
	} `json:"counters"`
	IndexStatusURI      string `json:"indexStatusURI"`
	CheckPermissionsURI string `json:"checkPermissionsURI"`
	ServerGroupsURI     string `json:"serverGroupsUri"`
	ClusterName         string `json:"clusterName"`
	Balanced            bool   `json:"balanced"`
	MemoryQuota         int    `json:"memoryQuota"`
	IndexMemoryQuota    int    `json:"indexMemoryQuota"`
	FtsMemoryQuota      int    `json:"ftsMemoryQuota"`
	CbasMemoryQuota     int    `json:"cbasMemoryQuota"`
	EventingMemoryQuota int    `json:"eventingMemoryQuota"`
	AuditUID            string `json:"auditUid"`
	StorageTotals       struct {
		RAM struct {
			Total             int64 `json:"total"`
			QuotaTotal        int64 `json:"quotaTotal"`
			QuotaUsed         int   `json:"quotaUsed"`
			Used              int64 `json:"used"`
			UsedByData        int   `json:"usedByData"`
			QuotaUsedPerNode  int   `json:"quotaUsedPerNode"`
			QuotaTotalPerNode int64 `json:"quotaTotalPerNode"`
		} `json:"ram"`
		Hdd struct {
			Total      int64 `json:"total"`
			QuotaTotal int64 `json:"quotaTotal"`
			Used       int64 `json:"used"`
			UsedByData int   `json:"usedByData"`
			Free       int64 `json:"free"`
		} `json:"hdd"`
	} `json:"storageTotals"`
}

func (pd *PoolsDefault) Quota() bool{
	return float64(pd.StorageTotals.RAM.QuotaTotal) / float64(pd.StorageTotals.RAM.Total) > 0.5
}

func (pd *PoolsDefault) DiskSpace() bool{
	return float64(pd.StorageTotals.Hdd.Used) / float64(pd.StorageTotals.Hdd.Total) > 0.7
}

func THP() bool {

	thpLocations := [4]string{
		"/sys/kernel/mm/transparent_hugepage/enabled",
		"/sys/kernel/mm/transparent_hugepage/defrag",
		"/sys/kernel/mm/redhat_transparent_hugepage/enabled",
		"/sys/kernel/mm/redhat_transparent_hugepage/defrag"}

	thpResult := ""
	for _, loc := range thpLocations {
		var err error
		thpResult, err = readFileSingleLine(loc)
		if err != nil {
			log.Println(fmt.Sprintf("Error reading THP setting at %s: \n %v", loc, err))
		} else if thpResult != "" {
			break
		}
	}

	if thpResult != "" {
		// always [madvise] never
		if strings.Contains(thpResult, "[always]"){
			return true
		}
	}
	return false
}