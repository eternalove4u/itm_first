package port

import (
	"fmt"
	"regexp"
	"strings"
)

type Service struct {
	prod producer
	pres presenter
}

func NewService(prod producer, pres presenter) (*Service, error) {
	if prod == nil || pres == nil {
		return nil, fmt.Errorf("prod && presenter is nil")
	}
	return &Service{prod: prod, pres: pres}, nil
}

func (s *Service) Run() error {
	messages, err := s.prod.Produce()
	if err != nil {
		return fmt.Errorf("prod produce: %w", err)
	}
	for i, m := range messages {
		messages[i] = s.MaskLinksInMessage(m)
	}
	return s.pres.Present(messages)
}

var httpRegexp = regexp.MustCompile(`(h|H)(t|T)(t|T)(p|P)(s|S)?:\/\/(\w*\/*\.*)*`)

// MaskLinksInMessage checks only "http://"-case
func (s *Service) MaskLinksInMessage(message string) string {
	res := []byte(message)
	finded := httpRegexp.FindAllString(message, -1)
	for _, f := range finded {
		// бага
		pos := strings.Index(message, f)
		for idx := pos + len(f) - 1; idx >= pos+7; idx = idx - 1 {
			res[idx] = '*'
		}
	}
	return string(res)
}
