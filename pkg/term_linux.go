package gurl

import (
	"exp/terminal"
)

func TermWidth() (int, error) {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		return 0, err
	}
	return width, nil
}
