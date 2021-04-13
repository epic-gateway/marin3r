package generators

import (
	"fmt"

	operatorv1alpha1 "github.com/3scale-ops/marin3r/apis/operator.marin3r/v1alpha1"
	"github.com/3scale-ops/marin3r/pkg/envoy"
	defaults "github.com/3scale-ops/marin3r/pkg/envoy/bootstrap/defaults"
	"github.com/3scale-ops/marin3r/pkg/reconcilers/lockedresources"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (cfg *GeneratorOptions) Deployment(hash string, replicas *int32) lockedresources.GeneratorFunction {

	return func() client.Object {

		dep := &appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Deployment",
				APIVersion: appsv1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      cfg.resourceName(),
				Namespace: cfg.Namespace,
				Labels:    cfg.labels(),
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: replicas,
				Selector: &metav1.LabelSelector{
					MatchLabels: cfg.labels(),
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						CreationTimestamp: metav1.Time{},
						Labels: func() (labels map[string]string) {
							labels = cfg.labels()
							labels[operatorv1alpha1.EnvoyDeploymentBootstrapConfigHashLabelKey] = hash
							return
						}(),
					},
					Spec: corev1.PodSpec{
						Affinity: cfg.PodAffinity,
						Volumes: []corev1.Volume{
							{
								Name: defaults.DeploymentTLSVolume,
								VolumeSource: corev1.VolumeSource{
									Secret: &corev1.SecretVolumeSource{
										SecretName: fmt.Sprintf("%s-%s", defaults.DeploymentClientCertificate, cfg.InstanceName),
									},
								},
							},
							{
								Name: defaults.DeploymentConfigVolume,
								VolumeSource: corev1.VolumeSource{
									ConfigMap: &corev1.ConfigMapVolumeSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: func() string {
												if cfg.EnvoyAPIVersion == envoy.APIv2 {
													return fmt.Sprintf("%s-%s", defaults.DeploymentBootstrapConfigMapV2, cfg.InstanceName)
												}
												return fmt.Sprintf("%s-%s", defaults.DeploymentBootstrapConfigMapV3, cfg.InstanceName)
											}(),
										},
									},
								},
							},
						},
						Containers: []corev1.Container{
							{
								Name:    defaults.DeploymentContainerName,
								Image:   cfg.DeploymentImage,
								Command: []string{"envoy"},
								Args: func() []string {
									args := []string{"-c",
										fmt.Sprintf("%s/%s", defaults.EnvoyConfigBasePath, defaults.EnvoyConfigFileName),
										"--service-node",
										cfg.EnvoyNodeID,
										"--service-cluster",
										cfg.EnvoyClusterID,
									}
									// TODO: validate that user is not overwriting 'service-node' or 'service-cluster'
									if len(cfg.ExtraArgs) > 0 {
										args = append(args, cfg.ExtraArgs...)
									}
									return args
								}(),
								Resources: cfg.DeploymentResources,
								Ports: func() []corev1.ContainerPort {
									ports := make([]corev1.ContainerPort, len(cfg.ExposedPorts)+1)
									for i := 0; i < len(cfg.ExposedPorts); i++ {
										p := corev1.ContainerPort{
											Name:          cfg.ExposedPorts[i].Name,
											ContainerPort: cfg.ExposedPorts[i].Port,
										}
										if cfg.ExposedPorts[i].Protocol != nil {
											p.Protocol = *cfg.ExposedPorts[i].Protocol
										}
										ports[i] = p
									}
									ports[len(cfg.ExposedPorts)] = corev1.ContainerPort{
										Name:          "admin",
										ContainerPort: cfg.AdminPort,
									}
									return ports
								}(),
								VolumeMounts: []corev1.VolumeMount{
									{
										Name:      defaults.DeploymentTLSVolume,
										ReadOnly:  true,
										MountPath: defaults.EnvoyTLSBasePath,
									},
									{
										Name:      defaults.DeploymentConfigVolume,
										ReadOnly:  true,
										MountPath: defaults.EnvoyConfigBasePath,
									},
								},
								LivenessProbe: &corev1.Probe{
									Handler: corev1.Handler{
										HTTPGet: &corev1.HTTPGetAction{
											Path: "/ready",
											Port: intstr.IntOrString{IntVal: cfg.AdminPort},
										},
									},
									InitialDelaySeconds: cfg.LivenessProbe.InitialDelaySeconds,
									TimeoutSeconds:      cfg.LivenessProbe.TimeoutSeconds,
									PeriodSeconds:       cfg.LivenessProbe.PeriodSeconds,
									SuccessThreshold:    cfg.LivenessProbe.SuccessThreshold,
									FailureThreshold:    cfg.LivenessProbe.FailureThreshold,
								},
								ReadinessProbe: &corev1.Probe{
									Handler: corev1.Handler{
										HTTPGet: &corev1.HTTPGetAction{
											Path: "/ready",
											Port: intstr.IntOrString{IntVal: cfg.AdminPort},
										},
									},
									InitialDelaySeconds: cfg.ReadinessProbe.InitialDelaySeconds,
									TimeoutSeconds:      cfg.ReadinessProbe.TimeoutSeconds,
									PeriodSeconds:       cfg.ReadinessProbe.PeriodSeconds,
									SuccessThreshold:    cfg.ReadinessProbe.SuccessThreshold,
									FailureThreshold:    cfg.ReadinessProbe.FailureThreshold,
								},
								TerminationMessagePath:   corev1.TerminationMessagePathDefault,
								TerminationMessagePolicy: corev1.TerminationMessageReadFile,
								ImagePullPolicy:          corev1.PullIfNotPresent,
							},
						},
						RestartPolicy:                 corev1.RestartPolicyAlways,
						TerminationGracePeriodSeconds: pointer.Int64Ptr(corev1.DefaultTerminationGracePeriodSeconds),
						DNSPolicy:                     corev1.DNSClusterFirst,
						ServiceAccountName:            "default",
						DeprecatedServiceAccount:      "default",
						SecurityContext:               &corev1.PodSecurityContext{},
						SchedulerName:                 corev1.DefaultSchedulerName,
					},
				},
				Strategy: appsv1.DeploymentStrategy{
					Type: appsv1.RollingUpdateDeploymentStrategyType,
					RollingUpdate: &appsv1.RollingUpdateDeployment{
						MaxUnavailable: &intstr.IntOrString{
							Type:   intstr.String,
							StrVal: "25%",
						},
						MaxSurge: &intstr.IntOrString{
							Type:   intstr.String,
							StrVal: "25%",
						},
					},
				},
				RevisionHistoryLimit:    pointer.Int32Ptr(10),
				ProgressDeadlineSeconds: pointer.Int32Ptr(600),
			},
		}

		// Enforce a fixed number of replicas if static replicas is active
		if cfg.Replicas.Static != nil {
			dep.Spec.Replicas = cfg.Replicas.Static
		}

		return dep
	}
}