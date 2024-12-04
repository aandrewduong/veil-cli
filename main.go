package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
	"veil-v2/tasks"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

func main() {

	var dnsServers = []string{"8.8.8.8", "8.8.4.4", "1.1.1.1", "1.0.0.1"}
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
		tls_client.WithClientProfile(profiles.Chrome_131), // Updated for Chrome 131 Profile, previous was 121
		tls_client.WithCookieJar(jar),
		tls_client.WithDialer(dialer),
	}
	t.Client, _ = tls_client.NewHttpClient(tls_client.NewLogger(), client_options...)
	t.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

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
		if len(row) < 8 {
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
		var registrationTime = row[7]

		if t.Mode == "Release" {
			t.Mode = "Signup"
			pattern := regexp.MustCompile(`\d{2}/\d{2}/\d{4} \d{2}:\d{2} [APM]{2}`)
			matches := pattern.FindAllString(registrationTime, -1)
			if len(matches) == 0 {
				fmt.Printf("Invalid Registration Time Format")
			}

			location, _ := time.LoadLocation("America/Los_Angeles")
			targetTime, _ := time.ParseInLocation("01/02/2006 03:04 PM", matches[0], location)
			now := time.Now().In(location)
			timeToWait := targetTime.Sub(now) - 5*time.Minute

			if now.Before(targetTime) {
				fmt.Printf("Will continue in: %s\n", timeToWait.String())
				time.Sleep(timeToWait)
			}
		}
	}

	t.Run()
}
