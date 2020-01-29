package hook

import (
	common2 "github.com/argoproj/argo-cd/engine/pkg/utils/kube/sync/common"
	resource2 "github.com/argoproj/argo-cd/engine/pkg/utils/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/argoproj/argo-cd/common"
	helmhook "github.com/argoproj/argo-cd/engine/pkg/utils/kube/sync/hook/helm"
)

func DeletePolicies(obj *unstructured.Unstructured) []common2.HookDeletePolicy {
	var policies []common2.HookDeletePolicy
	for _, text := range resource2.GetAnnotationCSVs(obj, common.AnnotationKeyHookDeletePolicy) {
		p, ok := common2.NewHookDeletePolicy(text)
		if ok {
			policies = append(policies, p)
		}
	}
	for _, p := range helmhook.DeletePolicies(obj) {
		policies = append(policies, p.DeletePolicy())
	}
	if len(policies) == 0 {
		policies = append(policies, common2.HookDeletePolicyBeforeHookCreation)
	}
	return policies
}
