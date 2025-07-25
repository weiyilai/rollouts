/*
Copyright 2019 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mutating

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/openkruise/rollouts/pkg/util"
	"github.com/openkruise/rollouts/pkg/webhook/types"
)

// +kubebuilder:webhook:path=/mutate-apps-kruise-io-v1alpha1-cloneset,mutating=true,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1;v1beta1,groups=apps.kruise.io,resources=clonesets,verbs=update,versions=v1alpha1,name=mcloneset.kb.io
// +kubebuilder:webhook:path=/mutate-apps-kruise-io-v1alpha1-daemonset,mutating=true,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1;v1beta1,groups=apps.kruise.io,resources=daemonsets,verbs=update,versions=v1alpha1,name=mdaemonset.kb.io
// +kubebuilder:webhook:path=/mutate-apps-v1-deployment,mutating=true,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1;v1beta1,groups=apps,resources=deployments,verbs=update,versions=v1,name=mdeployment.kb.io
// +kubebuilder:webhook:path=/mutate-unified-workload,mutating=true,failurePolicy=fail,sideEffects=None,admissionReviewVersions=v1;v1beta1,groups=*,resources=*,verbs=create;update,versions=*,name=munifiedworload.kb.io

var (
	// HandlerMap contains admission webhook handlers
	HandlerMap = map[string]types.HandlerGetter{
		"mutate-apps-kruise-io-v1alpha1-cloneset": func(mgr manager.Manager) admission.Handler {
			decoder := admission.NewDecoder(mgr.GetScheme())
			return &WorkloadHandler{Decoder: decoder, Client: mgr.GetClient(), Finder: util.NewControllerFinder(mgr.GetClient())}
		},
		"mutate-apps-v1-deployment": func(mgr manager.Manager) admission.Handler {
			decoder := admission.NewDecoder(mgr.GetScheme())
			return &WorkloadHandler{Decoder: decoder, Client: mgr.GetClient(), Finder: util.NewControllerFinder(mgr.GetClient())}
		},

		"mutate-apps-kruise-io-v1alpha1-daemonset": func(mgr manager.Manager) admission.Handler {
			decoder := admission.NewDecoder(mgr.GetScheme())
			return &WorkloadHandler{Decoder: decoder, Client: mgr.GetClient(), Finder: util.NewControllerFinder(mgr.GetClient())}
		},
		"mutate-unified-workload": func(mgr manager.Manager) admission.Handler {
			decoder := admission.NewDecoder(mgr.GetScheme())
			return &UnifiedWorkloadHandler{Decoder: decoder, Client: mgr.GetClient(), Finder: util.NewControllerFinder(mgr.GetClient())}
		},
	}
)
