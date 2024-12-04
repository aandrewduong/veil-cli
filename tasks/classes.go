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
	headers := [][2]string{
		{"accept", "*/*"},
		{"accept-language", "en-US,en;q=0.9"},
		{"content-type", "application/x-www-form-urlencoded"},
		{"user-agent", t.UserAgent},
	}

	values := url.Values{
		"term":            {t.TermID},
		"studyPath":       {},
		"startDatepicker": {},
		"endDatepicker":   {},
		"uniqueSessionId": {t.Session.UniqueSessionId},
	}

	response, err := t.DoReq(t.MakeReq("POST", "https://reg-prod.ec.fhda.edu/StudentRegistrationSsb/ssb/term/search?mode=search", headers, []byte(values.Encode())), "Submitting Term", true)
	if err != nil {
		discardResp(response)
		return err
	}

	return nil
}

func (t *Task) GetCourses() error {
	headers := [][2]string{
		{"accept", "application/json"},
		{"accept-language", "en-US,en;q=0.9"},
		{"user-agent", t.UserAgent},
	}

	url := fmt.Sprintf("https://reg-prod.ec.fhda.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=%s&txt_term=%s&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=100&sortColumn=subjectDescription&sortDirection=asc", t.Subject, t.TermID)
	response, err := t.DoReq(t.MakeReq("POST", url, headers, nil), fmt.Sprintf("Getting Courses (%s)", t.Subject), true)
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
		fmt.Println("No Courses Found")
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
	currentTime := time.Now()
	fileName := fmt.Sprintf("%s.csv", currentTime.Format("2006-01-02_15-04-05"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Term", "Course Reference Number", "Subject", "Course Number", "Sequence Number", "Course Title", "Display Name", "Begin Time", "End Time", "Start Date", "End Date", "Meeting Type", "Room", "Maximum Enrollment", "Enrollment", "Seats Available", "Waitlist Available"}
	fmt.Printf("Writing %s\n", fileName)
	err = writer.Write(header)
	if err != nil {
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
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}
	fmt.Println("Exported Search Data")
	return nil
}

func (t *Task) Classes() error {
	t.GenSessionId()
	t.SubmitTerm()
	t.GetCourses()
	return nil
}
