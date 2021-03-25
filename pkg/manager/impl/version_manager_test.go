package impl

import (
	"context"
	"testing"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/stretchr/testify/assert"
)

var (
	build      = "fa70a14"
	buildTime  = "2021-03-25"
	appversion = "v0.3.7-47-gfa70a14"
)

func TestVersionManager_GetVersion(t *testing.T) {
	vmanager := &VersionManager{
		Build:     build,
		Version:   appversion,
		BuildTime: buildTime,
	}
	v, err := vmanager.GetVersion(context.Background(), &admin.GetVersionRequest{})
	assert.Nil(t, err)
	assert.Equal(t, v.ControlPlaneVersion.BuildTime, buildTime)
	assert.Equal(t, v.ControlPlaneVersion.Build, build)
	assert.Equal(t, v.ControlPlaneVersion.Version, appversion)
}
