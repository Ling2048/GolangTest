package Common

import (
	"strings"
)

func Test(s string) string {
	var ss string = "s"
	return strings.Join([]string{"test", s}, ",")
}
