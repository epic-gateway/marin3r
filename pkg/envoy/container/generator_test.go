package container

import (
	"fmt"
	"reflect"
	"testing"

	operatorv1alpha1 "github.com/3scale-ops/marin3r/apis/operator.marin3r/v1alpha1"
	"github.com/3scale-ops/marin3r/pkg/envoy/container/defaults"
	"github.com/3scale-ops/marin3r/pkg/envoy/container/shutdownmanager"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestContainerConfig_Container(t *testing.T) {
	tests := []struct {
		name string
		cc   ContainerConfig
		want []corev1.Container
	}{
		{
			name: "Generates an Envoy container for the given config",
			cc: ContainerConfig{
				Name:               "envoy",
				Image:              "envoy:test",
				BootstrapConfigMap: "bootstrap-configmap",
				ConfigBasePath:     "/config",
				ConfigFileName:     "config.json",
				ConfigVolume:       "config",
				TLSBasePath:        "/tls",
				TLSVolume:          "tls",
				NodeID:             "test-id",
				ClusterID:          "test-id",
				ClientCertSecret:   "client-secret",
				ExtraArgs:          []string{"--some-arg", "some-value"},
				Resources:          corev1.ResourceRequirements{},
				AdminPort:          5000,
				Ports: []corev1.ContainerPort{
					{
						Name:          "udp",
						ContainerPort: 6000,
						Protocol:      corev1.Protocol("UDP"),
					},
					{
						Name:          "https",
						ContainerPort: 8443,
					},
				},
				LivenessProbe: operatorv1alpha1.ProbeSpec{
					InitialDelaySeconds: 1,
					TimeoutSeconds:      1,
					PeriodSeconds:       1,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
				ReadinessProbe: operatorv1alpha1.ProbeSpec{
					InitialDelaySeconds: 1,
					TimeoutSeconds:      1,
					PeriodSeconds:       1,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
			},
			want: []corev1.Container{{
				Name:    "envoy",
				Image:   "envoy:test",
				Command: []string{"envoy"},
				Args: []string{
					"-c",
					"/config/config.json",
					"--service-node",
					"test-id",
					"--service-cluster",
					"test-id",
					"--some-arg",
					"some-value",
				},
				Ports: []corev1.ContainerPort{
					{
						Name:          "udp",
						ContainerPort: 6000,
						Protocol:      corev1.Protocol("UDP"),
					},
					{
						Name:          "https",
						ContainerPort: 8443,
					},
					{
						Name:          "admin",
						ContainerPort: 5000,
					},
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "tls",
						ReadOnly:  true,
						MountPath: "/tls",
					},
					{
						Name:      "config",
						ReadOnly:  true,
						MountPath: "/config",
					},
				},
				LivenessProbe: &corev1.Probe{
					Handler: corev1.Handler{
						HTTPGet: &corev1.HTTPGetAction{
							Path: "/ready",
							Port: intstr.IntOrString{IntVal: 5000},
						},
					},
					InitialDelaySeconds: 1,
					TimeoutSeconds:      1,
					PeriodSeconds:       1,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
				ReadinessProbe: &corev1.Probe{
					Handler: corev1.Handler{
						HTTPGet: &corev1.HTTPGetAction{
							Path: "/ready",
							Port: intstr.IntOrString{IntVal: 5000},
						},
					},
					InitialDelaySeconds: 1,
					TimeoutSeconds:      1,
					PeriodSeconds:       1,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
				TerminationMessagePath:   corev1.TerminationMessagePathDefault,
				TerminationMessagePolicy: corev1.TerminationMessageReadFile,
				ImagePullPolicy:          corev1.PullIfNotPresent,
			}},
		},
		{
			name: "Generates containers for the given config (shtdnmgr enabled)",
			cc: ContainerConfig{
				Name:               "envoy",
				Image:              "envoy:test",
				BootstrapConfigMap: "bootstrap-configmap",
				ConfigBasePath:     "/config",
				ConfigFileName:     "config.json",
				ConfigVolume:       "config",
				TLSBasePath:        "/tls",
				TLSVolume:          "tls",
				NodeID:             "test-id",
				ClusterID:          "test-id",
				ClientCertSecret:   "client-secret",
				ExtraArgs:          []string{"--some-arg", "some-value"},
				Resources:          corev1.ResourceRequirements{},
				AdminPort:          5000,
				Ports: []corev1.ContainerPort{
					{
						Name:          "udp",
						ContainerPort: 6000,
						Protocol:      corev1.Protocol("UDP"),
					},
					{
						Name:          "https",
						ContainerPort: 8443,
					},
				},
				LivenessProbe: operatorv1alpha1.ProbeSpec{
					InitialDelaySeconds: 1,
					TimeoutSeconds:      1,
					PeriodSeconds:       1,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
				ReadinessProbe: operatorv1alpha1.ProbeSpec{
					InitialDelaySeconds: 1,
					TimeoutSeconds:      1,
					PeriodSeconds:       1,
					SuccessThreshold:    1,
					FailureThreshold:    1,
				},
				ShutdownManagerEnabled: true,
				ShutdownManagerPort:    30000,
				ShutdownManagerImage:   "image:shtdnmgr",
			},
			want: []corev1.Container{
				{
					Name:    "envoy",
					Image:   "envoy:test",
					Command: []string{"envoy"},
					Args: []string{
						"-c",
						"/config/config.json",
						"--service-node",
						"test-id",
						"--service-cluster",
						"test-id",
						"--some-arg",
						"some-value",
					},
					Ports: []corev1.ContainerPort{
						{
							Name:          "udp",
							ContainerPort: 6000,
							Protocol:      corev1.Protocol("UDP"),
						},
						{
							Name:          "https",
							ContainerPort: 8443,
						},
						{
							Name:          "admin",
							ContainerPort: 5000,
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "tls",
							ReadOnly:  true,
							MountPath: "/tls",
						},
						{
							Name:      "config",
							ReadOnly:  true,
							MountPath: "/config",
						},
					},
					LivenessProbe: &corev1.Probe{
						Handler: corev1.Handler{
							HTTPGet: &corev1.HTTPGetAction{
								Path: "/ready",
								Port: intstr.IntOrString{IntVal: 5000},
							},
						},
						InitialDelaySeconds: 1,
						TimeoutSeconds:      1,
						PeriodSeconds:       1,
						SuccessThreshold:    1,
						FailureThreshold:    1,
					},
					ReadinessProbe: &corev1.Probe{
						Handler: corev1.Handler{
							HTTPGet: &corev1.HTTPGetAction{
								Path: "/ready",
								Port: intstr.IntOrString{IntVal: 5000},
							},
						},
						InitialDelaySeconds: 1,
						TimeoutSeconds:      1,
						PeriodSeconds:       1,
						SuccessThreshold:    1,
						FailureThreshold:    1,
					},
					TerminationMessagePath:   corev1.TerminationMessagePathDefault,
					TerminationMessagePolicy: corev1.TerminationMessageReadFile,
					ImagePullPolicy:          corev1.PullIfNotPresent,
					Lifecycle: &corev1.Lifecycle{
						PreStop: &corev1.Handler{
							HTTPGet: &corev1.HTTPGetAction{
								Path:   shutdownmanager.ShutdownEndpoint,
								Port:   intstr.FromInt(30000),
								Scheme: corev1.URISchemeHTTP,
							},
						},
					},
				},
				{
					Name:  "envoy-shtdn-mgr",
					Image: "image:shtdnmgr",
					Args: []string{
						"shutdown-manager",
						"--port",
						fmt.Sprintf("%d", 30000),
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse(defaults.ShtdnMgrDefaultCPURequests),
							corev1.ResourceMemory: resource.MustParse(defaults.ShtdnMgrDefaultMemoryRequests),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse(defaults.ShtdnMgrDefaultCPULimits),
							corev1.ResourceMemory: resource.MustParse(defaults.ShtdnMgrDefaultMemoryLimits),
						},
					},
					LivenessProbe: &corev1.Probe{
						Handler: corev1.Handler{
							HTTPGet: &corev1.HTTPGetAction{
								Path:   shutdownmanager.HealthEndpoint,
								Port:   intstr.FromInt(30000),
								Scheme: corev1.URISchemeHTTP,
							},
						},
						InitialDelaySeconds: 3,
						PeriodSeconds:       10,
					},
					Lifecycle: &corev1.Lifecycle{
						PreStop: &corev1.Handler{
							HTTPGet: &corev1.HTTPGetAction{
								Path:   shutdownmanager.DrainEndpoint,
								Port:   intstr.FromInt(30000),
								Scheme: corev1.URISchemeHTTP,
							},
						},
					},
					TerminationMessagePath:   corev1.TerminationMessagePathDefault,
					TerminationMessagePolicy: corev1.TerminationMessageReadFile,
					ImagePullPolicy:          corev1.PullIfNotPresent,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.cc.Containers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContainerConfig.Container() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainerConfig_Volumes(t *testing.T) {
	tests := []struct {
		name string
		cc   ContainerConfig
		want []corev1.Volume
	}{
		{
			name: "Generates required volumes for an Envoy container with the given config",
			cc: ContainerConfig{
				BootstrapConfigMap: "bootstrap-configmap",
				ConfigVolume:       "config",
				TLSVolume:          "tls",
				ClientCertSecret:   "client-secret",
			},
			want: []corev1.Volume{
				{
					Name: "tls",
					VolumeSource: corev1.VolumeSource{
						Secret: &corev1.SecretVolumeSource{
							SecretName: "client-secret",
						},
					},
				},
				{
					Name: "config",
					VolumeSource: corev1.VolumeSource{
						ConfigMap: &corev1.ConfigMapVolumeSource{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "bootstrap-configmap",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cc.Volumes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContainerConfig.Volumes() = %v, want %v", got, tt.want)
			}
		})
	}
}