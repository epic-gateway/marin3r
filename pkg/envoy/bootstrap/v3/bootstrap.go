package envoy

import (
	"time"

	"github.com/3scale-ops/marin3r/pkg/envoy"
	envoy_bootstrap_options "github.com/3scale-ops/marin3r/pkg/envoy/bootstrap/options"
	envoy_resources "github.com/3scale-ops/marin3r/pkg/envoy/resources"
	envoy_serializer_v3 "github.com/3scale-ops/marin3r/pkg/envoy/serializer/v3"
	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	envoy_service_discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Config is a struct with options and methods to generate an envoy bootstrap config
type Config struct {
	Options envoy_bootstrap_options.ConfigOptions
}

func (c *Config) getAdminAddress() string { return stringOrDefault(c.Options.AdminAddress, "0.0.0.0") }
func (c *Config) getAdminPort() uint32    { return intOrDefault(c.Options.AdminPort, 9001) }
func (c *Config) getAdminAccessLogPath() string {
	return stringOrDefault(c.Options.AdminAccessLogPath, "/dev/null")
}

// GenerateStatic returns the json serialized representation of an envoy
// bootstrap object that can be passed as the configuration file to an envoy proxy
// so it can connect to the discovery service.
func (c *Config) GenerateStatic() (string, error) {

	tlsContext := &envoy_extensions_transport_sockets_tls_v3.UpstreamTlsContext{
		CommonTlsContext: &envoy_extensions_transport_sockets_tls_v3.CommonTlsContext{
			TlsCertificateSdsSecretConfigs: []*envoy_extensions_transport_sockets_tls_v3.SdsSecretConfig{
				{
					Name: "xds_client_certificate",
					SdsConfig: &envoy_config_core_v3.ConfigSource{
						ConfigSourceSpecifier: &envoy_config_core_v3.ConfigSource_Path{
							Path: c.Options.SdsConfigSourcePath,
						},
					},
				},
			},
		},
	}

	serializedTLSContext, err := anypb.New(tlsContext)
	if err != nil {
		return "", err
	}

	cfg := &envoy_config_bootstrap_v3.Bootstrap{
		Admin: &envoy_config_bootstrap_v3.Admin{
			AccessLogPath: c.getAdminAccessLogPath(),
			Address: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Address: c.getAdminAddress(),
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: c.getAdminPort(),
						},
					},
				},
			},
		},
		Node: &envoy_config_core_v3.Node{
			Id:      c.Options.NodeID,
			Cluster: c.Options.Cluster,
		},
		DynamicResources: &envoy_config_bootstrap_v3.Bootstrap_DynamicResources{
			AdsConfig: &envoy_config_core_v3.ApiConfigSource{
				ApiType:             envoy_config_core_v3.ApiConfigSource_GRPC,
				TransportApiVersion: envoy_config_core_v3.ApiVersion_V3,
				GrpcServices: []*envoy_config_core_v3.GrpcService{
					{
						TargetSpecifier: &envoy_config_core_v3.GrpcService_EnvoyGrpc_{
							EnvoyGrpc: &envoy_config_core_v3.GrpcService_EnvoyGrpc{
								ClusterName: envoy_bootstrap_options.XdsClusterName,
							},
						},
					},
				},
			},
			CdsConfig: &envoy_config_core_v3.ConfigSource{
				ResourceApiVersion: envoy_config_core_v3.ApiVersion_V3,
				ConfigSourceSpecifier: &envoy_config_core_v3.ConfigSource_Ads{
					Ads: &envoy_config_core_v3.AggregatedConfigSource{},
				},
			},
			LdsConfig: &envoy_config_core_v3.ConfigSource{
				ResourceApiVersion: envoy_config_core_v3.ApiVersion_V3,
				ConfigSourceSpecifier: &envoy_config_core_v3.ConfigSource_Ads{
					Ads: &envoy_config_core_v3.AggregatedConfigSource{},
				},
			},
		},
		StaticResources: &envoy_config_bootstrap_v3.Bootstrap_StaticResources{
			Clusters: []*envoy_config_cluster_v3.Cluster{
				{
					Name:           envoy_bootstrap_options.XdsClusterName,
					ConnectTimeout: durationpb.New(1 * time.Second),
					ClusterDiscoveryType: &envoy_config_cluster_v3.Cluster_Type{
						Type: envoy_config_cluster_v3.Cluster_STRICT_DNS,
					},
					Http2ProtocolOptions: &envoy_config_core_v3.Http2ProtocolOptions{},
					LoadAssignment: &envoy_config_endpoint_v3.ClusterLoadAssignment{
						ClusterName: envoy_bootstrap_options.XdsClusterName,
						Endpoints: []*envoy_config_endpoint_v3.LocalityLbEndpoints{
							{
								LbEndpoints: []*envoy_config_endpoint_v3.LbEndpoint{
									{
										HostIdentifier: &envoy_config_endpoint_v3.LbEndpoint_Endpoint{
											Endpoint: &envoy_config_endpoint_v3.Endpoint{
												Address: &envoy_config_core_v3.Address{
													Address: &envoy_config_core_v3.Address_SocketAddress{
														SocketAddress: &envoy_config_core_v3.SocketAddress{
															Address: c.Options.XdsHost,
															PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
																PortValue: c.Options.XdsPort,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					TransportSocket: &envoy_config_core_v3.TransportSocket{
						Name: wellknown.TransportSocketTls,
						ConfigType: &envoy_config_core_v3.TransportSocket_TypedConfig{
							TypedConfig: serializedTLSContext,
						},
					},
				},
				{
					Name:           "eds-server",
					ConnectTimeout: ptypes.DurationProto(1 * time.Second),
					ClusterDiscoveryType: &envoy_config_cluster_v3.Cluster_Type{
						Type: envoy_config_cluster_v3.Cluster_STRICT_DNS,
					},
					Http2ProtocolOptions: &envoy_config_core_v3.Http2ProtocolOptions{},
					LoadAssignment: &envoy_config_endpoint_v3.ClusterLoadAssignment{
						ClusterName: "eds-server",
						Endpoints: []*envoy_config_endpoint_v3.LocalityLbEndpoints{
							{
								LbEndpoints: []*envoy_config_endpoint_v3.LbEndpoint{
									{
										HostIdentifier: &envoy_config_endpoint_v3.LbEndpoint_Endpoint{
											Endpoint: &envoy_config_endpoint_v3.Endpoint{
												Address: &envoy_config_core_v3.Address{
													Address: &envoy_config_core_v3.Address_SocketAddress{
														SocketAddress: &envoy_config_core_v3.SocketAddress{
															Address: "eds-server.epic",
															PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
																PortValue: c.Options.XdsPort,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					TransportSocket: &envoy_config_core_v3.TransportSocket{
						Name: wellknown.TransportSocketTls,
						ConfigType: &envoy_config_core_v3.TransportSocket_TypedConfig{
							TypedConfig: serializedTLSContext,
						},
					},
					UpstreamConnectionOptions: &envoy_config_cluster_v3.UpstreamConnectionOptions{
						TcpKeepalive: &envoy_config_core_v3.TcpKeepalive{
							KeepaliveProbes:   &wrapperspb.UInt32Value{Value: 4},
							KeepaliveTime:     &wrapperspb.UInt32Value{Value: 10},
							KeepaliveInterval: &wrapperspb.UInt32Value{Value: 5},
						},
					},
				},
			},
		},
		LayeredRuntime: &envoy_config_bootstrap_v3.LayeredRuntime{
			Layers: []*envoy_config_bootstrap_v3.RuntimeLayer{{
				Name: c.Options.RtdsLayerResourceName,
				LayerSpecifier: &envoy_config_bootstrap_v3.RuntimeLayer_RtdsLayer_{
					RtdsLayer: &envoy_config_bootstrap_v3.RuntimeLayer_RtdsLayer{
						Name: c.Options.RtdsLayerResourceName,
						RtdsConfig: &envoy_config_core_v3.ConfigSource{
							ResourceApiVersion: envoy_config_core_v3.ApiVersion_V3,
							ConfigSourceSpecifier: &envoy_config_core_v3.ConfigSource_Ads{
								Ads: &envoy_config_core_v3.AggregatedConfigSource{},
							},
						},
					},
				},
			}},
		},
	}

	if len(c.Options.Metadata) > 0 {
		cfg.Node.Metadata = &structpb.Struct{Fields: map[string]*structpb.Value{}}
		for key, value := range c.Options.Metadata {
			cfg.Node.Metadata.Fields[key] = &structpb.Value{
				Kind: &structpb.Value_StringValue{
					StringValue: value,
				},
			}
		}
	}

	json, err := envoy_serializer_v3.JSON{}.Marshal(cfg)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GenerateSdsResources generates the envoy static config required for
// filesystem discovery of certificates.
func (c *Config) GenerateSdsResources() (map[string]string, error) {

	generator := envoy_resources.NewGenerator(envoy.APIv3)
	secret := generator.NewSecretFromPath("xds_client_certificate", c.Options.XdsClientCertificatePath, c.Options.XdsClientCertificateKeyPath)

	a, err := anypb.New(secret)
	if err != nil {
		return nil, err
	}
	cfg := &envoy_service_discovery_v3.DiscoveryResponse{
		Resources: []*anypb.Any{a},
	}

	json, err := envoy_serializer_v3.JSON{}.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		envoy_bootstrap_options.TlsCertificateSdsSecretFileName: string(json),
	}, nil
}

func stringOrDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}

func intOrDefault(i, def uint32) uint32 {
	if i == 0 {
		return def
	}
	return i
}
