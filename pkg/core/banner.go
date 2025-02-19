package core

import (
	"fmt"
	"strings"
)

func Banner() string {
	banner := "\n\t" + fmt.Sprintf(`%v`, ICON) + fmt.Sprintf(` %v: `, strings.ToUpper(string(BINARY[0]))+BINARY[1:]) + fmt.Sprintf(`%v`, DESC)
	banner += "\n\t\t\x20\x20\x20\x20\x20\x20\x20" + fmt.Sprintf(`%v`, VERSION) + fmt.Sprintf(` by %v`, AUTHOR) + "\n\n"
	return banner
}
