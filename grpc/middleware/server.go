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

package middleware

import (
	"context"

	"github.com/Navegos/errors"
	"github.com/Navegos/errors/extgrpc"
	"github.com/gogo/status"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, err
	}

	st, ok := status.FromError(err)
	if !ok {
		code := extgrpc.GetGrpcCode(err)
		st = status.New(code, err.Error())
		enc := errors.EncodeError(ctx, err)
		st, err = st.WithDetails(&enc)
		if err != nil {

			// https://jbrandhorst.com/post/grpc-errors/
			// "If this errored, it will always error
			// here, so better panic so we can figure
			// out why than have this silently passing."
			//
			// More specifically, an error here is from ptypes.MarshalAny(detail), which probably
			// means that your proto.Message is not registered with gogoproto.  (Make sure that
			// your error's .pb.go file imports "github.com/gogo/protobuf/proto".)
			//
			// By panicking, we either take down the service or (if it has a recovery middleware) cause
			// the call to fail dramatically.  Either case will draw attention to get it fixed.
			//
			// If we simply returned an errors.AssertionFailed, our entire error stack would vanish
			// as it crosses the network boundary.  A client would receive a grpc status with code.Internal,
			// and the stringification of the error.  This change in behavior could induce subtle bugs
			// in the client since none of the usual errors are being returned.
			//
			// We could also log the error here via whatever appropriate mechanism, but the truth is
			// that the service was seriously misconfigured and shouldn't be running at all.
			//
			panic(err)
		}
	}

	return resp, st.Err()
}
