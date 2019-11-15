package platform

import (
	"fmt"
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"github.com/artbegolli/spec-gen/pkg/generate"
)

type Python struct {
	Base    string
	BaseTag string
	Cmds    []string
}

func (p Python) Generate(imageName string, tag string) error {

	var pythonCmds []string
	for _, c := range p.Cmds {
		pythonCmds = append(pythonCmds, fmt.Sprintf("pip install %s", c))
	}

	template := v1alpha1.BuildTemplate{
		Base:      p.Base,
		BaseTag:   p.BaseTag,
		Cmds:      pythonCmds,
		ImageName: imageName,
		Tag:       tag,
	}
	return generate.Generate(template)
}

func NewPythonSpec(cmds []string) Platform {
	return Python{
		Base:    "python",
		BaseTag: "3",
		Cmds:    cmds,
	}
}
