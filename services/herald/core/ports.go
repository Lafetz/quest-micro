package herald

type EmailSender interface {
	SendEmail(Email) error
}
type HeraldServiceApi interface {
	SendQuestAssignmentEmail(email Email) error
}
