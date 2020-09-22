package profile_control_plane

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package:  "../../../../../stacks/openshift/components/profile-control-plane",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}
