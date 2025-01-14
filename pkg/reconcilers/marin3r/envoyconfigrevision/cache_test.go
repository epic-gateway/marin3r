package reconcilers

import (
	"context"
	"reflect"
	"testing"

	marin3rv1alpha1 "github.com/3scale-ops/marin3r/apis/marin3r/v1alpha1"
	xdss "github.com/3scale-ops/marin3r/pkg/discoveryservice/xdss"
	xdss_v3 "github.com/3scale-ops/marin3r/pkg/discoveryservice/xdss/v3"
	envoy "github.com/3scale-ops/marin3r/pkg/envoy"
	envoy_resources "github.com/3scale-ops/marin3r/pkg/envoy/resources"
	envoy_resources_v3 "github.com/3scale-ops/marin3r/pkg/envoy/resources/v3"
	envoy_serializer "github.com/3scale-ops/marin3r/pkg/envoy/serializer"
	testutil "github.com/3scale-ops/marin3r/pkg/util/test"
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	envoy_service_runtime_v3 "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	cache_types "github.com/envoyproxy/go-control-plane/pkg/cache/types"
	cache_v3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNewCacheReconciler(t *testing.T) {
	type args struct {
		ctx       context.Context
		logger    logr.Logger
		client    client.Client
		xdsCache  xdss.Cache
		decoder   envoy_serializer.ResourceUnmarshaller
		generator envoy_resources.Generator
	}
	tests := []struct {
		name string
		args args
		want CacheReconciler
	}{
		{
			name: "Returns a CacheReconciler (v3)",
			args: args{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			want: CacheReconciler{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCacheReconciler(tt.args.ctx, tt.args.logger, tt.args.client, tt.args.xdsCache, tt.args.decoder, tt.args.generator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCacheReconciler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheReconciler_Reconcile(t *testing.T) {
	type fields struct {
		ctx       context.Context
		logger    logr.Logger
		client    client.Client
		xdsCache  xdss.Cache
		decoder   envoy_serializer.ResourceUnmarshaller
		generator envoy_resources.Generator
	}
	type args struct {
		req       types.NamespacedName
		resources *marin3rv1alpha1.EnvoyResources
		nodeID    string
		version   string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *marin3rv1alpha1.VersionTracker
		wantErr     bool
		wantSnap    xdss.Snapshot
		wantVersion string
	}{
		{
			name: "Reconciles cache (v3)",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Endpoints: []marin3rv1alpha1.EnvoyResource{
						{Name: "endpoint", Value: "{\"cluster_name\": \"endpoint\"}"},
					}},
				version: "xxxx",
				nodeID:  "node2",
			},

			want:    &marin3rv1alpha1.VersionTracker{Endpoints: "845f965864"},
			wantErr: false,
			wantSnap: xdss_v3.NewSnapshot(&cache_v3.Snapshot{
				Resources: [7]cache_v3.Resources{
					{Version: "845f965864", Items: map[string]cache_types.ResourceWithTtl{
						"endpoint": {Resource: &envoy_config_endpoint_v3.ClusterLoadAssignment{ClusterName: "endpoint"}}}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
				}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CacheReconciler{
				ctx:       tt.fields.ctx,
				logger:    tt.fields.logger,
				client:    tt.fields.client,
				xdsCache:  tt.fields.xdsCache,
				decoder:   tt.fields.decoder,
				generator: tt.fields.generator,
			}
			got, err := r.Reconcile(tt.args.req, tt.args.resources, tt.args.nodeID, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("CacheReconciler.Reconcile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CacheReconciler.Reconcile() = %v, want %v", got, tt.want)
			}
			gotSnap, _ := r.xdsCache.GetSnapshot(tt.args.nodeID)
			if !testutil.SnapshotsAreEqual(gotSnap, tt.wantSnap) {
				t.Errorf("CacheReconciler.Reconcile() Snapshot = E:%s C:%s R:%s L:%s S:%s RU:%s, want E:%s C:%s R:%s L:%s S:%s RU:%s",
					gotSnap.GetVersion(envoy.Endpoint), gotSnap.GetVersion(envoy.Cluster), gotSnap.GetVersion(envoy.Route), gotSnap.GetVersion(envoy.Listener), gotSnap.GetVersion(envoy.Secret), gotSnap.GetVersion(envoy.Runtime),
					tt.wantSnap.GetVersion(envoy.Endpoint), tt.wantSnap.GetVersion(envoy.Cluster), tt.wantSnap.GetVersion(envoy.Route), tt.wantSnap.GetVersion(envoy.Listener), tt.wantSnap.GetVersion(envoy.Secret), tt.wantSnap.GetVersion(envoy.Runtime),
				)
			}
		})
	}
}

func TestCacheReconciler_GenerateSnapshot(t *testing.T) {
	type fields struct {
		ctx       context.Context
		logger    logr.Logger
		client    client.Client
		xdsCache  xdss.Cache
		decoder   envoy_serializer.ResourceUnmarshaller
		generator envoy_resources.Generator
	}
	type args struct {
		req       types.NamespacedName
		resources *marin3rv1alpha1.EnvoyResources
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    xdss.Snapshot
		wantErr bool
	}{
		{
			name: "Loads v3 resources into the snapshot",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Endpoints: []marin3rv1alpha1.EnvoyResource{
						{Name: "endpoint", Value: "{\"cluster_name\": \"endpoint\"}"},
					},
					Clusters: []marin3rv1alpha1.EnvoyResource{
						{Name: "cluster", Value: "{\"name\": \"cluster\"}"},
					},
					Routes: []marin3rv1alpha1.EnvoyResource{
						{Name: "route", Value: "{\"name\": \"route\"}"},
					},
					Listeners: []marin3rv1alpha1.EnvoyResource{
						{Name: "listener", Value: "{\"name\": \"listener\"}"},
					},
					Runtimes: []marin3rv1alpha1.EnvoyResource{
						{Name: "runtime", Value: "{\"name\": \"runtime\"}"},
					}},
			},
			want: xdss_v3.NewSnapshot(&cache_v3.Snapshot{
				Resources: [7]cache_v3.Resources{
					{Version: "845f965864", Items: map[string]cache_types.ResourceWithTtl{
						"endpoint": {Resource: &envoy_config_endpoint_v3.ClusterLoadAssignment{ClusterName: "endpoint"}},
					}},
					{Version: "568989d74c", Items: map[string]cache_types.ResourceWithTtl{
						"cluster": {Resource: &envoy_config_cluster_v3.Cluster{Name: "cluster"}},
					}},
					{Version: "6645547657", Items: map[string]cache_types.ResourceWithTtl{
						"route": {Resource: &envoy_config_route_v3.RouteConfiguration{Name: "route"}},
					}},
					{Version: "7cb77864cf", Items: map[string]cache_types.ResourceWithTtl{
						"listener": {Resource: &envoy_config_listener_v3.Listener{Name: "listener"}},
					}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "7456685887", Items: map[string]cache_types.ResourceWithTtl{
						"runtime": {Resource: &envoy_service_runtime_v3.Runtime{Name: "runtime"}},
					}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
				},
			}),
			wantErr: false,
		},
		{
			name: "Error, bad endpoint value",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Endpoints: []marin3rv1alpha1.EnvoyResource{
						{Name: "endpoint", Value: "giberish"},
					}},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
		{
			name: "Error, bad cluster value",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Clusters: []marin3rv1alpha1.EnvoyResource{
						{Name: "cluster", Value: "giberish"},
					}},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
		{
			name: "Error, bad route value",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Routes: []marin3rv1alpha1.EnvoyResource{
						{Name: "route", Value: "giberish"},
					}},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
		{
			name: "Error, bad listener value",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Listeners: []marin3rv1alpha1.EnvoyResource{
						{Name: "listener", Value: "giberish"},
					}},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
		{
			name: "Error, bad runtime value",
			fields: fields{
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				client:    fake.NewClientBuilder().Build(),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Runtimes: []marin3rv1alpha1.EnvoyResource{
						{Name: "runtime", Value: "giberish"},
					}},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
		{
			name: "Loads secret resources into the snapshot (v3)",
			fields: fields{
				ctx:    context.TODO(),
				logger: ctrl.Log.WithName("test"),
				client: fake.NewFakeClient(&corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{Name: "secret", Namespace: "xx"},
					Type:       corev1.SecretTypeTLS,
					Data:       map[string][]byte{"tls.crt": []byte("cert"), "tls.key": []byte("key")},
				}),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Secrets: []marin3rv1alpha1.EnvoySecretResource{{Name: "secret"}},
				},
			},
			wantErr: false,
			want: xdss_v3.NewSnapshot(&cache_v3.Snapshot{
				Resources: [7]cache_v3.Resources{
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "56c6b8dc45", Items: map[string]cache_types.ResourceWithTtl{
						"secret": {Resource: &envoy_extensions_transport_sockets_tls_v3.Secret{
							Name: "secret",
							Type: &envoy_extensions_transport_sockets_tls_v3.Secret_TlsCertificate{
								TlsCertificate: &envoy_extensions_transport_sockets_tls_v3.TlsCertificate{
									PrivateKey: &envoy_config_core_v3.DataSource{
										Specifier: &envoy_config_core_v3.DataSource_InlineBytes{InlineBytes: []byte("key")},
									},
									CertificateChain: &envoy_config_core_v3.DataSource{
										Specifier: &envoy_config_core_v3.DataSource_InlineBytes{InlineBytes: []byte("cert")},
									}}}}}}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
					{Version: "", Items: map[string]cache_types.ResourceWithTtl{}},
				},
			}),
		},
		{
			name: "Fails with wrong secret type",
			fields: fields{
				client: fake.NewFakeClient(&corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{Name: "secret", Namespace: "xx"},
					Type:       corev1.SecretTypeBasicAuth,
					Data:       map[string][]byte{"tls.crt": []byte("cert"), "tls.key": []byte("key")},
				}),
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Secrets: []marin3rv1alpha1.EnvoySecretResource{
						{Name: "secret"}},
				},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
		{
			name: "Fails when secret does not exist",
			fields: fields{
				client:    fake.NewClientBuilder().Build(),
				ctx:       context.TODO(),
				logger:    ctrl.Log.WithName("test"),
				xdsCache:  xdss_v3.NewCache(cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)),
				decoder:   envoy_serializer.NewResourceUnmarshaller(envoy_serializer.JSON, envoy.APIv3),
				generator: envoy_resources_v3.Generator{},
			},
			args: args{
				req: types.NamespacedName{Name: "xx", Namespace: "xx"},
				resources: &marin3rv1alpha1.EnvoyResources{
					Secrets: []marin3rv1alpha1.EnvoySecretResource{
						{Name: "secret"}},
				},
			},
			wantErr: true,
			want:    xdss_v3.NewSnapshot(&cache_v3.Snapshot{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CacheReconciler{
				ctx:       tt.fields.ctx,
				logger:    tt.fields.logger,
				client:    tt.fields.client,
				xdsCache:  tt.fields.xdsCache,
				decoder:   tt.fields.decoder,
				generator: tt.fields.generator,
			}
			got, err := r.GenerateSnapshot(tt.args.req, tt.args.resources)
			if (err != nil) != tt.wantErr {
				t.Errorf("CacheReconciler.GenerateSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !testutil.SnapshotsAreEqual(got, tt.want) {
				t.Errorf("CacheReconciler.GenerateSnapshot() = E:%s C:%s R:%s L:%s S:%s RU:%s, want E:%s C:%s R:%s L:%s S:%s RU:%s",
					got.GetVersion(envoy.Endpoint), got.GetVersion(envoy.Cluster), got.GetVersion(envoy.Route), got.GetVersion(envoy.Listener), got.GetVersion(envoy.Secret), got.GetVersion(envoy.Runtime),
					tt.want.GetVersion(envoy.Endpoint), tt.want.GetVersion(envoy.Cluster), tt.want.GetVersion(envoy.Route), tt.want.GetVersion(envoy.Listener), tt.want.GetVersion(envoy.Secret), tt.want.GetVersion(envoy.Runtime),
				)
			}
		})
	}
}
