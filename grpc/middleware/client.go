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
	"github.com/gogo/status"

	"google.golang.org/grpc"
)

func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	err := invoker(ctx, method, req, reply, cc, opts...)

	st := status.Convert(err)
	var reconstituted error
	for _, det := range st.Details() {
		switch t := det.(type) {
		case *errors.EncodedError:
			reconstituted = errors.DecodeError(ctx, *t)
		}
	}

	if reconstituted != nil {
		err = reconstituted
	}

	return err
}
