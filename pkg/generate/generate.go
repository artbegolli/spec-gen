package generate

import (
	"bytes"
	"fmt"
	"html/template"
	
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"github.com/gobuffalo/packr"
)

func Generate(buildTemplate v1alpha1.BuildTemplate) ([]byte, error) {
	var buf bytes.Buffer
	box := packr.NewBox("../../configs/templates")

	file, err := box.Find("ocibuilder")
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("generatedSpec").Parse(string(file))
	if err != nil {
		return nil, err
	}

	if err = tmpl.Execute(&buf, buildTemplate); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
