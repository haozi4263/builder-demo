/*
Copyright 2021 jude.

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

	webappv1 "github.com/haozi4263/builder-demo/api/v1"
)

// BookReconciler reconciles a Book object
type BookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.shimo.im,resources=books,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.shimo.im,resources=books/status,verbs=get;update;patch

func (r *BookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	log := r.Log.WithValues("book", req.NamespacedName)

	//查询book是否存在，存在需要查看spec是否有变化，有变化做一下更新等操作
	// 实现自己等业务逻辑
	log.Info("crd book update")

	return ctrl.Result{}, nil
}

func (r *BookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Book{}).
		Complete(r)
}
