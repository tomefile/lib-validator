package libvalidator

import (
	"fmt"
	"strings"
)

type InternalError struct {
	Format string
	Args   []any
}

func (err InternalError) Error() string {
	return fmt.Sprintf(err.Format, err.Args...)
}

type UndefinedError struct {
	Scope   string
	Name    string
	Allowed []string
}

func (err UndefinedError) Error() string {
	if err.Allowed == nil {
		return fmt.Sprintf("undefined %s: %q", err.Scope, err.Name)
	}

	return fmt.Sprintf(
		"undefined %s: %q\nDefined values are: %s",
		err.Scope,
		err.Name,
		strings.Join(err.Allowed, ", "),
	)
}

type MismatchedError struct {
	Token       string
	ExpectedArg string
	ActualArg   string
}

func (err MismatchedError) Error() string {
	return fmt.Sprintf(
		"mismatched arguments of %q: expected %q got %q",
		err.Token,
		err.ExpectedArg,
		err.ActualArg,
	)
}
