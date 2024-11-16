package service

type MailServiceImpl interface {
	Send(email string) error
}

type mailService struct{}

func NewMailService() *mailService {
	return &mailService{}
}

func (s *mailService) Send(email string) error {
	return nil
}
