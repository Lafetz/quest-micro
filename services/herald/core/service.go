package herald

type HeraldService struct {
	Sender EmailSender
}

func (es *HeraldService) SendQuestAssignmentEmail(email Email) error {

	return es.Sender.SendEmail(email)
}
func NewHeraldService(sender EmailSender) *HeraldService {
	return &HeraldService{
		Sender: sender,
	}
}
