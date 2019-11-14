package generate

import (
	"bytes"
	"fmt"
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"github.com/gobuffalo/packr"
	"html/template"
)

type buildTemplate struct {
	Base      string
	BaseTag   string
	Cmds      []string
	ImageName string
	Tag       string
}

func Generate(buildArgs v1alpha1.BuildArgs) error {
	var buf bytes.Buffer
	box := packr.NewBox("../../configs/templates")

	file, err := box.Find("ocibuilder")
	if err != nil {
		return err
	}

	tmpl, err := template.New("generatedSpec").Parse(string(file))
	if err != nil {
		return err
	}

	var cmds []string
	for _, cmd := range buildArgs.Cmds {
		cmds = append(cmds, cmd.String())
	}

	buildTemplate := buildTemplate{
		buildArgs.Base,
		buildArgs.BaseTag,
		cmds,
		buildArgs.ImageName,
		buildArgs.Tag,
	}

	if err = tmpl.Execute(&buf, buildTemplate); err != nil {
		return err
	}

	fmt.Println(buf.String())

	return nil
}
