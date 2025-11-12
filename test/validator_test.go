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
	assert.Assert(test, err != nil)
	assert.ErrorContains(test, err, `$PATH`)

	_, err = libvalidator.Validate(&libparser.ExecNode{
		Binary: "/bin/go",
	})
	assert.Assert(test, err == nil)
}

func TestValidatorDirective(test *testing.T) {
	node, err := libvalidator.Validate(&libparser.DirectiveNode{
		Name: libvalidator.DIR_EXPORT,
		NodeArgs: []libparser.Node{
			&libparser.StringNode{Contents: "hello_world"},
			&libparser.StringNode{Contents: "123"},
		},
	})
	assert.Assert(test, err == nil)
	assert.DeepEqual(test, node, &libvalidator.ValidDirectiveNode{
		Name: libvalidator.DIR_EXPORT,
		NodeArgs: []libparser.Node{
			&libvalidator.ValidStringNode{
				Original: "hello_world",
				Segments: []libparser.Segment{&libparser.LiteralNode{Contents: "hello_world"}},
			},
			&libvalidator.ValidStringNode{
				Original: "123",
				Segments: []libparser.Segment{&libparser.LiteralNode{Contents: "123"}},
			},
		},
		Arguments: map[string]*libvalidator.ValidStringNode{
			"env_name": {
				Original: "hello_world",
				Segments: []libparser.Segment{&libparser.LiteralNode{Contents: "hello_world"}},
			},
			"value": {
				Original: "123",
				Segments: []libparser.Segment{&libparser.LiteralNode{Contents: "123"}},
			},
		},
	})

	_, err = libvalidator.Validate(&libparser.DirectiveNode{
		Name:     libvalidator.DIR_EXPORT,
		NodeArgs: []libparser.Node{},
	})
	assert.Assert(test, err != nil)
	assert.ErrorContains(test, err, `expected`)

	_, err = libvalidator.Validate(&libparser.DirectiveNode{
		Name:     libvalidator.DIR_EXPORT + "DOES_NOT_EXIST",
		NodeArgs: []libparser.Node{},
	})
	assert.Assert(test, err != nil)
	assert.ErrorContains(test, err, `define`)
}

func TestValidatorArgs(test *testing.T) {
	node := &libparser.DirectiveNode{
		Name: libvalidator.DIR_UNSET,
		NodeArgs: []libparser.Node{
			&libparser.StringNode{Contents: "Hello ${world?}"},
		},
	}
	_, err := libvalidator.Validate(node)
	assert.Assert(test, err == nil)
	assert.DeepEqual(test, node.NodeArgs, libparser.NodeArgs{
		&libvalidator.ValidStringNode{
			Original: "Hello ${world?}",
			Segments: []libparser.Segment{
				&libparser.LiteralNode{Contents: "Hello "},
				&libparser.VariableSegment{
					Name:       "world",
					Modifier:   nil,
					IsOptional: true,
				},
			},
		},
	})

	_, err = libvalidator.Validate(&libparser.DirectiveNode{
		Name: libvalidator.DIR_UNSET,
		NodeArgs: []libparser.Node{
			&libparser.StringNode{Contents: "Hello ${world"},
		},
	})
	assert.Assert(test, err != nil)
	assert.ErrorContains(test, err, "unexpected end of file")
}
