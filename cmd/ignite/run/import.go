package run

import (
	api "github.com/weaveworks/ignite/pkg/apis/ignite"
	meta "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	"github.com/weaveworks/ignite/pkg/metadata"
	"github.com/weaveworks/ignite/pkg/operations"
	"github.com/weaveworks/ignite/pkg/providers"
	"github.com/weaveworks/ignite/pkg/util"
)

func ImportImage(source string) (image *api.Image, err error) {
	ociRef, err := meta.NewOCIImageRef(source)
	if err != nil {
		return
	}

	image, err = operations.FindOrImportImage(providers.Client, ociRef)
	if err != nil {
		return
	}
	defer util.DeferErr(&err, func() error { return metadata.Cleanup(image, false) })

	err = metadata.Success(image)

	return
}

func ImportKernel(source string) (kernel *api.Kernel, err error) {
	ociRef, err := meta.NewOCIImageRef(source)
	if err != nil {
		return
	}

	kernel, err = operations.FindOrImportKernel(providers.Client, ociRef)
	if err != nil {
		return
	}
	defer util.DeferErr(&err, func() error { return metadata.Cleanup(kernel, false) })

	err = metadata.Success(kernel)

	return
}
