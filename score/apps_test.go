package score

import (
	"testing"
)

func TestDeploymentHasPodAntiAffinityPreffered(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "deployment-host-antiaffinity-preffered.yaml", "Deployment has host PodAntiAffinity", 10)
}

func TestDeploymentHasPodAntiAffinityPrefferedNoSelectorMatch(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "deployment-host-antiaffinity-preffered-selector-no-match.yaml", "Deployment has host PodAntiAffinity", 5)
}

func TestDeploymentHasPodAntiAffinityPrefferedSelectorExpression(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "deployment-host-antiaffinity-preffered-selector-expression.yaml", "Deployment has host PodAntiAffinity", 10)
}

func TestDeploymentHasPodAntiAffinityRequired(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "deployment-host-antiaffinity-required.yaml", "Deployment has host PodAntiAffinity", 10)
}

func TestDeploymentHasPodAntiAffinityNotSet(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "deployment-host-antiaffinity-not-set.yaml", "Deployment has host PodAntiAffinity", 5)
}

func TestDeploymentHasPodAntiAffinityOneReplica(t *testing.T) {
	t.Parallel()
	// skipped
	testExpectedScore(t, "deployment-host-antiaffinity-1-replica.yaml", "Deployment has host PodAntiAffinity", 0)
}

func TestStatefulSetHasPodAntiAffinityPreffered(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "statefulset-host-antiaffinity-preffered.yaml", "StatefulSet has host PodAntiAffinity", 10)
}

func TestStatefulSetHasPodAntiAffinityRequired(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "statefulset-host-antiaffinity-required.yaml", "StatefulSet has host PodAntiAffinity", 10)
}

func TestStatefulSetHasPodAntiAffinityNotSet(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "statefulset-host-antiaffinity-not-set.yaml", "StatefulSet has host PodAntiAffinity", 5)
}

func TestStatefulSetHasPodAntiAffinityOneReplica(t *testing.T) {
	t.Parallel()
	// skipped
	testExpectedScore(t, "statefulset-host-antiaffinity-1-replica.yaml", "StatefulSet has host PodAntiAffinity", 0)
}

func TestStatefulSetHasPodAntiAffinityUndefinedReplicas(t *testing.T) {
	t.Parallel()
	testExpectedScore(t, "statefulset-host-antiaffinity-undefined-replicas.yaml", "StatefulSet has host PodAntiAffinity", 5)
}
