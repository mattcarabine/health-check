package health

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type RestClient struct {
	httpClient *http.Client
	username string
	password string
	cluster string
}

var client *RestClient

func InitClient(username string, password string, cluster string) {
	client = &RestClient{httpClient: &http.Client{}, username:  username, password: password, cluster: cluster}
}

func (c *RestClient) CMGet(path string) []byte {
	url := fmt.Sprintf("http://%s%s", c.cluster, path)
	req, err := http.NewRequest("GET", url, nil)
	if err !=  nil {
		panic(err)
	}
	req.SetBasicAuth(c.username, c.password)
	response, err := c.httpClient.Do(req)
	var data []byte
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ = ioutil.ReadAll(response.Body)
	}
	return data
}

func (c *RestClient) Get(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err !=  nil {
		panic(err)
	}
	response, err := c.httpClient.Do(req)
	var data []byte
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ = ioutil.ReadAll(response.Body)
	}
	return data
}