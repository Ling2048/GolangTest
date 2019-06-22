package Common

import "strings"

func Test(s string) string {
	return strings.Join([]string{"test", s}, ",")
}
