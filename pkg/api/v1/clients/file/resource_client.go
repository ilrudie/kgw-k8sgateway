package file

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"time"
	"github.com/gogo/protobuf/proto"
	"strconv"
	"fmt"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/fileutils"
	"path/filepath"
	"reflect"
	"os"
)

type ResourceClient struct {
	dir         string
	refreshRate time.Duration
}

func NewResourceClient(dir string, refreshRate time.Duration) *ResourceClient {
	return &ResourceClient{
		dir:         dir,
		refreshRate: refreshRate,
	}
}

var _ clients.ResourceClient = &ResourceClient{}

func (rc *ResourceClient) Register() error {
	return nil
}

func (rc *ResourceClient) Read(name string, into resources.Resource, opts clients.GetOptions) error {
	if err := resources.ValidateName(name); err != nil {
		return errors.Wrapf(err, "validation error")
	}
	if opts.Namespace == "" {
		opts.Namespace = clients.DefaultNamespace
	}
	return fileutils.ReadFileInto(rc.filename(opts.Namespace, name), into)
}

func (rc *ResourceClient) Write(resource resources.Resource, opts clients.WriteOptions) (resources.Resource, error) {
	if err := resources.Validate(resource); err != nil {
		return nil, errors.Wrapf(err, "validation error")
	}

	meta := resource.GetMetadata()
	if meta.Namespace == "" {
		meta.Namespace = clients.DefaultNamespace
	}

	if !opts.OverwriteExisting {
		empty := reflect.New(reflect.TypeOf(resource)).Elem().Interface().(resources.Resource)
		if err := rc.Read(resource.GetMetadata().Name, empty, clients.GetOptions{
			Ctx:       opts.Ctx,
			Namespace: meta.Namespace,
		}); err == nil {
			return nil, errors.NewAlreadyExistsErr(resource.GetMetadata())
		}
	}

	// mutate and return clone
	clone := proto.Clone(resource).(resources.Resource)
	// initialize or increment resource version
	meta.ResourceVersion = newOrIncrementResourceVer(meta.ResourceVersion)
	clone.SetMetadata(meta)

	path := rc.filename(meta.Namespace, meta.Name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil && !os.IsExist(err) {
		return nil, errors.Wrapf(err, "creating directory")
	}
	if err := fileutils.WriteToFile(path, clone); err != nil {
		return nil, errors.Wrapf(err, "writing file")
	}
	return clone, nil
}

func (rc *ResourceClient) Delete(name string, opts clients.DeleteOptions) error { panic("yay") }

func (rc *ResourceClient) List(opts clients.ListOptions) ([]resources.Resource, error) { panic("yay") }

func (rc *ResourceClient) Watch(opts clients.WatchOptions) (<-chan []resources.Resource, error) { panic("yay") }

func (rc *ResourceClient) filename(namespace, name string) string {
	return filepath.Join(rc.dir, namespace, name)
}

// util methods
func newOrIncrementResourceVer(resourceVersion string) string {
	curr, err := strconv.Atoi(resourceVersion)
	if err != nil {
		curr = 1
	}
	return fmt.Sprintf("%v", curr)
}
