package tests

import (
	"testing"

	pods "github.com/mv11l/kubenet-ai/internal/k8s"
	"github.com/stretchr/testify/assert"
)

func TestGetPods(t *testing.T) {
	t.Run("GetPods return arrays of pods with valid ns", func(t *testing.T) {
		t.Parallel()
		pods := pods.GetPodsFromNs("kube-system")
		assert.NotEmpty(t, pods)
	})
	t.Run("GetPods return empty with invalid ns", func(t *testing.T) {
		t.Parallel()
		pods := pods.GetPodsFromNs("doestexistns")
		assert.Empty(t, pods)
	})
}
