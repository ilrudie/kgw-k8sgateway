// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewLicenseKey(namespace, name string) *LicenseKey {
	return &LicenseKey{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *LicenseKey) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *LicenseKey) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Key,
		r.Type,
	)
}

type LicenseKeyList []*LicenseKey
type LicensesByNamespace map[string]LicenseKeyList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list LicenseKeyList) Find(namespace, name string) (*LicenseKey, error) {
	for _, licenseKey := range list {
		if licenseKey.Metadata.Name == name {
			if namespace == "" || licenseKey.Metadata.Namespace == namespace {
				return licenseKey, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find licenseKey %v.%v", namespace, name)
}

func (list LicenseKeyList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, licenseKey := range list {
		ress = append(ress, licenseKey)
	}
	return ress
}

func (list LicenseKeyList) Names() []string {
	var names []string
	for _, licenseKey := range list {
		names = append(names, licenseKey.Metadata.Name)
	}
	return names
}

func (list LicenseKeyList) NamespacesDotNames() []string {
	var names []string
	for _, licenseKey := range list {
		names = append(names, licenseKey.Metadata.Namespace+"."+licenseKey.Metadata.Name)
	}
	return names
}

func (list LicenseKeyList) Sort() LicenseKeyList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list LicenseKeyList) Clone() LicenseKeyList {
	var licenseKeyList LicenseKeyList
	for _, licenseKey := range list {
		licenseKeyList = append(licenseKeyList, proto.Clone(licenseKey).(*LicenseKey))
	}
	return licenseKeyList
}

func (list LicenseKeyList) Each(f func(element *LicenseKey)) {
	for _, licenseKey := range list {
		f(licenseKey)
	}
}

func (list LicenseKeyList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *LicenseKey) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace LicensesByNamespace) Add(licenseKey ...*LicenseKey) {
	for _, item := range licenseKey {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace LicensesByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace LicensesByNamespace) List() LicenseKeyList {
	var list LicenseKeyList
	for _, licenseKeyList := range byNamespace {
		list = append(list, licenseKeyList...)
	}
	return list.Sort()
}

func (byNamespace LicensesByNamespace) Clone() LicensesByNamespace {
	cloned := make(LicensesByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &LicenseKey{}

// Kubernetes Adapter for LicenseKey

func (o *LicenseKey) GetObjectKind() schema.ObjectKind {
	t := LicenseKeyCrd.TypeMeta()
	return &t
}

func (o *LicenseKey) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*LicenseKey)
}

var LicenseKeyCrd = crd.NewCrd("licensing.solo.io",
	"licenses",
	"licensing.solo.io",
	"v1",
	"LicenseKey",
	"lc",
	false,
	&LicenseKey{})
