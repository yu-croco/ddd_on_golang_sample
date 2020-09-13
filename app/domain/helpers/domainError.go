package helpers

import "errors"

type DomainError struct {
	errors []error `json:errors`
}

func NewDomainError(message string) DomainError {
	var errorResult []error

	err := errors.New(message)
	errorResult = append(errorResult, err)

	return DomainError{errors: errorResult}
}

func (domainErr *DomainError) HasErrors() bool {
	return len(domainErr.errors) >= 1
}
