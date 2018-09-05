package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewSchema(namespace, name string) *Schema {
	return &Schema{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Schema) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Schema) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type SchemaList []*Schema
type SchemasByNamespace map[string]SchemaList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list SchemaList) Find(namespace, name string) (*Schema, error) {
	for _, schema := range list {
		if schema.Metadata.Name == name {
			if namespace == "" || schema.Metadata.Namespace == namespace {
				return schema, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find schema %v.%v", namespace, name)
}

func (list SchemaList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, schema := range list {
		ress = append(ress, schema)
	}
	return ress
}

func (list SchemaList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, schema := range list {
		ress = append(ress, schema)
	}
	return ress
}

func (list SchemaList) Names() []string {
	var names []string
	for _, schema := range list {
		names = append(names, schema.Metadata.Name)
	}
	return names
}

func (list SchemaList) NamespacesDotNames() []string {
	var names []string
	for _, schema := range list {
		names = append(names, schema.Metadata.Namespace+"."+schema.Metadata.Name)
	}
	return names
}

func (list SchemaList) Sort() SchemaList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list SchemaList) Clone() SchemaList {
	var schemaList SchemaList
	for _, schema := range list {
		schemaList = append(schemaList, proto.Clone(schema).(*Schema))
	}
	return schemaList
}

func (list SchemaList) ByNamespace() SchemasByNamespace {
	byNamespace := make(SchemasByNamespace)
	for _, schema := range list {
		byNamespace.Add(schema)
	}
	return byNamespace
}

func (byNamespace SchemasByNamespace) Add(schema ...*Schema) {
	for _, item := range schema {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace SchemasByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace SchemasByNamespace) List() SchemaList {
	var list SchemaList
	for _, schemaList := range byNamespace {
		list = append(list, schemaList...)
	}
	return list.Sort()
}

func (byNamespace SchemasByNamespace) Clone() SchemasByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Schema{}

// Kubernetes Adapter for Schema

func (o *Schema) GetObjectKind() schema.ObjectKind {
	t := SchemaCrd.TypeMeta()
	return &t
}

func (o *Schema) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Schema)
}

var SchemaCrd = crd.NewCrd("sqoop.solo.io",
	"schemas",
	"sqoop.solo.io",
	"v1",
	"Schema",
	"sc",
	&Schema{})
