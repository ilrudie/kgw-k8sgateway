package generate

import glooGen "github.com/solo-io/gloo/install/helm/gloo/generate"

type HelmConfig struct {
	Global                 interface{}          `json:"global,omitempty"`
	Enabled                *bool                `json:"enabled,omitempty" desc:"If true, deploy federation service (default true)."`
	CreateLicenseSecret    *bool                `json:"create_license_secret,omitempty"`
	LicenseSecretName      *string              `json:"license_secret_name,omitempty"`
	LicenseKey             *string              `json:"license_key,omitempty"`
	EnableMultiClusterRbac *bool                `json:"enableMultiClusterRbac,omitempty"`
	GlooFedApiServer       *ApiServerDeployment `json:"glooFedApiserver,omitempty"`
	GlooFed                *GlooFedDeployment   `json:"glooFed,omitempty"`
	Rbac                   *RbacConfiguration   `json:"rbac,omitempty"`
	RbacWebhook            *RbacWebhook         `json:"rbacWebhook,omitempty"`
}

type GlooFedDeployment struct {
	Replicas  *int                          `json:"replicas,omitempty"`
	Resources *glooGen.ResourceRequirements `json:"resources,omitempty"`
	Image     *glooGen.Image                `json:"image,omitempty"`
	Stats     *glooGen.Stats                `json:"stats,omitempty"`
}

type ApiServerDeployment struct {
	Enable          *bool                         `json:"enable,omitempty"`
	Replicas        *int                          `json:"replicas,omitempty"`
	Image           *glooGen.Image                `json:"image,omitempty"`
	Port            *int                          `json:"port,omitempty"`
	HealthCheckPort *int                          `json:"healthCheckPort,omitempty"`
	Resources       *glooGen.ResourceRequirements `json:"resources,omitempty"`
	Stats           *glooGen.Stats                `json:"stats,omitempty"`
	FloatingUserId  *bool                         `json:"floatingUserId,omitempty"`
	RunAsUser       *float64                      `json:"runAsUser,omitempty"`
	Console         *ConsoleContainer             `json:"console,omitempty"`
	Envoy           *EnvoyContainer               `json:"envoy,omitempty"`
}

type ConsoleContainer struct {
	Image     *glooGen.Image                `json:"image,omitempty"`
	Port      *int                          `json:"port,omitempty"`
	Resources *glooGen.ResourceRequirements `json:"resources,omitempty"`
}

type EnvoyContainer struct {
	Image                 *glooGen.Image                `json:"image,omitempty"`
	Resources             *glooGen.ResourceRequirements `json:"resources,omitempty"`
	BoostrapConfiguration *EnvoyBootstrapConfiguration  `json:"bootstrapConfig,omitempty""`
}

type EnvoyBootstrapConfiguration struct {
	ConfigMapName *string `json:"configMapName,omitempty"`
}

type RbacWebhook struct {
	Image     *glooGen.Image                `json:"image,omitempty"`
	Resources *glooGen.ResourceRequirements `json:"resources,omitempty"`
}

type RbacConfiguration struct {
	Create *bool `json:"create,omitempty"`
}
