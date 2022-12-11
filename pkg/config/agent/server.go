// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Package agent implements a gnmi server to mock a device with YANG models.
package agent

import (
	api "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/wangxn2015/onos-lib-go/pkg/logging"
	"github.com/wangxn2015/onos-lib-go/pkg/northbound"
	"github.com/wangxn2015/onos-ric-sdk-go/pkg/config/configurable"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = logging.GetLogger("gnmi", "agent")

// NewService creates a new gnmi service
func NewService(configurable configurable.Configurable) GnmiService {
	return newService(configurable)
}

// GnmiService :
type GnmiService interface {
	northbound.Service
}

func newService(configurable configurable.Configurable) GnmiService {
	server := &Server{
		configurable: configurable,
	}

	return &Service{
		server: server,
	}
}

// Service is a Service implementation for gnmi service.
type Service struct {
	server *Server
}

// Register registers the Service with the gRPC server.
func (s *Service) Register(r *grpc.Server) {
	server := s.server
	api.RegisterGNMIServer(r, server)
	reflection.Register(r)
}

// GnmiService :
var _ GnmiService = &Service{}
