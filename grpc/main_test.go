// Copyright (C) 2025 @Navegos & @DevelVitorF Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package actor provides the structures for representing an actor who has
// access to resources.

package grpc

import (
	"net"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/Navegos/errors/grpc/middleware"
	"github.com/hydrogen18/memlistener"
)

var (
	Client EchoerClient
)

func TestMain(m *testing.M) {

	srv := &EchoServer{}

	lis := memlistener.NewMemoryListener()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryServerInterceptor))
	RegisterEchoerServer(grpcServer, srv)

	go grpcServer.Serve(lis)

	dialOpts := []grpc.DialOption{
		grpc.WithDialer(func(target string, d time.Duration) (net.Conn, error) {
			return lis.Dial("", "")
		}),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(middleware.UnaryClientInterceptor),
	}

	clientConn, err := grpc.Dial("", dialOpts...)
	if err != nil {
		panic(err)
	}

	Client = NewEchoerClient(clientConn)

	code := m.Run()

	grpcServer.Stop()
	clientConn.Close()

	os.Exit(code)
}
