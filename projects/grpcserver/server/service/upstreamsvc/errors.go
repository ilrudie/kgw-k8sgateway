package upstreamsvc

import (
	"github.com/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var (
	FailedToReadUpstreamError = func(err error, ref *core.ResourceRef) error {
		return errors.Wrapf(err, "Failed to read upstream %v.%v", ref.GetNamespace(), ref.GetName())
	}

	FailedToListUpstreamsError = func(err error, namespace string) error {
		return errors.Wrapf(err, "Failed to list upstreams in %v", namespace)
	}

	FailedToCreateUpstreamError = func(err error, ref *core.ResourceRef) error {
		return errors.Wrapf(err, "Failed to create upstream %v.%v", ref.GetNamespace(), ref.GetName())
	}

	FailedToUpdateUpstreamError = func(err error, ref *core.ResourceRef) error {
		return errors.Wrapf(err, "Failed to update upstream %v.%v", ref.GetNamespace(), ref.GetName())
	}

	FailedToParseUpstreamFromYamlError = func(err error, ref *core.ResourceRef) error {
		return errors.Wrapf(err, "Failed to parse upstream  %s.%s from YAML", ref.GetNamespace(), ref.GetName())
	}

	FailedToDeleteUpstreamError = func(err error, ref *core.ResourceRef) error {
		return errors.Wrapf(err, "Failed to delete upstream %v.%v", ref.GetNamespace(), ref.GetName())
	}

	FailedToCheckIsUpstreamReferencedError = func(err error, ref *core.ResourceRef) error {
		return errors.Wrapf(err, "Failed to verify that upstream %v.%v is not referenced in a virtual service", ref.GetNamespace(), ref.GetName())
	}

	CannotDeleteReferencedUpstreamError = func(upstreamRef *core.ResourceRef, virtualServiceRefs []*core.ResourceRef) error {
		return errors.Errorf("Upstream %v.%v is referenced in virtual service(s) %v",
			upstreamRef.GetNamespace(),
			upstreamRef.GetName(),
			virtualServiceRefs,
		)
	}
)
