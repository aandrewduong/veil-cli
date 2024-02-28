package tasks

import (
	"encoding/json"
	"fmt"
	"time"
)

func (t *Task) GetTerms() error {
	headers := [][2]string{
		{"accept", "application/json"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/json"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}
	response, err := t.DoReq(t.MakeReq("GET", fmt.Sprintf("https://reg-prod.ec.fhda.edu/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerm=&offset=1&max=10&_=%v", time.Now().UnixNano()/int64(time.Millisecond)), headers, nil), "Getting Terms", true)
	if err != nil {
		fmt.Println(err)
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	terms := Terms{}
	if err := json.Unmarshal(body, &terms); err != nil {
		return err
	}

	t.Terms = make(map[string]string)
	for _, term := range terms {
		t.Terms[term.Description] = term.Code
	}
	return nil
}

func (t *Task) GetTermByName(term string) {
	t.GetTerms()
	t.TermID = t.Terms[term]
	if len(t.TermID) > 0 {

	}
}
