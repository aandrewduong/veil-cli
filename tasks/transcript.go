package tasks

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TranscriptSession struct {
	Name              string
	UserId            string
	Degree            string
	DegreeDescription string
	SchoolKey         string
	SchoolDescription string
}

func (t *Task) GetStudentData() error {
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	response, err := t.DoReq(t.MakeReq("GET", "https://dw-prod.ec.fhda.edu/responsiveDashboard/api/students/myself", headers, nil), "Getting Student Data", true)
	if err != nil {
		fmt.Println(err)
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	userInfo := UserInfo{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return err
	}

	var transcriptSession TranscriptSession

	for _, student := range userInfo.Embedded.Students {
		transcriptSession.Name = student.Name
		transcriptSession.UserId = student.ID
		transcriptSession.SchoolKey = student.Goals[0].School.Key
		transcriptSession.SchoolDescription = student.Goals[0].Degree.Key
		transcriptSession.Degree = student.Goals[0].Degree.Key
		transcriptSession.DegreeDescription = student.Goals[0].Degree.Description
	}

	t.GetAudit(transcriptSession)
	return nil
}

func (t *Task) GetAudit(transcriptSession TranscriptSession) error {
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"},
	}

	response, err := t.DoReq(t.MakeReq("GET", fmt.Sprintf("https://dw-prod.ec.fhda.edu/responsiveDashboard/api/audit?studentId=%s&school=%s&degree=%s&is-process-new=false&audit-type=AA&auditId=&include-inprogress=true&include-preregistered=true&aid-term=", transcriptSession.UserId, transcriptSession.SchoolKey, transcriptSession.Degree), headers, nil), "Getting Audit", true)
	if err != nil {
		fmt.Println(err)
		discardResp(response)
		return err
	}
	body, _ := readBody(response)

	var auditInfo []AuditInfo
	audit := Audit{}
	if err := json.Unmarshal(body, &audit); err != nil {
		return err
	}

	for _, class := range audit.ClassInformation.ClassArray {
		classInfo := AuditInfo{
			Term:        class.TermLiteralLong,
			Subject:     class.Discipline,
			Number:      class.Number,
			CourseTitle: class.CourseTitle,
			LetterGrade: class.LetterGrade,
			Credits:     class.Credits,
		}
		auditInfo = append(auditInfo, classInfo)
	}
	t.ExportTranscriptData(transcriptSession, auditInfo)
	return nil
}

func (t *Task) ExportTranscriptData(transcriptSession TranscriptSession, auditInfo []AuditInfo) error {
	currentTime := time.Now()
	fileName := fmt.Sprintf("%s-%s-%s.csv", transcriptSession.Name, transcriptSession.Degree, currentTime.Format("2006-01-02_15-04-05"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Term", "Subject", "Number", "Course Title", "Letter Grade", "Credits"}
	fmt.Printf("Writing %s\n", fileName)
	err = writer.Write(header)
	if err != nil {
		return err
	}
	for _, audit := range auditInfo {
		record := []string{
			audit.Term,
			audit.Subject,
			audit.Number,
			audit.CourseTitle,
			audit.LetterGrade,
			audit.Credits,
		}
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}
	fmt.Println("Exported Transcript Data")
	return nil
}

func (t *Task) Transcript() error {
	t.GenSession()
	t.GetStudentData()
	return nil
}
