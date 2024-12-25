package trc

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetAccount(t *testing.T) {
	url := "https://api.shasta.trongrid.io/wallet/getaccount"

	payload := strings.NewReader("{\"address\":\"TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func TestGetAccountBalance(t *testing.T) {

	url := "https://api.shasta.trongrid.io/wallet/getaccountbalance"

	payload := strings.NewReader("{\"account_identifier\":{\"address\":\"TKKib32o2zPBoXWKbTUQoBtNFM5LYFKb4a\"}, }}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
