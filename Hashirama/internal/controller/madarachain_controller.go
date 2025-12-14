/*
Copyright 2025.

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

package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	batchv1alpha1 "github.com/example/hashirama/api/v1alpha1"
)

// MadaraChainReconciler reconciles a MadaraChain object
type MadaraChainReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=batch.starknet.l3,resources=madarachains,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.starknet.l3,resources=madarachains/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch.starknet.l3,resources=madarachains/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *MadaraChainReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	// Fetch the MadaraChain instance
	madaraChain := &batchv1alpha1.MadaraChain{}
	err := r.Get(ctx, req.NamespacedName, madaraChain)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Define StatefulSet
	sts := r.statefulSetForMadaraChain(madaraChain)

	// Check if StatefulSet exists
	foundSts := &appsv1.StatefulSet{}
	err = r.Get(ctx, client.ObjectKey{Name: sts.Name, Namespace: sts.Namespace}, foundSts)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating a new StatefulSet", "StatefulSet.Namespace", sts.Namespace, "StatefulSet.Name", sts.Name)
		err = r.Create(ctx, sts)
		if err != nil {
			return ctrl.Result{}, err
		}
		// Requeue to ensure status update
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// Update StatefulSet if needed (e.g. replicas, image, args, ports)
	if *foundSts.Spec.Replicas != *sts.Spec.Replicas ||
		foundSts.Spec.Template.Spec.Containers[0].Image != sts.Spec.Template.Spec.Containers[0].Image {

		foundSts.Spec.Replicas = sts.Spec.Replicas
		foundSts.Spec.Template.Spec.Containers[0].Image = sts.Spec.Template.Spec.Containers[0].Image
		foundSts.Spec.Template.Spec.Containers[0].Args = sts.Spec.Template.Spec.Containers[0].Args
		foundSts.Spec.Template.Spec.Containers[0].Ports = sts.Spec.Template.Spec.Containers[0].Ports
		log.Info("Updating StatefulSet", "StatefulSet.Namespace", foundSts.Namespace, "StatefulSet.Name", foundSts.Name)
		err = r.Update(ctx, foundSts)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	// Define Service
	svc := r.serviceForMadaraChain(madaraChain)

	// Check if Service exists
	foundSvc := &corev1.Service{}
	err = r.Get(ctx, client.ObjectKey{Name: svc.Name, Namespace: svc.Namespace}, foundSvc)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating a new Service", "Service.Namespace", svc.Namespace, "Service.Name", svc.Name)
		err = r.Create(ctx, svc)
		if err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// Update Service if needed
	if foundSvc.Spec.Ports[0].TargetPort != svc.Spec.Ports[0].TargetPort {
		foundSvc.Spec.Ports = svc.Spec.Ports
		log.Info("Updating Service", "Service.Namespace", foundSvc.Namespace, "Service.Name", foundSvc.Name)
		err = r.Update(ctx, foundSvc)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	// Update Status
	if foundSts.Status.ReadyReplicas != madaraChain.Status.NodesRunning {
		madaraChain.Status.NodesRunning = foundSts.Status.ReadyReplicas
		err = r.Status().Update(ctx, madaraChain)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *MadaraChainReconciler) statefulSetForMadaraChain(m *batchv1alpha1.MadaraChain) *appsv1.StatefulSet {
	labels := map[string]string{"app": "madara", "madara_cr": m.Name}
	replicas := m.Spec.Replicas

	sts := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "madara",
						Image: "nginx:latest", // Temporary placeholder for ARM64 verification
						// Args removed for nginx
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80, // nginx default port
							Name:          "http",
						}},
					}},
				},
			},
		},
	}
	// Set owner reference
	ctrl.SetControllerReference(m, sts, r.Scheme)
	return sts
}

func (r *MadaraChainReconciler) serviceForMadaraChain(m *batchv1alpha1.MadaraChain) *corev1.Service {
	labels := map[string]string{"app": "madara", "madara_cr": m.Name}

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name + "-service",
			Namespace: m.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{{
				Port:       m.Spec.Port,
				TargetPort: intstr.FromInt(80), // Map to nginx port
				Protocol:   corev1.ProtocolTCP,
			}},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
	// Set owner reference
	ctrl.SetControllerReference(m, svc, r.Scheme)
	return svc
}

// SetupWithManager sets up the controller with the Manager.
func (r *MadaraChainReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1alpha1.MadaraChain{}).
		Owns(&appsv1.StatefulSet{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
