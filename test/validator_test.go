package libvalidator_test

import (
	"testing"

	libparser "github.com/tomefile/lib-parser"
	libvalidator "github.com/tomefile/lib-validator"
	"gotest.tools/assert"
)

func TestValidatorExec(test *testing.T) {
	_, err := libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_EXEC,
		Literal: "what_are_the_chances_this_binary_exists_and_the_test_doesnt_pass",
	})
	assert.ErrorContains(test, err, `undefined binary`)
}

func TestValidatorDirective(test *testing.T) {
	_, err := libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_DIRECTIVE,
		Literal: libvalidator.DIR_EXPORT,
		Args:    []any{"hello_world", "123"},
	})
	assert.NilError(test, err)

	_, err = libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_DIRECTIVE,
		Literal: libvalidator.DIR_EXPORT,
	})
	assert.ErrorContains(test, err, `mismatched arguments`)

	_, err = libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_DIRECTIVE,
		Literal: libvalidator.DIR_EXPORT + "DOES_NOT_EXIST",
	})
	assert.ErrorContains(test, err, `undefined directive`)
}

func TestValidatorMacro(test *testing.T) {
	macro := "example_macro"
	libvalidator.ValidMacros = append(libvalidator.ValidMacros, macro)

	_, err := libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_MACRO,
		Literal: macro,
	})
	assert.NilError(test, err)

	_, err = libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_MACRO,
		Literal: "does_not_exist",
	})
	assert.ErrorContains(test, err, `undefined macro`)
}

func TestValidatorArgs(test *testing.T) {
	_, err := libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_DIRECTIVE,
		Literal: libvalidator.DIR_UNSET,
		Args: []any{
			"Hello ${world}",
		},
	})
	assert.NilError(test, err)

	_, err = libvalidator.Validate(&libparser.Node{
		Type:    libparser.NODE_DIRECTIVE,
		Literal: libvalidator.DIR_UNSET,
		Args: []any{
			"Hello ${world",
		},
	})
	assert.ErrorContains(test, err, "unexpected end of file")
}
