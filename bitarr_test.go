package bitarray

import (
	"testing"
	"fmt"
)
var _ = fmt.Printf // For debugging; delete when done.


func TestBitArray(t *testing.T) {
	mlen := uint32(1000000000) //save 1 billion value. will alloc about 238.42mb 
	//Every value will occupy 2-bit. 
	//For 2-bit width value have 4 choices.00 01 10 11
	bm := NewBitArray(mlen, 2)  
	bm.SetB(4, 0x03) // 0000 0011
	bm.SetB(5, 0x01)
	bm.SetB(7, 0x01)
	bm.SetB(8, 0x01)
	if bm.GetB(5)&bm.GetB(7)&bm.GetB(8) != 0x01 {
		t.Errorf("bitmap get b 5 error")
	}
	
	if bm.GetB(4) != 0x03 {
		t.Errorf("bitmap get b 4 error",bm.GetB(4))
	}
	
	if bm.GetB(6) != 0x00 {
		t.Errorf("bitmap get b 6 error")
	}
	bm.SetB(mlen, 0x01)
	if bm.GetB(mlen) != 0x01 {
		t.Errorf("bitmap get b %d error ", mlen)
	}
	al := bm.GetAllocLen()
	fmt.Printf("bm alloc %d bytes. %.2f kb. %.2f mb\n", al, float32(al)/1024, float32(al)*1.0/1024/1024)
	fmt.Printf("Or just int array will occupied: %d bytes. %.2fkb. %.2fmb\n", mlen*4,float32(mlen)*4/1024,float32(mlen)*4/1024/1024)
}