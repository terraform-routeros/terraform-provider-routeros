package routeros

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-routeros/routeros"
	"net/http"
	"os"
	"testing"
	"time"
)

func newApiClient(ctx context.Context, hostUrl, user, pass string, useTLS bool) (*ApiClient, error) {
	api := &ApiClient{
		ctx:       ctx,
		HostURL:   hostUrl,
		Username:  user,
		Password:  pass,
		Transport: TransportAPI,
	}

	tlsConf := tls.Config{
		InsecureSkipVerify: true,
	}

	var err error

	if useTLS {
		api.Client, err = routeros.DialTLS(api.HostURL, api.Username, api.Password, &tlsConf)
	} else {
		api.Client, err = routeros.Dial(api.HostURL, api.Username, api.Password)
	}
	if err != nil {
		return nil, err
	}
	api.Async()

	return api, nil
}

func newRestClient(ctx context.Context, hostUrl, user, pass string) *RestClient {
	return &RestClient{
		ctx:       ctx,
		HostURL:   hostUrl,
		Username:  user,
		Password:  pass,
		Transport: TransportREST,
		Client: &http.Client{
			Timeout: time.Minute,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

func TestClientTransport_SendRequest(t *testing.T) {
	testAccPreCheck(t)
	ctx := context.Background()
	host := reHost.FindStringSubmatch(os.Getenv("ROS_HOSTURL"))[1]
	user := os.Getenv("ROS_USERNAME")
	pass := os.Getenv("ROS_PASSWORD")
	api, err := newApiClient(ctx, host+":8728", user, pass, false)
	if err != nil {
		t.Fatal(err)
	}
	apis, err := newApiClient(ctx, host+":8729", user, pass, true)
	if err != nil {
		t.Fatal(err)
	}
	rest := newRestClient(ctx, "https://"+host+":443", user, pass)

	type fields struct {
		Transport TransportType
		Client    interface{}
	}
	type args struct {
		method crudMethod
		url    *URL
		item   MikrotikItem
		result interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Test API connection",
			fields:  fields{TransportAPI, api},
			args:    args{crudRead, &URL{Path: "/system/resource"}, nil, &[]MikrotikItem{}},
			wantErr: false,
		},
		{
			name:    "Test APIs connection",
			fields:  fields{TransportAPI, apis},
			args:    args{crudRead, &URL{Path: "/system/resource"}, nil, &[]MikrotikItem{}},
			wantErr: false,
		},
		{
			name:    "Test REST connection",
			fields:  fields{TransportREST, rest},
			args:    args{crudRead, &URL{Path: "/system/resource"}, nil, &MikrotikItem{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields.Client.(Client)
			if err := c.SendRequest(tt.args.method, tt.args.url, tt.args.item, tt.args.result); (err != nil) != tt.wantErr {
				t.Fatalf("SendRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
			var info MikrotikItem
			if tt.fields.Transport == TransportAPI {
				if len(*tt.args.result.(*[]MikrotikItem)) == 0 {
					t.Fatalf("Response is empty.")
				}
				info = (*tt.args.result.(*[]MikrotikItem))[0]
			} else {
				info = *tt.args.result.(*MikrotikItem)
			}
			fmt.Printf("\t\t::: %v %v, RouterOS: %v at %v :::\n", info["platform"], info["board-name"], info["version"], info["build-time"])
		})
	}
}
