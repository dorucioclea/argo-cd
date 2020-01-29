package helm

import (
	"github.com/argoproj/argo-cd/engine/pkg/utils/kube/sync/common"
	resource2 "github.com/argoproj/argo-cd/engine/pkg/utils/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type Type string

const (
	PreInstall  Type = "pre-install"
	PreUpgrade  Type = "pre-upgrade"
	PostUpgrade Type = "post-upgrade"
	PostInstall Type = "post-install"
)

func NewType(t string) (Type, bool) {
	return Type(t),
		t == string(PreInstall) ||
			t == string(PreUpgrade) ||
			t == string(PostUpgrade) ||
			t == string(PostInstall)
}

var hookTypes = map[Type]common.HookType{
	PreInstall:  common.HookTypePreSync,
	PreUpgrade:  common.HookTypePreSync,
	PostUpgrade: common.HookTypePostSync,
	PostInstall: common.HookTypePostSync,
}

func (t Type) HookType() common.HookType {
	return hookTypes[t]
}

func Types(obj *unstructured.Unstructured) []Type {
	var types []Type
	for _, text := range resource2.GetAnnotationCSVs(obj, "helm.sh/hook") {
		t, ok := NewType(text)
		if ok {
			types = append(types, t)
		}
	}
	return types
}
