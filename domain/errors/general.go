package domain_errors

import "errors"

var (
	ErrNotFound = DomainError{
		err: errors.New("not found"),
	}
)

var (
	notFoundResourceError = map[DomainError]bool{
		ErrNotFound:               true,
		ErrUserNotFound:           true,
		ErrProfileNotFound:        true,
		ErrPremiumPackageNotFound: true,
	}

	validationError = map[DomainError]bool{
		ErrUserAlreadyExists: true,
		ErrCannotDoSwipe:     true,
	}
)

type DomainError struct {
	err error
}

var _ error = &DomainError{}

func (err DomainError) IsResourceNotFound() bool {
	return notFoundResourceError[err]
}

func (err DomainError) IsValidationError() bool {
	return validationError[err]
}

func (err DomainError) Error() string {
	return err.err.Error()
}

// users
var (
	ErrUserNotFound = DomainError{
		err: errors.New("user not found"),
	}
	ErrUserAlreadyExists = DomainError{
		err: errors.New("user already exists"),
	}
)

// profiles
var (
	ErrProfileNotFound = DomainError{
		err: errors.New("profile not found"),
	}
)

// swipes
var (
	ErrCannotDoSwipe = DomainError{
		err: errors.New("cannot do swipe"),
	}
)

// premium package
var (
	ErrPremiumPackageNotFound = DomainError{
		err: errors.New("premium package not found"),
	}
)
