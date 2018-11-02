package protoc

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	"github.com/solo-io/solo-kit/pkg/errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/solo-io/solo-kit/pkg/code-generator/codegen"
	"github.com/solo-io/solo-kit/pkg/utils/log"
)

// plugin is an implementation of protokit.Plugin
type Plugin struct{}

func parseParamsString(paramString string) codegen.Params {
	var params codegen.Params
	pairs := strings.Split(paramString, ",")
	for _, pair := range pairs {
		split := strings.Split(pair, "=")
		key := split[0]
		val := split[1]
		switch key {
		case "project_file":
			params.ProjectFile = val
		case "collection_run":
			params.CollectionRun = val == "true"
		}
	}
	return params
}

func (p *Plugin) Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	log.DefaultOut = &bytes.Buffer{}
	if false {
		log.DefaultOut = os.Stderr
	}

	log.Printf("parsing request %v", req.FileToGenerate, req.GetParameter())
	paramString := req.GetParameter()
	if paramString == "" {
		return nil, errors.Errorf(`must provide project params via --solo-kit_out=project_file=${PWD}/project.json,collection_run=true:${OUT}`)
	}

	params := parseParamsString(paramString)

	// append descriptors if they already exist
	collectedDescriptorsFile := params.ProjectFile + ".descriptors"
	descriptorBytes, err := ioutil.ReadFile(collectedDescriptorsFile)
	if err == nil {
		var previousDescriptors plugin_go.CodeGeneratorRequest
		if err := proto.Unmarshal(descriptorBytes, &previousDescriptors); err != nil {
			return nil, errors.Wrapf(err, "unmarshalling previous request")
		}
		// append all unique
		for _, prev := range previousDescriptors.ProtoFile {
			var duplicate bool
			for _, desc := range req.ProtoFile {
				if desc.GetName() == prev.GetName() {
					duplicate = true
					break
				}
			}
			if duplicate {
				continue
			}
			req.ProtoFile = append(req.ProtoFile, prev)
			for _, genFile := range previousDescriptors.FileToGenerate {
				if genFile == prev.GetName() {
					req.FileToGenerate = append(req.FileToGenerate, genFile)
					break
				}
			}
		}
	}

	collectedDescriptorsBytes, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal %+v", req)
	}
	if err := ioutil.WriteFile(collectedDescriptorsFile, collectedDescriptorsBytes, 0644); err != nil {
		return nil, errors.Wrapf(err, "failed to write %v", collectedDescriptorsFile)
	}
	// only purpose in the collection run is to build up our request
	return new(plugin_go.CodeGeneratorResponse), nil
	if params.CollectionRun {
	} else {
		return &plugin_go.CodeGeneratorResponse{}, nil
	}

	project, err := codegen.ParseRequest(params, req)
	if err != nil {
		return nil, err
	}

	code, err := codegen.GenerateFiles(project)
	if err != nil {
		return nil, err
	}

	log.Printf("%v", project)
	log.Printf("%v", code)

	resp := new(plugin_go.CodeGeneratorResponse)

	for _, file := range code {
		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(file.Filename),
			Content: proto.String(file.Content),
		})
	}

	return resp, nil
}
