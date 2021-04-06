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

package v1alpha3

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

/*
This setup doubles as setup for our conversion webhooks: as long as our
types implement the
[Hub](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Hub) and
[Convertible](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Convertible)
interfaces, a conversion webhook will be registered.
*/
func (r *TrafficTarget) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}
