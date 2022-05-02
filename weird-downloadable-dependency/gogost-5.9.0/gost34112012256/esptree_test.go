// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2021 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost34112012256

import (
	"bytes"
	"testing"
)

func TestESPTree(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0xB6, 0x18, 0x0C, 0x14, 0x5C, 0x51, 0x2D, 0xBD,
			0x69, 0xD9, 0xCE, 0xA9, 0x2C, 0xAC, 0x1B, 0x5C,
			0xE1, 0xBC, 0xFA, 0x73, 0x79, 0x2D, 0x61, 0xAF,
			0x0B, 0x44, 0x0D, 0x84, 0xB5, 0x22, 0xCC, 0x38,
		})
		is := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x2F, 0xF1, 0xC9, 0x0E, 0xDE, 0x78, 0x6E, 0x06,
			0x1E, 0x17, 0xB3, 0x74, 0xD7, 0x82, 0xAF, 0x7B,
			0xD8, 0x80, 0xBD, 0x52, 0x7C, 0x66, 0xA2, 0xBA,
			0xDC, 0x3E, 0x56, 0x9A, 0xAB, 0x27, 0x1D, 0xA4,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("2", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0xB6, 0x18, 0x0C, 0x14, 0x5C, 0x51, 0x2D, 0xBD,
			0x69, 0xD9, 0xCE, 0xA9, 0x2C, 0xAC, 0x1B, 0x5C,
			0xE1, 0xBC, 0xFA, 0x73, 0x79, 0x2D, 0x61, 0xAF,
			0x0B, 0x44, 0x0D, 0x84, 0xB5, 0x22, 0xCC, 0x38,
		})
		is := []byte{0x00, 0x00, 0x01, 0x00, 0x01}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x9A, 0xBA, 0xC6, 0x57, 0x78, 0x18, 0x0E, 0x6F,
			0x2A, 0xF6, 0x1F, 0xB8, 0xD5, 0x71, 0x62, 0x36,
			0x66, 0xC2, 0xF5, 0x13, 0x0D, 0x54, 0xE2, 0x11,
			0x6C, 0x7D, 0x53, 0x0E, 0x6E, 0x7D, 0x48, 0xBC,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("3", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0x5B, 0x50, 0xBF, 0x33, 0x78, 0x87, 0x02, 0x38,
			0xF3, 0xCA, 0x74, 0x0F, 0xD1, 0x24, 0xBA, 0x6C,
			0x22, 0x83, 0xEF, 0x58, 0x9B, 0xE6, 0xF4, 0x6A,
			0x89, 0x4A, 0xA3, 0x5D, 0x5F, 0x06, 0xB2, 0x03,
		})
		is := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x25, 0x65, 0x21, 0xE2, 0x70, 0xB7, 0x4A, 0x16,
			0x4D, 0xFC, 0x26, 0xE6, 0xBF, 0x0C, 0xCA, 0x76,
			0x5E, 0x9D, 0x41, 0x02, 0x7D, 0x4B, 0x7B, 0x19,
			0x76, 0x2B, 0x1C, 0xC9, 0x01, 0xDC, 0xDE, 0x7F,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("4", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0x5B, 0x50, 0xBF, 0x33, 0x78, 0x87, 0x02, 0x38,
			0xF3, 0xCA, 0x74, 0x0F, 0xD1, 0x24, 0xBA, 0x6C,
			0x22, 0x83, 0xEF, 0x58, 0x9B, 0xE6, 0xF4, 0x6A,
			0x89, 0x4A, 0xA3, 0x5D, 0x5F, 0x06, 0xB2, 0x03,
		})
		is := []byte{0x00, 0x00, 0x01, 0x00, 0x01}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x20, 0xE0, 0x46, 0xD4, 0x09, 0x83, 0x9B, 0x23,
			0xF0, 0x66, 0xA5, 0x0A, 0x7A, 0x06, 0x5B, 0x4A,
			0x39, 0x24, 0x4F, 0x0E, 0x29, 0xEF, 0x1E, 0x6F,
			0x2E, 0x5D, 0x2E, 0x13, 0x55, 0xF5, 0xDA, 0x08,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("5", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0x98, 0xBD, 0x34, 0xCE, 0x3B, 0xE1, 0x9A, 0x34,
			0x65, 0xE4, 0x87, 0xC0, 0x06, 0x48, 0x83, 0xF4,
			0x88, 0xCC, 0x23, 0x92, 0x63, 0xDC, 0x32, 0x04,
			0x91, 0x9B, 0x64, 0x3F, 0xE7, 0x57, 0xB2, 0xBE,
		})
		is := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x98, 0xF1, 0x03, 0x01, 0x81, 0x0A, 0x04, 0x1C,
			0xDA, 0xDD, 0xE1, 0xBD, 0x85, 0xA0, 0x8F, 0x21,
			0x8B, 0xAC, 0xB5, 0x7E, 0x00, 0x35, 0xE2, 0x22,
			0xC8, 0x31, 0xE3, 0xE4, 0xF0, 0xA2, 0x0C, 0x8F,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("6", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0x98, 0xBD, 0x34, 0xCE, 0x3B, 0xE1, 0x9A, 0x34,
			0x65, 0xE4, 0x87, 0xC0, 0x06, 0x48, 0x83, 0xF4,
			0x88, 0xCC, 0x23, 0x92, 0x63, 0xDC, 0x32, 0x04,
			0x91, 0x9B, 0x64, 0x3F, 0xE7, 0x57, 0xB2, 0xBE,
		})
		is := []byte{0x00, 0x00, 0x00, 0x00, 0x01}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x02, 0xC5, 0x41, 0x87, 0x7C, 0xC6, 0x23, 0xF3,
			0xF1, 0x35, 0x91, 0x9A, 0x75, 0x13, 0xB6, 0xF8,
			0xA8, 0xA1, 0x8C, 0xB2, 0x63, 0x99, 0x86, 0x2F,
			0x50, 0x81, 0x4F, 0x52, 0x91, 0x01, 0x67, 0x84,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("7", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0xD0, 0x65, 0xB5, 0x30, 0xFA, 0x20, 0xB8, 0x24,
			0xC7, 0x57, 0x0C, 0x1D, 0x86, 0x2A, 0xE3, 0x39,
			0x2C, 0x1C, 0x07, 0x6D, 0xFA, 0xDA, 0x69, 0x75,
			0x74, 0x4A, 0x07, 0xA8, 0x85, 0x7D, 0xBD, 0x30,
		})
		is := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0x4C, 0x61, 0x45, 0x99, 0xA0, 0xA0, 0x67, 0xF1,
			0x94, 0x87, 0x24, 0x0A, 0xE1, 0x00, 0xE1, 0xB7,
			0xEA, 0xF2, 0x3E, 0xDA, 0xF8, 0x7E, 0x38, 0x73,
			0x50, 0x86, 0x1C, 0x68, 0x3B, 0xA4, 0x04, 0x46,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
	t.Run("8", func(t *testing.T) {
		espTree := NewESPTree([]byte{
			0xD0, 0x65, 0xB5, 0x30, 0xFA, 0x20, 0xB8, 0x24,
			0xC7, 0x57, 0x0C, 0x1D, 0x86, 0x2A, 0xE3, 0x39,
			0x2C, 0x1C, 0x07, 0x6D, 0xFA, 0xDA, 0x69, 0x75,
			0x74, 0x4A, 0x07, 0xA8, 0x85, 0x7D, 0xBD, 0x30,
		})
		is := []byte{0x00, 0x00, 0x00, 0x00, 0x01}
		got := espTree.Derive(is)
		if bytes.Compare(got, []byte{
			0xB4, 0xF3, 0xF9, 0x0D, 0xC4, 0x87, 0xFA, 0xB8,
			0xC4, 0xAF, 0xD0, 0xEB, 0x45, 0x49, 0xF2, 0xF0,
			0xE4, 0x36, 0x32, 0xB6, 0x79, 0x19, 0x37, 0x2E,
			0x1E, 0x96, 0x09, 0xEA, 0xF0, 0xB8, 0xE2, 0x28,
		}) != 0 {
			t.FailNow()
		}
		if _, cached := espTree.DeriveCached(is); !cached {
			t.FailNow()
		}
	})
}
