package registry

import (
	"context"
	"log"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/pkg/errors"
)

func ExistsImage(ctx context.Context, image, user, password string) (bool, error) {
	ref, err := name.ParseReference(image)
	if err != nil {
		return false, err
	}
	auth := authn.Anonymous
	if user == "AWS" && password != "" {
		auth = &authn.Bearer{Token: password}
	}
	log.Println("auth:", auth)
	_, err = remote.Image(ref, remote.WithAuth(auth), remote.WithContext(ctx), remote.WithJobs(1))
	if err != nil {
		return false, errors.Wrapf(err, "failed to check remote image %s", image)
	}
	return true, nil
}
