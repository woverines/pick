package safe

import (
	"github.com/bndw/pick/errors"
)

func (s *Safe) Add(name, username, password string) (*Account, error) {
	if _, exists := s.Accounts[name]; exists {
		account := s.Accounts[name]
		return &account, &errors.AccountExists{}
	}

	account := NewAccount(name, username, password)
	s.Accounts[name] = account

	if err := s.save(); err != nil {
		return nil, err
	}

	return &account, nil
}
