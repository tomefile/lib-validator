package libvalidator

import (
	"bufio"
	"fmt"
	"strings"

	libparser "github.com/tomefile/lib-parser"
)

func Validate(node libparser.Node) (libparser.Node, *libparser.DetailedError) {
	switch node := node.(type) {

	case *libparser.StringNode:
		// TODO: Store result into the node, so you don't have to compute it twice.
		formatter := libparser.NewStringFormatter(bufio.NewReader(strings.NewReader(node.Contents)))
		_, err := formatter.Format()
		if err != nil {
			return node, err
		}

	case *libparser.ExecNode:
		path, is_found := FindExec(node.Binary)
		if !is_found {
			return node, &libparser.DetailedError{
				Name:    "Validation Error",
				Details: fmt.Sprintf("could not find %q in $PATH", node.Binary),
				Trace:   []libparser.TraceItem{},
				Context: fmt.Sprintf("$ %s", node.Binary),
			}
		}
		node.Binary = path
		if err := ValidateChildren(libparser.NodeChildren(node.NodeArgs)); err != nil {
			return node, err
		}

	case *libparser.CallNode:
		_, is_found := ValidMacros[node.Macro]
		if !is_found {
			return node, &libparser.DetailedError{
				Name:    "Validation Error",
				Details: fmt.Sprintf("macro %q is not defined", node.Macro),
				Trace:   []libparser.TraceItem{},
				Context: fmt.Sprintf("$ %s!", node.Macro),
			}
		}
		if err := ValidateChildren(libparser.NodeChildren(node.NodeArgs)); err != nil {
			return node, err
		}

	case *libparser.DirectiveNode:
		required_args, is_found := ValidDirectives[node.Name]
		if !is_found {
			return node, &libparser.DetailedError{
				Name:    "Validation Error",
				Details: fmt.Sprintf("directive %q is not defined", node.Name),
				Trace:   []libparser.TraceItem{},
				Context: fmt.Sprintf(":%s", node.Name),
			}
		}
		if len(required_args) == 0 {
			if len(node.NodeArgs) != 0 {
				return node, &libparser.DetailedError{
					Name:    "Validation Error",
					Details: fmt.Sprintf("expected no arguments but got %q.", node.NodeArgs[0]),
					Trace:   []libparser.TraceItem{},
					Context: fmt.Sprintf(":%s", node.Name),
				}
			}
			break
		}
		for i, arg := range required_args {
			if isOptionalArg(arg) {
				break
			}
			if i >= len(node.NodeArgs) {
				return node, &libparser.DetailedError{
					Name:    "Validation Error",
					Details: fmt.Sprintf("expected an argument %q.", arg),
					Trace:   []libparser.TraceItem{},
					Context: fmt.Sprintf(":%s", node.Name),
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

func ValidateChildren(children libparser.NodeChildren) *libparser.DetailedError {
	for _, arg := range children {
		_, ctx_err := Validate(arg)
		if ctx_err != nil {
			return ctx_err
		}
	}

	return nil
}

func isOptionalArg(arg string) bool {
	return strings.HasSuffix(arg, "?")
}
