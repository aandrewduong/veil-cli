package tasks

import (
	"fmt"
	"goquery"
	"net/url"
	"strings"
	"time"
)

type Session struct {
	LoginAttempts   int
	SAMLResponse    string
	RelayState      string
	SignupSession   SignupSession
	UniqueSessionId string
}

func (t *Task) GenSessionId() error {
	t.Session.UniqueSessionId = fmt.Sprintf("%s%v", strings.ToLower(generateRandomString(5)), time.Now().UnixNano()/int64(time.Millisecond))
	return nil
}

func (t *Task) VisitHomepage() error {

	headers := [][2]string{
		{"accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"},
		{"accept-language", "en-US,en;q=0.9"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	response, err := t.DoReq(t.MakeReq("GET", t.HomepageURL, headers, nil), "Gen Session", true)
	if err != nil {
		discardResp(response)
		return err
	}
	return nil
}

func (t *Task) Login() error {
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/x-www-form-urlencoded"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	t.Session.LoginAttempts++

	values := url.Values{}
	values.Set("j_username", t.Username)
	values.Set("j_password", t.Password)
	values.Set("_eventId_proceed", "")
	response, err := t.DoReq(t.MakeReq("POST", fmt.Sprintf("https://ssoshib.fhda.edu/idp/profile/SAML2/Redirect/SSO?execution=e1s%d", t.Session.LoginAttempts), headers, []byte(values.Encode())), "Logging In", true)
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
	var message string
	document.Find("div[class='alert alert-danger']").Each(func(index int, element *goquery.Selection) {
		message = strings.TrimSpace(element.Text())
	})

	switch message {
	case "The username you entered cannot be identified.":
		fmt.Println("Invalid Username")
	case "The password you entered was incorrect.":
		fmt.Println("Invalid Password")
		time.Sleep(2 * time.Second)
		t.Login()
	case "You may be seeing this page because you used the Back button while browsing a secure web site or application. Alternatively, you may have mistakenly bookmarked the web login form instead of the actual web site you wanted to bookmark or used a link created by somebody else who made the same mistake.  Left unchecked, this can cause errors on some browsers or result in you returning to the web site you tried to leave, so this page is presented instead.":
		fmt.Println("Bad Session")
		t.GenSession()
	case "":
		break
	default:
		fmt.Println(message)
		time.Sleep(2 * time.Second)
		t.Login()
	}

	relayState := getSelectorAttr(document, "input[name='RelayState']", "value")
	samlResponse := getSelectorAttr(document, "input[name='SAMLResponse']", "value")

	t.Session.RelayState = relayState
	t.Session.SAMLResponse = samlResponse
	return nil
}

func (t *Task) SubmitCommonAuth() error {

	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/x-www-form-urlencoded"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	values := url.Values{
		"RelayState":   {t.Session.RelayState},
		"SAMLResponse": {t.Session.SAMLResponse},
	}

	response, err := t.DoReq(t.MakeReq("POST", "https://eis-prod.ec.fhda.edu/commonauth", headers, []byte(values.Encode())), "Submitting Common Auth", true)
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
	var message string
	document.Find("div[class='retry-msg-text text_right_custom']").Each(func(index int, element *goquery.Selection) {
		message = strings.TrimSpace(element.Text())
	})
	if strings.Contains(message, "Authentication Error!") {
		fmt.Println("")
	}

	relayState := getSelectorAttr(document, "input[name='RelayState']", "value")
	samlResponse := getSelectorAttr(document, "input[name='SAMLResponse']", "value")

	t.Session.RelayState = relayState
	t.Session.SAMLResponse = samlResponse
	return nil
}

func (t *Task) SubmitSSOManager() error {
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/x-www-form-urlencoded"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	values := url.Values{
		"RelayState":   {t.Session.RelayState},
		"SAMLResponse": {t.Session.SAMLResponse},
	}

	response, err := t.DoReq(t.MakeReq("POST", t.SSOManagerURL, headers, []byte(values.Encode())), "Submitting SSO Manager", true)
	if err != nil {
		discardResp(response)
		return err
	}
	return nil
}

func (t *Task) GenSession() {
	t.GenSessionId()
	t.VisitHomepage()
	t.Login()
	t.SubmitCommonAuth()
	t.SubmitSSOManager()
	t.RegisterPostSignIn()
	t.SubmitSamIsso()
	t.SubmitSSBSp()
}
