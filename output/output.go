package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(val any) {
	switch t := val.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	case nil:
		return
	default:
		color.Red("Неизвестный тип ошибки")
		fmt.Print(val)
	}
}

func PrintWarning(val any) {
	switch t := val.(type) {
	case string:
		color.Yellow(t)
	case int:
		color.Yellow("Код предупреждения: %d", t)
	default:
		color.Yellow("Неизвестный тип предупреждения")
	}
}

func PrintSuccess(val any) {
	switch t := val.(type) {
	case string:
		color.Green(t)
	case int:
		color.Green("Операция успешно завершена с кодом: %d", t)
	default:
		color.Green("Операция успешно завершена")
	}
}
