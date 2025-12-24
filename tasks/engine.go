package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"regexp"
	"sync"
	"time"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

type Engine struct {
	Config *Config
}

func NewEngine(config *Config) *Engine {
	return &Engine{Config: config}
}

func (e *Engine) Run() {
	var wg sync.WaitGroup

	for i := range e.Config.Tasks {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			e.runTask(&e.Config.Tasks[idx])
		}(i)
	}

	wg.Wait()
}

func (e *Engine) runTask(tc *TaskConfig) {
	dnsServers := []string{"8.8.8.8", "8.8.4.4", "1.1.1.1", "1.0.0.1"}
	dnsServer := dnsServers[rand.Intn(len(dnsServers))]

	jar := tls_client.NewCookieJar()
	dialer := net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: 5 * time.Second,
				}
				return d.DialContext(ctx, "udp", net.JoinHostPort(dnsServer, "53"))
			},
		},
	}

	clientOptions := []tls_client.HttpClientOption{
		tls_client.WithClientProfile(profiles.Chrome_117),
		tls_client.WithCookieJar(jar),
		tls_client.WithDialer(dialer),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewLogger(), clientOptions...)
	if err != nil {
		fmt.Printf("[%s] Error creating client: %v\n", tc.Username, err)
		return
	}

	t := &Task{
		Username:   tc.Username,
		Password:   tc.Password,
		Subject:    tc.Subject,
		Mode:       tc.Mode,
		CRNs:       tc.CRNs,
		WebhookURL: tc.WebhookURL,
		Client:     client,
	}

	t.GetTermByName(tc.Term)

	if t.Mode == "Release" {
		t.Mode = "Signup"
		pattern := regexp.MustCompile(`\d{2}/\d{2}/\d{4} \d{2}:\d{2} [APM]{2}`)
		matches := pattern.FindAllString(tc.RegistrationTime, -1)
		if len(matches) == 0 {
			fmt.Printf("[%s] Invalid Registration Time Format\n", tc.Username)
			return
		}

		location, _ := time.LoadLocation("America/Los_Angeles")
		targetTime, _ := time.ParseInLocation("01/02/2006 03:04 PM", matches[0], location)
		now := time.Now().In(location)

		if now.Before(targetTime) {
			timeToWait := targetTime.Sub(now) - 5*time.Minute
			if timeToWait > 0 {
				fmt.Printf("[%s] Will continue in: %s\n", tc.Username, timeToWait.String())
				time.Sleep(timeToWait)
			}
		}
	}

	if err := t.Run(); err != nil {
		fmt.Printf("[%s] Task execution failed: %v\n", tc.Username, err)
	}
}
