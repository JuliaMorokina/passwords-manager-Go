package input

import (
	"fmt"
)

func InputData(messages ...any) string {
	for i, line := range messages {
		if i == len(messages)-1 {
			fmt.Printf("%v: \n", line)
			break
		}
		fmt.Println(line)
	}

	var res string
	fmt.Scanln(&res)
	return res
}
