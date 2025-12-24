package tasks

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (t *Task) GetTerms() error {
	url := fmt.Sprintf("%s%s?searchTerm=&offset=1&max=100&_=%v", BaseRegURL, PathGetTerms, time.Now().UnixMilli())
	response, err := t.DoReq(t.MakeReq("GET", url, t.GetHeaders("html"), nil), "Getting Terms", true)
	if err != nil {
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
	data := strings.Fields(term)
	if len(data) < 2 {
		return ""
	}

	year, quarter := data[0], data[1]
	var campus string
	if len(data) >= 4 {
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
	if err := t.GetTerms(); err != nil {
		fmt.Printf("[%s] Error getting terms: %v\n", t.Username, err)
	}
	t.TermID = t.Terms[term]
	if t.TermID == "" {
		t.TermID = BuildTermId(term)
	}
}
