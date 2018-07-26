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
	crypto_rand "crypto/rand"
	"encoding/hex"
	"fmt"
)

const (
	// the size of the id token
	tokenIDBytes = 3
	// the size of the secret
	tokenSecretBytes = 8
)

// RandomBytes generates some random bytes
func RandomBytes(length int) (string, error) {
	b := make([]byte, length)
	_, err := crypto_rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// NewToken creates and returns a new token
func NewToken() (*Token, error) {
	id, err := RandomBytes(tokenIDBytes)
	if err != nil {
		return nil, err
	}

	secret, err := RandomBytes(tokenSecretBytes)
	if err != nil {
		return nil, err
	}

	return &Token{ID: id, Secret: secret}, nil
}

// Name returns the secret name
func (t *Token) Name() string {
	return fmt.Sprintf("%s%s", bootstrapTokenSecretPrefix, t.ID)
}

// String returns the encoded secret
func (t *Token) String() string {
	return fmt.Sprintf("%s.%s", t.ID, t.Secret)
}
