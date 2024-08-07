package service

type Service struct {
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

func (s *Service) Run() error {
	messages, err := s.prod.Produce()
	if err != nil {
		return err
	}
	for i := range messages {
		messages[i] = s.MaskLinksInMessage(messages[i])
	}
	return s.pres.Present(messages)
}

// MaskLinksInMessage checks only "http://"-case
func (s *Service) MaskLinksInMessage(message string) string {
	const linkPrefix = "http://"
	maskedMessage := []byte(message)
	for i := 0; i < len(message)-len(linkPrefix); i++ {
		if message[i:i+len(linkPrefix)] == linkPrefix {
			i += len(linkPrefix)
			for i < len(message) {
				if message[i] == ' ' {
					break
				}
				maskedMessage[i] = '*'
				i++
			}
		}
	}
	return string(maskedMessage)
}
