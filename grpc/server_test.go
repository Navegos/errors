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
	"context"

	"github.com/Navegos/errors"
	"github.com/Navegos/errors/grpc/status"

	"google.golang.org/grpc/codes"
)

var ErrCantEcho = errors.New("unable to echo")
var ErrTooLong = errors.New("text is too long")
var ErrInternal = errors.New("internal error!")

type EchoServer struct {
}

func (srv *EchoServer) Echo(ctx context.Context, req *EchoRequest) (*EchoReply, error) {
	msg := req.Text
	switch {
	case msg == "noecho":
		return nil, ErrCantEcho
	case len(msg) > 10:
		return nil, errors.WithMessage(ErrTooLong, msg+" is too long")
	case msg == "reverse":
		return nil, status.Error(codes.Unimplemented, "reverse is not implemented")
	case msg == "internal":
		return nil, status.WrapErr(codes.Internal, "there was a problem", ErrInternal)
	}
	return &EchoReply{
		Reply: "echoing: " + msg,
	}, nil
}
