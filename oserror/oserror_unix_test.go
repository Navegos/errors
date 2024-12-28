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

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package oserror

import (
	"syscall"
	"testing"

	"github.com/Navegos/errors"
	"github.com/Navegos/errors/testutils"
)

func TestErrorPredicatesUnix(t *testing.T) {
	tt := testutils.T{T: t}

	tt.Check(IsPermission(errors.Wrap(syscall.EACCES, "woo")))
	tt.Check(IsExist(errors.Wrap(syscall.ENOTEMPTY, "woo")))
	tt.Check(IsNotExist(errors.Wrap(syscall.ENOENT, "woo")))
	tt.Check(IsTimeout(errors.Wrap(syscall.EAGAIN, "woo")))
}