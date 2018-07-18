package resources

import (
	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"v/k8s.io/apimachinery@v0.0.0-20180328184639-0ed326127d30/pkg/util/validation"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type Resource interface {
	proto.Message
	Equal(that interface{}) bool
	GetStatus() core.Status
	GetMetadata() core.Metadata
	SetStatus(status core.Status)
	SetMetadata(meta core.Metadata)
}

func Validate(resource Resource) error {
	return ValidateName(resource.GetMetadata().Name)
}

func ValidateName(name string) error {
	errs := validation.IsDNS1035Label(name)
	if len(name) < 1 {
		errs = append(errs, "name cannot be empty")
	}
	if len(name) > 253 {
		errs = append(errs, "name has a max length of 253 characters")
	}
	if len(errs) > 0 {
		return errors.Errors(errs)
	}
	return nil
}