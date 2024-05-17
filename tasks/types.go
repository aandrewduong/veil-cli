package tasks

type Terms []struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

var QuarterCodes = map[string]int{
	"Summer": 1,
	"Fall":   2,
	"Winter": 3,
	"Spring": 4,
}

var CampusCodes = map[string]int{
	"Foothill": 1,
	"De Anza":  2,
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
			CampusOverride                      any    `json:"campusOverride"`
			GradingMode                         any    `json:"gradingMode"`
			SelectedCreditHour                  any    `json:"selectedCreditHour"`
			Block                               any    `json:"block"`
			GradeComment                        any    `json:"gradeComment"`
			CreditHourInitial                   any    `json:"creditHourInitial"`
			SequenceNumber                      any    `json:"sequenceNumber"`
			CourseRegistrationStatusDescription any    `json:"courseRegistrationStatusDescription"`
			TuitionWaiverIndicator              any    `json:"tuitionWaiverIndicator"`
			RepeatOverride                      any    `json:"repeatOverride"`
			PreqOverride                        any    `json:"preqOverride"`
			CreditHour                          any    `json:"creditHour"`
			MajorOverride                       any    `json:"majorOverride"`
			CourseReferenceNumber               string `json:"courseReferenceNumber"`
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

type Courses struct {
	Success    bool `json:"success"`
	TotalCount int  `json:"totalCount"`
	Data       []struct {
		ID                      int    `json:"id"`
		Term                    string `json:"term"`
		TermDesc                string `json:"termDesc"`
		CourseReferenceNumber   string `json:"courseReferenceNumber"`
		PartOfTerm              string `json:"partOfTerm"`
		CourseNumber            string `json:"courseNumber"`
		Subject                 string `json:"subject"`
		SubjectDescription      string `json:"subjectDescription"`
		SequenceNumber          string `json:"sequenceNumber"`
		CampusDescription       string `json:"campusDescription"`
		ScheduleTypeDescription string `json:"scheduleTypeDescription"`
		CourseTitle             string `json:"courseTitle"`
		CreditHours             any    `json:"creditHours"`
		MaximumEnrollment       int    `json:"maximumEnrollment"`
		Enrollment              int    `json:"enrollment"`
		SeatsAvailable          int    `json:"seatsAvailable"`
		WaitCapacity            int    `json:"waitCapacity"`
		WaitCount               int    `json:"waitCount"`
		WaitAvailable           int    `json:"waitAvailable"`
		CrossList               any    `json:"crossList"`
		CrossListCapacity       any    `json:"crossListCapacity"`
		CrossListCount          any    `json:"crossListCount"`
		CrossListAvailable      any    `json:"crossListAvailable"`
		CreditHourHigh          any    `json:"creditHourHigh"`
		CreditHourLow           any    `json:"creditHourLow"`
		CreditHourIndicator     any    `json:"creditHourIndicator"`
		OpenSection             bool   `json:"openSection"`
		LinkIdentifier          any    `json:"linkIdentifier"`
		IsSectionLinked         bool   `json:"isSectionLinked"`
		SubjectCourse           string `json:"subjectCourse"`
		Faculty                 []struct {
			BannerID              string `json:"bannerId"`
			Category              any    `json:"category"`
			Class                 string `json:"class"`
			CourseReferenceNumber string `json:"courseReferenceNumber"`
			DisplayName           string `json:"displayName"`
			EmailAddress          any    `json:"emailAddress"`
			PrimaryIndicator      bool   `json:"primaryIndicator"`
			Term                  string `json:"term"`
		} `json:"faculty"`
		MeetingsFaculty []struct {
			Category              string `json:"category"`
			Class                 string `json:"class"`
			CourseReferenceNumber string `json:"courseReferenceNumber"`
			Faculty               []any  `json:"faculty"`
			MeetingTime           struct {
				BeginTime              string  `json:"beginTime"`
				Building               string  `json:"building"`
				BuildingDescription    string  `json:"buildingDescription"`
				Campus                 string  `json:"campus"`
				CampusDescription      string  `json:"campusDescription"`
				Category               string  `json:"category"`
				Class                  string  `json:"class"`
				CourseReferenceNumber  string  `json:"courseReferenceNumber"`
				CreditHourSession      float64 `json:"creditHourSession"`
				EndDate                string  `json:"endDate"`
				EndTime                string  `json:"endTime"`
				Friday                 bool    `json:"friday"`
				HoursWeek              float64 `json:"hoursWeek"`
				MeetingScheduleType    string  `json:"meetingScheduleType"`
				MeetingType            string  `json:"meetingType"`
				MeetingTypeDescription string  `json:"meetingTypeDescription"`
				Monday                 bool    `json:"monday"`
				Room                   string  `json:"room"`
				Saturday               bool    `json:"saturday"`
				StartDate              string  `json:"startDate"`
				Sunday                 bool    `json:"sunday"`
				Term                   string  `json:"term"`
				Thursday               bool    `json:"thursday"`
				Tuesday                bool    `json:"tuesday"`
				Wednesday              bool    `json:"wednesday"`
			} `json:"meetingTime"`
			Term string `json:"term"`
		} `json:"meetingsFaculty"`
		ReservedSeatSummary any `json:"reservedSeatSummary"`
		SectionAttributes   []struct {
			Class                 string `json:"class"`
			Code                  string `json:"code"`
			CourseReferenceNumber string `json:"courseReferenceNumber"`
			Description           string `json:"description"`
			IsZTCAttribute        bool   `json:"isZTCAttribute"`
			TermCode              string `json:"termCode"`
		} `json:"sectionAttributes"`
		InstructionalMethod            string `json:"instructionalMethod"`
		InstructionalMethodDescription string `json:"instructionalMethodDescription"`
	} `json:"data"`
	PageOffset           int    `json:"pageOffset"`
	PageMaxSize          int    `json:"pageMaxSize"`
	SectionsFetchedCount int    `json:"sectionsFetchedCount"`
	PathMode             string `json:"pathMode"`
	SearchResultsConfigs []struct {
		Config   string `json:"config"`
		Display  string `json:"display"`
		Title    string `json:"title"`
		Required bool   `json:"required"`
		Width    string `json:"width"`
	} `json:"searchResultsConfigs"`
	ZtcEncodedImage string `json:"ztcEncodedImage"`
}

type CourseInfo struct {
	TermDesc              string
	CourseReferenceNumber string
	Subject               string
	CourseNumber          string
	SequenceNumber        string
	CourseTitle           string
	DisplayName           string
	BeginTime             string
	EndTime               string
	StartDate             string
	EndDate               string
	MeetingType           string
	Room                  string
	MaximumEnrollment     int
	Enrollment            int
	SeatsAvailable        int
	WaitAvailable         int
}

type UserInfo struct {
	Embedded struct {
		Students []struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			ActiveTerm    string `json:"activeTerm"`
			BridgeRefresh struct {
				Date string `json:"date"`
				Time string `json:"time"`
			} `json:"bridgeRefresh"`
			BridgeChanged struct {
				Date string `json:"date"`
				Time string `json:"time"`
			} `json:"bridgeChanged"`
			Goals []struct {
				School struct {
					Key         string `json:"key"`
					Description string `json:"description"`
				} `json:"school"`
				Degree struct {
					Key         string `json:"key"`
					Description string `json:"description"`
				} `json:"degree"`
				Level struct {
					Key         string `json:"key"`
					Description string `json:"description"`
				} `json:"level"`
				CatalogYear struct {
					Key         string `json:"key"`
					Description string `json:"description"`
				} `json:"catalogYear"`
				Details []struct {
					Code struct {
						Key         string `json:"key"`
						Description string `json:"description"`
					} `json:"code"`
					Value struct {
						Key         string `json:"key"`
						Description string `json:"description"`
					} `json:"value"`
				} `json:"details"`
			} `json:"goals"`
			Custom []struct {
				Code   string `json:"code"`
				Value  string `json:"value"`
				School string `json:"school,omitempty"`
				Degree string `json:"degree,omitempty"`
			} `json:"custom"`
		} `json:"students"`
	} `json:"_embedded"`
	Page struct {
		Size          string `json:"size"`
		TotalElements string `json:"totalElements"`
		TotalPages    string `json:"totalPages"`
		Number        string `json:"number"`
	} `json:"page"`
}

type AuditInfo struct {
	Term        string
	Subject     string
	Number      string
	CourseTitle string
	LetterGrade string
	Credits     string
}

type Audit struct {
	Refresh struct {
		Bridged struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"bridged"`
		Changed struct {
			Date string `json:"date"`
			Time string `json:"time"`
		} `json:"changed"`
	} `json:"refresh"`
	AuditHeader struct {
		AuditID                        string `json:"auditId"`
		StudentID                      string `json:"studentId"`
		AuditType                      string `json:"auditType"`
		StudentName                    string `json:"studentName"`
		StudentEmail                   string `json:"studentEmail"`
		FreezeType                     string `json:"freezeType"`
		FreezeTypeDescription          string `json:"freezeTypeDescription"`
		FreezeDate                     string `json:"freezeDate"`
		FreezeUserName                 string `json:"freezeUserName"`
		AuditDescription               string `json:"auditDescription"`
		DateYear                       string `json:"dateYear"`
		DateMonth                      string `json:"dateMonth"`
		DateDay                        string `json:"dateDay"`
		TimeHour                       string `json:"timeHour"`
		TimeMinute                     string `json:"timeMinute"`
		StudentSystemGpa               string `json:"studentSystemGpa"`
		DegreeworksGpa                 string `json:"degreeworksGpa"`
		PercentComplete                string `json:"percentComplete"`
		Version                        string `json:"version"`
		InProgress                     string `json:"inProgress"`
		WhatIf                         string `json:"whatIf"`
		ResidentApplied                string `json:"residentApplied"`
		ResidentAppliedInProgress      string `json:"residentAppliedInProgress"`
		TransferApplied                string `json:"transferApplied"`
		ExamAppliedCredits             string `json:"examAppliedCredits"`
		ResidentOverTheLimit           string `json:"residentOverTheLimit"`
		ResidentOverTheLimitInProgress string `json:"residentOverTheLimitInProgress"`
		TransferOverTheLimit           string `json:"transferOverTheLimit"`
		ExamOverTheLimit               string `json:"examOverTheLimit"`
	} `json:"auditHeader"`
	BlockArray []struct {
		RequirementID    string `json:"requirementId"`
		RequirementType  string `json:"requirementType"`
		RequirementValue string `json:"requirementValue"`
		Title            string `json:"title"`
		PercentComplete  string `json:"percentComplete"`
		CatalogYearStart string `json:"catalogYearStart"`
		CatalogYearStop  string `json:"catalogYearStop"`
		CatalogYear      string `json:"catalogYear"`
		CatalogYearLit   string `json:"catalogYearLit"`
		Degree           string `json:"degree,omitempty"`
		Gpa              string `json:"gpa"`
		ClassesApplied   string `json:"classesApplied"`
		CreditsApplied   string `json:"creditsApplied"`
		GpaGradePoints   string `json:"gpaGradePoints"`
		GpaCredits       string `json:"gpaCredits"`
		Header           struct {
			QualifierArray []struct {
				NodeID         string   `json:"nodeId"`
				NodeType       string   `json:"nodeType"`
				Satisfied      string   `json:"satisfied,omitempty"`
				Applied        string   `json:"applied,omitempty"`
				ClassesApplied string   `json:"classesApplied,omitempty"`
				CreditsApplied string   `json:"creditsApplied,omitempty"`
				Name           string   `json:"name"`
				Credits        string   `json:"credits,omitempty"`
				Text           string   `json:"text"`
				MinGPA         string   `json:"minGPA,omitempty"`
				SubTextList    []string `json:"subTextList,omitempty"`
				JustAdded      string   `json:"justAdded,omitempty"`
				ECAOverall     struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-Overall,omitempty"`
				ECAErrorArray []struct {
					Text string `json:"text"`
				} `json:"ECA-ErrorArray,omitempty"`
				ECABlocksRequired struct {
					Text string `json:"text"`
				} `json:"ECA-BlocksRequired,omitempty"`
				ECABlocksNotRequired struct {
					Text string `json:"text"`
				} `json:"ECA-BlocksNotRequired,omitempty"`
				ECABlocksRequiredIncluded struct {
					Text string `json:"text"`
				} `json:"ECA-BlocksRequiredIncluded,omitempty"`
				ECABlocksum struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-Blocksum,omitempty"`
				ECAShared struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-Shared,omitempty"`
				ECAECA struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-ECA,omitempty"`
				ECARequiredCreditsApplied struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-RequiredCreditsApplied,omitempty"`
				ECANonrequiredCreditsApplied struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-NonrequiredCreditsApplied,omitempty"`
				ECAOverflow struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"ECA-Overflow,omitempty"`
				CreditsAppliedTowardsDegree struct {
					Credits string `json:"credits"`
					Text    string `json:"text"`
				} `json:"creditsAppliedTowardsDegree,omitempty"`
			} `json:"qualifierArray"`
		} `json:"header,omitempty"`
		RuleArray []struct {
			PercentComplete   string `json:"percentComplete"`
			RuleID            string `json:"ruleId"`
			NodeID            string `json:"nodeId"`
			NodeType          string `json:"nodeType"`
			IndentLevel       string `json:"indentLevel"`
			RuleType          string `json:"ruleType"`
			Label             string `json:"label"`
			BooleanEvaluation string `json:"booleanEvaluation,omitempty"`
			Requirement       struct {
				LeftCondition struct {
					Connector     string `json:"connector"`
					LeftCondition struct {
						RelationalOperator struct {
							Left       string `json:"left"`
							Operator   string `json:"operator"`
							Right      string `json:"right"`
							Evaluation string `json:"evaluation"`
						} `json:"relationalOperator"`
					} `json:"leftCondition"`
				} `json:"leftCondition"`
				IfPart struct {
					RuleArray []struct {
						Label           string `json:"label"`
						LabelTag        string `json:"labelTag"`
						PercentComplete string `json:"percentComplete"`
						RuleID          string `json:"ruleId"`
						NodeID          string `json:"nodeId"`
						NodeType        string `json:"nodeType"`
						IndentLevel     string `json:"indentLevel"`
						RuleType        string `json:"ruleType"`
						IfElsePart      string `json:"ifElsePart"`
						Requirement     struct {
							NumBlocks string `json:"numBlocks"`
							Type      string `json:"type"`
							Value     string `json:"value"`
						} `json:"requirement"`
						Advice struct {
							BlockID string `json:"blockId"`
						} `json:"advice"`
					} `json:"ruleArray"`
				} `json:"ifPart"`
				ElsePart struct {
					RuleArray []struct {
						PercentComplete   string `json:"percentComplete"`
						RuleID            string `json:"ruleId"`
						NodeID            string `json:"nodeId"`
						NodeType          string `json:"nodeType"`
						IndentLevel       string `json:"indentLevel"`
						RuleType          string `json:"ruleType"`
						Label             string `json:"label"`
						IfElsePart        string `json:"ifElsePart"`
						BooleanEvaluation string `json:"booleanEvaluation"`
						Requirement       struct {
							LeftCondition struct {
								Connector     string `json:"connector"`
								LeftCondition struct {
									RelationalOperator struct {
										Left       string `json:"left"`
										Operator   string `json:"operator"`
										Right      string `json:"right"`
										Evaluation string `json:"evaluation"`
									} `json:"relationalOperator"`
								} `json:"leftCondition"`
							} `json:"leftCondition"`
							IfPart struct {
								RuleArray []struct {
									Label           string `json:"label"`
									LabelTag        string `json:"labelTag"`
									PercentComplete string `json:"percentComplete"`
									RuleID          string `json:"ruleId"`
									NodeID          string `json:"nodeId"`
									NodeType        string `json:"nodeType"`
									IndentLevel     string `json:"indentLevel"`
									RuleType        string `json:"ruleType"`
									IfElsePart      string `json:"ifElsePart"`
									Requirement     struct {
										NumBlocks string `json:"numBlocks"`
										Type      string `json:"type"`
										Value     string `json:"value"`
									} `json:"requirement"`
									Advice struct {
										BlockID string `json:"blockId"`
									} `json:"advice"`
								} `json:"ruleArray"`
							} `json:"ifPart"`
							ElsePart struct {
								RuleArray []struct {
									Label           string `json:"label"`
									LabelTag        string `json:"labelTag"`
									PercentComplete string `json:"percentComplete"`
									RuleID          string `json:"ruleId"`
									NodeID          string `json:"nodeId"`
									NodeType        string `json:"nodeType"`
									IndentLevel     string `json:"indentLevel"`
									RuleType        string `json:"ruleType"`
									IfElsePart      string `json:"ifElsePart"`
									Requirement     struct {
										NumberOfGroups string `json:"numberOfGroups"`
										NumberOfRules  string `json:"numberOfRules"`
									} `json:"requirement"`
									RuleArray []struct {
										Label           string `json:"label"`
										LabelTag        string `json:"labelTag"`
										PercentComplete string `json:"percentComplete"`
										RuleID          string `json:"ruleId"`
										NodeID          string `json:"nodeId"`
										NodeType        string `json:"nodeType"`
										IndentLevel     string `json:"indentLevel"`
										RuleType        string `json:"ruleType"`
										Requirement     struct {
											NumBlocks string `json:"numBlocks"`
											Type      string `json:"type"`
											Value     string `json:"value"`
										} `json:"requirement"`
										Advice struct {
											BlockID string `json:"blockId"`
										} `json:"advice,omitempty"`
										LastRuleInGroup string `json:"lastRuleInGroup,omitempty"`
									} `json:"ruleArray"`
								} `json:"ruleArray"`
							} `json:"elsePart"`
						} `json:"requirement"`
					} `json:"ruleArray"`
				} `json:"elsePart"`
			} `json:"requirement,omitempty"`
			LabelTag     string `json:"labelTag,omitempty"`
			Requirement0 struct {
				NumBlocktypes string `json:"numBlocktypes"`
				Type          string `json:"type"`
			} `json:"requirement,omitempty"`
			Advice struct {
				TitleList []string `json:"titleList"`
			} `json:"advice,omitempty"`
		} `json:"ruleArray"`
		Header0 struct {
			QualifierArray []struct {
				NodeID   string `json:"nodeId"`
				NodeType string `json:"nodeType"`
				Applied  string `json:"applied"`
				Name     string `json:"name"`
				MinGrade string `json:"minGrade"`
				Text     string `json:"text"`
			} `json:"qualifierArray"`
			Remark struct {
				TextList []string `json:"textList"`
			} `json:"remark"`
		} `json:"header,omitempty"`
		Major1 string `json:"major1,omitempty"`
	} `json:"blockArray"`
	ClassInformation struct {
		ClassArray []struct {
			Discipline            string `json:"discipline"`
			Number                string `json:"number"`
			Credits               string `json:"credits"`
			LetterGrade           string `json:"letterGrade"`
			ID                    string `json:"id"`
			CourseTitle           string `json:"courseTitle"`
			Term                  string `json:"term"`
			TermLiteral           string `json:"termLiteral"`
			TermLiteralLong       string `json:"termLiteralLong"`
			RecordType            string `json:"recordType"`
			RecordSequence        string `json:"recordSequence"`
			ClassErrorNumber      string `json:"classErrorNumber"`
			Status                string `json:"status"`
			StudentSystemCredits  string `json:"studentSystemCredits"`
			ReasonInsufficient    string `json:"reasonInsufficient"`
			ForceInsufficient     string `json:"forceInsufficient"`
			ForceToFallthrough    string `json:"forceToFallthrough"`
			InProgress            string `json:"inProgress"`
			FailedCountInMajorGpa string `json:"failedCountInMajorGpa"`
			Incomplete            string `json:"incomplete"`
			Passfail              string `json:"passfail"`
			Passed                string `json:"passed"`
			AtCode                string `json:"atCode"`
			GradePoints           string `json:"gradePoints"`
			NumericGrade          string `json:"numericGrade"`
			GpaGradePoints        string `json:"gpaGradePoints"`
			GpaCredits            string `json:"gpaCredits"`
			Section               string `json:"section"`
			School                string `json:"school"`
			GradeType             string `json:"gradeType"`
			RepeatDiscipline      string `json:"repeatDiscipline"`
			RepeatNumber          string `json:"repeatNumber"`
			RepeatPolicy          string `json:"repeatPolicy"`
			Transfer              string `json:"transfer"`
			TransferCode          string `json:"transferCode"`
			TransferType          string `json:"transferType"`
			WithData              string `json:"withData"`
			EquivalenceExists     string `json:"equivalenceExists"`
			AttributeArray        []struct {
				Code  string `json:"code"`
				Value string `json:"value"`
			} `json:"attributeArray"`
			LocArray []struct {
				RequirementID     string `json:"requirementId"`
				NodeLocation      string `json:"nodeLocation"`
				Level             string `json:"level"`
				Rank              string `json:"rank"`
				ReasonRemoved     string `json:"reasonRemoved,omitempty"`
				ReasonComparisons string `json:"reasonComparisons,omitempty"`
				ReasonRemovedList struct {
					ReasonList         string `json:"reasonList"`
					ReasonRemovedArray []struct {
						Code string `json:"code"`
					} `json:"reasonRemovedArray"`
				} `json:"reasonRemovedList,omitempty"`
			} `json:"locArray,omitempty"`
		} `json:"classArray"`
	} `json:"classInformation"`
	FallThrough struct {
		Classes    string `json:"classes"`
		Credits    string `json:"credits"`
		Noncourses string `json:"noncourses"`
		ClassArray []struct {
			Discipline  string `json:"discipline"`
			Number      string `json:"number"`
			Credits     string `json:"credits"`
			LetterGrade string `json:"letterGrade"`
			ID          string `json:"id"`
			Code        string `json:"code"`
		} `json:"classArray"`
	} `json:"fallThrough"`
	OverTheLimit struct {
		Classes    string `json:"classes"`
		Credits    string `json:"credits"`
		ClassArray []struct {
			Discipline  string `json:"discipline"`
			Number      string `json:"number"`
			Credits     string `json:"credits"`
			LetterGrade string `json:"letterGrade"`
			ID          string `json:"id"`
			Reason      string `json:"reason"`
			Reason2     string `json:"reason2"`
		} `json:"classArray"`
	} `json:"overTheLimit"`
	Insufficient struct {
		Classes    string `json:"classes"`
		Credits    string `json:"credits"`
		ClassArray []struct {
			Discipline         string `json:"discipline"`
			Number             string `json:"number"`
			Credits            string `json:"credits"`
			LetterGrade        string `json:"letterGrade"`
			ID                 string `json:"id"`
			ReasonInsufficient string `json:"reasonInsufficient"`
			ForceInsufficient  string `json:"forceInsufficient"`
			GpaGradePoints     string `json:"gpaGradePoints"`
			GpaCredits         string `json:"gpaCredits"`
		} `json:"classArray"`
	} `json:"insufficient"`
	InProgress struct {
		Classes    string `json:"classes"`
		Credits    string `json:"credits"`
		ClassArray []struct {
			Discipline  string `json:"discipline"`
			Number      string `json:"number"`
			Credits     string `json:"credits"`
			LetterGrade string `json:"letterGrade"`
			ID          string `json:"id"`
		} `json:"classArray"`
	} `json:"inProgress"`
	FitList struct {
		ClassArray []struct {
			Discipline  string `json:"discipline"`
			Number      string `json:"number"`
			Credits     string `json:"credits"`
			LetterGrade string `json:"letterGrade"`
			ID          string `json:"id"`
			LocArray    []struct {
				RequirementID string `json:"requirementId"`
				NodeLocation  string `json:"nodeLocation"`
				Level         string `json:"level"`
				Rank          string `json:"rank"`
			} `json:"locArray"`
		} `json:"classArray"`
	} `json:"fitList"`
	SplitCredits struct {
		Classes string `json:"classes"`
		Credits string `json:"credits"`
	} `json:"splitCredits"`
	DegreeInformation struct {
		DegreeDataArray []struct {
			Degree                                      string `json:"degree"`
			School                                      string `json:"school"`
			CatalogYear                                 string `json:"catalogYear"`
			ActiveTerm                                  string `json:"activeTerm"`
			StudentLevel                                string `json:"studentLevel"`
			DegreeTerm                                  string `json:"degreeTerm"`
			StudentSystemCumulativeGradePointsAttempted string `json:"studentSystemCumulativeGradePointsAttempted"`
			StudentSystemCumulativeGradePointsEarned    string `json:"studentSystemCumulativeGradePointsEarned"`
			StudentSystemCumulativeGpa                  string `json:"studentSystemCumulativeGpa"`
			StudentSystemCumulativeTotalCreditsEarned   string `json:"studentSystemCumulativeTotalCreditsEarned"`
			StudentSystemCumulativeCreditsEarned        string `json:"studentSystemCumulativeCreditsEarned"`
			DegreeSource                                string `json:"degreeSource"`
			DegreeLiteral                               string `json:"degreeLiteral"`
			SchoolLiteral                               string `json:"schoolLiteral"`
			StudentLevelLiteral                         string `json:"studentLevelLiteral"`
			CatalogYearLit                              string `json:"catalogYearLit"`
			ActiveTermLiteral                           string `json:"activeTermLiteral"`
		} `json:"degreeDataArray"`
		GoalArray []struct {
			Value        string `json:"value"`
			CatalogYear  string `json:"catalogYear"`
			Code         string `json:"code"`
			ValueLiteral string `json:"valueLiteral"`
		} `json:"goalArray"`
	} `json:"degreeInformation"`
	ExceptionList struct {
	} `json:"exceptionList"`
	Notes struct {
	} `json:"notes"`
	Flags struct {
		Cfg020DAP14    string `json:"cfg020DAP14"`
		Cfg020TIEBREAK string `json:"cfg020TIEBREAK"`
	} `json:"flags"`
}
