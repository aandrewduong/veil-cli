package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

type Task struct {
	Mode          string
	Terms         map[string]string
	Username      string
	Password      string
	Subject       string
	CRNs          []string
	TermID        string
	UserAgent     string
	Client        tls_client.HttpClient
	Session       Session
	WebhookURL    string
	HomepageURL   string
	SSOManagerURL string
	WaitlistTask  bool
}

func (t *Task) GetHeaders(contentType string) [][2]string {
	headers := [][2]string{
		{"accept-language", "en-US,en;q=0.9"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"},
	}

	switch contentType {
	case "html":
		headers = append(headers, [2]string{"accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"})
	case "json":
		headers = append(headers, [2]string{"accept", "application/json"})
		headers = append(headers, [2]string{"content-type", "application/json"})
	case "form":
		headers = append(headers, [2]string{"accept", "*/*"})
		headers = append(headers, [2]string{"content-type", "application/x-www-form-urlencoded"})
	default:
		headers = append(headers, [2]string{"accept", "*/*"})
	}

	return headers
}

func (t *Task) MakeReq(method string, url string, headers [][2]string, body []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("[%s] Error creating request: %v\n", t.Username, err)
	}
	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}
	return req
}

func (t *Task) DoReq(req *http.Request, stage string, useDefaultResponseHandling bool) (*http.Response, error) {
	fmt.Printf("[%s] %s\n", t.Username, stage)

	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		resp, err := t.Client.Do(req)
		if err != nil {
			return nil, err
		}

		if useDefaultResponseHandling {
			if resp.StatusCode >= 400 {
				fmt.Printf("[%s] Error %s [%d]. Retrying (%d/%d)...\n", t.Username, stage, resp.StatusCode, i+1, maxRetries)
				discardResp(resp)
				time.Sleep(time.Second * 2)
				continue
			}
		}
		return resp, nil
	}

	return nil, fmt.Errorf("failed after %d retries", maxRetries)
}

func (t *Task) SendNotification(action string, message string) {
	payload := WebhookPayload{
		Username: "veil-v2",
		Embeds: []Embed{
			{
				Title:       action,
				Description: message,
				Footer: &Footer{
					Text: "veil-v2",
				},
				Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
			},
		},
	}

	jsonData, _ := json.Marshal(payload)

	headers := [][2]string{
		{"accept", "application/json"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/json"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"},
	}

	t.DoReq(t.MakeReq("POST", t.WebhookURL, headers, jsonData), "Sending Notification", false)
}

func (t *Task) Run() error {
	if t.Mode == "Signup" {
		return t.Signup()
	} else if t.Mode == "Classes" {
		return t.Classes()
	} else if t.Mode == "Transcript" {
		t.HomepageURL = "https://dw-prod.ec.fhda.edu/responsiveDashboard/worksheets/WEB31"
		return t.Transcript()
	} else if t.Mode == "Watch" {
		return t.Watch()
	}
	return nil
}
