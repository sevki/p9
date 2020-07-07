// +build plan9
package p9

import (
	"syscall"
)

type QID = syscall.QID

const (
	// TypeDir represents a directory type.
	TypeDir QIDType = syscall.QTDIR

	// TypeAppendOnly represents an append only file.
	TypeAppendOnly QIDType = syscall.QTAPPEND

	// TypeExclusive represents an exclusive-use file.
	TypeExclusive QIDType = syscall.QTEXCL

	// TypeMount represents a mounted channel.
	TypeMount QIDType = syscall.QTMOUNT

	// TypeAuth represents an authentication file.
	TypeAuth QIDType = syscall.QTAUTH

	// TypeTemporary represents a temporary file.
	TypeTemporary QIDType = syscall.QTTMP

	// TypeSymlink represents a symlink.
	//	TypeSymlink QIDType = 0x02 these aren't things in plan9

	// TypeLink represents a hard link.
	//	TypeLink QIDType = 0x01

	// TypeRegular represents a regular file.
	TypeRegular QIDType = syscall.QTFILE
)
