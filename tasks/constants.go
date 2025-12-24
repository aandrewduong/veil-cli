package tasks

const (
	BaseRegURL        = "https://reg.oci.fhda.edu"
	BaseEISURL        = "https://eis-prod.ec.fhda.edu"
	BaseSSOURL        = "https://ssoshib.fhda.edu"
	BaseDWURL         = "https://dw-prod.ec.fhda.edu"
	BaseSSBManagerURL = "https://ssb-prod.ec.fhda.edu"
)

const (
	PathAuthAjax            = "/StudentRegistrationSsb/login/authAjax"
	PathRegisterPostSignIn  = "/StudentRegistrationSsb/ssb/registration/registerPostSignIn"
	PathSamlSSO             = "/samlsso"
	PathSamlSSOSp           = "/StudentRegistrationSsb/saml/SSO/alias/registrationssb-prod-sp"
	PathSectionDetails      = "/StudentRegistrationSsb/ssb/classRegistration/getSectionDetailsFromCRN"
	PathTermSearch          = "/StudentRegistrationSsb/ssb/term/search"
	PathClassRegistration   = "/StudentRegistrationSsb/ssb/classRegistration/classRegistration"
	PathAddRegistrationItem = "/StudentRegistrationSsb/ssb/classRegistration/addRegistrationItem"
	PathSubmitRegistration  = "/StudentRegistrationSsb/ssb/classRegistration/submitRegistration/batch"
	PathGetTerms            = "/StudentRegistrationSsb/ssb/classSearch/getTerms"
	PathSearchResults       = "/StudentRegistrationSsb/ssb/searchResults/searchResults"
	PathEnrollmentInfo      = "/StudentRegistrationSsb/ssb/searchResults/getEnrollmentInfo"
	PathDWStudentMyself     = "/responsiveDashboard/api/students/myself"
	PathDWAudit             = "/responsiveDashboard/api/audit"
)
