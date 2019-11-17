package v1alpha1

import (
	"github.com/artbegolli/spec-gen/pkg/platform"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}

func (s *Server) SpecGen(ctx context.Context, in *SpecMessage) (*SpecResponse, error) {
	log.Printf("Receive message %s", in.ImageName)
	platformSpec, _ := platform.GetPlatformSpec(Platform(in.Platform), in.Cmds)
	spec, err := platformSpec.Generate(in.ImageName, in.Tag)
	if err != nil {
		return nil, err
	}

	return &SpecResponse{
		ImageName: in.ImageName,
		BuildSpec: spec,
	}, nil
}
