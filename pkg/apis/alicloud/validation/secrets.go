// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"fmt"

	"github.com/gardener/gardener-extension-provider-alicloud/pkg/alicloud"
	corev1 "k8s.io/api/core/v1"
)

const (
	accessKeyIDMinLen     = 16
	accessKeyIDMaxLen     = 128
	accessKeySecretMinLen = 30
)

// ValidateCloudProviderSecret checks whether the given secret contains a valid Alicloud access keys.
func ValidateCloudProviderSecret(secret *corev1.Secret) error {
	secretRef := fmt.Sprintf("%s/%s", secret.Namespace, secret.Name)

	accessKeyID, ok := secret.Data[alicloud.AccessKeyID]
	if !ok {
		return fmt.Errorf("missing %q field in secret %s", alicloud.AccessKeyID, secretRef)
	}

	if len(accessKeyID) < accessKeyIDMinLen {
		return fmt.Errorf("field %q in secret %s must have at least %d characters", alicloud.AccessKeyID, secretRef, accessKeyIDMinLen)
	}

	if len(accessKeyID) > accessKeyIDMaxLen {
		return fmt.Errorf("field %q in secret %s cannot be longer than %d characters", alicloud.AccessKeyID, secretRef, accessKeyIDMaxLen)
	}

	secretAccessKey, ok := secret.Data[alicloud.AccessKeySecret]
	if !ok {
		return fmt.Errorf("missing %q field in secret %s", alicloud.AccessKeySecret, secretRef)
	}

	if len(secretAccessKey) < accessKeySecretMinLen {
		return fmt.Errorf("field %q in secret %s must have at least %d characters", alicloud.AccessKeySecret, secretRef, accessKeySecretMinLen)
	}

	return nil
}
