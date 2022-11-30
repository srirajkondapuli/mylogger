package newlog

import (
	"fmt"
)

var Version string = "1.0"

func NewLog(mess string) {
	fmt.Println("[New LOG] " + mess)
}
