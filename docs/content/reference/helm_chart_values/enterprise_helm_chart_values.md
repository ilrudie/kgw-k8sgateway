---
title: "Enterprise Gloo Edge"
description: Listing of Helm values for the Enterprise Gloo Edge chart
weight: 30
---

The table below describes all the values that you can override in your custom values file when working with the Helm 
chart for Enterprise Gloo Edge. More information on using a Helm chart to install the Gloo Edge can be found 
[here]({{% versioned_link_path fromRoot="/installation/gateway/kubernetes/#installing-on-kubernetes-with-helm" %}}).

{{% notice warning %}}
Because the Gloo Edge Enterprise Helm chart uses the open source chart as a dependency, 
you must add the `gloo.` prefix to all open source Gloo Edge chart values.

This applies to all values except for `global.*`. For example, when you install Gloo Edge Enterprise, `ingress.deployment.nodeName` must be changed to `gloo.ingress.deployment.nodeName`, but `global.glooRbac.create` remains unchanged.
{{% /notice %}}

{{< readfile file="static/content/glooe-values.docgen" markdown="true" >}}

## Helm Chart KubeResourceOverrides

Most changes that need to be made to the default helm chart are supported by the helm values above.
However, there may be a case where the helm values do not cover a necessary change.
Helm values which expose the kubernetes API (ie Service, Deployment) now include a `KubeResourceOverride` field.
The following example uses `KubeResourceOverride` to add
labels to a deployment, which is not explicitly implemented in the helm chart:

Gloo deployment:
```yaml
apiVersion: v1
kind: Deployment
metadata:
  labels:
    gloo: gloo
  name: gloo
```

We want to add the `resource-owner: infra-team` label to the Deployment. We can do so by specifying the `KubeResourceOverride` in the helm values file:
```yaml
gloo:
  deployment:
    kubeResourceOverride:
      metadata:
        labels:
          resource-owner: infra-team
```

Yaml under the `kubeResourceOverride` is merged in to the deployment yaml, to create the resulting kube resource:
```yaml
apiVersion: v1
kind: Deployment
metadata:
  labels:
    gloo: gloo
    resource-owner: infra-team
```

{{% notice note %}}
`kubeResourceOverride` does not support merging in lists, and a list in the override will replace the list in the original resource.
{{% /notice %}}

##### KubeResourceOverride vs Kustomize

[Kustomize](https://kustomize.io/) is another solution for patching resources generated by helm. The following table enumerates differences between the offered solutions.

| Kustomize                                                                                          | Helm KubeResourceOverride                                                                                                          |
|----------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
| - Finer control over merges, e.g. merging lists <br>- Requires additional files (overlays)<br>- Unsupported in older versions of flux helm operator  | - Can be done only with helm, does not require CI/CD pipeline changes <br>- Specified through values.yaml file or command line arguments |