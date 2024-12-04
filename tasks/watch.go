package tasks

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Update 12/4/2024: Updated CheckEnrollmentData to use DoReqWithNewSession to circumvent session blocks as it creates a new session (no cookies) per call
func (t *Task) CheckEnrollmentData(CRN string) error {
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/x-www-form-urlencoded"},
		{"user-agent", t.UserAgent},
	}

	values := url.Values{
		"term":                  {t.TermID},
		"courseReferenceNumber": {CRN},
	}

	response, err := t.DoReqWithNewSession(t.MakeReq("POST", "https://reg-prod.ec.fhda.edu/StudentRegistrationSsb/ssb/searchResults/getEnrollmentInfo", headers, []byte(values.Encode())), fmt.Sprintf("Getting Enrollment Data (%s)", CRN), true)
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

	var enrollmentSeatsAvailable, waitlistCapacity, waitlistActual, waitlistSeatsAvailable string

	document.Find("span.status-bold").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "Enrollment Seats Available:") {
			enrollmentSeatsAvailable = s.Next().Text()
		} else if strings.Contains(s.Text(), "Waitlist Seats Available:") {
			waitlistSeatsAvailable = s.Next().Text()
		} else if strings.Contains(s.Text(), "Waitlist Capacity:") {
			waitlistCapacity = s.Next().Text()
		} else if strings.Contains(s.Text(), "Waitlist Actual:") {
			waitlistActual = s.Next().Text()
		}
	})

	numEnrollmentSeatsAvailable, _ := strconv.Atoi(enrollmentSeatsAvailable)
	numWaitlistCapacity, _ := strconv.Atoi(waitlistCapacity)
	numWaitlistActual, _ := strconv.Atoi(waitlistActual)
	numWaitlistSeatsAvailable, _ := strconv.Atoi(waitlistSeatsAvailable)

	if numWaitlistCapacity > numWaitlistActual && (numWaitlistSeatsAvailable > 0) || (numEnrollmentSeatsAvailable > 0 && numWaitlistSeatsAvailable > 0) {
		t.SendNotification("Watch Task", fmt.Sprintf("[%s] %s Waitlist spot(s) is now Available", CRN, waitlistSeatsAvailable))
		t.WaitlistTask = true
		t.CRNs = []string{CRN}
		t.Signup()
	} else {
		if numEnrollmentSeatsAvailable >= 1 && numWaitlistSeatsAvailable == 0 {
			fmt.Printf("[%s] - (Waitlist Opening Soon)\n", CRN)
		} else {
			fmt.Printf("[%s] - (Not Available)\n", CRN)
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2501)+1500))
		return t.CheckEnrollmentData(CRN)
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
