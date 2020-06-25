package nodeconfigcache

import (
	"context"
	"reflect"
	"testing"

	marin3rv1alpha1 "github.com/3scale/marin3r/pkg/apis/marin3r/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/status"
	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestReconcileNodeConfigCache_ensureNodeConfigRevision(t *testing.T) {

	t.Run("Creates a new NodeConfigRevision if one does not exist", func(t *testing.T) {
		ncc := &marin3rv1alpha1.NodeConfigCache{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncc",
				Namespace: "default",
			},
			Spec: marin3rv1alpha1.NodeConfigCacheSpec{
				NodeID: "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{
					Endpoints: []marin3rv1alpha1.EnvoyResource{
						{Name: "endpoint", Value: "{\"cluster_name\": \"endpoint\"}"},
					}},
			}}

		cl := fake.NewFakeClient(ncc)
		r := &ReconcileNodeConfigCache{client: cl, scheme: s, adsCache: fakeTestCache()}

		gotErr := r.ensureNodeConfigRevision(context.TODO(), ncc, "xxxx")
		if gotErr != nil {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() error = %v", gotErr)
			return
		}

		ncrList := &marin3rv1alpha1.NodeConfigRevisionList{}
		selector, _ := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
			MatchLabels: map[string]string{
				nodeIDTag:  ncc.Spec.NodeID,
				versionTag: "xxxx",
			},
		})
		r.client.List(context.TODO(), ncrList, &client.ListOptions{LabelSelector: selector})
		if len(ncrList.Items) != 1 {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() - no NodeConfigRevision was created")
			return
		}

		if !apiequality.Semantic.DeepEqual(ncrList.Items[0].Spec.Resources, ncc.Spec.Resources) {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() - resources '%v', want '%v'", &ncrList.Items[0].Spec.Resources, ncc.Spec.Resources)
			return
		}
	})

	t.Run("Publishes an already existent revision", func(t *testing.T) {
		ncr := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr",
				Namespace: "default",
				Labels: map[string]string{
					nodeIDTag:  "node1",
					versionTag: "xxxx",
				},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "xxxx",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
		}
		ncc := &marin3rv1alpha1.NodeConfigCache{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncc",
				Namespace: "default",
			},
			Spec: marin3rv1alpha1.NodeConfigCacheSpec{
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			}}

		cl := fake.NewFakeClient(ncc, ncr)
		r := &ReconcileNodeConfigCache{client: cl, scheme: s, adsCache: fakeTestCache()}

		gotErr := r.ensureNodeConfigRevision(context.TODO(), ncc, "xxxx")
		if gotErr != nil {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() error = %v", gotErr)
			return
		}

		ncrList := &marin3rv1alpha1.NodeConfigRevisionList{}
		selector, _ := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
			MatchLabels: map[string]string{
				nodeIDTag:  ncc.Spec.NodeID,
				versionTag: "xxxx",
			},
		})
		r.client.List(context.TODO(), ncrList, &client.ListOptions{LabelSelector: selector})
		if len(ncrList.Items) != 1 {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() got '%v' ncr objects, expected 1", len(ncrList.Items))
			return
		}
	})

}

func TestReconcileNodeConfigCache_consolidateRevisionList(t *testing.T) {
	t.Run("Consolidates the revision list in the ncc status", func(t *testing.T) {
		ncr := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr",
				Namespace: "default",
				Labels: map[string]string{
					nodeIDTag:  "node1",
					versionTag: "xxxx",
				},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "xxxx",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
		}
		ncc := &marin3rv1alpha1.NodeConfigCache{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncc",
				Namespace: "default",
			},
			Spec: marin3rv1alpha1.NodeConfigCacheSpec{
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
		}

		cl := fake.NewFakeClient(ncc, ncr)
		r := &ReconcileNodeConfigCache{client: cl, scheme: s, adsCache: fakeTestCache()}

		gotErr := r.consolidateRevisionList(context.TODO(), ncc, "xxxx")
		if gotErr != nil {
			t.Errorf("TestReconcileNodeConfigCache_consolidateRevisionList() error = %v", gotErr)
			return
		}

		gotNCC := &marin3rv1alpha1.NodeConfigCache{}
		wantConfigRevisions := []marin3rv1alpha1.ConfigRevisionRef{
			{Version: "xxxx", Ref: corev1.ObjectReference{Name: "ncr", Namespace: "default"}},
		}
		r.client.Get(context.TODO(), types.NamespacedName{Name: "ncc", Namespace: "default"}, gotNCC)

		if !apiequality.Semantic.DeepEqual(gotNCC.Status.ConfigRevisions, wantConfigRevisions) {
			t.Errorf("TestReconcileNodeConfigCache_consolidateRevisionList() got '%v', want '%v'", gotNCC.Status.ConfigRevisions, wantConfigRevisions)
			return
		}
	})

	t.Run("Moves the published revision to the last position of the list", func(t *testing.T) {
		ncr1 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr1",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "1"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "1",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
		}
		ncr2 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr2",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "2"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "2",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
		}
		ncr3 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr3",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "3"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "3",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
		}
		ncc := &marin3rv1alpha1.NodeConfigCache{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncc",
				Namespace: "default",
			},
			Spec: marin3rv1alpha1.NodeConfigCacheSpec{
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
			Status: marin3rv1alpha1.NodeConfigCacheStatus{
				ConfigRevisions: []marin3rv1alpha1.ConfigRevisionRef{
					{Version: "1", Ref: corev1.ObjectReference{Name: "ncr1", Namespace: "default"}},
					{Version: "2", Ref: corev1.ObjectReference{Name: "ncr2", Namespace: "default"}},
					{Version: "3", Ref: corev1.ObjectReference{Name: "ncr3", Namespace: "default"}},
				},
			},
		}

		cl := fake.NewFakeClient(ncc, ncr1, ncr2, ncr3)
		r := &ReconcileNodeConfigCache{client: cl, scheme: s, adsCache: fakeTestCache()}

		gotErr := r.consolidateRevisionList(context.TODO(), ncc, "1")
		if gotErr != nil {
			t.Errorf("TestReconcileNodeConfigCache_consolidateRevisionList() error = %v", gotErr)
			return
		}

		gotNCC := &marin3rv1alpha1.NodeConfigCache{}
		wantConfigRevisions := []marin3rv1alpha1.ConfigRevisionRef{
			{Version: "2", Ref: corev1.ObjectReference{Name: "ncr2", Namespace: "default"}},
			{Version: "3", Ref: corev1.ObjectReference{Name: "ncr3", Namespace: "default"}},
			{Version: "1", Ref: corev1.ObjectReference{Name: "ncr1", Namespace: "default"}},
		}
		r.client.Get(context.TODO(), types.NamespacedName{Name: "ncc", Namespace: "default"}, gotNCC)

		if !apiequality.Semantic.DeepEqual(gotNCC.Status.ConfigRevisions, wantConfigRevisions) {
			t.Errorf("TestReconcileNodeConfigCache_consolidateRevisionList() got '%v', want '%v'", gotNCC.Status.ConfigRevisions, wantConfigRevisions)
			return
		}
	})
}

func TestReconcileNodeConfigCache_deleteUnreferencedRevisions(t *testing.T) {
	type args struct {
		ctx context.Context
		ncc *marin3rv1alpha1.NodeConfigCache
	}
	tests := []struct {
		name    string
		r       *ReconcileNodeConfigCache
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.deleteUnreferencedRevisions(tt.args.ctx, tt.args.ncc); (err != nil) != tt.wantErr {
				t.Errorf("ReconcileNodeConfigCache.deleteUnreferencedRevisions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReconcileNodeConfigCache_markRevisionPublished(t *testing.T) {
	t.Run("Keeps current revision published", func(t *testing.T) {
		ncr1 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr1",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "1"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "1",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
			Status: marin3rv1alpha1.NodeConfigRevisionStatus{
				Conditions: []status.Condition{{Type: marin3rv1alpha1.RevisionPublishedCondition, Status: corev1.ConditionFalse}},
			},
		}

		ncr2 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr2",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "2"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "2",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
			Status: marin3rv1alpha1.NodeConfigRevisionStatus{
				Conditions: []status.Condition{{Type: marin3rv1alpha1.RevisionPublishedCondition, Status: corev1.ConditionTrue}},
			},
		}

		cl := fake.NewFakeClient(ncr1, ncr2)
		r := &ReconcileNodeConfigCache{client: cl, scheme: s, adsCache: fakeTestCache()}

		gotErr := r.markRevisionPublished(context.TODO(), "node1", "2", "reason", "msg")
		if gotErr != nil {
			t.Errorf("TestReconcileNodeConfigCache_markRevisionPublished() error = %v", gotErr)
			return
		}

		ncr := &marin3rv1alpha1.NodeConfigRevision{}

		// ncr2 should still be marked as published
		r.client.Get(context.TODO(), types.NamespacedName{Name: "ncr2", Namespace: "default"}, ncr)
		if !ncr.Status.Conditions.IsTrueFor(marin3rv1alpha1.RevisionPublishedCondition) {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() - ncr2 RevisionPublishedCondition != True or missing")
		}

		// ncr1 should not be marked as published
		r.client.Get(context.TODO(), types.NamespacedName{Name: "ncr1", Namespace: "default"}, ncr)
		if ncr.Status.Conditions.IsTrueFor(marin3rv1alpha1.RevisionPublishedCondition) {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() - ncr1 RevisionPublishedCondition == True")
		}
	})

	t.Run("Changes the published revision", func(t *testing.T) {
		ncr1 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr1",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "1"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "1",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
			Status: marin3rv1alpha1.NodeConfigRevisionStatus{
				Conditions: []status.Condition{{Type: marin3rv1alpha1.RevisionPublishedCondition, Status: corev1.ConditionFalse}},
			},
		}

		ncr2 := &marin3rv1alpha1.NodeConfigRevision{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "ncr2",
				Namespace: "default",
				Labels:    map[string]string{nodeIDTag: "node1", versionTag: "2"},
			},
			Spec: marin3rv1alpha1.NodeConfigRevisionSpec{
				Version:   "2",
				NodeID:    "node1",
				Resources: &marin3rv1alpha1.EnvoyResources{},
			},
			Status: marin3rv1alpha1.NodeConfigRevisionStatus{
				Conditions: []status.Condition{{Type: marin3rv1alpha1.RevisionPublishedCondition, Status: corev1.ConditionTrue}},
			},
		}

		cl := fake.NewFakeClient(ncr1, ncr2)
		r := &ReconcileNodeConfigCache{client: cl, scheme: s, adsCache: fakeTestCache()}

		gotErr := r.markRevisionPublished(context.TODO(), "node1", "1", "reason", "msg")
		if gotErr != nil {
			t.Errorf("TestReconcileNodeConfigCache_markRevisionPublished() error = %v", gotErr)
			return
		}

		ncr := &marin3rv1alpha1.NodeConfigRevision{}

		// ncr2 should not be marked as published
		r.client.Get(context.TODO(), types.NamespacedName{Name: "ncr2", Namespace: "default"}, ncr)
		if ncr.Status.Conditions.IsTrueFor(marin3rv1alpha1.RevisionPublishedCondition) {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() - ncr2 RevisionPublishedCondition == True")
		}

		// ncr1 should not be marked as published
		r.client.Get(context.TODO(), types.NamespacedName{Name: "ncr1", Namespace: "default"}, ncr)
		if !ncr.Status.Conditions.IsTrueFor(marin3rv1alpha1.RevisionPublishedCondition) {
			t.Errorf("TestReconcileNodeConfigCache_ensureNodeConfigRevision() - ncr1 RevisionPublishedCondition != True or missing")
		}
	})
}

func Test_trimRevisions(t *testing.T) {
	type args struct {
		list []marin3rv1alpha1.ConfigRevisionRef
		max  int
	}
	tests := []struct {
		name string
		args args
		want []marin3rv1alpha1.ConfigRevisionRef
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimRevisions(tt.args.list, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimRevisions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRevisionIndex(t *testing.T) {
	type args struct {
		version   string
		revisions []marin3rv1alpha1.ConfigRevisionRef
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRevisionIndex(tt.args.version, tt.args.revisions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRevisionIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveRevisionToLast(t *testing.T) {
	type args struct {
		list []marin3rv1alpha1.ConfigRevisionRef
		idx  int
	}
	tests := []struct {
		name string
		args args
		want []marin3rv1alpha1.ConfigRevisionRef
	}{
		{
			name: "Moves the revision to the last position in the list",
			args: args{
				list: []marin3rv1alpha1.ConfigRevisionRef{
					{Version: "1", Ref: corev1.ObjectReference{}},
					{Version: "2", Ref: corev1.ObjectReference{}},
					{Version: "3", Ref: corev1.ObjectReference{}},
					{Version: "4", Ref: corev1.ObjectReference{}},
					{Version: "5", Ref: corev1.ObjectReference{}},
					{Version: "6", Ref: corev1.ObjectReference{}},
				},
				idx: 3,
			},
			want: []marin3rv1alpha1.ConfigRevisionRef{
				{Version: "1", Ref: corev1.ObjectReference{}},
				{Version: "2", Ref: corev1.ObjectReference{}},
				{Version: "3", Ref: corev1.ObjectReference{}},
				{Version: "5", Ref: corev1.ObjectReference{}},
				{Version: "6", Ref: corev1.ObjectReference{}},
				{Version: "4", Ref: corev1.ObjectReference{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveRevisionToLast(tt.args.list, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveRevisionToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
