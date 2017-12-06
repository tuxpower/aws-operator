package namespacev1

import (
	"context"

	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"

	"github.com/giantswarm/aws-operator/service/keyv1"
)

func (r *Resource) GetCurrentState(ctx context.Context, obj interface{}) (interface{}, error) {
	customObject, err := keyv1.ToCustomObject(obj)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, "debug", "looking for the namespace in the Kubernetes API")

	// Lookup the current state of the namespace.
	var namespace *apiv1.Namespace
	{
		manifest, err := r.k8sClient.CoreV1().Namespaces().Get(keyv1.ClusterNamespace(customObject), apismetav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			r.logger.LogCtx(ctx, "debug", "did not find the namespace in the Kubernetes API")
			// fall through
		} else if err != nil {
			return nil, microerror.Mask(err)
		} else {
			r.logger.LogCtx(ctx, "debug", "found the namespace in the Kubernetes API")
			namespace = manifest
		}
	}

	return namespace, nil
}
