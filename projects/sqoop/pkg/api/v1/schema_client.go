package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type SchemaClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Schema, error)
	Write(resource *Schema, opts clients.WriteOpts) (*Schema, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (SchemaList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan SchemaList, <-chan error, error)
}

type schemaClient struct {
	rc clients.ResourceClient
}

func NewSchemaClient(rcFactory factory.ResourceClientFactory) (SchemaClient, error) {
	return NewSchemaClientWithToken(rcFactory, "")
}

func NewSchemaClientWithToken(rcFactory factory.ResourceClientFactory, token string) (SchemaClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Schema{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Schema resource client")
	}
	return &schemaClient{
		rc: rc,
	}, nil
}

func (client *schemaClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *schemaClient) Register() error {
	return client.rc.Register()
}

func (client *schemaClient) Read(namespace, name string, opts clients.ReadOpts) (*Schema, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Schema), nil
}

func (client *schemaClient) Write(schema *Schema, opts clients.WriteOpts) (*Schema, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(schema, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Schema), nil
}

func (client *schemaClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *schemaClient) List(namespace string, opts clients.ListOpts) (SchemaList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToSchema(resourceList), nil
}

func (client *schemaClient) Watch(namespace string, opts clients.WatchOpts) (<-chan SchemaList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	schemasChan := make(chan SchemaList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				schemasChan <- convertToSchema(resourceList)
			case <-opts.Ctx.Done():
				close(schemasChan)
				return
			}
		}
	}()
	return schemasChan, errs, nil
}

func convertToSchema(resources resources.ResourceList) SchemaList {
	var schemaList SchemaList
	for _, resource := range resources {
		schemaList = append(schemaList, resource.(*Schema))
	}
	return schemaList
}
