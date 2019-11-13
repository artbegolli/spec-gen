package v1alpha1

type Cmd struct {
	Name    string
	Command string
	Args    []string
}

type BuildArgs struct {
	Cmds      []Cmd
	ImageName string
	Base      string
	BaseTag   string
}
