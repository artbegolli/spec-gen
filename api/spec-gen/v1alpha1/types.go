package v1alpha1

import (
	"fmt"
	"strings"
)

type Platform string

const (
	Python Platform = "python"
	Java   Platform = "java"
	Go     Platform = "go"
)

type Cmd struct {
	Name    string
	Command string
	Args    []string
}

func (c Cmd) String() string {
	return fmt.Sprintf("%s %s %s", c.Name, c.Command, strings.Join(c.Args, " "))
}

type BuildArgs struct {
	Base      string
	BaseTag   string
	Cmds      []Cmd
	ImageName string
	Tag       string
}

type BuildTemplate struct {
	Base      string
	BaseTag   string
	Cmds      []string
	ImageName string
	Tag       string
}
