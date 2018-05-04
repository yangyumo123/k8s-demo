package validation

import (
	"github.com/yangyumo123/k8s-demo/pkg/util"
)

var supportedPortProtocols = util.NewStringSet("TCP", "UDP")
var supportedManifestVersions = util.NewStringSet("v1beta1", "v1beta2")
