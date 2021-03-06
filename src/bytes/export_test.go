// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

// Export func for testing
var IndexBytePortable = indexBytePortable
var EqualPortable = equalPortable

func (b *Buffer) Cap() int {
	return cap(b.buf)
}
