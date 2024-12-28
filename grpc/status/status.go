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

package status

import (
	"github.com/Navegos/errors"
	"github.com/Navegos/errors/extgrpc"

	"google.golang.org/grpc/codes"
)

func Error(c codes.Code, msg string) error {
	return extgrpc.WrapWithGrpcCode(errors.New(msg), c)
}

func Errorf(c codes.Code, format string, args ...interface{}) error {
	return extgrpc.WrapWithGrpcCode(errors.Newf(format, args...), c)
}

func WrapErr(c codes.Code, msg string, err error) error {
	return extgrpc.WrapWithGrpcCode(errors.WrapWithDepth(1, err, msg), c)
}

func WrapErrf(c codes.Code, err error, format string, args ...interface{}) error {
	return extgrpc.WrapWithGrpcCode(errors.WrapWithDepthf(1, err, format, args...), c)
}

func Code(err error) codes.Code {
	return extgrpc.GetGrpcCode(err)
}
