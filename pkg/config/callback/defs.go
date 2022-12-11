// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package callback

import (
	"github.com/wangxn2015/onos-lib-go/pkg/logging"
	"github.com/wangxn2015/onos-ric-sdk-go/pkg/config/configurable"
	"github.com/wangxn2015/onos-ric-sdk-go/pkg/config/store"
)

var log = logging.GetLogger("config", "callback")

// Config :
type Config struct {
	config *store.ConfigStore
}

// InitConfig :
func (c *Config) InitConfig(config *store.ConfigStore) {
	c.config = config

}

var _ configurable.Configurable = &Config{}
