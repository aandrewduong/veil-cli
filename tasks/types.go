package tasks

type Terms []struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type RegistrationStatus struct {
	StudentEligValid    bool     `json:"studentEligValid"`
	StudentEligFailures []string `json:"studentEligFailures"`
	FwdURL              string   `json:"fwdURL"`
}

type Course struct {
	Subject         string `json:"subject"`
	CourseTitle     string `json:"courseTitle"`
	SequenceNumber  string `json:"sequenceNumber"`
	CourseNumber    string `json:"courseNumber"`
	ResponseDisplay string `json:"responseDisplay"`
	Olr             bool   `json:"olr"`
	ProjectionError bool   `json:"projectionError"`
	Success         bool   `json:"success"`
}

type AddCourse struct {
	Success bool                   `json:"success"`
	Model   map[string]interface{} `json:"model"`
	Message string                 `json:"message"`
}

type Batch struct {
	Create          []map[string]interface{} `json:"create"`
	Update          []map[string]interface{} `json:"update"`
	Destroy         []map[string]interface{} `json:"destroy"`
	UniqueSessionId string                   `json:"uniqueSessionId"`
}

type ChangesCRNError struct {
	Class       string `json:"class"`
	ErrorFlag   string `json:"errorFlag"`
	Message     string `json:"message"`
	MessageType string `json:"messageType"`
}

type Changes struct {
	Success bool `json:"success"`
	Data    struct {
		Create  []any `json:"create"`
		Destroy []any `json:"destroy"`
		Update  []struct {
			RegistrationToDate                 any    `json:"registrationToDate"`
			ErrorFlag                          string `json:"errorFlag"`
			ReservedKey                        any    `json:"reservedKey"`
			CourseContinuingEducationIndicator string `json:"courseContinuingEducationIndicator"`
			MessageType                        any    `json:"messageType"`
			CourseRegistrationStatus           string `json:"courseRegistrationStatus"`
			ApprovalOverride                   string `json:"approvalOverride"`
			SubjectDescription                 string `json:"subjectDescription"`
			StudentAttributeOverride           any    `json:"studentAttributeOverride"`
			SelectedLevel                      struct {
				Class       string `json:"class"`
				Description any    `json:"description"`
				Level       any    `json:"level"`
			} `json:"selectedLevel"`
			CampusOverride                      any `json:"campusOverride"`
			GradingMode                         any `json:"gradingMode"`
			SelectedCreditHour                  any `json:"selectedCreditHour"`
			Block                               any `json:"block"`
			GradeComment                        any `json:"gradeComment"`
			CreditHourInitial                   any `json:"creditHourInitial"`
			SequenceNumber                      any `json:"sequenceNumber"`
			CourseRegistrationStatusDescription any `json:"courseRegistrationStatusDescription"`
			TuitionWaiverIndicator              any `json:"tuitionWaiverIndicator"`
			RepeatOverride                      any `json:"repeatOverride"`
			PreqOverride                        any `json:"preqOverride"`
			CreditHour                          any `json:"creditHour"`
			MajorOverride                       any `json:"majorOverride"`
			CourseReferenceNumber               any `json:"courseReferenceNumber"`
			CreditHours                         struct {
				Class               string `json:"class"`
				CreditHourHigh      any    `json:"creditHourHigh"`
				CreditHourIndicator any    `json:"creditHourIndicator"`
				CreditHourList      any    `json:"creditHourList"`
				CreditHourLow       any    `json:"creditHourLow"`
			} `json:"creditHours"`
			StructuredRegistrationHeaderSequence any    `json:"structuredRegistrationHeaderSequence"`
			ScheduleType                         string `json:"scheduleType"`
			ApprovalReceivedIndicator            any    `json:"approvalReceivedIndicator"`
			Grade                                any    `json:"grade"`
			ProgramOverride                      any    `json:"programOverride"`
			StartDate                            string `json:"startDate"`
			AddAuthorizationCrnMessage           any    `json:"addAuthorizationCrnMessage"`
			OriginalCourseRegistrationStatus     string `json:"originalCourseRegistrationStatus"`
			RegistrationLevels                   []any  `json:"registrationLevels"`
			AuthorizationCode                    any    `json:"authorizationCode"`
			AttemptedHours                       any    `json:"attemptedHours"`
			SelectedGradingMode                  struct {
				Class       string `json:"class"`
				Description any    `json:"description"`
				GradingMode any    `json:"gradingMode"`
			} `json:"selectedGradingMode"`
			WaivHour                             any    `json:"waivHour"`
			LevelDescription                     string `json:"levelDescription"`
			SelectedStartEndDate                 any    `json:"selectedStartEndDate"`
			SelectedOverride                     any    `json:"selectedOverride"`
			OriginalRecordStatus                 string `json:"originalRecordStatus"`
			TestOverride                         any    `json:"testOverride"`
			SubmitResultIndicator                any    `json:"submitResultIndicator"`
			DataOrigin                           string `json:"dataOrigin"`
			Term                                 string `json:"term"`
			MexcOverride                         string `json:"mexcOverride"`
			RegistrationOverrides                []any  `json:"registrationOverrides"`
			LevelOverride                        string `json:"levelOverride"`
			StructuredRegistrationDetailSequence any    `json:"structuredRegistrationDetailSequence"`
			RegistrationFromDate                 any    `json:"registrationFromDate"`
			WaitCapacity                         any    `json:"waitCapacity"`
			BlockRuleSequenceNumber              any    `json:"blockRuleSequenceNumber"`
			BillHours                            struct {
				Class               string `json:"class"`
				CreditHourHigh      any    `json:"creditHourHigh"`
				CreditHourIndicator any    `json:"creditHourIndicator"`
				CreditHourList      any    `json:"creditHourList"`
				CreditHourLow       any    `json:"creditHourLow"`
			} `json:"billHours"`
			BlockPermitOverride       any    `json:"blockPermitOverride"`
			StudyPathName             any    `json:"studyPathName"`
			AddDate                   string `json:"addDate"`
			StudyPathKeySequence      any    `json:"studyPathKeySequence"`
			AddAuthorizationCrnStatus any    `json:"addAuthorizationCrnStatus"`
			RecordStatus              string `json:"recordStatus"`
			NumberOfUnits             any    `json:"numberOfUnits"`
			ScheduleDescription       string `json:"scheduleDescription"`
			TimeStatusHours           any    `json:"timeStatusHours"`
			SelectedBillHour          any    `json:"selectedBillHour"`
			GradeDate                 any    `json:"gradeDate"`
			RegistrationActions       []struct {
				Class                    string `json:"class"`
				CourseRegistrationStatus any    `json:"courseRegistrationStatus"`
				Description              string `json:"description"`
				RegistrationStatusDate   any    `json:"registrationStatusDate"`
				Remove                   bool   `json:"remove"`
				SubActions               any    `json:"subActions"`
				VoiceType                any    `json:"voiceType"`
			} `json:"registrationActions"`
			ErrorLink                           any               `json:"errorLink"`
			MaxEnrollment                       any               `json:"maxEnrollment"`
			StatusIndicator                     string            `json:"statusIndicator"`
			Subject                             string            `json:"subject"`
			RemoveIndicator                     any               `json:"removeIndicator"`
			BillHour                            any               `json:"billHour"`
			PartOfTerm                          string            `json:"partOfTerm"`
			SectionCourseTitle                  string            `json:"sectionCourseTitle"`
			OverrideDurationIndicator           bool              `json:"overrideDurationIndicator"`
			CourseTitle                         string            `json:"courseTitle"`
			DuplicateOverride                   string            `json:"duplicateOverride"`
			DurationUnit                        any               `json:"durationUnit"`
			CensusEnrollmentDate                any               `json:"censusEnrollmentDate"`
			CorqOverride                        string            `json:"corqOverride"`
			GradeMid                            any               `json:"gradeMid"`
			Level                               string            `json:"level"`
			InstructionalMethodDescription      string            `json:"instructionalMethodDescription"`
			Campus                              string            `json:"campus"`
			DegreeOverride                      any               `json:"degreeOverride"`
			WaitOverride                        any               `json:"waitOverride"`
			NewBlock                            any               `json:"newBlock"`
			RpthOverride                        string            `json:"rpthOverride"`
			NewBlockRuleSequenceNumber          any               `json:"newBlockRuleSequenceNumber"`
			GradingModeDescription              string            `json:"gradingModeDescription"`
			PartOfTermDescription               string            `json:"partOfTermDescription"`
			TimeOverride                        string            `json:"timeOverride"`
			RegistrationAuthorizationActiveCode any               `json:"registrationAuthorizationActiveCode"`
			LinkOverride                        string            `json:"linkOverride"`
			RegistrationStatusDate              string            `json:"registrationStatusDate"`
			DepartmentOverride                  any               `json:"departmentOverride"`
			CrnErrors                           []ChangesCRNError `json:"crnErrors"`
			RegistrationGradingModes            []struct {
				Class       string `json:"class"`
				Description string `json:"description"`
				GradingMode string `json:"gradingMode"`
			} `json:"registrationGradingModes"`
			SelectedStudyPath struct {
				Class             string `json:"class"`
				Description       any    `json:"description"`
				KeySequenceNumber any    `json:"keySequenceNumber"`
			} `json:"selectedStudyPath"`
			ApprovalReceivedIndicatorHold any    `json:"approvalReceivedIndicatorHold"`
			CapcOverride                  string `json:"capcOverride"`
			SpecialApproval               any    `json:"specialApproval"`
			CourseNumber                  string `json:"courseNumber"`
			Message                       any    `json:"message"`
			CohortOverride                any    `json:"cohortOverride"`
			StatusDescription             string `json:"statusDescription"`
			SelectedAction                struct {
				Class                    string `json:"class"`
				CourseRegistrationStatus any    `json:"courseRegistrationStatus"`
				Description              any    `json:"description"`
				RegistrationStatusDate   any    `json:"registrationStatusDate"`
				Remove                   bool   `json:"remove"`
				SubActions               any    `json:"subActions"`
				VoiceType                any    `json:"voiceType"`
			} `json:"selectedAction"`
			CollegeOverride                 string `json:"collegeOverride"`
			OriginalVoiceResponseStatusType string `json:"originalVoiceResponseStatusType"`
			CompletionDate                  string `json:"completionDate"`
			PermitOverrideUpdate            any    `json:"permitOverrideUpdate"`
			VoiceResponseStatusType         string `json:"voiceResponseStatusType"`
			RegistrationStudyPaths          []any  `json:"registrationStudyPaths"`
			ID                              any    `json:"id"`
			Messages                        []struct {
				Message string `json:"message"`
				Type    string `json:"type"`
				Field   any    `json:"field"`
			} `json:"messages"`
		} `json:"update"`
	} `json:"data"`
	Structures      any    `json:"structures"`
	RegisteredHours string `json:"registeredHours"`
	BillingHours    string `json:"billingHours"`
	CeuHours        string `json:"ceuHours"`
	MaxHours        string `json:"maxHours"`
	MinHours        string `json:"minHours"`
	Message         string `json:"message"`
}

type WebhookPayload struct {
	Username string  `json:"username,omitempty"`
	Content  string  `json:"content,omitempty"`
	Embeds   []Embed `json:"embeds,omitempty"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Color       int     `json:"color,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Footer      *Footer `json:"footer,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
}
