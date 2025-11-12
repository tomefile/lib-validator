package libvalidator

import (
	"fmt"
	"strings"

	liberrors "github.com/tomefile/lib-errors"
	libparser "github.com/tomefile/lib-parser"
)

func Validate(node libparser.Node) (libparser.Node, *liberrors.DetailedError) {
	switch node := node.(type) {

	case *libparser.StringNode:
		formatter := libparser.NewStringFormatter(node.Contents)
		segments, derr := formatter.Format()
		if derr != nil {
			return node, derr
		}
		return &ValidStringNode{
			Original: node.Contents,
			Segments: segments,
		}, nil

	case *libparser.ExecNode:
		path, is_found := FindExec(node.Binary)
		if !is_found {
			return node, &liberrors.DetailedError{
				Name:    liberrors.ERROR_VALIDATION,
				Details: fmt.Sprintf("could not find %q in $PATH", node.Binary),
				Trace:   []liberrors.TraceItem{},
				Context: liberrors.Context{},
			}
		}
		node.Binary = path
		if err := ValidateChildren(libparser.NodeChildren(node.NodeArgs)); err != nil {
			return node, err
		}

	case *libparser.CallNode:
		_, is_found := ValidMacros[node.Macro]
		if !is_found {
			return node, &liberrors.DetailedError{
				Name:    liberrors.ERROR_VALIDATION,
				Details: fmt.Sprintf("macro %q is not defined", node.Macro),
				Trace:   []liberrors.TraceItem{},
				Context: liberrors.Context{},
			}
		}
		if err := ValidateChildren(libparser.NodeChildren(node.NodeArgs)); err != nil {
			return node, err
		}

	case *libparser.DirectiveNode:
		required_args, is_found := ValidDirectives[node.Name]
		if !is_found {
			return node, &liberrors.DetailedError{
				Name:    liberrors.ERROR_VALIDATION,
				Details: fmt.Sprintf("directive %q is not defined", node.Name),
				Trace:   []liberrors.TraceItem{},
				Context: liberrors.Context{},
			}
		}
		if len(required_args) == 0 {
			if len(node.NodeArgs) != 0 {
				return node, &liberrors.DetailedError{
					Name:    liberrors.ERROR_VALIDATION,
					Details: fmt.Sprintf("expected no arguments but got %q.", node.NodeArgs[0]),
					Trace:   []liberrors.TraceItem{},
					Context: liberrors.Context{},
				}
			}
			break
		}
		for i, arg := range required_args {
			if isOptionalArg(arg) {
				break
			}
			if i >= len(node.NodeArgs) {
				return node, &liberrors.DetailedError{
					Name:    liberrors.ERROR_VALIDATION,
					Details: fmt.Sprintf("expected an argument %q.", arg),
					Trace:   []liberrors.TraceItem{},
					Context: liberrors.Context{},
				}
			}
		}
		if err := ValidateChildren(libparser.NodeChildren(node.NodeArgs)); err != nil {
			return node, err
		}
		if err := ValidateChildren(node.NodeChildren); err != nil {
			return node, err
		}

	}

	return node, nil
}

func ValidateChildren(children libparser.NodeChildren) *liberrors.DetailedError {
	for i, arg := range children {
		node, ctx_err := Validate(arg)
		if ctx_err != nil {
			return ctx_err
		}
		children[i] = node
	}

	return nil
}

func isOptionalArg(arg string) bool {
	return strings.HasSuffix(arg, "?")
}
