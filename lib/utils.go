/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lib

import (
	"fmt"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

func ParseEndpoint(ep string) (string, string, error) {
	if strings.HasPrefix(strings.ToLower(ep), "unix://") || strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		return s[0], s[1], nil
	}
	return "", "", fmt.Errorf("Invalid endpoint: %v", ep)
}

func GetVersionString(v *csi.Version) string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func GetVersionFromString(v string) (*csi.Version, error) {
	var major, minor, patch uint32

	n, err := fmt.Sscanf(v, "%d.%d.%d", &major, &minor, &patch)
	if err != nil {
		return nil, err
	}
	if n != 3 {
		return nil, fmt.Errorf("Invalid format. Specify version in x.y.z format")
	}

	return &csi.Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func NewVolumeCapabilityAccessMode(mode csi.VolumeCapability_AccessMode_Mode) *csi.VolumeCapability_AccessMode {
	return &csi.VolumeCapability_AccessMode{mode}
}

func NewDefaultNodeServer(d *CSIDriver) *NodeServerDefaults {
	return &NodeServerDefaults{
		Driver: d,
	}
}

func NewDefaultIdentityServer(d *CSIDriver) *IdentityServerDefaults {
	return &IdentityServerDefaults{
		Driver: d,
	}
}

func NewDefaultControllerServer(d *CSIDriver) *ControllerServerDefaults {
	return &ControllerServerDefaults{
		Driver: d,
	}
}

func NewControllerServiceCapability(cap csi.ControllerServiceCapability_RPC_Type) *csi.ControllerServiceCapability {
	return &csi.ControllerServiceCapability{
		&csi.ControllerServiceCapability_Rpc{
			&csi.ControllerServiceCapability_RPC{
				cap,
			},
		},
	}
}
