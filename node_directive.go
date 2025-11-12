package libvalidator

import (
	"fmt"

	libparser "github.com/tomefile/lib-parser"
)

// Validated version of [libparser.DirectiveNode]
type ValidDirectiveNode struct {
	Name string

	libparser.NodeArgs
	libparser.NodeChildren

	// Mapped arguments
	Arguments map[string]*ValidStringNode
}

func (node *ValidDirectiveNode) Node() string {
	return fmt.Sprintf("[directive %q]", node.Name)
}
