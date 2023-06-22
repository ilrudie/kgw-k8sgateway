package checker

import (
	"context"
	"fmt"

	"github.com/solo-io/go-utils/contextutils"
	sk_sets "github.com/solo-io/skv2/contrib/pkg/sets/v2"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	corev1 "k8s.io/api/core/v1"
)

// Get a summary of pods in the given namespace and cluster. To bypass the cluster check (e.g. for single-cluster
// use, pass in "" for the cluster.
func GetPodsSummary(ctx context.Context, set sk_sets.ResourceSet[*corev1.Pod], namespace, cluster string) *Summary {
	summary := &Summary{}
	for _, podIter := range set.List() {
		pod := podIter
		if (cluster != "" && ezkube.GetClusterName(pod) != cluster) || pod.Namespace != namespace {
			continue
		}

		summary.Total += 1
		for _, condition := range pod.Status.Conditions {
			var message string

			if condition.Message != "" {
				message = fmt.Sprintf(" Message: %s", condition.Message)
			}

			// if condition is not met and the pod is not completed
			conditionNotMet := condition.Status != corev1.ConditionTrue && condition.Reason != "PodCompleted"

			// possible condition types listed at https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-conditions
			switch condition.Type {
			case corev1.PodScheduled:
				if conditionNotMet {
					summary.Errors = append(
						summary.Errors,
						&ResourceReport{
							Ref: &v1.ObjectRef{
								Namespace: pod.Namespace,
								Name:      pod.Name,
							},
							Message: fmt.Sprintf("Pod is not yet scheduled!%s\n", message),
						},
					)
				}
			case corev1.PodReady:
				if conditionNotMet {
					summary.Errors = append(
						summary.Errors,
						&ResourceReport{
							Ref: &v1.ObjectRef{
								Namespace: pod.Namespace,
								Name:      pod.Name,
							},
							Message: fmt.Sprintf("Pod is not ready!%s\n", message),
						},
					)
				}
			case corev1.PodInitialized:
				if conditionNotMet {
					summary.Errors = append(
						summary.Errors,
						&ResourceReport{
							Ref: &v1.ObjectRef{
								Namespace: pod.Namespace,
								Name:      pod.Name,
							},
							Message: fmt.Sprintf("Pod is not yet initialized!%s\n", message),
						},
					)
				}
			case corev1.PodReasonUnschedulable:
				if conditionNotMet {
					summary.Errors = append(
						summary.Errors,
						&ResourceReport{
							Ref: &v1.ObjectRef{
								Namespace: pod.Namespace,
								Name:      pod.Name,
							},
							Message: fmt.Sprintf("Pod is unschedulable!%s\n", message),
						},
					)
				}
			case corev1.ContainersReady:
				if conditionNotMet {
					summary.Errors = append(
						summary.Errors,
						&ResourceReport{
							Ref: &v1.ObjectRef{
								Namespace: pod.Namespace,
								Name:      pod.Name,
							},
							Message: fmt.Sprintf("Not all containers are ready!%s\n", message),
						},
					)
				}
			default:
				contextutils.LoggerFrom(ctx).Debugw("Note: Unhandled pod condition %s", condition.Type)
			}
		}
	}

	SortLists(summary)
	return summary
}