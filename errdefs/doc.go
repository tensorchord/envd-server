// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package errdefs defines a set of error interfaces that packages should use for communicating classes of errors.
// Errors that cross the package boundary should implement one (and only one) of these interfaces.
//
// Packages should not reference these interfaces directly, only implement them.
// To check if a particular error implements one of these interfaces, there are helper
// functions provided (e.g. `Is<SomeError>`) which can be used rather than asserting the interfaces directly.
// If you must assert on these interfaces, be sure to check the causal chain (`err.Cause()`).
package errdefs // import "github.com/docker/docker/errdefs"
