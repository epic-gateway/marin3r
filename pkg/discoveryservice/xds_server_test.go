// Copyright 2020 rvazquez@redhat.com
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package discoveryservice

import (
	"context"
	"crypto/tls"
	"reflect"
	"sync"
	"testing"

	xdss "github.com/3scale-ops/marin3r/pkg/discoveryservice/xdss"
	"github.com/3scale-ops/marin3r/pkg/discoveryservice/xdss/stats"
	xdss_v3 "github.com/3scale-ops/marin3r/pkg/discoveryservice/xdss/v3"
	envoy "github.com/3scale-ops/marin3r/pkg/envoy"
	cache_v3 "github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	server_v3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	snapshotCacheV3 = cache_v3.NewSnapshotCache(true, cache_v3.IDHash{}, nil)
)

func TestNewXdsServer(t *testing.T) {

	type args struct {
		ctx       context.Context
		adsPort   uint
		tlsConfig *tls.Config
		logger    logr.Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Returns a new XdsServer from the given params",
			args{context.Background(), 10000, &tls.Config{}, ctrl.Log},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewXdsServer(tt.args.ctx, tt.args.adsPort, tt.args.tlsConfig, tt.args.logger)
			if got.snapshotCacheV3 == nil || got.serverV3 == nil || got.callbacksV3 == nil {
				t.Errorf("TestNewXdsServer = expected non-empty caches")
			}
		})
	}
}

func TestXdsServer_Start(t *testing.T) {

	tests := []struct {
		name string
		xdss *XdsServer
	}{
		{
			"Runs the ads server",
			&XdsServer{
				context.Background(),
				10000,
				&tls.Config{},
				server_v3.NewServer(context.Background(), snapshotCacheV3, &xdss_v3.Callbacks{Logger: ctrl.Log}),
				snapshotCacheV3,
				&xdss_v3.Callbacks{Logger: ctrl.Log},
				stats.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wait sync.WaitGroup
			stopCh := make(chan struct{})
			wait.Add(1)
			go func() {
				if err := tt.xdss.Start(stopCh); err != nil {
					t.Errorf("TestXdsServer_Start = non nil error: '%s'", err)
				}
				wait.Done()
			}()
			close(stopCh)
			wait.Wait()
		})
	}
}

func TestXdsServer_GetCache(t *testing.T) {
	tests := []struct {
		name    string
		xdss    *XdsServer
		want    xdss.Cache
		version envoy.APIVersion
	}{
		{
			"Gets the server's Cache",
			&XdsServer{
				context.Background(),
				10000,
				&tls.Config{},
				server_v3.NewServer(context.Background(), snapshotCacheV3, &xdss_v3.Callbacks{Logger: ctrl.Log}),
				snapshotCacheV3,
				&xdss_v3.Callbacks{Logger: ctrl.Log},
				stats.New(),
			},
			xdss_v3.NewCache(snapshotCacheV3),
			envoy.APIv3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.xdss.GetCache(tt.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XdsServer.GetCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
