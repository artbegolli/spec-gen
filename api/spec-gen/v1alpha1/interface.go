package v1alpha1

type Platform interface {
	Generate() error
}
