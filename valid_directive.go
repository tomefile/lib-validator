package libvalidator

type DirArgs []string

const (
	DIR_INCLUDE     = "include"
	DIR_SECTION     = "section"
	DIR_TOME        = "tome"
	DIR_REQUIRE     = "require"
	DIR_DEPEND_CALL = "dcall"
	DIR_CALL        = "call"
	DIR_SET         = "set"
	DIR_UNSET       = "unset"
	DIR_IF          = "if"
	DIR_ELIF        = "elif"
	DIR_ELSE        = "else"
	DIR_TRY         = "try"
	DIR_EXPORT      = "export"
)

var ValidDirectives = map[string]DirArgs{
	DIR_INCLUDE:     []string{"path"},
	DIR_SECTION:     []string{"description?"},
	DIR_TOME:        []string{"tome_name", "description?"},
	DIR_REQUIRE:     []string{"variable_name", "default_value?"},
	DIR_DEPEND_CALL: []string{"tome_name", "...paths?"},
	DIR_CALL:        []string{"tome_name"},
	DIR_SET:         []string{"variable_name", "value"},
	DIR_UNSET:       []string{"variable_name"},
	DIR_IF:          []string{"variable_name", "sign?", "value?"},
	DIR_ELIF:        []string{"variable_name", "sign?", "value?"},
	DIR_ELSE:        []string{},
	DIR_TRY:         []string{"command", "...args?"},
	DIR_EXPORT:      []string{"env_name", "value"},
}
