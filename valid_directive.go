package libvalidator

type DirArgs []string

const (
	DIR_INCLUDE     = "include"
	DIR_SECTION     = "section"
	DIR_DEFINE      = "define"
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

const (
	ARG_GENERIC       = "arg"
	ARG_PATH          = "path"
	ARG_DESCRIPTION   = "description"
	ARG_MACRO_NAME    = "macro_name"
	ARG_TOME_NAME     = "tome_name"
	ARG_VARIABLE_NAME = "variable_name"
	ARG_ENV_NAME      = "env_name"
	ARG_DEFAULT_VALUE = "default_value"
	ARG_VALUE         = "value"
	ARG_SIGN          = "sign"
	ARG_COMMAND       = "command"
)

func ArgVariadic(arg string) string {
	return ArgOptional("..." + arg)
}

func ArgOptional(arg string) string {
	return arg + "?"
}

var ValidDirectives = map[string]DirArgs{
	DIR_INCLUDE:     []string{ARG_PATH},
	DIR_SECTION:     []string{ArgOptional(ARG_DESCRIPTION)},
	DIR_DEFINE:      []string{ARG_MACRO_NAME},
	DIR_TOME:        []string{ARG_TOME_NAME, ArgOptional(ARG_DESCRIPTION)},
	DIR_REQUIRE:     []string{ARG_VARIABLE_NAME, ArgOptional(ARG_DEFAULT_VALUE)},
	DIR_DEPEND_CALL: []string{ARG_TOME_NAME, ArgVariadic(ARG_PATH)},
	DIR_CALL:        []string{ARG_TOME_NAME},
	DIR_SET:         []string{ARG_VARIABLE_NAME, ARG_VALUE},
	DIR_UNSET:       []string{ARG_VARIABLE_NAME},
	DIR_IF:          []string{ARG_VARIABLE_NAME, ArgOptional(ARG_SIGN), ArgOptional(ARG_VALUE)},
	DIR_ELIF:        []string{ARG_VARIABLE_NAME, ArgOptional(ARG_SIGN), ArgOptional(ARG_VALUE)},
	DIR_ELSE:        []string{},
	DIR_TRY:         []string{ARG_COMMAND, ArgVariadic(ARG_GENERIC)},
	DIR_EXPORT:      []string{ARG_ENV_NAME, ARG_VALUE},
}
