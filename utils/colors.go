package utils

type Color string

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	Gray   Color = "\033[37m"
	White  Color = "\033[97m"
)

func Colorize(color Color, text string) string {
	return string(color) + text + string(Reset)
}
