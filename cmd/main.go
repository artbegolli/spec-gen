package main

import (
	"github.com/artbegolli/spec-gen/pkg/platform"
)

func main() {

	platformSpec, _ := platform.GetPlatformSpec("python", []string{"jupyter", "tensorflow"})
	platformSpec.Generate("test-image", "v0.1.0")
}
