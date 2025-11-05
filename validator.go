package libvalidator

import (
	"fmt"
	"slices"
	"strings"

	libparser "github.com/tomefile/lib-parser"
)

func Validate(node *libparser.Node) (*libparser.Node, error) {
	switch node.Type {

	case libparser.NODE_NULL,
		libparser.NODE_COMMENT,
		libparser.NODE_ERROR_READ,
		libparser.NODE_ERROR_SYNTAX:
		return node, nil

	case libparser.NODE_DIRECTIVE:
		required_args, is_found := ValidDirectives[node.Literal]
		if !is_found {
			return node, UndefinedError{
				Scope:   "directive",
				Name:    node.Literal,
				Allowed: mapKeys(ValidDirectives),
			}
		}
		if len(required_args) == 0 {
			if len(node.Args) != 0 {
				return node, MismatchedError{
					Token:       node.Literal,
					ExpectedArg: "<empty>",
					ActualArg:   fmt.Sprint(node.Args),
				}
			}
			break
		}
		for i, arg := range required_args {
			if isOptionalArg(arg) {
				break
			}
			if i >= len(node.Args) {
				return node, MismatchedError{
					Token:       node.Literal,
					ExpectedArg: arg,
					ActualArg:   "<empty>",
				}
			}
		}

	case libparser.NODE_EXEC:
		path, is_found := FindExec(node.Literal)
		if !is_found {
			return node, UndefinedError{
				Scope:   "binary",
				Name:    node.Literal,
				Allowed: nil,
			}
		}
		node.Literal = path

	case libparser.NODE_MACRO:
		if !slices.Contains(ValidMacros, node.Literal) {
			return node, UndefinedError{
				Scope:   "macro",
				Name:    node.Literal,
				Allowed: ValidMacros,
			}
		}

	default:
		return node, InternalError{
			Format: "unexpected libparser.NodeType: %#v",
			Args:   []any{node.Type},
		}
	}

	for _, arg := range node.Args {
		switch arg := arg.(type) {

		case *libparser.Node:
			ctx_node, ctx_err := Validate(arg)
			if ctx_err != nil {
				return ctx_node, ctx_err
			}

		case string:
			formatter := libparser.NewFormatter(strings.NewReader(arg))
			_, err := formatter.Parse()
			if err != nil {
				return node, err
			}
		}
	}

	for _, child := range node.Children {
		ctx_node, ctx_err := Validate(child)
		if ctx_err != nil {
			return ctx_node, ctx_err
		}
	}

	return node, nil
}

func mapKeys[T any](in map[string]T) []string {
	out := make([]string, 0, len(in))

	for key := range in {
		out = append(out, key)
	}

	return out
}

func isOptionalArg(arg string) bool {
	return strings.HasSuffix(arg, "?")
}
