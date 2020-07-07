// +build !plan9

package p9

const (
	// TypeDir represents a directory type.
	TypeDir QIDType = 0x80

	// TypeAppendOnly represents an append only file.
	TypeAppendOnly QIDType = 0x40

	// TypeExclusive represents an exclusive-use file.
	TypeExclusive QIDType = 0x20

	// TypeMount represents a mounted channel.
	TypeMount QIDType = 0x10

	// TypeAuth represents an authentication file.
	TypeAuth QIDType = 0x08

	// TypeTemporary represents a temporary file.
	TypeTemporary QIDType = 0x04

	// TypeSymlink represents a symlink.
	TypeSymlink QIDType = 0x02

	// TypeLink represents a hard link.
	TypeLink QIDType = 0x01

	// TypeRegular represents a regular file.
	TypeRegular QIDType = 0x00
)

// QID is a unique file identifier.
//
// This may be embedded in other requests and responses.
type QID struct {
	// Type is the highest order byte of the file mode.
	Type QIDType

	// Version is an arbitrary server version number.
	Version uint32

	// Path is a unique server identifier for this path (e.g. inode).
	Path uint64
}
