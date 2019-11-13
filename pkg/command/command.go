package command

import (
	"fmt"
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"strings"
)

type Command struct {
	v1alpha1.Cmd
}

func (c Command) ParseCmd() string {
	return fmt.Sprintf("%s %s %s", c.Name, c.Command, strings.Join(c.Args[:], " "))
}
