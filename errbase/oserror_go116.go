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

//go:build go1.16
// +build go1.16

package errbase

import "io/fs"

func registerOsPathErrorMigration() {
	// The os.PathError type was migrated to io.fs.PathError in Go 1.16.
	RegisterTypeMigration("os", "*os.PathError", &fs.PathError{})
}