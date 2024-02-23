package main

import (
	"context"
	"net"
	"time"
	"veil-v2/tasks"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func main() {

	t := &tasks.Task{}
	jar := tls_client.NewCookieJar()

	dialer := net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(context context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(5000) * time.Millisecond,
				}
				return d.DialContext(context, "udp", net.JoinHostPort("8.8.8.8", "53"))
			},
		},
	}
	client_options := []tls_client.HttpClientOption{
		tls_client.WithClientProfile(profiles.Chrome_117),
		tls_client.WithCookieJar(jar),
		tls_client.WithDialer(dialer),
	}
	t.Client, _ = tls_client.NewHttpClient(tls_client.NewLogger(), client_options...)

	t.CRNs = []string{"48467", "47836", "48591"}
	t.WebhookURL = "https://discord.com/api/webhooks/1022240016786800761/lGBemtv7h9G0QrxZeOJ1pwWeOeetVMY42vA9vd75ipFDeyz9c3emrwWOLVKM04txPoVZ"
	t.GetTermByName("2024 Spring De Anza")
	t.Signup()
}
