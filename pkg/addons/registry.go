// SPDX-FileCopyrightText: The kubectl-gather authors
// SPDX-License-Identifier: Apache-2.0

package addons

import (
	"net/http"

	"k8s.io/client-go/rest"

	"github.com/nirs/kubectl-gather/pkg/gather"
)

func Registry(config *rest.Config, client *http.Client, out *gather.Output, opts *gather.Options, q gather.Queuer) (map[string]gather.Addon, error) {
	logsAddon, err := NewLogsAddon(config, client, out, opts, q)
	if err != nil {
		return nil, err
	}

	rookAddon, err := NewRookCephAddon(config, client, out, opts, q)
	if err != nil {
		return nil, err
	}

	registry := map[string]gather.Addon{
		"pods":                      logsAddon,
		"cephclusters.ceph.rook.io": rookAddon,
	}

	return registry, nil
}
