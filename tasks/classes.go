package tasks

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"
)

func (t *Task) SubmitTerm() error {
	values := url.Values{
		"term":            {t.TermID},
		"studyPath":       {},
		"startDatepicker": {},
		"endDatepicker":   {},
		"uniqueSessionId": {t.Session.UniqueSessionId},
	}

	response, err := t.DoReq(t.MakeReq("POST", BaseRegURL+PathTermSearch+"?mode=search", t.GetHeaders("form"), []byte(values.Encode())), "Submitting Term", true)
	if err != nil {
		discardResp(response)
		return err
	}
	return nil
}

func (t *Task) GetCourses() error {
	url := fmt.Sprintf("%s%s?txt_subject=%s&txt_term=%s&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=100&sortColumn=subjectDescription&sortDirection=asc", BaseRegURL, PathSearchResults, t.Subject, t.TermID)
	response, err := t.DoReq(t.MakeReq("POST", url, t.GetHeaders("json"), nil), fmt.Sprintf("Getting Courses (%s)", t.Subject), true)
	if err != nil {
		discardResp(response)
		return err
	}
	body, _ := readBody(response)
	courses := Courses{}
	if err := json.Unmarshal(body, &courses); err != nil {
		return err
	}

	if courses.TotalCount == 0 {
		fmt.Printf("[%s] No Courses Found\n", t.Username)
		return nil
	}

	var coursesInfo []CourseInfo
	for _, section := range courses.Data {
		for _, faculty := range section.Faculty {
			for _, meetingfaculty := range section.MeetingsFaculty {
				course := CourseInfo{
					TermDesc:              section.TermDesc,
					CourseReferenceNumber: faculty.CourseReferenceNumber,
					Subject:               section.Subject,
					CourseNumber:          section.CourseNumber,
					SequenceNumber:        section.SequenceNumber,
					CourseTitle:           section.CourseTitle,
					DisplayName:           faculty.DisplayName,
					BeginTime:             meetingfaculty.MeetingTime.BeginTime,
					EndTime:               meetingfaculty.MeetingTime.EndTime,
					StartDate:             meetingfaculty.MeetingTime.StartDate,
					EndDate:               meetingfaculty.MeetingTime.EndDate,
					MeetingType:           meetingfaculty.MeetingTime.MeetingTypeDescription,
					Room:                  meetingfaculty.MeetingTime.Room,
					MaximumEnrollment:     section.MaximumEnrollment,
					Enrollment:            section.Enrollment,
					SeatsAvailable:        section.SeatsAvailable,
					WaitAvailable:         section.WaitAvailable,
				}
				coursesInfo = append(coursesInfo, course)
			}
		}
	}

	t.ExportCourseData(coursesInfo)
	return nil
}

func (t *Task) ExportCourseData(courses []CourseInfo) error {
	fileName := fmt.Sprintf("%s_%s.csv", t.Username, time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Term", "Course Reference Number", "Subject", "Course Number", "Sequence Number", "Course Title", "Display Name", "Begin Time", "End Time", "Start Date", "End Date", "Meeting Type", "Room", "Maximum Enrollment", "Enrollment", "Seats Available", "Waitlist Available"}
	fmt.Printf("[%s] Writing %s\n", t.Username, fileName)
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, course := range courses {
		record := []string{
			course.TermDesc,
			course.CourseReferenceNumber,
			course.Subject,
			course.CourseNumber,
			course.SequenceNumber,
			course.CourseTitle,
			course.DisplayName,
			course.BeginTime,
			course.EndTime,
			course.StartDate,
			course.EndDate,
			course.MeetingType,
			course.Room,
			strconv.Itoa(course.MaximumEnrollment),
			strconv.Itoa(course.Enrollment),
			strconv.Itoa(course.SeatsAvailable),
			strconv.Itoa(course.WaitAvailable),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	fmt.Printf("[%s] Exported Search Data\n", t.Username)
	return nil
}

func (t *Task) Classes() error {
	t.HomepageURL = BaseRegURL + "/StudentRegistrationSsb/saml/login"
	if err := t.CheckAuthSession(); err != nil {
		fmt.Printf("[%s] Authentication failed: %v\n", t.Username, err)
		return err
	}
	t.SubmitTerm()
	t.GetCourses()
	return nil
}
