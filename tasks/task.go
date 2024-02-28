package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"math/rand"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

type Task struct {
	Terms      map[string]string
	Username   string
	Password   string
	Subject    string
	CRNs       []string
	TermID     string
	UserAgent  string
	Client     tls_client.HttpClient
	Session    Session
	WebhookURL string
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
		if resp.StatusCode == 400 || resp.StatusCode >= 500 {
			body, _ := readBody(resp)
			reader := strings.NewReader(string(body))
			document, err := goquery.NewDocumentFromReader(reader)
			if err != nil {
				discardResp(resp)
				fmt.Println(err)
			}
			message := getSelectorAttr(document, "meta[name='errorMessage']", "content")
			fmt.Printf("Error %s [%d] %s\n", stage, resp.StatusCode, message)
			time.Sleep(time.Second * 2)
			return t.DoReq(req, stage, useDefaultResponseHandling)
		}
	}
	return resp, err
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
				Timestamp: time.Now().Format("2006-01-02T15:04:05.000Z"),
			},
		},
	}

	jsonData, _ := json.Marshal(payload)

	headers := [][2]string{
		{"accept", "application/json"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/json"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
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
