package generate

import (
	"bytes"
	"fmt"
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"github.com/gobuffalo/packr"
	"html/template"
)

func Generate(buildTemplate v1alpha1.BuildTemplate) error {
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

	if err = tmpl.Execute(&buf, buildTemplate); err != nil {
		return err
	}

	fmt.Println(buf.String())
	return nil
}
