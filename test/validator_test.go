package libvalidator_test

import (
	"testing"

	libparser "github.com/tomefile/lib-parser"
	libvalidator "github.com/tomefile/lib-validator"
	"gotest.tools/assert"
)

func TestValidatorExec(test *testing.T) {
	_, err := libvalidator.Validate(&libparser.ExecNode{
		Binary: "what_are_the_chances_this_binary_exists_and_the_test_doesnt_pass",
	})
	assert.ErrorContains(test, err, `$PATH`)
}

func TestValidatorDirective(test *testing.T) {
	_, err := libvalidator.Validate(&libparser.DirectiveNode{
		Name: libvalidator.DIR_EXPORT,
		NodeArgs: []libparser.Node{
			&libparser.StringNode{Contents: "hello_world"},
			&libparser.StringNode{Contents: "123"},
		},
	})
	assert.Assert(test, err == nil)

	_, err = libvalidator.Validate(&libparser.DirectiveNode{
		Name:     libvalidator.DIR_EXPORT,
		NodeArgs: []libparser.Node{},
	})
	assert.ErrorContains(test, err, `expected`)

	_, err = libvalidator.Validate(&libparser.DirectiveNode{
		Name:     libvalidator.DIR_EXPORT + "DOES_NOT_EXIST",
		NodeArgs: []libparser.Node{},
	})
	assert.ErrorContains(test, err, `define`)
}

func TestValidatorMacro(test *testing.T) {
	macro := "example_macro"
	libvalidator.ValidMacros[macro] = nil

	_, err := libvalidator.Validate(&libparser.CallNode{
		Macro: macro,
	})
	assert.Assert(test, err == nil)

	_, err = libvalidator.Validate(&libparser.CallNode{
		Macro: "does_not_exist",
	})
	assert.ErrorContains(test, err, `define`)
}

func TestValidatorArgs(test *testing.T) {
	_, err := libvalidator.Validate(&libparser.DirectiveNode{
		Name: libvalidator.DIR_UNSET,
		NodeArgs: []libparser.Node{
			&libparser.StringNode{Contents: "Hello ${world}"},
		},
	})
	assert.Assert(test, err == nil)

	_, err = libvalidator.Validate(&libparser.DirectiveNode{
		Name: libvalidator.DIR_UNSET,
		NodeArgs: []libparser.Node{
			&libparser.StringNode{Contents: "Hello ${world"},
		},
	})
	assert.ErrorContains(test, err, "unexpected end of file")
}
