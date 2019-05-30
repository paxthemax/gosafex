package safexdrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/safex/gosafex/pkg/safex"
	"github.com/tidwall/gjson"
)

// Type declarations for building JSON-like object
type JSONElement = map[string]interface{}
type JSONArray = []interface{}

type Client struct {
	Port       uint
	Host       string
	ID         uint
	httpClient http.Client
}

// must panics in the case of error.
func must(err error) {
	if err == nil {
		return
	}

	log.Panicln(err)
}

// InitClient creates and initializes RPC client and returns client object
// takes host and port as arguments
func InitClient(host string, port uint) (client *Client) {

	client = &Client{
		Port: port,
		Host: host,
		ID:   0,
	}

	// Create config
	trConfig := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client.httpClient = http.Client{Transport: trConfig}
	return client
}

type JSONResult struct {
	count  int    `json:"count"`
	status string `json:"status"`
}

type JSONResponse struct {
	Id      string     `json:"id"`
	JSONRpc string     `json:"jsonrpc"`
	Result  JSONResult `json:"result"`
}

//Close destroys RPC client
func (c *Client) Close() {

}

func (c Client) JSONSafexdCall(method string, params interface{}) ([]byte, error) {
	body := map[string]interface{}{"jsonrpc": "2.0", "id": 1, "method": method, "params": params}
	url := "http://" + c.Host + ":" + strconv.Itoa(int(c.Port)) + "/json_rpc"

	jsonBuff, _ := json.Marshal(body)

	fmt.Println(string(jsonBuff))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBuff))
	must(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	must(err)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	must(err)

	errorJson := gjson.Get(string(resBody), "error.message")
	if errorJson.Str != "" {
		err = errors.New(errorJson.Str)
		return nil, err
	}
	return resBody, err
}

func (c Client) SafexdCall(method string, params interface{}, httpMethod string) ([]byte, error) {
	var body []byte
	var err error
	if params == nil {
		body = []byte("")
	} else {
		body, err = json.Marshal(params)
	}

	must(err)
	url := "http://" + c.Host + ":" + strconv.Itoa(int(c.Port)) + "/" + method

	fmt.Println(string(body))

	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	must(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	must(err)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	must(err)

	errorJson := gjson.Get(string(resBody), "error.message")
	if errorJson.Str != "" {
		err = errors.New(errorJson.Str)
		return nil, err
	}
	return resBody, err

}

//GetBlockCount gets current node latest block number
func (c Client) GetBlockCount() (count uint64, err error) {

	result, err := c.JSONSafexdCall("get_block_count", JSONElement{})
	must(err)
	count = uint64(gjson.GetBytes(result, "result.count").Num)
	return count, err
}

//OnGetBlockHash returns hash of block with provide height
func (c Client) OnGetBlockHash(height uint64) (hash string, err error) {

	result, err := c.JSONSafexdCall("on_get_block_hash", JSONArray{height})
	must(err)
	var jsonObj interface{}
	json.Unmarshal(result, &jsonObj)
	return jsonObj.(JSONElement)["result"].(string), err
}

func getSliceForPath(input []byte, path string) []byte {
	temp := gjson.GetBytes(input, path)
	return input[temp.Index : temp.Index+len(temp.Raw)]
}

func (c Client) GetDaemonInfo() (info safex.DaemonInfo, err error) {
	result, err := c.SafexdCall("get_info", nil, "POST")
	must(err)
	err = json.Unmarshal(result, &info)
	must(err)
	return info, err
}

func (c Client) GetHardForkInfo() (info safex.HardForkInfo, err error) {
	result, err := c.JSONSafexdCall("hard_fork_info", JSONElement{})
	must(err)

	err = json.Unmarshal(getSliceForPath(result, "result"), &info)
	must(err)
	return info, err
}

func (c Client) GetTransactions(hashes []string) (txs safex.Transactions, err error) {
	result, err := c.SafexdCall("proto/get_transactions", JSONElement{"txs_hashes": hashes}, "POST")
	err = proto.Unmarshal(result, &txs)
	must(err)
	return txs, err
}

func (c Client) GetBlocks(start uint64, end uint64) (blcks safex.Blocks, err error) {
	result, err := c.SafexdCall("proto/get_blocks", JSONElement{"start_height": start, "end_height": end}, "POST")
	err = proto.Unmarshal(result, &blcks)
	must(err)
	return blcks, err
}
