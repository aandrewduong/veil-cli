package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
	"veil-v2/tasks"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func main() {

	// https://www.lifewire.com/free-and-public-dns-servers-2626062
	var dnsServers = []string{"8.8.8.8", "8.8.4.4", "76.76.2.0", "76.76.10.0", "1.1.1.1", "1.0.0.1"}
	randomIndex := rand.Intn(len(dnsServers))

	dnsServer := dnsServers[randomIndex]

	t := &tasks.Task{}
	jar := tls_client.NewCookieJar()
	dialer := net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(context context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(5) * time.Second,
				}
				return d.DialContext(context, "udp", net.JoinHostPort(dnsServer, "53"))
			},
		},
	}
	client_options := []tls_client.HttpClientOption{
		tls_client.WithClientProfile(profiles.Chrome_117),
		tls_client.WithCookieJar(jar),
		tls_client.WithDialer(dialer),
		//tls_client.WithDebug(),
	}
	t.Client, _ = tls_client.NewHttpClient(tls_client.NewLogger(), client_options...)

	file, err := os.Open("settings.csv")
	if err != nil {
		fmt.Println("Error Opening File:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error Reading Header:", err)
		return
	}
	for {
		row, err := reader.Read()
		if err != nil {
			if err != csv.ErrFieldCount {
				fmt.Println("Finished Reading Configuration File")
			} else {
				fmt.Println("Error Reading Row: ", err)
			}
			break
		}
		if len(row) < 7 {
			fmt.Println("Invalid Configuration File")
			continue
		}
		t.Username = row[0]
		t.Password = row[1]
		t.GetTermByName(row[2])
		t.Subject = row[3]
		t.Mode = row[4]
		t.CRNs = strings.Split(row[5], ",")
		t.WebhookURL = row[6]
	}

	t.Run()
}
