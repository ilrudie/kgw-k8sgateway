package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	"github.com/solo-io/go-utils/installutils/kuberesource"
	"github.com/solo-io/solo-projects/install/helm/gloo-ee/generate"
	"github.com/solo-io/solo-projects/pkg/install"
	jobsv1 "k8s.io/api/batch/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/intstr"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/go-utils/manifesttestutils"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

var _ = Describe("Helm Test", func() {
	var (
		version string

		normalPromAnnotations = map[string]string{
			"prometheus.io/path":   "/metrics",
			"prometheus.io/port":   "9091",
			"prometheus.io/scrape": "true",
		}

		statsEnvVar = v1.EnvVar{
			Name:  "START_STATS_SERVER",
			Value: "true",
		}
	)

	Describe("gloo-ee helm tests", func() {
		var (
			labels        map[string]string
			selector      map[string]string
			getPullPolicy func() v1.PullPolicy
			manifestYaml  string
		)

		BeforeEach(func() {
			version = os.Getenv("TAGGED_VERSION")
			if version == "" {
				version = "dev"
				getPullPolicy = func() v1.PullPolicy { return v1.PullAlways }
			} else {
				version = version[1:]
				getPullPolicy = func() v1.PullPolicy { return v1.PullIfNotPresent }
			}
			manifestYaml = ""
		})

		AfterEach(func() {
			if manifestYaml != "" {
				err := os.Remove(manifestYaml)
				Expect(err).ToNot(HaveOccurred())
			}
		})

		Context("observability", func() {
			var (
				observabilityDeployment *appsv1.Deployment
				grafanaDeployment       *appsv1.Deployment
			)
			BeforeEach(func() {
				labels = map[string]string{
					"app":  "gloo",
					"gloo": "observability",
				}
				selector = map[string]string{
					"app":  "gloo",
					"gloo": "observability",
				}

				rb := ResourceBuilder{
					Namespace: namespace,
					Name:      "observability",
					Labels:    labels,
				}
				observabilityDeployment = rb.GetDeploymentAppsv1()

				observabilityDeployment.Spec.Template.Spec.Volumes = []v1.Volume{
					{
						Name: "upstream-dashboard-template",
						VolumeSource: v1.VolumeSource{
							ConfigMap: &v1.ConfigMapVolumeSource{
								LocalObjectReference: v1.LocalObjectReference{Name: "glooe-observability-config"},
								Items: []v1.KeyToPath{
									{
										Key:  "DASHBOARD_JSON_TEMPLATE",
										Path: "dashboard-template.json",
									},
								},
							},
						},
					},
				}
				observabilityDeployment.Spec.Template.Spec.Containers = []v1.Container{
					{
						Name:  "observability",
						Image: "quay.io/solo-io/observability-ee:dev",
						EnvFrom: []v1.EnvFromSource{
							{ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "glooe-observability-config"}}},
							{SecretRef: &v1.SecretEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "glooe-observability-secrets"}}},
						},
						VolumeMounts: []v1.VolumeMount{
							{
								Name:      "upstream-dashboard-template",
								ReadOnly:  true,
								MountPath: "/observability",
							},
						},
						Env: []v1.EnvVar{
							{
								Name: "GLOO_LICENSE_KEY",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{
											Name: "license",
										},
										Key: "license-key",
									},
								},
							},
							{
								Name: "POD_NAMESPACE",
								ValueFrom: &v1.EnvVarSource{
									FieldRef: &v1.ObjectFieldSelector{
										FieldPath: "metadata.namespace",
									},
								},
							},
							statsEnvVar,
						},
						Resources:       v1.ResourceRequirements{},
						ImagePullPolicy: "Always",
					},
				}
				observabilityDeployment.Spec.Template.Spec.ServiceAccountName = "observability"
				observabilityDeployment.Spec.Strategy = appsv1.DeploymentStrategy{}
				observabilityDeployment.Spec.Selector.MatchLabels = selector
				observabilityDeployment.Spec.Template.ObjectMeta.Labels = selector
				observabilityDeployment.Spec.Template.ObjectMeta.Annotations = normalPromAnnotations

				grafanaBuilder := ResourceBuilder{
					Namespace: "", // grafana installs to empty namespace during tests
					Name:      "release-name-grafana",
					Labels:    labels,
				}
				grafanaDeployment = grafanaBuilder.GetDeploymentAppsv1()
			})

			It("has valid default dashboards", func() {
				dashboardsDir := "../helm/gloo-ee/dashboards/"
				files, err := ioutil.ReadDir(dashboardsDir)
				Expect(err).NotTo(HaveOccurred(), "Should be able to list files")
				Expect(files).NotTo(HaveLen(0), "Should have dashboard files")
				for _, f := range files {
					bytes, err := ioutil.ReadFile(path.Join(dashboardsDir, f.Name()))
					Expect(err).NotTo(HaveOccurred(), "Should be able to read the Envoy dashboard json file")
					err = json.Unmarshal(bytes, &map[string]interface{}{})
					Expect(err).NotTo(HaveOccurred(), "Should be able to successfully unmarshal the envoy dashboard json")
				}
			})

			Context("observability deployment", func() {
				It("is installed by default", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{})
					Expect(err).NotTo(HaveOccurred())

					testManifest.ExpectDeploymentAppsV1(observabilityDeployment)
				})

				It("is not installed when grafana is disabled", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
						valuesArgs: []string{"grafana.defaultInstallationEnabled=false"},
					})
					Expect(err).NotTo(HaveOccurred())

					testManifest.Expect(observabilityDeployment.Kind, observabilityDeployment.Namespace, observabilityDeployment.Name).To(BeNil())
				})

				It("is installed when a custom grafana instance is present", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
						valuesArgs: []string{"observability.customGrafana.enabled=true"},
					})
					Expect(err).NotTo(HaveOccurred())

					testManifest.ExpectDeploymentAppsV1(observabilityDeployment)
				})
			})

			Context("observability RBAC rule", func() {

				It("allows correct operations on upstreams", func() {
					labels = map[string]string{
						"app":  "gloo",
						"gloo": "observability",
					}
					rb := ResourceBuilder{
						Name:   "observability-upstream-role-gloo-system",
						Labels: labels,
					}

					observabilityClusterRole := rb.GetClusterRole()
					observabilityClusterRole.Rules = []rbacv1.PolicyRule{
						{
							Verbs:     []string{"get", "list", "watch"},
							APIGroups: []string{"gloo.solo.io"},
							Resources: []string{"upstreams"},
						},
					}

					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{})
					Expect(err).NotTo(HaveOccurred())

					clusterRoles := testManifest.SelectResources(func(unstructured *unstructured.Unstructured) bool {
						return unstructured.GetKind() == "ClusterRole" && unstructured.GetLabels()["gloo"] == "observability"
					})

					clusterRoles.ExpectClusterRole(observabilityClusterRole)
				})
			})

			Context("grafana deployment", func() {
				It("is not installed when grafana is disabled", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
						valuesArgs: []string{"grafana.defaultInstallationEnabled=false"},
					})
					Expect(err).NotTo(HaveOccurred())

					testManifest.Expect(grafanaDeployment.Kind, grafanaDeployment.Namespace, grafanaDeployment.Name).To(BeNil())
				})

				It("is not installed when using a custom grafana instance", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
						valuesArgs: []string{
							"grafana.defaultInstallationEnabled=false",
							"observability.customGrafana.enabled=true",
						},
					})
					Expect(err).NotTo(HaveOccurred())

					testManifest.Expect(grafanaDeployment.Kind, grafanaDeployment.Namespace, grafanaDeployment.Name).To(BeNil())
				})
			})
		})

		Context("external auth server", func() {

			var expectedDeployment *appsv1.Deployment

			BeforeEach(func() {
				labels = map[string]string{
					"app":  "gloo",
					"gloo": "extauth",
				}
				selector = map[string]string{
					"gloo": "extauth",
				}

				rb := ResourceBuilder{
					Namespace: namespace,
					Name:      "extauth",
					Labels:    labels,
				}
				expectedDeployment = rb.GetDeploymentAppsv1()

				expectedDeployment.Spec.Replicas = aws.Int32(1)
				expectedDeployment.Spec.Template.Spec.Containers = []v1.Container{
					{
						Name:            "extauth",
						Image:           "quay.io/solo-io/extauth-ee:dev",
						ImagePullPolicy: "Always",
						Env: []v1.EnvVar{
							{
								Name: "POD_NAMESPACE",
								ValueFrom: &v1.EnvVarSource{
									FieldRef: &v1.ObjectFieldSelector{
										FieldPath: "metadata.namespace",
									},
								},
							},
							{
								Name:  "SERVICE_NAME",
								Value: "ext-auth",
							},
							{
								Name:  "GLOO_ADDRESS",
								Value: "gloo:9977",
							},
							{
								Name: "SIGNING_KEY",
								ValueFrom: &v1.EnvVarSource{
									SecretKeyRef: &v1.SecretKeySelector{
										LocalObjectReference: v1.LocalObjectReference{
											Name: "extauth-signing-key",
										},
										Key: "signing-key",
									},
								},
							},
							{
								Name:  "SERVER_PORT",
								Value: "8083",
							},
							{
								Name:  "USER_ID_HEADER",
								Value: "x-user-id",
							},
							statsEnvVar,
						},
						ReadinessProbe: &v1.Probe{
							Handler: v1.Handler{
								Exec: &v1.ExecAction{
									Command: []string{"/bin/sh", "-c", "nc -z localhost 8083"},
								},
							},
							InitialDelaySeconds: 1,
							FailureThreshold:    3,
							SuccessThreshold:    1,
						},
						Resources: v1.ResourceRequirements{},
					},
				}
				expectedDeployment.Spec.Strategy = appsv1.DeploymentStrategy{}
				expectedDeployment.Spec.Selector.MatchLabels = selector
				expectedDeployment.Spec.Template.ObjectMeta.Labels = selector
				expectedDeployment.Spec.Template.ObjectMeta.Annotations = normalPromAnnotations

				expectedDeployment.Spec.Template.Spec.Affinity = &v1.Affinity{
					PodAffinity: &v1.PodAffinity{
						PreferredDuringSchedulingIgnoredDuringExecution: []v1.WeightedPodAffinityTerm{
							{
								Weight: 100,
								PodAffinityTerm: v1.PodAffinityTerm{
									LabelSelector: &k8s.LabelSelector{
										MatchLabels: map[string]string{
											"gloo": "gateway-proxy",
										},
									},
									TopologyKey: "kubernetes.io/hostname",
								},
							},
						},
					}}
			})

			It("produces expected default deployment", func() {
				testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{})
				Expect(err).NotTo(HaveOccurred())

				actualDeployment := testManifest.SelectResources(func(unstructured *unstructured.Unstructured) bool {
					return unstructured.GetKind() == "Deployment" && unstructured.GetLabels()["gloo"] == "extauth"
				})

				actualDeployment.ExpectDeploymentAppsV1(expectedDeployment)
			})

			It("allows setting the number of replicas for the deployment", func() {
				testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
					valuesArgs: []string{"global.extensions.extAuth.deployment.replicas=3"},
				})
				Expect(err).NotTo(HaveOccurred())

				actualDeployment := testManifest.SelectResources(func(unstructured *unstructured.Unstructured) bool {
					return unstructured.GetKind() == "Deployment" && unstructured.GetLabels()["gloo"] == "extauth"
				})

				expectedDeployment.Spec.Replicas = aws.Int32(3)
				actualDeployment.ExpectDeploymentAppsV1(expectedDeployment)
			})

			It("allows multiple extauth plugins", func() {
				helmOverrideFileContents := `
global:
  extensions:
    extAuth:
      plugins:
        first-plugin:
          image:
            repository: ext-auth-plugins
            registry: quay.io/solo-io
            pullPolicy: IfNotPresent
            tag: 1.2.3
        second-plugin:
          image:
            repository: foo
            registry: bar
            pullPolicy: IfNotPresent
            tag: 1.2.3`
				helmOverrideFile := "helm-override.yaml"
				err := ioutil.WriteFile(helmOverrideFile, []byte(helmOverrideFileContents), 0644)
				Expect(err).NotTo(HaveOccurred())
				defer os.Remove(helmOverrideFile)
				testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
					valuesFile: helmOverrideFile,
				})
				Expect(err).NotTo(HaveOccurred())

				actualDeployment := testManifest.SelectResources(func(unstructured *unstructured.Unstructured) bool {
					return unstructured.GetKind() == "Deployment" && unstructured.GetLabels()["gloo"] == "extauth"
				})

				authPluginVolumeMount := []v1.VolumeMount{
					v1.VolumeMount{
						Name:      "auth-plugins",
						MountPath: "/auth-plugins",
					},
				}
				expectedDeployment.Spec.Template.Spec.InitContainers = []v1.Container{
					v1.Container{
						Name:            "plugin-first-plugin",
						Image:           "quay.io/solo-io/ext-auth-plugins:1.2.3",
						ImagePullPolicy: v1.PullIfNotPresent,
						VolumeMounts:    authPluginVolumeMount,
					},
					v1.Container{
						Name:            "plugin-second-plugin",
						Image:           "bar/foo:1.2.3",
						ImagePullPolicy: v1.PullIfNotPresent,
						VolumeMounts:    authPluginVolumeMount,
					},
				}
				expectedDeployment.Spec.Template.Spec.Volumes = []v1.Volume{
					v1.Volume{
						Name: "auth-plugins",
						VolumeSource: v1.VolumeSource{
							EmptyDir: &v1.EmptyDirVolumeSource{},
						},
					},
				}
				for i, _ := range expectedDeployment.Spec.Template.Spec.Containers {
					expectedDeployment.Spec.Template.Spec.Containers[i].VolumeMounts =
						append(expectedDeployment.Spec.Template.Spec.Containers[i].VolumeMounts, authPluginVolumeMount...)
				}
				actualDeployment.ExpectDeploymentAppsV1(expectedDeployment)
			})
		})

		Context("gateway", func() {
			BeforeEach(func() {
				labels = map[string]string{
					"gloo":             "gateway-proxy",
					"gateway-proxy-id": defaults.GatewayProxyName,
					"app":              "gloo",
				}
				selector = map[string]string{
					"gateway-proxy": "live",
				}
			})

			Context("gateway-proxy deployment", func() {
				var (
					gatewayProxyDeployment *appsv1.Deployment
				)

				includeStatConfig := func() {
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-stats"] = "/stats"
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-ready"] = "/ready"
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-config_dump"] = "/config_dump"
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-port"] = "8082"
				}

				BeforeEach(func() {
					selector = map[string]string{
						"gloo":             "gateway-proxy",
						"gateway-proxy-id": defaults.GatewayProxyName,
					}
					podLabels := map[string]string{
						"gloo":             "gateway-proxy",
						"gateway-proxy-id": defaults.GatewayProxyName,
						"gateway-proxy":    "live",
					}
					podAnnotations := map[string]string{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   "8081",
						"prometheus.io/scrape": "true",
					}
					podname := v1.EnvVar{
						Name: "POD_NAME",
						ValueFrom: &v1.EnvVarSource{
							FieldRef: &v1.ObjectFieldSelector{
								FieldPath: "metadata.name",
							},
						},
					}

					container := GetQuayContainerSpec("gloo-ee-envoy-wrapper", version, GetPodNamespaceEnvVar(), podname)
					container.Name = defaults.GatewayProxyName
					container.Args = []string{"--disable-hot-restart"}

					rb := ResourceBuilder{
						Namespace:  namespace,
						Name:       defaults.GatewayProxyName,
						Labels:     labels,
						Containers: []ContainerSpec{container},
					}
					deploy := rb.GetDeploymentAppsv1()
					deploy.Spec.Selector = &k8s.LabelSelector{
						MatchLabels: selector,
					}
					deploy.Spec.Template.ObjectMeta.Labels = podLabels
					deploy.Spec.Template.ObjectMeta.Annotations = podAnnotations
					deploy.Spec.Template.Spec.Volumes = []v1.Volume{{
						Name: "envoy-config",
						VolumeSource: v1.VolumeSource{
							ConfigMap: &v1.ConfigMapVolumeSource{
								LocalObjectReference: v1.LocalObjectReference{
									Name: "gateway-proxy-envoy-config",
								},
							},
						},
					}}
					deploy.Spec.Template.Spec.Containers[0].ImagePullPolicy = getPullPolicy()
					deploy.Spec.Template.Spec.Containers[0].Ports = []v1.ContainerPort{
						{Name: "http", ContainerPort: 8080, Protocol: "TCP"},
						{Name: "https", ContainerPort: 8443, Protocol: "TCP"},
					}
					deploy.Spec.Template.Spec.Containers[0].VolumeMounts = []v1.VolumeMount{{
						Name:      "envoy-config",
						ReadOnly:  false,
						MountPath: "/etc/envoy",
						SubPath:   "",
					}}
					truez := true
					falsez := false
					deploy.Spec.Template.Spec.Containers[0].SecurityContext = &v1.SecurityContext{
						Capabilities: &v1.Capabilities{
							Add:  []v1.Capability{"NET_BIND_SERVICE"},
							Drop: []v1.Capability{"ALL"},
						},
						ReadOnlyRootFilesystem:   &truez,
						AllowPrivilegeEscalation: &falsez,
					}

					deploy.Spec.Template.Spec.ServiceAccountName = "gateway-proxy"

					gatewayProxyDeployment = deploy
				})

				It("creates a deployment without envoy config annotations", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{})
					Expect(err).NotTo(HaveOccurred())
					testManifest.ExpectDeploymentAppsV1(gatewayProxyDeployment)
				})

				It("creates a deployment with envoy config annotations", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
						valuesArgs: []string{
							"gloo.gatewayProxies.gatewayProxy.readConfig=true",
						},
					})
					Expect(err).NotTo(HaveOccurred())
					includeStatConfig()
					testManifest.ExpectDeploymentAppsV1(gatewayProxyDeployment)
				})

				It("creates a deployment without extauth sidecar", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{})
					Expect(err).NotTo(HaveOccurred())
					testManifest.ExpectDeploymentAppsV1(gatewayProxyDeployment)
				})

				It("creates a deployment with extauth sidecar", func() {
					testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
						valuesArgs: []string{
							"global.extensions.extAuth.envoySidecar=true",
						},
					})
					Expect(err).NotTo(HaveOccurred())

					gatewayProxyDeployment.Spec.Template.Spec.Volumes = append(
						gatewayProxyDeployment.Spec.Template.Spec.Volumes,
						v1.Volume{
							Name: "shared-data",
							VolumeSource: v1.VolumeSource{
								EmptyDir: &v1.EmptyDirVolumeSource{},
							},
						})

					gatewayProxyDeployment.Spec.Template.Spec.Containers[0].VolumeMounts = append(
						gatewayProxyDeployment.Spec.Template.Spec.Containers[0].VolumeMounts,
						v1.VolumeMount{
							Name:      "shared-data",
							MountPath: "/usr/share/shared-data",
						})

					gatewayProxyDeployment.Spec.Template.Spec.Containers = append(
						gatewayProxyDeployment.Spec.Template.Spec.Containers,
						v1.Container{
							Name:            "extauth",
							Image:           "quay.io/solo-io/extauth-ee:dev",
							Ports:           nil,
							ImagePullPolicy: getPullPolicy(),
							Env: []v1.EnvVar{
								{
									Name: "POD_NAMESPACE",
									ValueFrom: &v1.EnvVarSource{
										FieldRef: &v1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								{
									Name:  "SERVICE_NAME",
									Value: "ext-auth",
								},
								{
									Name:  "GLOO_ADDRESS",
									Value: "gloo:9977",
								},
								{
									Name: "SIGNING_KEY",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											LocalObjectReference: v1.LocalObjectReference{
												Name: "extauth-signing-key",
											},
											Key: "signing-key",
										},
									},
								},
								{
									Name:  "SERVER_PORT",
									Value: "8083",
								},
								{
									Name:  "UDS_ADDR",
									Value: "/usr/share/shared-data/.sock",
								},
								{
									Name:  "USER_ID_HEADER",
									Value: "x-user-id",
								},
								{
									Name:  "START_STATS_SERVER",
									Value: "true",
								},
							},
							VolumeMounts: []v1.VolumeMount{
								{
									Name:      "shared-data",
									MountPath: "/usr/share/shared-data",
								},
							},
						})

					testManifest.ExpectDeploymentAppsV1(gatewayProxyDeployment)
				})

				Context("apiserver deployment", func() {
					var expectedDeployment *appsv1.Deployment

					BeforeEach(func() {
						labels = map[string]string{
							"gloo": "apiserver-ui",
							"app":  "gloo",
						}
						selector = map[string]string{
							"app":  "gloo",
							"gloo": "apiserver-ui",
						}
						grpcPortEnvVar := v1.EnvVar{
							Name:  "GRPC_PORT",
							Value: "10101",
						}
						noAuthEnvVar := v1.EnvVar{
							Name:  "NO_AUTH",
							Value: "1",
						}
						licenseEnvVar := v1.EnvVar{
							Name: "GLOO_LICENSE_KEY",
							ValueFrom: &v1.EnvVarSource{
								SecretKeyRef: &v1.SecretKeySelector{
									LocalObjectReference: v1.LocalObjectReference{
										Name: "license",
									},
									Key: "license-key",
								},
							},
						}
						uiContainer := v1.Container{
							Name:            "apiserver-ui",
							Image:           "quay.io/solo-io/grpcserver-ui:" + version,
							ImagePullPolicy: v1.PullAlways,
							VolumeMounts: []v1.VolumeMount{
								{Name: "empty-cache", MountPath: "/var/cache/nginx"},
								{Name: "empty-run", MountPath: "/var/run"},
							},
							Ports: []v1.ContainerPort{{Name: "static", ContainerPort: 8080, Protocol: v1.ProtocolTCP}},
						}
						grpcServerContainer := v1.Container{
							Name:            "apiserver",
							Image:           "quay.io/solo-io/grpcserver-ee:" + version,
							ImagePullPolicy: v1.PullAlways,
							Ports:           []v1.ContainerPort{{Name: "grpcport", ContainerPort: 10101, Protocol: v1.ProtocolTCP}},
							Env: []v1.EnvVar{
								GetPodNamespaceEnvVar(),
								grpcPortEnvVar,
								statsEnvVar,
								noAuthEnvVar,
								licenseEnvVar,
							},
						}
						envoyContainer := v1.Container{
							Name:            "gloo-grpcserver-envoy",
							Image:           "quay.io/solo-io/grpcserver-envoy:" + version,
							ImagePullPolicy: v1.PullAlways,
							ReadinessProbe: &v1.Probe{
								Handler: v1.Handler{HTTPGet: &v1.HTTPGetAction{
									Path: "/",
									Port: intstr.IntOrString{IntVal: 8080},
								}},
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
							},
						}

						rb := ResourceBuilder{
							Namespace: namespace,
							Name:      "api-server",
							Labels:    labels,
						}
						expectedDeployment = rb.GetDeploymentAppsv1()
						expectedDeployment.Spec.Selector.MatchLabels = selector
						expectedDeployment.Spec.Template.ObjectMeta.Labels = selector
						expectedDeployment.Spec.Template.Spec.Volumes = []v1.Volume{
							{Name: "empty-cache", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}},
							{Name: "empty-run", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}},
						}
						expectedDeployment.Spec.Template.Spec.Containers = []v1.Container{uiContainer, grpcServerContainer, envoyContainer}
						expectedDeployment.Spec.Template.Spec.ServiceAccountName = "apiserver-ui"
						expectedDeployment.Spec.Template.ObjectMeta.Annotations = normalPromAnnotations
					})

					It("is there by default", func() {
						testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{})
						Expect(err).NotTo(HaveOccurred())
						testManifest.ExpectDeploymentAppsV1(expectedDeployment)
					})

					It("correctly sets resource limits", func() {
						testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
							valuesArgs: []string{
								"apiServer.deployment.ui.resources.limits.cpu=300m",
								"apiServer.deployment.ui.resources.limits.memory=300Mi",
								"apiServer.deployment.ui.resources.requests.cpu=30m",
								"apiServer.deployment.ui.resources.requests.memory=30Mi",
								"apiServer.deployment.envoy.resources.limits.cpu=100m",
								"apiServer.deployment.envoy.resources.limits.memory=100Mi",
								"apiServer.deployment.envoy.resources.requests.cpu=10m",
								"apiServer.deployment.envoy.resources.requests.memory=10Mi",
								"apiServer.deployment.server.resources.limits.cpu=200m",
								"apiServer.deployment.server.resources.limits.memory=200Mi",
								"apiServer.deployment.server.resources.requests.cpu=20m",
								"apiServer.deployment.server.resources.requests.memory=20Mi",
							},
						})
						Expect(err).NotTo(HaveOccurred())

						// UI
						expectedDeployment.Spec.Template.Spec.Containers[0].Resources = v1.ResourceRequirements{
							Limits: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("300m"),
								v1.ResourceMemory: resource.MustParse("300Mi"),
							},
							Requests: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("30m"),
								v1.ResourceMemory: resource.MustParse("30Mi"),
							},
						}

						// Server
						expectedDeployment.Spec.Template.Spec.Containers[1].Resources = v1.ResourceRequirements{
							Limits: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("200m"),
								v1.ResourceMemory: resource.MustParse("200Mi"),
							},
							Requests: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("20m"),
								v1.ResourceMemory: resource.MustParse("20Mi"),
							},
						}

						// Envoy
						expectedDeployment.Spec.Template.Spec.Containers[2].Resources = v1.ResourceRequirements{
							Limits: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("100m"),
								v1.ResourceMemory: resource.MustParse("100Mi"),
							},
							Requests: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("10m"),
								v1.ResourceMemory: resource.MustParse("10Mi"),
							},
						}

						testManifest.ExpectDeploymentAppsV1(expectedDeployment)
					})
				})
			})
		})

		Context("gloo mtls settings", func() {
			var (
				glooMtlsSecretVolume = v1.Volume{
					Name: "gloo-mtls-certs",
					VolumeSource: v1.VolumeSource{
						Secret: &v1.SecretVolumeSource{
							SecretName:  "gloo-mtls-certs",
							Items:       nil,
							DefaultMode: proto.Int(420),
						},
					},
				}

				haveEnvoySidecar = func(containers []v1.Container) bool {
					for _, c := range containers {
						if c.Name == "envoy-sidecar" {
							return true
						}
					}
					return false
				}

				haveSdsSidecar = func(containers []v1.Container) bool {
					for _, c := range containers {
						if c.Name == "sds" {
							return true
						}
					}
					return false
				}

				haveEnvVariable = func(containers []v1.Container, containerName, env, value string) bool {
					for _, c := range containers {
						if c.Name == containerName {
							Expect(c.Env).To(ContainElement(v1.EnvVar{Name: env, Value: value}))
							return true
						}
					}
					return false
				}
			)

			It("should add or change the correct components in the resulting helm chart", func() {
				testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
					valuesArgs: []string{"global.glooMtls.enabled=true"},
				})
				Expect(err).NotTo(HaveOccurred())

				foundGlooMtlsCertgenJob := false
				testManifest.SelectResources(func(resource *unstructured.Unstructured) bool {
					return resource.GetKind() == "Job"
				}).ExpectAll(func(job *unstructured.Unstructured) {
					jobObject, err := kuberesource.ConvertUnstructured(job)
					Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Job %+v should be able to convert from unstructured", job))
					structuredDeployment, ok := jobObject.(*jobsv1.Job)
					Expect(ok).To(BeTrue(), fmt.Sprintf("Job %+v should be able to cast to a structured job", job))

					if structuredDeployment.GetName() == "gloo-mtls-certgen" {
						foundGlooMtlsCertgenJob = true
					}
				})
				Expect(foundGlooMtlsCertgenJob).To(BeTrue(), "Did not find the gloo-mtls-certgen job")

				testManifest.SelectResources(func(resource *unstructured.Unstructured) bool {
					return resource.GetKind() == "Deployment"
				}).ExpectAll(func(deployment *unstructured.Unstructured) {
					deploymentObject, err := kuberesource.ConvertUnstructured(deployment)
					Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Deployment %+v should be able to convert from unstructured", deployment))
					structuredDeployment, ok := deploymentObject.(*appsv1.Deployment)
					Expect(ok).To(BeTrue(), fmt.Sprintf("Deployment %+v should be able to cast to a structured deployment", deployment))

					if structuredDeployment.GetName() == "gloo" {
						Ω(haveEnvoySidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Ω(haveSdsSidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Expect(structuredDeployment.Spec.Template.Spec.Volumes).To(ContainElement(glooMtlsSecretVolume))
					}

					if structuredDeployment.GetName() == "gateway-proxy" {
						Ω(haveSdsSidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Expect(structuredDeployment.Spec.Template.Spec.Volumes).To(ContainElement(glooMtlsSecretVolume))
					}

					// should add envoy, sds sidecars to the Extauth and Rate-Limit Deployment
					if structuredDeployment.GetName() == "rate-limit" {
						Ω(haveEnvoySidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Ω(haveSdsSidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Ω(haveEnvVariable(structuredDeployment.Spec.Template.Spec.Containers,
							"rate-limit", "GLOO_ADDRESS", "127.0.0.1:9955")).To(BeTrue())
						Expect(structuredDeployment.Spec.Template.Spec.Volumes).To(ContainElement(glooMtlsSecretVolume))
					}

					if structuredDeployment.GetName() == "extauth" {
						Ω(haveEnvoySidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Ω(haveSdsSidecar(structuredDeployment.Spec.Template.Spec.Containers)).To(BeTrue())
						Ω(haveEnvVariable(structuredDeployment.Spec.Template.Spec.Containers,
							"extauth", "GLOO_ADDRESS", "127.0.0.1:9955")).To(BeTrue())
						Ω(haveEnvVariable(structuredDeployment.Spec.Template.Spec.Containers,
							"extauth", "SERVER_PORT", "8084")).To(BeTrue())
						Expect(structuredDeployment.Spec.Template.Spec.Volumes).To(ContainElement(glooMtlsSecretVolume))
					}
				})
			})

			It("should add an additional listener to the gateway-proxy-envoy-config for extauth sidecar", func() {
				testManifest, err := BuildTestManifest(install.GlooEnterpriseChartName, namespace, helmValues{
					valuesArgs: []string{"global.glooMtls.enabled=true,global.extensions.extAuth.envoySidecar=true"},
				})
				Expect(err).NotTo(HaveOccurred())

				testManifest.SelectResources(func(resource *unstructured.Unstructured) bool {
					return resource.GetKind() == "ConfigMap"
				}).ExpectAll(func(configMap *unstructured.Unstructured) {
					configMapObject, err := kuberesource.ConvertUnstructured(configMap)
					Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("ConfigMap %+v should be able to convert from unstructured", configMap))
					structuredConfigMap, ok := configMapObject.(*v1.ConfigMap)
					Expect(ok).To(BeTrue(), fmt.Sprintf("ConfigMap %+v should be able to cast to a structured config map", configMap))

					if structuredConfigMap.GetName() == "gateway-proxy-envoy-config" {
						expectedGlooMtlsListener := "    - name: gloo_xds_mtls_listener"
						Expect(structuredConfigMap.Data["envoy.yaml"]).To(ContainSubstring(expectedGlooMtlsListener))
					}
				})
			})
		})
	})

	Describe("gloo with read-only ui helm tests", func() {
		var (
			labels           map[string]string
			selector         map[string]string
			manifestYaml     string
			glooOsVersion    string
			glooOsPullPolicy v1.PullPolicy
		)

		BeforeEach(func() {

			var err error
			var glooEGenerationFiles = &generate.GenerationFiles{
				Artifact:             generate.GlooE,
				RequirementsTemplate: "../../install/helm/gloo-ee/requirements-template.yaml",
			}
			var glooOsWithReadOnlyUiGenerationFiles = &generate.GenerationFiles{
				Artifact:             generate.GlooWithRoUi,
				RequirementsTemplate: "../../install/helm/gloo-os-with-ui/requirements-template.yaml",
			}
			glooOsVersion, err = generate.GetGlooOsVersion(glooEGenerationFiles, glooOsWithReadOnlyUiGenerationFiles)
			Expect(err).NotTo(HaveOccurred())
			glooOsPullPolicy = v1.PullAlways

			version = os.Getenv("TAGGED_VERSION")
			if version == "" {
				version = "dev"
			} else {
				version = version[1:]
			}
			manifestYaml = ""
		})

		AfterEach(func() {
			if manifestYaml != "" {
				err := os.Remove(manifestYaml)
				Expect(err).ToNot(HaveOccurred())
			}
		})

		Context("gateway", func() {
			BeforeEach(func() {
				labels = map[string]string{
					"gloo":             "gateway-proxy",
					"gateway-proxy-id": defaults.GatewayProxyName,
					"app":              "gloo",
				}
				selector = map[string]string{
					"gateway-proxy": "live",
				}
			})

			Context("gateway-proxy deployment", func() {
				var (
					gatewayProxyDeployment *appsv1.Deployment
				)

				includeStatConfig := func() {
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-stats"] = "/stats"
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-ready"] = "/ready"
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-config_dump"] = "/config_dump"
					gatewayProxyDeployment.Spec.Template.ObjectMeta.Annotations["readconfig-port"] = "8082"
				}

				BeforeEach(func() {
					selector = map[string]string{
						"gloo":             "gateway-proxy",
						"gateway-proxy-id": defaults.GatewayProxyName,
					}
					podLabels := map[string]string{
						"gloo":             "gateway-proxy",
						"gateway-proxy-id": defaults.GatewayProxyName,
						"gateway-proxy":    "live",
					}
					podAnnotations := map[string]string{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   "8081",
						"prometheus.io/scrape": "true",
					}
					podname := v1.EnvVar{
						Name: "POD_NAME",
						ValueFrom: &v1.EnvVarSource{
							FieldRef: &v1.ObjectFieldSelector{
								FieldPath: "metadata.name",
							},
						},
					}

					container := GetQuayContainerSpec("gloo-envoy-wrapper", glooOsVersion, GetPodNamespaceEnvVar(), podname)
					container.Name = defaults.GatewayProxyName
					container.Args = []string{"--disable-hot-restart"}

					rb := ResourceBuilder{
						Namespace:  namespace,
						Name:       defaults.GatewayProxyName,
						Labels:     labels,
						Containers: []ContainerSpec{container},
					}
					deploy := rb.GetDeploymentAppsv1()
					deploy.Spec.Selector = &k8s.LabelSelector{
						MatchLabels: selector,
					}
					deploy.Spec.Template.ObjectMeta.Labels = podLabels
					deploy.Spec.Template.ObjectMeta.Annotations = podAnnotations
					deploy.Spec.Template.Spec.Volumes = []v1.Volume{{
						Name: "envoy-config",
						VolumeSource: v1.VolumeSource{
							ConfigMap: &v1.ConfigMapVolumeSource{
								LocalObjectReference: v1.LocalObjectReference{
									Name: "gateway-proxy-envoy-config",
								},
							},
						},
					}}
					deploy.Spec.Template.Spec.Containers[0].ImagePullPolicy = glooOsPullPolicy
					deploy.Spec.Template.Spec.Containers[0].Ports = []v1.ContainerPort{
						{Name: "http", ContainerPort: 8080, Protocol: "TCP"},
						{Name: "https", ContainerPort: 8443, Protocol: "TCP"},
					}
					deploy.Spec.Template.Spec.Containers[0].VolumeMounts = []v1.VolumeMount{{
						Name:      "envoy-config",
						ReadOnly:  false,
						MountPath: "/etc/envoy",
						SubPath:   "",
					}}
					truez := true
					falsez := false
					deploy.Spec.Template.Spec.Containers[0].SecurityContext = &v1.SecurityContext{
						Capabilities: &v1.Capabilities{
							Add:  []v1.Capability{"NET_BIND_SERVICE"},
							Drop: []v1.Capability{"ALL"},
						},
						ReadOnlyRootFilesystem:   &truez,
						AllowPrivilegeEscalation: &falsez,
					}

					deploy.Spec.Template.Spec.ServiceAccountName = "gateway-proxy"

					gatewayProxyDeployment = deploy
				})

				It("creates a deployment without envoy config annotations", func() {
					testManifest, err := BuildTestManifest(install.GlooOsWithUiChartName, namespace, helmValues{})
					Expect(err).NotTo(HaveOccurred())
					testManifest.ExpectDeploymentAppsV1(gatewayProxyDeployment)
				})

				It("creates a deployment with envoy config annotations", func() {
					testManifest, err := BuildTestManifest(install.GlooOsWithUiChartName, namespace, helmValues{
						valuesArgs: []string{"gloo.gatewayProxies.gatewayProxy.readConfig=true"},
					})
					Expect(err).NotTo(HaveOccurred())
					includeStatConfig()
					testManifest.ExpectDeploymentAppsV1(gatewayProxyDeployment)
				})

				Context("apiserver deployment", func() {
					var deploy *appsv1.Deployment

					BeforeEach(func() {
						labels = map[string]string{
							"gloo": "apiserver-ui",
							"app":  "gloo",
						}
						selector = map[string]string{
							"app":  "gloo",
							"gloo": "apiserver-ui",
						}
						grpcPortEnvVar := v1.EnvVar{
							Name:  "GRPC_PORT",
							Value: "10101",
						}
						noAuthEnvVar := v1.EnvVar{
							Name:  "NO_AUTH",
							Value: "1",
						}
						uiContainer := v1.Container{
							Name:            "apiserver-ui",
							Image:           "quay.io/solo-io/grpcserver-ui:" + version,
							ImagePullPolicy: v1.PullAlways,
							VolumeMounts: []v1.VolumeMount{
								{Name: "empty-cache", MountPath: "/var/cache/nginx"},
								{Name: "empty-run", MountPath: "/var/run"},
							},
							Ports: []v1.ContainerPort{{Name: "static", ContainerPort: 8080, Protocol: v1.ProtocolTCP}},
						}
						grpcServerContainer := v1.Container{
							Name:            "apiserver",
							Image:           "quay.io/solo-io/grpcserver-ee:" + version,
							ImagePullPolicy: v1.PullAlways,
							Ports:           []v1.ContainerPort{{Name: "grpcport", ContainerPort: 10101, Protocol: v1.ProtocolTCP}},
							Env: []v1.EnvVar{
								GetPodNamespaceEnvVar(),
								grpcPortEnvVar,
								statsEnvVar,
								noAuthEnvVar,
							},
						}
						envoyContainer := v1.Container{
							Name:            "gloo-grpcserver-envoy",
							Image:           "quay.io/solo-io/grpcserver-envoy:" + version,
							ImagePullPolicy: v1.PullAlways,
							ReadinessProbe: &v1.Probe{
								Handler: v1.Handler{HTTPGet: &v1.HTTPGetAction{
									Path: "/",
									Port: intstr.IntOrString{IntVal: 8080},
								}},
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
							},
						}

						rb := ResourceBuilder{
							Namespace: namespace,
							Name:      "api-server",
							Labels:    labels,
						}
						deploy = rb.GetDeploymentAppsv1()
						deploy.Spec.Selector.MatchLabels = selector
						deploy.Spec.Template.ObjectMeta.Labels = selector
						deploy.Spec.Template.Spec.Volumes = []v1.Volume{
							{Name: "empty-cache", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}},
							{Name: "empty-run", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}},
						}
						deploy.Spec.Template.Spec.Containers = []v1.Container{uiContainer, grpcServerContainer, envoyContainer}
						deploy.Spec.Template.Spec.ServiceAccountName = "apiserver-ui"
						deploy.Spec.Template.ObjectMeta.Annotations = normalPromAnnotations
					})

					It("is there by default", func() {
						testManifest, err := BuildTestManifest(install.GlooOsWithUiChartName, namespace, helmValues{})
						Expect(err).NotTo(HaveOccurred())
						testManifest.ExpectDeploymentAppsV1(deploy)
					})
				})
			})

		})
	})
})
