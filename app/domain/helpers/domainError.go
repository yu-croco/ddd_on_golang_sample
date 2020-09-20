package helpers

import (
	"errors"
)

type DomainError struct {
	Errors []string `json:"errors"`
}

func NewDomainError(message string) DomainError {
	var errorResult []string

	err := errors.New(message)
	errorResult = append(errorResult, err.Error())

	return DomainError{Errors: errorResult}
}

func (domainErr *DomainError) HasErrors() bool {
	if domainErr == nil {
		return false
	}
	return len(domainErr.Errors) >= 1
}
