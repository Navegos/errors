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

package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Navegos/errors"
	"github.com/Navegos/errors/testutils"
)

func TestUnwrap(t *testing.T) {
	tt := testutils.T{t}

	e := fmt.Errorf("foo %w %w", fmt.Errorf("bar"), fmt.Errorf("baz"))

	// Compatibility with go 1.20: Unwrap() on a multierror returns nil
	// (per API documentation)
	tt.Check(errors.Unwrap(e) == nil)
}

// More detailed testing of Join is in datadriven_test.go. Here we make
// sure that the public API includes the stacktrace wrapper.
func TestJoin(t *testing.T) {
	e := errors.Join(errors.New("abc123"), errors.New("def456"))
	printed := fmt.Sprintf("%+v", e)
	expected := `Error types: (1) *withstack.withStack (2) *join.joinError (3) *withstack.withStack (4) *errutil.leafError (5) *withstack.withStack (6) *errutil.leafError`
	if !strings.Contains(printed, expected) {
		t.Errorf("Expected: %s to contain: %s", printed, expected)
	}
}
