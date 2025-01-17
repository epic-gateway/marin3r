package generators

import (
	"testing"

	operatorv1alpha1 "github.com/3scale-ops/marin3r/apis/operator.marin3r/v1alpha1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestGeneratorOptions_PDB(t *testing.T) {
	tests := []struct {
		name string
		opts GeneratorOptions
		want client.Object
	}{
		{
			name: "Generate an HPA",
			opts: GeneratorOptions{
				InstanceName: "instance",
				Namespace:    "default",
				PodDisruptionBudget: operatorv1alpha1.PodDisruptionBudgetSpec{
					MinAvailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
				},
			},
			want: &policyv1beta1.PodDisruptionBudget{
				TypeMeta: metav1.TypeMeta{
					Kind:       "PodDisruptionBudget",
					APIVersion: policyv1beta1.SchemeGroupVersion.String(),
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "marin3r-envoydeployment-instance",
					Namespace: "default",
					Labels: map[string]string{
						"app.kubernetes.io/name":       "marin3r",
						"app.kubernetes.io/managed-by": "marin3r-operator",
						"app.kubernetes.io/component":  "envoy-deployment",
						"app.kubernetes.io/instance":   "instance",
					},
				},
				Spec: policyv1beta1.PodDisruptionBudgetSpec{
					MinAvailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 1},
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app.kubernetes.io/name":       "marin3r",
							"app.kubernetes.io/managed-by": "marin3r-operator",
							"app.kubernetes.io/component":  "envoy-deployment",
							"app.kubernetes.io/instance":   "instance",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.opts
			if got := cfg.PDB()(); !equality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("GeneratorOptions.PDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
