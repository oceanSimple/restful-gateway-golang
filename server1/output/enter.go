package output

func DyeText(text string, color string) string {
	return color + text + Reset
}

func Error() string {
	return DyeText("[Error] ", Red)
}

func Info() string {
	return DyeText("[Info] ", Green)
}

func Default() string {
	return DyeText("[Default] ", Grey)
}
