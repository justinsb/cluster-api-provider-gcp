/*
Copyright 2018 The Kubernetes Authors.

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

package bootstrap

import (
	"k8s.io/api/core/v1"
)

const (
	// bootstrapTokenIDKey is the id of this token. This can be transmitted in the
	// clear and encoded in the name of the secret. It must be a random 6 character
	// string that matches the regexp `^([a-z0-9]{6})$`. Required.
	bootstrapTokenIDKey = "token-id"

	// bootstrapTokenSecretKey is the actual secret. It must be a random 16 character
	// string that matches the regexp `^([a-z0-9]{16})$`. Required.
	bootstrapTokenSecretKey = "token-secret"

	// bootstrapTokenExpirationKey is when this token should be expired and no
	// longer used. A controller will delete this resource after this time. This
	// is an absolute UTC time using RFC3339. If this cannot be parsed, the token
	// should be considered invalid. Optional.
	bootstrapTokenExpirationKey = "expiration"

	// bootstrapTokenUsagePrefix is the prefix for the other usage constants that specifies different
	// functions of a bootstrap token
	bootstrapTokenUsagePrefix = "usage-bootstrap-"

	// bootstrapTokenSecretPrefix is the prefix for bootstrap token names.
	// Bootstrap tokens secrets must be named in the form
	// `bootstrap-token-<token-id>`.  This is the prefix to be used before the
	// token ID.
	bootstrapTokenSecretPrefix = "bootstrap-token-"

	// secretTypeBootstrapToken is used during the automated bootstrap process (first
	// implemented by kubeadm). It stores tokens that are used to sign well known
	// ConfigMaps. They may also eventually be used for authentication.
	secretTypeBootstrapToken v1.SecretType = "bootstrap.kubernetes.io/token"
)

// Token defines a bootstrap token
type Token struct {
	// ID is the id of the token
	ID string
	// Secret is the secret of the token
	Secret string
}
