package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BTCReq struct {
	Method string
}

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	method := jsonBody["method"]
	j, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.json", method))
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", string(j))
})

func TestFetchMetrics(t *testing.T) {
	ts := httptest.NewServer(testHandler)
	defer ts.Close()

	var bitcoin BitcoinPlugin
	bitcoin.Dest = strings.TrimPrefix(ts.URL, "http://")
	t.Log(ts.URL)
	bitcoin.User = "test"
	bitcoin.Password = "test"
	stat, err := bitcoin.FetchMetrics()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(stat)
	assert.EqualValues(t, 10, stat["height"])
	assert.EqualValues(t, 27, stat["in"])
	assert.EqualValues(t, 10, stat["out"])
	assert.EqualValues(t, 37, stat["total"])
	assert.EqualValues(t, 2000, stat["network.score.203_0_113_0"])
}
