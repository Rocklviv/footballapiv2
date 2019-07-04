package footballapiv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

const (
	// APIURL - URL of football-data version 2
	APIURL = "http://api.football-data.org/v2/"
)

// Client struct that handles
// `httpClient http.Client`
// `apikey     string` Can be set via `NewClient`
// `rpm        int` Can be set via NewClient(apiket).SetMaxRPM(30)
type Client struct {
	httpClient http.Client
	apikey     string
	rpm        int
}

// NewClient creates `footballapiv2` instance.
func NewClient(apikey string) *Client {
	var h http.Client
	return &Client{
		httpClient: h,
		apikey:     apikey,
		rpm:        10,
	}
}

// SetMaxRPM - sets maximum requests per minute.
// By defualt is 10 requests per minute that used by FREE plan.
func (c *Client) SetMaxRPM(rpm int) *Client {
	c.rpm = rpm
	return c
}

// Sends request to API and returns json.Decoder.
func (c *Client) doRequest(method, url string, values url.Values) (js *json.Decoder, err error) {
	req := &http.Request{
		Method: method,
		URL:    prepareURL(url, values),
		Header: http.Header{},
	}
	if len(c.apikey) < 0 {
		fmt.Println("Cannot find API KEY.")
		os.Exit(1)
	}

	req.Header.Set("X-Auth-Token", c.apikey)
	req.Header.Set("User-Agent", "go-football-data-api/2.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Status Code: %d, - Status: %s", resp.StatusCode, resp.Status)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)

	js = json.NewDecoder(buf)
	return
}

// Prepares URL with parameters.
func prepareURL(path string, values url.Values) *url.URL {
	if values == nil {
		values = url.Values{}
	}

	r, err := url.Parse(APIURL)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range values {
		for _, vv := range v {
			if reflect.TypeOf(vv).Kind() == reflect.String && vv == "" {
				values.Del(k)
			} else if reflect.TypeOf(vv).Kind() == reflect.Int && len(vv) >= 0 {
				values.Del(k)
			}
		}
	}

	updatedURL := &url.URL{Path: path, RawQuery: values.Encode()}
	return r.ResolveReference(updatedURL)
}
