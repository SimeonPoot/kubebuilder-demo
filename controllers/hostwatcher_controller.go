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

	"github.com/go-logr/logr"
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

	return ctrl.Result{}, nil
}

func (r *HostWatcherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hostsv1.HostWatcher{}).
		Complete(r)
}
