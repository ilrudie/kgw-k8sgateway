// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
)

type Artifact struct {
	Data     string   `json:"data"`
	Metadata Metadata `json:"metadata"`
}
type AwsDestinationSpec struct {
	LogicalName            string                   `json:"logicalName"`
	InvocationStyle        AwsLambdaInvocationStyle `json:"invocationStyle"`
	ResponseTransformation bool                     `json:"responseTransformation"`
}
type AwsLambdaFunction struct {
	LogicalName  string `json:"logicalName"`
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}
type AwsSecret struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}
type AwsUpstreamSpec struct {
	Region    string              `json:"region"`
	SecretRef ResourceRef         `json:"secretRef"`
	Functions []AwsLambdaFunction `json:"functions"`
}
type AzureDestinationSpec struct {
	FunctionName string `json:"functionName"`
}
type AzureFunction struct {
	FunctionName string           `json:"functionName"`
	AuthLevel    AzureFnAuthLevel `json:"authLevel"`
}
type AzureSecret struct {
	APIKeys *MapStringString `json:"apiKeys"`
}
type AzureUpstreamSpec struct {
	FunctionAppName string          `json:"functionAppName"`
	SecretRef       ResourceRef     `json:"secretRef"`
	Functions       []AzureFunction `json:"functions"`
}
type Destination interface{}
type DestinationSpec interface{}
type FieldResolver struct {
	FieldName string   `json:"fieldName"`
	Resolver  Resolver `json:"resolver"`
}
type GRPCServiceSpec struct {
	Empty *string `json:"empty"`
}
type GlooResolver struct {
	RequestTemplate  *RequestTemplate  `json:"requestTemplate"`
	ResponseTemplate *ResponseTemplate `json:"responseTemplate"`
	Destination      Destination       `json:"destination"`
}
type InputArtifact struct {
	Data     string        `json:"data"`
	Metadata InputMetadata `json:"metadata"`
}
type InputAwsDestinationSpec struct {
	LogicalName            string                   `json:"logicalName"`
	InvocationStyle        AwsLambdaInvocationStyle `json:"invocationStyle"`
	ResponseTransformation bool                     `json:"responseTransformation"`
}
type InputAwsLambdaFunction struct {
	LogicalName  string `json:"logicalName"`
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}
type InputAwsSecret struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}
type InputAwsUpstreamSpec struct {
	Region    string                   `json:"region"`
	SecretRef InputResourceRef         `json:"secretRef"`
	Functions []InputAwsLambdaFunction `json:"functions"`
}
type InputAzureDestinationSpec struct {
	FunctionName string `json:"functionName"`
}
type InputAzureFunction struct {
	FunctionName string `json:"functionName"`
	AuthLevel    string `json:"authLevel"`
}
type InputAzureSecret struct {
	APIKeys InputMapStringString `json:"apiKeys"`
}
type InputAzureUpstreamSpec struct {
	FunctionAppName string               `json:"functionAppName"`
	SecretRef       *InputResourceRef    `json:"secretRef"`
	Functions       []InputAzureFunction `json:"functions"`
}
type InputDestination struct {
	SingleDestination *InputSingleDestination `json:"singleDestination"`
	MultiDestination  *InputMultiDestination  `json:"multiDestination"`
}
type InputDestinationSpec struct {
	Aws     *InputAwsDestinationSpec    `json:"aws"`
	Azure   *InputAzureDestinationSpec  `json:"azure"`
	Swagger *InputSwaggerDestiationSpec `json:"swagger"`
}
type InputFieldResolver struct {
	FieldName string        `json:"fieldName"`
	Resolver  InputResolver `json:"resolver"`
}
type InputGRPCServiceSpec struct {
	Empty *string `json:"empty"`
}
type InputGlooResolver struct {
	RequestTemplate  *InputRequestTemplate  `json:"requestTemplate"`
	ResponseTemplate *InputResponseTemplate `json:"responseTemplate"`
	Destination      InputDestination       `json:"destination"`
}
type InputKeyValueMatcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}
type InputKubeUpstreamSpec struct {
	ServiceName      string                `json:"serviceName"`
	ServiceNamespace string                `json:"serviceNamespace"`
	ServicePort      int                   `json:"servicePort"`
	Selector         *InputMapStringString `json:"selector"`
	ServiceSpec      *InputServiceSpec     `json:"serviceSpec"`
}
type InputMapStringString struct {
	Values []InputValue `json:"values"`
}
type InputMatcher struct {
	PathMatch       string                 `json:"pathMatch"`
	PathMatchType   PathMatchType          `json:"pathMatchType"`
	Headers         []InputKeyValueMatcher `json:"headers"`
	QueryParameters []InputKeyValueMatcher `json:"queryParameters"`
	Methods         []string               `json:"methods"`
}
type InputMetadata struct {
	Name            string                `json:"name"`
	Namespace       string                `json:"namespace"`
	ResourceVersion string                `json:"resourceVersion"`
	Labels          *InputMapStringString `json:"labels"`
	Annotations     *InputMapStringString `json:"annotations"`
}
type InputMultiDestination struct {
	Destinations []InputWeightedDestination `json:"destinations"`
}
type InputNodeJSResolver struct {
	Empty *string `json:"empty"`
}
type InputRequestTemplate struct {
	Verb    string                `json:"verb"`
	Path    string                `json:"path"`
	Body    string                `json:"body"`
	Headers *InputMapStringString `json:"headers"`
}
type InputResolver struct {
	GlooResolver     *InputGlooResolver     `json:"glooResolver"`
	TemplateResolver *InputTemplateResolver `json:"templateResolver"`
	NodeResolver     *InputNodeJSResolver   `json:"nodeResolver"`
}
type InputResolverMap struct {
	Types    []InputTypeResolver `json:"types"`
	Metadata InputMetadata       `json:"metadata"`
}
type InputResourceRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type InputResponseTemplate struct {
	Body    string                `json:"body"`
	Headers *InputMapStringString `json:"headers"`
}
type InputRoute struct {
	Matcher     InputMatcher       `json:"matcher"`
	Destination InputDestination   `json:"destination"`
	Plugins     *InputRoutePlugins `json:"plugins"`
}
type InputRoutePlugins struct {
	Empty *string `json:"empty"`
}
type InputSchema struct {
	ResolverMap  InputResourceRef `json:"resolverMap"`
	InlineSchema string           `json:"inlineSchema"`
	Metadata     InputMetadata    `json:"metadata"`
}
type InputSecret struct {
	Kind     InputSecretKind `json:"kind"`
	Metadata InputMetadata   `json:"metadata"`
}
type InputSecretKind struct {
	Aws   *InputAwsSecret   `json:"aws"`
	Azure *InputAzureSecret `json:"azure"`
	TLS   *InputTlsSecret   `json:"tls"`
}
type InputServiceSpec struct {
	Swagger *InputSwaggerServiceSpec `json:"swagger"`
	Grpc    *InputGRPCServiceSpec    `json:"grpc"`
}
type InputSingleDestination struct {
	Upstream        InputResourceRef      `json:"upstream"`
	DestinationSpec *InputDestinationSpec `json:"destinationSpec"`
}
type InputSslConfig struct {
	SecretRef InputResourceRef `json:"secretRef"`
}
type InputStaticHost struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}
type InputStaticUpstreamSpec struct {
	Hosts       []InputStaticHost `json:"hosts"`
	ServiceSpec *InputServiceSpec `json:"serviceSpec"`
	UseTLS      bool              `json:"useTls"`
}
type InputStatus struct {
	State  State  `json:"state"`
	Reason string `json:"reason"`
}
type InputSwaggerDestiationSpec struct {
	FunctionName string `json:"functionName"`
}
type InputSwaggerServiceSpec struct {
	Empty *string `json:"empty"`
}
type InputTemplateResolver struct {
	Empty *string `json:"empty"`
}
type InputTlsSecret struct {
	CertChain  string `json:"certChain"`
	PrivateKey string `json:"privateKey"`
	RootCa     string `json:"rootCa"`
}
type InputTypeResolver struct {
	TypeName string               `json:"typeName"`
	Fields   []InputFieldResolver `json:"fields"`
}
type InputUpstream struct {
	Spec     InputUpstreamSpec `json:"spec"`
	Metadata InputMetadata     `json:"metadata"`
}
type InputUpstreamSpec struct {
	Aws    *InputAwsUpstreamSpec    `json:"aws"`
	Azure  *InputAzureUpstreamSpec  `json:"azure"`
	Kube   *InputKubeUpstreamSpec   `json:"kube"`
	Static *InputStaticUpstreamSpec `json:"static"`
}
type InputValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type InputVirtualService struct {
	Domains   []string                    `json:"domains"`
	Routes    []InputRoute                `json:"routes"`
	SslConfig *InputSslConfig             `json:"sslConfig"`
	Plugins   *InputVirtualServicePlugins `json:"plugins"`
	Metadata  InputMetadata               `json:"metadata"`
}
type InputVirtualServicePlugins struct {
	Empty *string `json:"empty"`
}
type InputWeightedDestination struct {
	Destination InputSingleDestination `json:"destination"`
	Weight      int                    `json:"weight"`
}
type KeyValueMatcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
}
type KubeUpstreamSpec struct {
	ServiceName      string           `json:"serviceName"`
	ServiceNamespace string           `json:"serviceNamespace"`
	ServicePort      int              `json:"servicePort"`
	Selector         *MapStringString `json:"selector"`
	ServiceSpec      ServiceSpec      `json:"serviceSpec"`
}
type MapStringString struct {
	Values []Value `json:"values"`
}
type Matcher struct {
	PathMatch       string            `json:"pathMatch"`
	PathMatchType   PathMatchType     `json:"pathMatchType"`
	Headers         []KeyValueMatcher `json:"headers"`
	QueryParameters []KeyValueMatcher `json:"queryParameters"`
	Methods         []string          `json:"methods"`
}
type Metadata struct {
	Name            string           `json:"name"`
	Namespace       string           `json:"namespace"`
	ResourceVersion string           `json:"resourceVersion"`
	Labels          *MapStringString `json:"labels"`
	Annotations     *MapStringString `json:"annotations"`
}
type MultiDestination struct {
	Destinations []WeightedDestination `json:"destinations"`
}
type NodeJSResolver struct {
	Empty *string `json:"empty"`
}
type OAuthEndpoint struct {
	URL        string `json:"url"`
	ClientName string `json:"clientName"`
}
type RequestTemplate struct {
	Verb    string           `json:"verb"`
	Path    string           `json:"path"`
	Body    string           `json:"body"`
	Headers *MapStringString `json:"headers"`
}
type Resolver interface{}
type ResolverMap struct {
	Types    []TypeResolver `json:"types"`
	Metadata Metadata       `json:"metadata"`
	Status   Status         `json:"status"`
}
type ResourceRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type ResponseTemplate struct {
	Body    string           `json:"body"`
	Headers *MapStringString `json:"headers"`
}
type Route struct {
	Matcher     Matcher       `json:"matcher"`
	Destination Destination   `json:"destination"`
	Plugins     *RoutePlugins `json:"plugins"`
}
type RoutePlugins struct {
	Empty *string `json:"empty"`
}
type Schema struct {
	ResolverMap  ResourceRef `json:"resolverMap"`
	InlineSchema string      `json:"inlineSchema"`
	Metadata     Metadata    `json:"metadata"`
	Status       Status      `json:"status"`
}
type Secret struct {
	Kind     SecretKind `json:"kind"`
	Metadata Metadata   `json:"metadata"`
}
type SecretKind interface{}
type ServiceSpec interface{}
type SingleDestination struct {
	Upstream        ResourceRef     `json:"upstream"`
	DestinationSpec DestinationSpec `json:"destinationSpec"`
}
type SslConfig struct {
	SecretRef ResourceRef `json:"secretRef"`
}
type StaticHost struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}
type StaticUpstreamSpec struct {
	Hosts       []StaticHost `json:"hosts"`
	ServiceSpec ServiceSpec  `json:"serviceSpec"`
	UseTLS      bool         `json:"useTls"`
}
type Status struct {
	State  State   `json:"state"`
	Reason *string `json:"reason"`
}
type SwaggerServiceSpec struct {
	Empty *string `json:"empty"`
}
type TemplateResolver struct {
	InlineTemplate *string `json:"inlineTemplate"`
}
type TlsSecret struct {
	CertChain  string `json:"certChain"`
	PrivateKey string `json:"privateKey"`
	RootCa     string `json:"rootCa"`
}
type TypeResolver struct {
	TypeName string          `json:"typeName"`
	Fields   []FieldResolver `json:"fields"`
}
type Upstream struct {
	Spec     UpstreamSpec `json:"spec"`
	Metadata Metadata     `json:"metadata"`
	Status   Status       `json:"status"`
}
type UpstreamSpec interface{}
type Value struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type VirtualService struct {
	Domains   []string               `json:"domains"`
	Routes    []Route                `json:"routes"`
	SslConfig *SslConfig             `json:"sslConfig"`
	Plugins   *VirtualServicePlugins `json:"plugins"`
	Metadata  Metadata               `json:"metadata"`
	Status    Status                 `json:"status"`
}
type VirtualServicePlugins struct {
	Empty *string `json:"empty"`
}
type WeightedDestination struct {
	Destination SingleDestination `json:"destination"`
	Weight      int               `json:"weight"`
}

type AwsLambdaInvocationStyle string

const (
	AwsLambdaInvocationStyleSync  AwsLambdaInvocationStyle = "SYNC"
	AwsLambdaInvocationStyleAsync AwsLambdaInvocationStyle = "ASYNC"
)

func (e AwsLambdaInvocationStyle) IsValid() bool {
	switch e {
	case AwsLambdaInvocationStyleSync, AwsLambdaInvocationStyleAsync:
		return true
	}
	return false
}

func (e AwsLambdaInvocationStyle) String() string {
	return string(e)
}

func (e *AwsLambdaInvocationStyle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AwsLambdaInvocationStyle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AwsLambdaInvocationStyle", str)
	}
	return nil
}

func (e AwsLambdaInvocationStyle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AzureFnAuthLevel string

const (
	AzureFnAuthLevelAnonymous AzureFnAuthLevel = "ANONYMOUS"
	AzureFnAuthLevelFunction  AzureFnAuthLevel = "FUNCTION"
	AzureFnAuthLevelAdmin     AzureFnAuthLevel = "ADMIN"
)

func (e AzureFnAuthLevel) IsValid() bool {
	switch e {
	case AzureFnAuthLevelAnonymous, AzureFnAuthLevelFunction, AzureFnAuthLevelAdmin:
		return true
	}
	return false
}

func (e AzureFnAuthLevel) String() string {
	return string(e)
}

func (e *AzureFnAuthLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AzureFnAuthLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AzureFnAuthLevel", str)
	}
	return nil
}

func (e AzureFnAuthLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PathMatchType string

const (
	PathMatchTypePrefix PathMatchType = "PREFIX"
	PathMatchTypeRegex  PathMatchType = "REGEX"
	PathMatchTypeExact  PathMatchType = "EXACT"
)

func (e PathMatchType) IsValid() bool {
	switch e {
	case PathMatchTypePrefix, PathMatchTypeRegex, PathMatchTypeExact:
		return true
	}
	return false
}

func (e PathMatchType) String() string {
	return string(e)
}

func (e *PathMatchType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PathMatchType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PathMatchType", str)
	}
	return nil
}

func (e PathMatchType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type State string

const (
	StatePending  State = "PENDING"
	StateAccepted State = "ACCEPTED"
	StateRejected State = "REJECTED"
)

func (e State) IsValid() bool {
	switch e {
	case StatePending, StateAccepted, StateRejected:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
