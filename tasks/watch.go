package tasks

import (
	"fmt"
	"goquery"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (t *Task) CheckEnrollmentData(CRN string) error {
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/x-www-form-urlencoded"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	values := url.Values{
		"term":                  {t.TermID},
		"courseReferenceNumber": {CRN},
	}

	response, err := t.DoReq(t.MakeReq("POST", "https://reg-prod.ec.fhda.edu/StudentRegistrationSsb/ssb/searchResults/getEnrollmentInfo", headers, []byte(values.Encode())), fmt.Sprintf("Getting Enrollment Data (%s)", CRN), true)
	if err != nil {
		fmt.Println(err)
		discardResp(response)
		return err
	}

	body, _ := readBody(response)
	reader := strings.NewReader(string(body))
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		discardResp(response)
		return err
	}

	var enrollmentSeatsAvailable, waitlistSeatsAvailable string

	document.Find("span.status-bold").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "Enrollment Seats Available:") {
			enrollmentSeatsAvailable = s.Next().Text()
		} else if strings.Contains(s.Text(), "Waitlist Seats Available:") {
			waitlistSeatsAvailable = s.Next().Text()
		}
	})

	numEnrollmentSeatsAvailable, _ := strconv.Atoi(enrollmentSeatsAvailable)
	numWaitlistSeatsAvailable, _ := strconv.Atoi(waitlistSeatsAvailable)

	if numEnrollmentSeatsAvailable == 0 {
		if numWaitlistSeatsAvailable > 0 {
			t.SendNotification("Watch Task", fmt.Sprintf("[%s] %s Waitlist spot(s) is now Available", CRN, waitlistSeatsAvailable))
		} else {
			time.Sleep(2 * time.Second)
			return t.CheckEnrollmentData(CRN)
		}
	} else {
		t.SendNotification("Watch Task", fmt.Sprintf("[%s] %s Enrollment Spot(s) is now Available", CRN, enrollmentSeatsAvailable))
	}
	return nil
}

func (t *Task) Watch() error {

	var waitGroup sync.WaitGroup
	errChan := make(chan error, len(t.CRNs))

	for _, course := range t.CRNs {
		waitGroup.Add(1)

		go func(course string) {
			defer waitGroup.Done()
			if err := t.CheckEnrollmentData(course); err != nil {
				errChan <- err
			}
		}(course)
	}

	waitGroup.Wait()
	close(errChan)
	return nil
}
