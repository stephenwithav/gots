/*
MIT License

Copyright 2016 Comcast Cable Communications Management, LLC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package psi

func PointerField(psi []byte) uint8 {
	return psi[0]
}

// TableID returns the psi table header table id
func TableID(psi []byte) uint8 {
	return tableID(psi[1+PointerField(psi):])
}

// SectionSyntaxIndicator returns true if the psi contains section syntax
func SectionSyntaxIndicator(psi []byte) bool {
	return sectionSyntaxIndicator(psi[1+PointerField(psi):])
}

// PrivateIndicator returns true if the psi contains private data
func PrivateIndicator(psi []byte) bool {
	return psi[2+PointerField(psi)]&0x40 != 0
}

// SectionLength returns the psi section length
func SectionLength(psi []byte) uint16 {
	return sectionLength(psi[1+PointerField(psi):])
}

// tableID returns the table id from the header of a section
func tableID(psi []byte) uint8 {
	return uint8(psi[0])
}

func sectionSyntaxIndicator(psi []byte) bool {
	return psi[1]&0x80 != 0
}

// sectionLength returns the length of a single psi section
func sectionLength(psi []byte) uint16 {
	return uint16(psi[1]&3)<<8 | uint16(psi[2])
}
