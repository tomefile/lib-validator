package libvalidator

type Modifier string

const (
	MOD_TO_LOWER           Modifier = "to_lower"
	MOD_TO_UPPER           Modifier = "to_upper"
	MOD_TO_SNAKE           Modifier = "to_snake"
	MOD_TO_SCREAMING_SNAKE Modifier = "to_screaming_snake"
	MOD_TO_KEBAB           Modifier = "to_kebab"
	MOD_TO_SCREAMING_KEBAB Modifier = "to_screaming_kebab"
	MOD_TO_CAMEL           Modifier = "to_camel"
	MOD_TO_PASCAL          Modifier = "to_pascal"
	MOD_TO_DELIMITED       Modifier = "to_delimited"
	MOD_TO_INVERTED        Modifier = "to_inverted"
	MOD_TRIM               Modifier = "trim"
	MOD_TRIM_LEFT          Modifier = "trim_left"
	MOD_TRIM_RIGHT         Modifier = "trim_right"
	MOD_PAD                Modifier = "pad"
	MOD_PAD_LEFT           Modifier = "pad_left"
	MOD_PAD_RIGHT          Modifier = "pad_right"
	MOD_LENGTH             Modifier = "length"
	MOD_QUOTED             Modifier = "quoted"
	MOD_REVERSE            Modifier = "reverse"
)

var ValidModifiers = []Modifier{
	MOD_TO_LOWER,
	MOD_TO_UPPER,
	MOD_TO_SNAKE,
	MOD_TO_SCREAMING_SNAKE,
	MOD_TO_KEBAB,
	MOD_TO_SCREAMING_KEBAB,
	MOD_TO_CAMEL,
	MOD_TO_PASCAL,
	MOD_TO_DELIMITED,
	MOD_TO_INVERTED,
	MOD_TRIM,
	MOD_TRIM_LEFT,
	MOD_TRIM_RIGHT,
	MOD_PAD,
	MOD_PAD_LEFT,
	MOD_PAD_RIGHT,
	MOD_LENGTH,
	MOD_QUOTED,
	MOD_REVERSE,
}
