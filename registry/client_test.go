package registry_test

import (
	"context"
	"testing"

	"github.com/kayac/ecspresso/registry"
)

var testImages = []string{
	"debian",
	"katsubushi/katsubushi:v1.6.0",
	"public.ecr.aws/mackerel/mackerel-container-agent:plugins",
	"gcr.io/kaniko-project/executor:v0.10.0",
	"ghcr.io/fujiwara/printenv:v0.0.2",
}

func TestImages(t *testing.T) {
	for _, img := range testImages {
		t.Logf("testing %s", img)
		if ok, err := registry.ExistsImage(context.Background(), img, "", ""); err != nil {
			t.Errorf("%s error %s", img, err)
		} else if !ok {
			t.Errorf("%s not found", img)
		}
	}
}
