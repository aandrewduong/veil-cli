package tasks

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Session struct {
	LoginAttempts   int
	SAMLResponse    string
	RelayState      string
	SignupSession   SignupSession
	UniqueSessionId string
	SAMLRequest     string
}

func (t *Task) GenSessionId() error {
	t.Session.UniqueSessionId = fmt.Sprintf("%s%v", strings.ToLower(generateRandomString(5)), time.Now().UnixNano()/int64(time.Millisecond))
	return nil
}

func (t *Task) VisitHomepage() error {
	response, err := t.DoReq(t.MakeReq("GET", t.HomepageURL, t.GetHeaders("html"), nil), "Gen Session", true)
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

	t.Session.SAMLRequest = getSelectorAttr(document, "input[name='SAMLRequest']", "value")
	return nil
}

func (t *Task) PreLoginSSO() error {
	values := url.Values{"SAMLRequest": {t.Session.SAMLRequest}}
	response, err := t.DoReq(t.MakeReq("POST", BaseSSOURL+"/idp/profile/SAML2/POST/SSO", t.GetHeaders("form"), []byte(values.Encode())), "Submitting SSO Request", true)
	if err != nil {
		discardResp(response)
		return err
	}
	return nil
}

func (t *Task) Login() error {
	for {
		t.Session.LoginAttempts++

		values := url.Values{}
		values.Set("j_username", t.Username)
		values.Set("j_password", t.Password)
		values.Set("_eventId_proceed", "")

		u := fmt.Sprintf("%s/idp/profile/SAML2/POST/SSO?execution=e1s%d", BaseSSOURL, t.Session.LoginAttempts)
		response, err := t.DoReq(t.MakeReq("POST", u, t.GetHeaders("form"), []byte(values.Encode())), "Logging In", true)
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

		var message string
		document.Find("div[class='alert alert-danger']").Each(func(index int, element *goquery.Selection) {
			message = strings.TrimSpace(element.Text())
		})

		if message == "" {
			t.Session.RelayState = getSelectorAttr(document, "input[name='RelayState']", "value")
			t.Session.SAMLResponse = getSelectorAttr(document, "input[name='SAMLResponse']", "value")
			return nil
		}

		switch {
		case strings.Contains(message, "cannot be identified"):
			fmt.Printf("[%s] Invalid Username\n", t.Username)
			return fmt.Errorf("invalid username")
		case strings.Contains(message, "incorrect"):
			fmt.Printf("[%s] Invalid Password\n", t.Username)
			return fmt.Errorf("invalid password")
		case strings.Contains(message, "Back button"):
			fmt.Printf("[%s] Bad session state (Back button error). Retrying...\n", t.Username)
			return fmt.Errorf("bad session state")
		default:
			fmt.Printf("[%s] Login Message: %s\n", t.Username, message)
			time.Sleep(time.Second * 2)
			// Continue loop to retry
		}
	}
}

func (t *Task) SubmitSSOManager() error {
	values := url.Values{"SAMLResponse": {t.Session.SAMLResponse}}
	response, err := t.DoReq(t.MakeReq("POST", BaseRegURL+"/StudentRegistrationSsb/saml/SSO", t.GetHeaders("form"), []byte(values.Encode())), "Submitting SSO Manager", true)
	if err != nil {
		discardResp(response)
		return err
	}
	return nil
}

func (t *Task) Check() error {
	response, err := t.DoReq(t.MakeReq("GET", BaseRegURL+"/StudentRegistrationSsb/ssb/registration", t.GetHeaders("html"), nil), "Checking Session", true)
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

	fullName := getSelectorAttr(document, "meta[name='fullName']", "content")
	if fullName != "" {
		fmt.Printf("[%s] Session active for %s.\n", t.Username, fullName)
	}
	return nil
}

func (t *Task) GenSession() error {
	t.GenSessionId()
	if err := t.VisitHomepage(); err != nil {
		return err
	}
	if err := t.PreLoginSSO(); err != nil {
		return err
	}
	if err := t.Login(); err != nil {
		return err
	}
	if err := t.SubmitSSOManager(); err != nil {
		return err
	}
	return t.Check()
}
