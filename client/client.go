package client

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	HostURL    string
	Username   string
	Password   string
	Insecure   bool
	HTTPClient *http.Client
}

type errorResponse struct {
	Detail  string `json:"detail"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func NewClient(hosturl string, username string, password string, insecure bool) *Client {
	return &Client{
		HostURL:  hosturl,
		Username: username,
		Password: password,
		Insecure: insecure,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
	}
}

func GetCredentialsFromEnvVar() (host string, username string, password string, insecure bool) {
	host = os.Getenv("ROS_HOSTURL")
	username = os.Getenv("ROS_USERNAME")
	password = os.Getenv("ROS_PASSWORD")
	insecure_str := os.Getenv("ROS_INSECURE")
	insecure, _ = strconv.ParseBool(insecure_str)
	return host, username, password, insecure
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	fmt.Println(req)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			fmt.Printf(errRes.Detail)
			return errors.New(errRes.Detail)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	body, _ := ioutil.ReadAll(res.Body)
	if len(body) != 0 {
		if err = json.Unmarshal(body, &v); err != nil {
			return err
		}
	}
	return nil
}
