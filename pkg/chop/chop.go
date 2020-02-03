// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chop

import (
	"github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	chopclientset "github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned"
)

type Chop struct {
	Version       string
	ConfigManager *ConfigManager
}

func NewChop(
	version string,
	chopClient *chopclientset.Clientset,
	initConfigFilePath string,

) *Chop {
	return &Chop{
		Version:       version,
		ConfigManager: NewConfigManager(chopClient, initConfigFilePath),
	}
}

func (c *Chop) Init() error {
	return c.ConfigManager.Init()
}

func (c *Chop) Config() *v1.OperatorConfig {
	return c.ConfigManager.Config()
}