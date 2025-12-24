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
	response, err := t.DoReq(t.MakeReq("GET", BaseDWURL+PathDWStudentMyself, t.GetHeaders("json"), nil), "Getting Student Data", true)
	if err != nil {
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
		if len(student.Goals) > 0 {
			transcriptSession.SchoolKey = student.Goals[0].School.Key
			transcriptSession.SchoolDescription = student.Goals[0].Degree.Key
			transcriptSession.Degree = student.Goals[0].Degree.Key
			transcriptSession.DegreeDescription = student.Goals[0].Degree.Description
		}
	}

	return t.GetAudit(transcriptSession)
}

func (t *Task) GetAudit(ts TranscriptSession) error {
	url := fmt.Sprintf("%s%s?studentId=%s&school=%s&degree=%s&is-process-new=false&audit-type=AA&auditId=&include-inprogress=true&include-preregistered=true&aid-term=", BaseDWURL, PathDWAudit, ts.UserId, ts.SchoolKey, ts.Degree)
	response, err := t.DoReq(t.MakeReq("GET", url, t.GetHeaders("json"), nil), "Getting Audit", true)
	if err != nil {
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
		auditInfo = append(auditInfo, AuditInfo{
			Term:        class.TermLiteralLong,
			Subject:     class.Discipline,
			Number:      class.Number,
			CourseTitle: class.CourseTitle,
			LetterGrade: class.LetterGrade,
			Credits:     class.Credits,
		})
	}
	return t.ExportTranscriptData(ts, auditInfo)
}

func (t *Task) ExportTranscriptData(ts TranscriptSession, auditInfo []AuditInfo) error {
	fileName := fmt.Sprintf("%s_%s_%s.csv", t.Username, ts.Name, time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Term", "Subject", "Number", "Course Title", "Letter Grade", "Credits"}
	fmt.Printf("[%s] Writing %s\n", t.Username, fileName)
	if err := writer.Write(header); err != nil {
		return err
	}
	for _, audit := range auditInfo {
		record := []string{audit.Term, audit.Subject, audit.Number, audit.CourseTitle, audit.LetterGrade, audit.Credits}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	fmt.Printf("[%s] Exported Transcript Data\n", t.Username)
	return nil
}

func (t *Task) Transcript() error {
	if err := t.GenSession(); err != nil {
		fmt.Printf("[%s] Authentication failed: %v\n", t.Username, err)
		return err
	}
	return t.GetStudentData()
}
