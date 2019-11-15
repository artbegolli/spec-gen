package main

import (
	"fmt"
	"github.com/artbegolli/spec-gen/pkg/platform"
)

func main() {
	fmt.Println("hello service mate")

	platformSpec, _ := platform.GetPlatformSpec("python", []string{"jupyter", "tensorflow"})
	platformSpec.Generate("test-image", "v0.1.0")
}
