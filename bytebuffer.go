package bytebufferpool

import "io"

// ByteBuffer provides byte buffer, which can be used for minimizing
// memory allocations.
//
// ByteBuffer may be used with functions appending data to the given []byte
// slice. See example code for details.
//
// Use Get for obtaining an empty byte buffer.
type ByteBuffer struct {

	// B is a byte buffer to use in append-like workloads.
	// See example code for details.
	B []byte
}

// Len returns the size of the byte buffer.
func (b *ByteBuffer) Len() int {
	_ = "STUB: not implemented"

	// ReadFrom implements io.ReaderFrom.
	//
	// The function appends all the data read from r to b.
	return 0
}

func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteTo implements io.WriterTo.
func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Bytes returns b.B, i.e. all the bytes accumulated in the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
func (b *ByteBuffer) Bytes() []byte {
	_ = "STUB: not implemented"

	// Write implements io.Writer - it appends p to ByteBuffer.B
	return nil
}

func (b *ByteBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteByte appends the byte c to the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
//
// The function always returns nil.
func (b *ByteBuffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

// WriteString appends s to ByteBuffer.B.
func (b *ByteBuffer) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Set sets ByteBuffer.B to p.
func (b *ByteBuffer) Set(p []byte) { _ = "STUB: not implemented"; return }

// SetString sets ByteBuffer.B to s.
func (b *ByteBuffer) SetString(s string) { _ = "STUB: not implemented"; return }

// String returns string representation of ByteBuffer.B.
func (b *ByteBuffer) String() string {
	_ = "STUB: not implemented"

	// Reset makes ByteBuffer.B empty.
	return ""
}

func (b *ByteBuffer) Reset() { _ = "STUB: not implemented"; return }
