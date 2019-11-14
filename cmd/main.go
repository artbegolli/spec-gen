package main

import (
	"fmt"
	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"github.com/artbegolli/spec-gen/pkg/generate"
)

func main() {
	fmt.Println("hello service mate")

	platform := "python"

	switch v1alpha1.PlatformType(platform) {

	case v1alpha1.Python:
		{
			buildArgs := v1alpha1.BuildArgs{
				ImageName: "first-image",
				Tag:       "v0.1.0.",
				Base:      "python",
				BaseTag:   "3",
				Cmds: []v1alpha1.Cmd{
					{
						Name:    "pip",
						Command: "install",
						Args:    []string{"jupyter"},
					},
					{
						Name:    "pip",
						Command: "install",
						Args:    []string{"tensorflow"},
					},
				},
			}
			generate.Generate(buildArgs)
		}
	}

}
