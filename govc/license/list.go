/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

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

package license

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/license"
	"golang.org/x/net/context"
)

var featureUsage = "List licenses with given feature"

type list struct {
	*flags.ClientFlag
	*flags.OutputFlag

	feature string
}

func init() {
	cli.Register("license.list", &list{})
}

func (cmd *list) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.feature, "feature", "", featureUsage)
}

func (cmd *list) Process() error { return nil }

func (cmd *list) Run(f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m := license.NewManager(client)
	result, err := m.List(context.TODO())
	if err != nil {
		return err
	}

	if cmd.feature != "" {
		result = result.WithFeature(cmd.feature)
	}

	return cmd.WriteResult(licenseOutput(result))
}
