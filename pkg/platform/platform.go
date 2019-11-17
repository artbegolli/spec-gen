package platform

import (
	"errors"

	"github.com/artbegolli/spec-gen/api/spec-gen/v1alpha1"
	"github.com/sirupsen/logrus"
)

type Platform interface {
	Generate(imageName string, tag string) ([]byte, error)
}

func GetPlatformSpec(platform v1alpha1.Platform, buildCmds []string) (Platform, error) {

	switch v1alpha1.Platform(platform) {

	case v1alpha1.Python:
		{
			return NewPythonSpec(buildCmds), nil
		}

	default:
		{
			logrus.WithField("platform", platform).Info("platform not currently supported")
			return nil, errors.New("platform not currently supported")
		}
	}

}
