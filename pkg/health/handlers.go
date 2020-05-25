package health

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func SummaryHandler(w http.ResponseWriter, r *http.Request){
	pd := PoolsDefault{}
	_ = json.Unmarshal(client.CMGet("/pools/default"), &pd)
	tooHigh := pd.Quota()
	quotaResult := "false"
	if tooHigh {
		quotaResult = "true"
	}

	thp := THP()
	thpResult := "false"
	if thp {
		thpResult = "true"
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"quotaTooHigh": %s, "thp": %s}`, quotaResult, thpResult)
}

func PingHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Query().Get("internal") == "true" {
		pr, _ := json.Marshal(pingAllNodes())
		fmt.Fprint(w, string(pr))
	} else {
		pr := pingAllNodes()
		prs := []PingReport{pr}
		for _, node := range pr.PingReports {
			host, _, _ := net.SplitHostPort(node.Node)
			healthUrl := fmt.Sprintf("http://%s:%d/ping?internal=true", host, node.HealthPort)
			pr := PingReport{}
			err := json.Unmarshal(client.Get(healthUrl), &pr)
			if err != nil {
				fmt.Println(err)
			}
			prs = append(prs, pr)
		}
		result, _ := json.Marshal(prs)
		fmt.Println(result)
		fmt.Fprintf(w, `{"pingReport": %s}`, string(result))
	}
}

func ConnectionsHandler(w http.ResponseWriter, r *http.Request){
	conns, _ := json.Marshal(GrabConnections())
	fmt.Fprint(w, string(conns))
}