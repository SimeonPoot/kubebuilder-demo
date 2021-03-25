/*


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

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hostsv1 "kubebuilder-demo/api/v1"
)

// HostWatcherReconciler reconciles a HostWatcher object
type HostWatcherReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=hosts.simp.io,resources=hostwatchers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hosts.simp.io,resources=hostwatchers/status,verbs=get;update;patch

// RBAC PODS
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods/status,verbs=get;update;patch

func (r *HostWatcherReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("hostwatcher", req.NamespacedName)

	// your logic here
	log.Info("fetching HostWatcher resource")
	hostWatcher := hostsv1.HostWatcher{}
	if err := r.Client.Get(ctx, req.NamespacedName, &hostWatcher); err != nil {
		log.Error(err, "failed to get HostWatcher resource")
		// Ignore NotFound errors as they will be retried automatically if the
		// resource is created in future.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("pod spec")

	// Using a typed object.
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "namespace",
			Name:      "created-by-kb-006",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Image: "nginx",
					Name:  "nginx",
				},
			},
		},
	}
	// c is a created client.
	_ = r.Client.Create(ctx, pod)

	podList := &corev1.PodList{}

	// 	_ = r.List(ctx, podList, client.InNamespace("namespace"))
	_ = r.List(ctx, podList)

	log.Info("start forloop")
	// count := 0
	// for i := 5; i < count; i++ {
	// 	fmt.Println(podList.Items[i].Name)
	// 	fmt.Println(podList.Items[i].ObjectMeta.UID)
	// 	log.Info("test")
	// }

	// // broken forloop: it's breaking at 2! perhaps the range in Items isn't correct.
	for i, pod := range podList.Items {
		// fmt.Println(pod.Items[i].Annotations)
		fmt.Println("Pod-index: ", i)
		fmt.Printf("PodName: %s, PodUID: %s", pod.Name, pod.UID)
	}

	log.Info("end forloop")

	instance := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "created-by-kb-006",
			Namespace: "namespace",
		},
		Spec: corev1.PodSpec{Containers: []corev1.Container{
			{
				Image: "nginx"},
		}},
		Status: corev1.PodStatus{},
	}
	// instances := &corev1.PodList{}

	_ = r.Client.Get(ctx, req.NamespacedName, instance)
	// if err := r.Client.Get(ctx, req.NamespacedName, instances); err != nil {
	// 	log.Error(err, "failed to get HostWatcher resource")
	// 	return ctrl.Result{}, client.IgnoreNotFound(err)
	// }
	// log.Info(instances.Kind)
	// pod := &corev1.Pod{
	// 	TypeMeta: v1.TypeMeta{},
	// 	ObjectMeta: v1.ObjectMeta{
	// 		Namespace: "kubebuilder-demo-system",
	// 	},
	// 	Spec:   corev1.PodSpec{},
	// 	Status: corev1.PodStatus{},
	// }

	// log.Info(string(instance.Status.Phase))
	// switch instance.Status.Phase {
	// case corev1.PodPending:
	// 	log.Info("Phase: PENDING")
	// case corev1.PodRunning:
	// 	log.Info("Phase: RUNNING", "PodIP", instance.Status.PodIP)
	// case corev1.PodSucceeded:
	// 	log.Info("Phase: SUCCEEDED")
	// case corev1.PodFailed:
	// 	log.Info("Phase: FAILED")
	// case corev1.PodUnknown:
	// 	log.Info("Phase: UNKNOWN")
	// }

	log.Info("here we are!")
	// log.Info(pod.Spec.Hostname)

	ingresses := networkingv1.NetworkPolicy{}
	log.Info(ingresses.Name)

	return ctrl.Result{}, nil
}

func (r *HostWatcherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hostsv1.HostWatcher{}).
		Complete(r)
}
