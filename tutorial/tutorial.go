package tutorial

import (
	"fmt"
	"regexp"
)

func Say(text string) {
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllString(text, "[32m$1[36m")
	fmt.Printf("[36m\n\n    %s\n\n[0m", text)
}
