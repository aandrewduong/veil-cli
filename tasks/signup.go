package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SignupSession struct {
	SAMLRequest string
	Model       map[string]interface{}
}

func (t *Task) CheckAuthSession() error {
	response, err := t.DoReq(t.MakeReq("GET", BaseRegURL+PathAuthAjax, t.GetHeaders("html"), nil), "Checking Auth Session", true)
	if err != nil {
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	if strings.Contains(string(body), "userNotLoggedIn") {
		return t.GenSession()
	}
	return nil
}

func (t *Task) RegisterPostSignIn() error {
	response, err := t.DoReq(t.MakeReq("GET", BaseRegURL+PathRegisterPostSignIn+"?mode=registration", t.GetHeaders("html"), nil), "Register Post Sign In", true)
	if err != nil {
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
	t.Session.SignupSession.SAMLRequest = getSelectorAttr(document, "input[name='SAMLRequest']", "value")
	return nil
}

func (t *Task) SubmitSamIsso() error {
	values := url.Values{"SAMLRequest": {t.Session.SignupSession.SAMLRequest}}
	response, err := t.DoReq(t.MakeReq("POST", BaseEISURL+PathSamlSSO, t.GetHeaders("form"), []byte(values.Encode())), "Submitting Sam Isso", true)
	if err != nil {
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
	t.Session.SAMLResponse = getSelectorAttr(document, "input[name='SAMLResponse']", "value")
	return nil
}

func (t *Task) SubmitSSBSp() error {
	values := url.Values{"SAMLResponse": {t.Session.SAMLResponse}}
	resp, err := t.DoReq(t.MakeReq("POST", BaseRegURL+PathSamlSSOSp, t.GetHeaders("form"), []byte(values.Encode())), "Submitting SSB SP", true)
	if err != nil {
		discardResp(resp)
		return err
	}
	return nil
}

func (t *Task) CheckCRN(course string) error {
	u := fmt.Sprintf("%s%s?courseReferenceNumber=%s&term=%s", BaseRegURL, PathSectionDetails, course, t.TermID)
	response, err := t.DoReq(t.MakeReq("GET", u, t.GetHeaders("html"), nil), fmt.Sprintf("Checking Course (%s)", course), true)
	if err != nil {
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	courseData := Course{}
	if err := json.Unmarshal(body, &courseData); err != nil {
		return err
	}
	if !courseData.Olr {
		fmt.Printf("[%s] %s\n", t.Username, courseData.ResponseDisplay)
	} else {
		fmt.Printf("[%s] [%s] - Unable to retrieve section data\n", t.Username, course)
	}
	return nil
}

func (t *Task) CheckCRNs() error {
	for _, course := range t.CRNs {
		t.CheckCRN(course)
	}
	return nil
}

func (t *Task) GetRegistrationStatus() error {
	values := url.Values{
		"term":            {t.TermID},
		"studyPath":       {},
		"startDatepicker": {},
		"endDatepicker":   {},
		"uniqueSessionId": {t.Session.UniqueSessionId},
	}

	response, err := t.DoReq(t.MakeReq("POST", BaseRegURL+PathTermSearch+"?mode=registration", t.GetHeaders("form"), []byte(values.Encode())), "Getting Registration Status", true)
	if err != nil {
		return err
	}

	body, _ := readBody(response)
	registrationStatus := RegistrationStatus{}
	if err := json.Unmarshal(body, &registrationStatus); err != nil {
		return err
	}

	var hasFailure, hasRegistrationTime bool
	var timeFailure string

	for _, failure := range registrationStatus.StudentEligFailures {
		fmt.Printf("[%s] Eligibility Message: %s\n", t.Username, failure)
		hasFailure = true
		if strings.Contains(failure, "You can register from") {
			hasRegistrationTime = true
			timeFailure = failure
			break
		}
	}

	if !hasFailure {
		return nil
	}

	if !hasRegistrationTime {
		return errors.New(registrationStatus.StudentEligFailures[len(registrationStatus.StudentEligFailures)-1])
	}

	pattern := regexp.MustCompile(`\d{2}/\d{2}/\d{4} \d{2}:\d{2} [APM]{2}`)
	matches := pattern.FindAllString(timeFailure, -1)

	if len(matches) > 0 {
		location, _ := time.LoadLocation("America/Los_Angeles")
		targetTime, _ := time.ParseInLocation("01/02/2006 03:04 PM", matches[0], location)
		now := time.Now().In(location)

		if now.After(targetTime) {
			fmt.Printf("[%s] Registration window is open. Proceeding...\n", t.Username)
			time.Sleep(2 * time.Second)
			return nil
		} else {
			t.CheckCRNs()
			timeToWait := targetTime.Sub(now)

			fmt.Printf("[%s] Waiting for registration window: %s\n", t.Username, targetTime.Format(time.RFC1123))
			fmt.Printf("[%s] Re-checking in %s\n", t.Username, formatDuration(targetTime.Sub(now)))

			// Wait in a separate goroutine to keep session alive
			stopChan := make(chan bool)
			go func() {
				ticker := time.NewTicker(5 * time.Minute)
				defer ticker.Stop()
				for {
					select {
					case <-ticker.C:
						if err := t.CheckAuthSession(); err != nil {
							fmt.Printf("[%s] Session heartbeat failed: %v\n", t.Username, err)
						}
					case <-stopChan:
						return
					}
				}
			}()

			time.Sleep(timeToWait)
			close(stopChan)
			if err := t.CheckAuthSession(); err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (t *Task) VisitClassRegistration() error {
	response, err := t.DoReq(t.MakeReq("HEAD", BaseRegURL+PathClassRegistration, t.GetHeaders("html"), nil), "Visiting Class Registration", true)
	if err != nil {
		discardResp(response)
		return err
	}
	return nil
}

func (t *Task) AddCourse(course string) error {
	u := fmt.Sprintf("%s%s?term=%s&courseReferenceNumber=%s&olr=false", BaseRegURL, PathAddRegistrationItem, t.TermID, course)
	response, err := t.DoReq(t.MakeReq("GET", u, t.GetHeaders("html"), nil), fmt.Sprintf("Adding Course (%s)", course), true)
	if err != nil {
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	addCourse := AddCourse{}

	if err := json.Unmarshal(body, &addCourse); err != nil {
		return err
	}
	if addCourse.Success {
		model, err := extractModel(body)
		if err != nil {
			return err
		}
		model["selectedAction"] = "WL"
		t.Session.SignupSession.Model = model
	} else {
		fmt.Printf("[%s] Registration Error (%s): %s\n", t.Username, course, addCourse.Message)
	}
	return nil
}

func (t *Task) AddCourses() error {
	for _, course := range t.CRNs {
		if err := t.AddCourse(course); err != nil {
			return err
		}
	}
	if len(t.Session.SignupSession.Model) == 0 {
		return errors.New("failed to add any courses")
	}
	return nil
}

func (t *Task) SendBatch() error {
	batch := Batch{
		Update:          []map[string]interface{}{t.Session.SignupSession.Model},
		UniqueSessionId: t.Session.UniqueSessionId,
	}

	batchJson, _ := json.MarshalIndent(batch, "", "  ")
	response, err := t.DoReq(t.MakeReq("POST", BaseRegURL+PathSubmitRegistration, t.GetHeaders("json"), batchJson), "Submitting Batch Update", true)
	if err != nil {
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	changes := Changes{}
	if err := json.Unmarshal(body, &changes); err != nil {
		return err
	}

	for _, data := range changes.Data.Update {
		for _, crn := range t.CRNs {
			if data.CourseReferenceNumber == crn {
				if data.StatusDescription == "Registered" {
					fmt.Printf("[%s] [%s] Successful Registration: %s %s - %s\n", t.Username, data.CourseReferenceNumber, data.Subject, data.CourseNumber, data.CourseTitle)
					t.SendNotification(data.CourseTitle, fmt.Sprintf("Registration Successful (%s)", data.CourseReferenceNumber))
				} else if data.StatusDescription == "Waitlisted" {
					fmt.Printf("[%s] [%s] Successful Waitlist: %s %s - %s\n", t.Username, data.CourseReferenceNumber, data.Subject, data.CourseNumber, data.CourseTitle)
					t.SendNotification(data.CourseTitle, fmt.Sprintf("Waitlist Successful (%s)", data.CourseReferenceNumber))
				} else if data.StatusDescription == "Errors Preventing Registration" {
					fmt.Printf("[%s] [%d] Registration errors for %s\n", t.Username, len(data.CrnErrors), data.CourseReferenceNumber)
					for _, err := range data.CrnErrors {
						fmt.Printf("[%s] Error: %s\n", t.Username, err.Message)
					}
				}
			}
		}
	}
	return nil
}

func (t *Task) Signup() error {
	t.HomepageURL = BaseRegURL + "/StudentRegistrationSsb/saml/login"
	t.SSOManagerURL = BaseSSBManagerURL + "/ssomanager/saml/SSO"
	if err := t.CheckAuthSession(); err != nil {
		fmt.Printf("[%s] Authentication failed: %v\n", t.Username, err)
		return err
	}
	if err := t.GetRegistrationStatus(); err != nil {
		return err
	}
	if err := t.VisitClassRegistration(); err != nil {
		return err
	}
	if err := t.AddCourses(); err != nil {
		return err
	}
	if err := t.SendBatch(); err != nil {
		return err
	}
	t.Client.CloseIdleConnections()
	return nil
}
