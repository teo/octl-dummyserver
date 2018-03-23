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
package octl_dummyserver

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"

	pb "github.com/teo/octl-dummyserver/protos"
	"google.golang.org/grpc/reflection"
	"runtime"
	"fmt"
	"time"
	"github.com/pborman/uuid"
)


func NewServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterOctlServer(s, &RpcServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s
}

func (m *RpcServer) logMethod() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return
	}
	fun := runtime.FuncForPC(pc)
	if fun == nil {
		return
	}
	fmt.Printf("handling RPC request %s\n", fun.Name())
}

// Implements interface pb.OctlServer
type RpcServer struct {}

func (m *RpcServer) TrackStatus(req *pb.StatusRequest, statusServer pb.Octl_TrackStatusServer) error {
	m.logMethod()
	for {
		err := statusServer.Send(&pb.StatusReply{
			State: "ok",
			StatusUpdates: []*pb.StatusUpdate{
				{
					Level: pb.StatusUpdate_INFO,
					Event: &pb.StatusUpdate_GenericMessage{
						GenericMessage: &pb.Event_GenericMessage{
							Message: fmt.Sprintf("sending update at %s", time.Now().String()),
						},
					},
				},
			},
		})
		if err != nil {
			fmt.Printf(err.Error())
			return err
		} else {
			fmt.Printf("stream update sent\n")
		}

		time.Sleep(1000*time.Millisecond)
	}
	return nil
}

func (m *RpcServer) GetFrameworkInfo(context.Context, *pb.GetFrameworkInfoRequest) (*pb.GetFrameworkInfoReply, error) {
	m.logMethod()
	return &pb.GetFrameworkInfoReply{
		State: "CONFIGURED",
		FrameworkId: uuid.NewUUID().String(),
		RolesCount: 42,
		EnvironmentsCount: 6,
	}, nil
}

func (m *RpcServer) Teardown(context.Context, *pb.TeardownRequest) (*pb.TeardownReply, error) {
	m.logMethod()
	return &pb.TeardownReply{}, nil
}

func (m *RpcServer) GetEnvironments(context.Context, *pb.GetEnvironmentsRequest) (*pb.GetEnvironmentsReply, error) {
	m.logMethod()
	return &pb.GetEnvironmentsReply{
		FrameworkId: uuid.NewUUID().String(),
		Environments: []*pb.EnvironmentInfo{
			{
				State: "CONFIGURED",
				Roles: []string{"flp1", "flp2", "epn3"},
				Id: uuid.NewUUID().String(),
				CreatedWhen: time.Now().String(),
			},
			{
				State: "RUNNING",
				Roles: []string{"flp3", "flp4", "epn1", "epn2"},
				Id: uuid.NewUUID().String(),
				CreatedWhen: time.Now().String(),
			},
		},
	}, nil
}

func (m *RpcServer) NewEnvironment(cxt context.Context, request *pb.NewEnvironmentRequest) (*pb.NewEnvironmentReply, error) {
	m.logMethod()
	return &pb.NewEnvironmentReply{
		Id: uuid.NewUUID().String(),
		State: "CONFIGURED",
	}, nil
}

func (m *RpcServer) GetEnvironment(context.Context, *pb.GetEnvironmentRequest) (*pb.GetEnvironmentReply, error) {
	m.logMethod()
	return &pb.GetEnvironmentReply{
		Environment: &pb.EnvironmentInfo{
			State: "CONFIGURED",
			Roles: []string{"flp1", "flp2", "epn3"},
			Id: uuid.NewUUID().String(),
			CreatedWhen: time.Now().String(),
		},
	}, nil
}

func (m *RpcServer) ControlEnvironment(context.Context, *pb.ControlEnvironmentRequest) (*pb.ControlEnvironmentReply, error) {
	m.logMethod()
	return &pb.ControlEnvironmentReply{}, nil
}

func (m *RpcServer) ModifyEnvironment(context.Context, *pb.ModifyEnvironmentRequest) (*pb.ModifyEnvironmentReply, error) {
	m.logMethod()
	return &pb.ModifyEnvironmentReply{
		FailedOperations: nil,
	}, nil
}

func (m *RpcServer) DestroyEnvironment(context.Context, *pb.DestroyEnvironmentRequest) (*pb.DestroyEnvironmentReply, error) {
	m.logMethod()
	return &pb.DestroyEnvironmentReply{}, nil
}

func (m *RpcServer) GetRoles(context.Context, *pb.GetRolesRequest) (*pb.GetRolesReply, error) {
	m.logMethod()
	return &pb.GetRolesReply{
		Roles: []*pb.RoleInfo{
			{
				Name: "flp1",
				State: "RUNNING",
				Condition: "ACTIVE",
				Hostname: "flp1.cern.ch",
			},
		},
	}, nil
}
