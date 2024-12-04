package tasks

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (t *Task) GetTerms() error {
	headers := [][2]string{
		{"accept", "application/json"},
		{"accept-language", "en-US,en;q=0.9"},
		{"user-agent", t.UserAgent},
	}

	response, err := t.DoReq(t.MakeReq("GET", fmt.Sprintf("https://reg-prod.ec.fhda.edu/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerm=&offset=1&max=100&_=%v", time.Now().UnixNano()/int64(time.Millisecond)), headers, nil), "Getting Terms", true)
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

func BuildTermId(term string) string {

	fmt.Println("Building Term ID (Offline)")
	var campus string
	data := strings.Fields(term)
	year, quarter := data[0], data[1]

	if len(data) == 4 {
		campus = "De Anza"
	}
	quarterCode := QuarterCodes[quarter]
	campusCode := CampusCodes[campus]

	yearInt, _ := strconv.Atoi(year)
	if quarter == "Summer" {
		yearInt++
	}

	return fmt.Sprintf("%d%d%d", yearInt, quarterCode, campusCode)
}

func (t *Task) GetTermByName(term string) {
	t.GetTerms()
	t.TermID = t.Terms[term]
	if t.Terms[term] == "" {
		t.TermID = BuildTermId(term)
	}
}
