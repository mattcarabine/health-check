package health

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/couchbase/cbauth"
	"log"
	"os/exec"
	"regexp"
)

type ConnectionData struct {
	AgentName  string `json:"agent_name"`
	Connection string `json:"connection"`
	Socket     int    `json:"socket"`
	Protocol   string `json:"protocol"`
	ParentPort int    `json:"parent_port"`
	Username   string `json:"username"`
	PeerName   string `json:"peername"`
	SockName   string `json:"sockname"`
	Internal   bool   `json:"internal"`
}

func GrabConnections() []ConnectionData {
	u, p, _ := cbauth.GetMemcachedServiceAuth("localhost:11210")
	stats, err := exec.Command("mcstat", "localhost:11210", "-u", u, "-P", p, "connections").Output()
	if err != nil {
		log.Println(fmt.Sprintf("Couldn't access mcstats output, reason: \n %v", err))
	}
	var connections []ConnectionData

	// Would rather fail silently than crash on this non-essential feature
	if err != nil {
		fmt.Println(err)
		return connections
	}
	re := regexp.MustCompile(`^[0-9]+\s+`)

	scanner := bufio.NewScanner(bytes.NewReader(stats))
	for scanner.Scan() {
		var conn *ConnectionData
		rawConn := re.ReplaceAll(scanner.Bytes(), []byte(""))
		err := json.Unmarshal(rawConn, &conn)
		if err != nil {
			log.Println(fmt.Sprintf("Unable to unmarshal connection, reason: %v", err))
		} else if conn.AgentName == "" || conn.Internal {
			// For now we only care about client connections. Skip if no agent name
		} else {
			connections = append(connections, *conn)
		}
	}
	return connections
}