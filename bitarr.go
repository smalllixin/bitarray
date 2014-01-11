package bitarray
import (
	"math"
	"fmt"
)
var _ = math.Ceil
var _ = fmt.Printf

/**
Bitmap algorithm. (BitSet With mulit-choiced bit width.)

*/

type BitArray struct {
	b []byte //we use bit
	valueBitWidth byte // how many bit present one value. only value 1,2,4,8 is supported
	countPerByte byte 
	bitmapLen uint32
}

/*
bitmapLen: how many bit we should save
*/
func NewBitArray(bitmapLen uint32, valueBitWidth byte) *BitArray {
	return new(BitArray).Init(bitmapLen, valueBitWidth)
}

func (s *BitArray) Init(bitmapLen uint32, valueBitWidth byte) *BitArray {
	validBitLen := false
	for i := uint(0); i < 4; i ++ {
		if valueBitWidth == 0x08>>i {
			validBitLen = true
			s.countPerByte = 0x08/valueBitWidth
			break
		}
	}
	if !validBitLen {
		panic("BitArray validBitLen only 1,2,4,8 is supported")
	}
	s.b = make([]byte, bitmapLen/uint32(s.countPerByte)+1)
	s.bitmapLen = bitmapLen
	s.valueBitWidth = valueBitWidth
	return s
}

func (s *BitArray) GetAllocLen() int {
	return len(s.b)
}

func (s *BitArray) SetB(pos uint32, val byte) {
	whichByte := pos/uint32(s.countPerByte)
	whichPos := pos%uint32(s.countPerByte)
	n := byte(whichPos)
	w := s.valueBitWidth
	oo := (byte(0xFF<<(8-w))>>(n*w))^0xFF
	zr := s.b[whichByte]&oo //something like [rr00 rrrr]
	sr := byte(val<<(8-w))>>(n*w) // [00ss 0000]
	s.b[whichByte] = zr | sr	
}

func (s *BitArray) GetBytes() []byte {
	return s.b
}

func (s *BitArray) GetB(pos uint32) byte {
	whichByte := pos/uint32(s.countPerByte)
	whichPos := pos%uint32(s.countPerByte)
	n := byte(whichPos)
	w := s.valueBitWidth
	
	oo := (byte(0xFF<<(8-w))>>(n*w)) // 0011 0000
	oorr := s.b[whichByte]&oo //00rr 0000
	return oorr>>(8-(n+1)*w)	
}