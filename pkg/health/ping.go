package health

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"time"
)

type NodeServices struct {
	Rev      int `json:"rev"`
	NodesExt []struct {
		Services struct {
			Mgmt               int `json:"mgmt"`
			MgmtSSL            int `json:"mgmtSSL"`
			Cbas               int `json:"cbas"`
			CbasSSL            int `json:"cbasSSL"`
			EventingAdminPort  int `json:"eventingAdminPort"`
			EventingSSL        int `json:"eventingSSL"`
			Fts                int `json:"fts"`
			FtsSSL             int `json:"ftsSSL"`
			FtsGRPC            int `json:"ftsGRPC"`
			FtsGRPCSSL         int `json:"ftsGRPCSSL"`
			IndexAdmin         int `json:"indexAdmin"`
			IndexScan          int `json:"indexScan"`
			IndexHTTP          int `json:"indexHttp"`
			IndexHTTPS         int `json:"indexHttps"`
			Kv                 int `json:"kv"`
			KvSSL              int `json:"kvSSL"`
			Capi               int `json:"capi"`
			CapiSSL            int `json:"capiSSL"`
			Projector          int `json:"projector"`
			N1Ql               int `json:"n1ql"`
			N1QlSSL            int `json:"n1qlSSL"`
			Health			   int `json:"health"`
		} `json:"services"`
		ThisNode bool `json:"thisNode"`
		Hostname string `json:"hostname"`
	} `json:"nodesExt"`
	ClusterCapabilitiesVer []int `json:"clusterCapabilitiesVer"`
	ClusterCapabilities    struct {
		N1Ql []string `json:"n1ql"`
	} `json:"clusterCapabilities"`
}

type PortDefinition struct {
	Name string `json:"name"`
	Port int `json:"port"`
}

type TargetPingReport struct {
	Node string `json:"node"`
	ReachablePorts []PortDefinition `json:"reachablePorts"`
	UnreachablePorts []PortDefinition `json:"unreachablePorts"`
	HealthPort int `json:"-"`
}

type PingReport struct {
	PingReports []TargetPingReport `json:"pingReports"`
	Node string `json:"node"`
}

type PingResult struct {
	Success bool
	Port PortDefinition
}

func pingAllNodes() PingReport{
	ns := NodeServices{}
	err := json.Unmarshal(client.CMGet("/pools/default/nodeServices"), &ns)
	if err != nil {
		fmt.Println(err)
	}
	var pingReports []TargetPingReport
	var srcHost string
	for _, node := range ns.NodesExt {
		if node.ThisNode {
			srcHost =  fmt.Sprintf("%s:%d", node.Hostname, node.Services.Mgmt)
			continue
		}
		var reachable []PortDefinition
		var unreachable []PortDefinition

		healthPort := node.Services.Health
		services := reflect.ValueOf(node.Services)
		c := make(chan PingResult)
		totalPorts := 0
		for i := 0; i < services.NumField(); i++ {
			port := PortDefinition{Name: services.Type().Field(i).Name, Port: int(services.Field(i).Int())}

			// If port is 0, it wasn't found in the rest call
			if port.Port == 0 {
				continue
			}
			totalPorts += 1

			go ping(node.Hostname, port, c)
		}

		for len(reachable) + len(unreachable) < totalPorts {
			result := <- c
			if result.Success {
				reachable = append(reachable, result.Port)
			} else {
				unreachable = append(unreachable, result.Port)
			}
		}

		hostname := fmt.Sprintf("%s:%d", node.Hostname, node.Services.Mgmt)
		tpr := TargetPingReport{Node: hostname, ReachablePorts: reachable, UnreachablePorts: unreachable, HealthPort: healthPort}
		pingReports = append(pingReports, tpr)
	}
	return PingReport{PingReports: pingReports,  Node: srcHost}
}

func ping(host string, port PortDefinition, c chan PingResult) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port.Port), time.Second * 2)
	c <- PingResult{Port: port, Success: err == nil}
}