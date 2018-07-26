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
	"time"

	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	tokenNamespace = "kube-system"
)

type tokens struct {
	client   kubernetes.Interface
	tokenTTL time.Duration
	usages   []string
}

func NewTokens(client kubernetes.Interface, tokenTTL time.Duration) (Tokens, error) {
	t := &tokens{
		client:   client,
		tokenTTL: tokenTTL,
		usages:   []string{"authentication", "signing"},
	}
	return t, nil
}

/*// safelyProvisionBootstrapToken is responsible for generating a bootstrap token for us
func (n *tokens) provisionBootstrapToken(ctx context.Context, request *NodeRegistration) error {
	usages := []string{"authentication", "signing"}

		token, err := n.createToken(n.config.TokenDuration, usages)
		if err != nil {
			return err
		}
	return token.String()

}*/

func (n *tokens) GetBootstrapToken() (string, error) {
	/*token, err := n.findToken()
	if err != nil {
		return "", err
	}*/
	var token *Token
	var err error

	if token == nil {
		token, err = n.createToken(n.tokenTTL, n.usages)
		if err != nil {
			return "", err
		}
	}

	// Temporary - for debugging
	glog.Infof("created token %q", token.String())

	return token.String(), nil
}

// createToken generates a token for the instance
func (n *tokens) createToken(expiration time.Duration, usages []string) (*Token, error) {
	for {
		// @step: generate a random token for them
		token, err := NewToken()
		if err != nil {
			return nil, err
		}

		// @step: check if the token already exist, remote but a possibility
		if exists, err := n.tokenExists(token.Name()); err != nil {
			return nil, err
		} else if exists {
			glog.Warningf("duplicate token found: %s, skipping", token.ID)
			time.Sleep(10 * time.Millisecond)
			continue
		}

		// @step: add the secret to the namespace
		v1secret := &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name: token.Name(),
				Labels: map[string]string{
					"name": token.Name(),
				},
			},
			Type: v1.SecretType(secretTypeBootstrapToken),
			Data: encodeTokenSecretData(token, usages, expiration),
		}

		if _, err := n.client.CoreV1().Secrets(tokenNamespace).Create(v1secret); err != nil {
			return nil, err
		}

		return token, err
	}
}

// hasToken checks if the tokens already exists
func (n *tokens) tokenExists(tokenName string) (bool, error) {
	resp, err := n.client.CoreV1().Secrets(tokenNamespace).List(metav1.ListOptions{
		LabelSelector: "name=" + tokenName,
		Limit:         1,
	})
	if err != nil {
		return false, err
	}

	return len(resp.Items) > 0, nil
}

// encodeTokenSecretData takes the token discovery object and an optional duration and returns the .Data for the Secret
func encodeTokenSecretData(token *Token, usages []string, ttl time.Duration) map[string][]byte {
	data := map[string][]byte{
		bootstrapTokenIDKey:     []byte(token.ID),
		bootstrapTokenSecretKey: []byte(token.Secret),
	}

	if ttl > 0 {
		expire := time.Now().Add(ttl).Format(time.RFC3339)
		data[bootstrapTokenExpirationKey] = []byte(expire)
	}

	for _, usage := range usages {
		data[bootstrapTokenUsagePrefix+usage] = []byte("true")
	}

	return data
}
