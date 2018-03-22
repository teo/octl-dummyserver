/*
 * === This file is part of octl <https://github.com/teo/octl> ===
 *
 * Copyright 2018 CERN and copyright holders of ALICE O².
 * Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this license CERN does not waive the privileges and
 * immunities granted to it by virtue of its status as an
 * Intergovernmental Organization or submit itself to any jurisdiction.
 */

//go:generate protoc --go_out=plugins=grpc:. protos/octlserver.proto
package core

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"

	pb "github.com/teo/octl-dummyserver/protos"
	"google.golang.org/grpc/reflection"
	"runtime"
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)


func NewServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterOctlServer(s, &RpcServer{
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s
}

func (m *RpcServer) logMethod() {
	if !m.state.config.verbose {
		return
	}
	pc, _, _, ok := runtime.Caller(0)
	if !ok {
		return
	}
	fun := runtime.FuncForPC(pc)
	if fun == nil {
		return
	}
	log.WithPrefix("rpcserver").
		WithField("method", fun.Name()).
		Debug("handling RPC request")
}

// Implements interface pb.OctlServer
type RpcServer struct {

}

func (*RpcServer) TrackStatus(*pb.StatusRequest, pb.Octl_TrackStatusServer) error {
	panic("implement me")
}

func (m *RpcServer) GetFrameworkInfo(context.Context, *pb.GetFrameworkInfoRequest) (*pb.GetFrameworkInfoReply, error) {
	panic("implement me")
}

func (*RpcServer) Teardown(context.Context, *pb.TeardownRequest) (*pb.TeardownReply, error) {
	panic("implement me")
}

func (m *RpcServer) GetEnvironments(context.Context, *pb.GetEnvironmentsRequest) (*pb.GetEnvironmentsReply, error) {
	panic("implement me")
}

func (m *RpcServer) NewEnvironment(cxt context.Context, request *pb.NewEnvironmentRequest) (*pb.NewEnvironmentReply, error) {
	panic("implement me")
}

func (*RpcServer) GetEnvironment(context.Context, *pb.GetEnvironmentRequest) (*pb.GetEnvironmentReply, error) {
	panic("implement me")
}

func (*RpcServer) ControlEnvironment(context.Context, *pb.ControlEnvironmentRequest) (*pb.ControlEnvironmentReply, error) {
	panic("implement me")
}

func (*RpcServer) ModifyEnvironment(context.Context, *pb.ModifyEnvironmentRequest) (*pb.ModifyEnvironmentReply, error) {
	panic("implement me")
}

func (*RpcServer) DestroyEnvironment(context.Context, *pb.DestroyEnvironmentRequest) (*pb.DestroyEnvironmentReply, error) {
	panic("implement me")
}

func (*RpcServer) GetRoles(context.Context, *pb.GetRolesRequest) (*pb.GetRolesReply, error) {
	panic("implement me")
}
