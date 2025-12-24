package tasks

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (t *Task) CheckEnrollmentData(CRN string) error {
	for {
		values := url.Values{
			"term":                  {t.TermID},
			"courseReferenceNumber": {CRN},
		}

		response, err := t.DoReq(t.MakeReq("POST", BaseRegURL+PathEnrollmentInfo, t.GetHeaders("form"), []byte(values.Encode())), fmt.Sprintf("Getting Enrollment Data (%s)", CRN), true)
		if err != nil {
			return err
		}

		body, _ := readBody(response)
		reader := strings.NewReader(string(body))
		document, err := goquery.NewDocumentFromReader(reader)
		if err != nil {
			discardResp(response)
			return err
		}

		var enrollmentSeatsAvailable, waitlistCapacity, waitlistActual, waitlistSeatsAvailable string
		document.Find("span.status-bold").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			nextText := s.Next().Text()
			if strings.Contains(text, "Enrollment Seats Available:") {
				enrollmentSeatsAvailable = nextText
			} else if strings.Contains(text, "Waitlist Seats Available:") {
				waitlistSeatsAvailable = nextText
			} else if strings.Contains(text, "Waitlist Capacity:") {
				waitlistCapacity = nextText
			} else if strings.Contains(text, "Waitlist Actual:") {
				waitlistActual = nextText
			}
		})

		numEnrollmentSeatsAvailable, _ := strconv.Atoi(enrollmentSeatsAvailable)
		numWaitlistCapacity, _ := strconv.Atoi(waitlistCapacity)
		numWaitlistActual, _ := strconv.Atoi(waitlistActual)
		numWaitlistSeatsAvailable, _ := strconv.Atoi(waitlistSeatsAvailable)

		if (numWaitlistCapacity > numWaitlistActual && numWaitlistSeatsAvailable > 0) || (numEnrollmentSeatsAvailable > 0) {
			t.SendNotification("Watch Task", fmt.Sprintf("[%s] Spot(s) available for %s", t.Username, CRN))
			t.WaitlistTask = true
			t.CRNs = []string{CRN}
			return t.Signup()
		}

		fmt.Printf("[%s] [%s] - Not Available (Waitlist: %d/%d)\n", t.Username, CRN, numWaitlistActual, numWaitlistCapacity)
		time.Sleep(time.Second * 10)
	}
}

func (t *Task) Watch() error {
	var wg sync.WaitGroup
	for _, course := range t.CRNs {
		wg.Add(1)
		go func(c string) {
			defer wg.Done()
			if err := t.CheckEnrollmentData(c); err != nil {
				fmt.Printf("[%s] Watch error for %s: %v\n", t.Username, c, err)
			}
		}(course)
	}
	wg.Wait()
	return nil
}
