// Copyright 2019 eBay Inc.
// Primary authors: Simon Fell, Diego Ongaro,
//                  Raymond Kroeker, and Sathish Kandasamy.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"github.com/ebay/beam/api"
	log "github.com/sirupsen/logrus"
)

func wipe(ctx context.Context, store api.BeamFactStoreClient, options *options) error {
	req := &api.WipeRequest{
		WaitFor: options.WipeWaitFor,
	}
	log.Infof("Invoking wipe: %+v", req)
	start := time.Now()
	res, err := store.Wipe(ctx, req)
	if err != nil {
		return err
	}
	log.Infof("Wipe returned: %+v", res)
	log.Infof("Wipe took %s", time.Since(start))
	return nil
}
