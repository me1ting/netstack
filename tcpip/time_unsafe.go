// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Check go:linkname function signatures when updating Go version.

package tcpip

import (
	_ "time"   // Used with go:linkname.
	_ "unsafe" // Required for go:linkname.
)

// StdClock implements Clock with the time package.
type StdClock struct{}

var _ Clock = (*StdClock)(nil)

//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

// NowNanoseconds implements Clock.NowNanoseconds.
func (*StdClock) NowNanoseconds() int64 {
	sec, nsec, _ := now()
	return sec*1e9 + int64(nsec)
}

// NowMonotonic implements Clock.NowMonotonic.
func (*StdClock) NowMonotonic() int64 {
	_, _, mono := now()
	return mono
}
