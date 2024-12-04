package tasks

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
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

func (t *Task) MakeReq(method string, url string, headers [][2]string, body []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}
	return req
}

func (t *Task) DoReq(req *http.Request, stage string, useDefaultResponseHandling bool) (*http.Response, error) {
	fmt.Println(stage)
	var resp *http.Response
	resp, err := t.Client.Do(req)

	if useDefaultResponseHandling {
		if resp.StatusCode >= 400 && resp.StatusCode <= 499 || resp.StatusCode >= 500 {
			body, _ := readBody(resp)
			reader := strings.NewReader(string(body))
			document, err := goquery.NewDocumentFromReader(reader)
			if err != nil {
				discardResp(resp)
				fmt.Println(err)
			}
			message := getSelectorAttr(document, "meta[name='errorMessage']", "content")
			fmt.Printf("Error %s [%d] %s\n", stage, resp.StatusCode, message)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1501)+1500))
			return t.DoReq(req, stage, useDefaultResponseHandling)
		}
	}
	return resp, err
}

// 12/4/2024: DoReqWithNewSession - Similar to DoReq, but creates a new client (blank state) with no cookies, only used for watch

func (t *Task) DoReqWithNewSession(req *http.Request, stage string, useDefaultResponseHandling bool) (*http.Response, error) {

	fmt.Println(stage)

	var dnsServers = []string{"8.8.8.8", "8.8.4.4", "1.1.1.1", "1.0.0.1"}
	randomIndex := rand.Intn(len(dnsServers))

	dnsServer := dnsServers[randomIndex]

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
		tls_client.WithClientProfile(profiles.Chrome_131),
		tls_client.WithDialer(dialer),
	}
	t.Client, _ = tls_client.NewHttpClient(tls_client.NewLogger(), client_options...)

	var resp *http.Response

	resp, err := t.Client.Do(req)

	if useDefaultResponseHandling {
		if resp.StatusCode >= 400 && resp.StatusCode <= 499 || resp.StatusCode >= 500 {
			body, _ := readBody(resp)
			reader := strings.NewReader(string(body))
			document, err := goquery.NewDocumentFromReader(reader)
			if err != nil {
				discardResp(resp)
				fmt.Println(err)
			}
			message := getSelectorAttr(document, "meta[name='errorMessage']", "content")
			fmt.Printf("Error %s [%d] %s\n", stage, resp.StatusCode, message)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1501)+1500))
			return t.DoReq(req, stage, useDefaultResponseHandling)
		}
	}
	return resp, err
}

func (t *Task) SendNotification(action string, message string) {
	payload := WebhookPayload{
		Username: "veil-cli",
		Embeds: []Embed{
			{
				Title:       action,
				Description: message,
				Footer: &Footer{
					Text: "veil-cli",
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
		{"user-agent", t.UserAgent},
	}

	t.DoReq(t.MakeReq("POST", t.WebhookURL, headers, []byte(string(jsonData))), "Sending Notification", false)
}

func discardResp(resp *http.Response) {
	if resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		defer resp.Body.Close()
	}
}

func readBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	return body, err
}

func getSelectorAttr(document *goquery.Document, selector string, attr string) string {
	value := ""
	document.Find(selector).Each(func(index int, element *goquery.Selection) {
		_value, exists := element.Attr(attr)
		if exists {
			value = _value
		}
	})
	return value
}

func extractModel(jsonData []byte) (map[string]interface{}, error) {
	var data AddCourse
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return data.Model, nil
}

func generateRandomString(length int) string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = characters[random.Intn(len(characters))]
	}
	return string(result)
}

func formatDuration(time time.Duration) string {
	totalSeconds := int64(time.Seconds())

	days := totalSeconds / (60 * 60 * 24)
	hours := (totalSeconds % (60 * 60 * 24)) / (60 * 60)
	minutes := (totalSeconds % (60 * 60)) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}

func saveRegistrationTime(registrationTime string) {
	file, err := os.Open("settings.csv")
	if err != nil {
		fmt.Println("Error Opening settings.csv:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error Reading settings.csv:", err)
		return
	}

	header := records[0]
	var timeIndex int
	found := false
	for i, field := range header {
		if field == "SavedRegistrationTime" {
			timeIndex = i
			found = true
			break
		}
	}

	if !found {
		fmt.Println("SavedRegistrationTime field not found in settings.csv")
		return
	}

	for i := 1; i < len(records); i++ {
		if len(records[i]) <= timeIndex {
			fmt.Println("Invalid Row, Missing SavedRegistrationTime field")
			continue
		}
		records[i][timeIndex] = registrationTime
	}

	outputFile, err := os.Create("settings.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error writing settings.csv :", err)
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing writer: ", err)
		return
	}

	fmt.Println("Saved Registration Time")
}

func (t *Task) Run() {
	if t.Mode == "Signup" {
		t.Signup()
	} else if t.Mode == "Classes" {
		t.Classes()
	} else if t.Mode == "Transcript" {
		t.HomepageURL = "https://dw-prod.ec.fhda.edu/responsiveDashboard/worksheets/WEB31"
		t.SSOManagerURL = "https://dw-prod.ec.fhda.edu/responsiveDashboard/saml/SSO"
		t.Transcript()
	} else if t.Mode == "Watch" {
		t.Watch()
	}
}
