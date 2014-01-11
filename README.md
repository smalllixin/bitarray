#bitarray
========

**Golang** bitarray implementation with variable bit-width. Like java bitset but a bit more powerful.

## Introduction

If there are 100,000,000 data. Just a few choice for these data(eg. true or false). Use bitarray to save the state is brilliant. Save your memory.


## Usage

Here show a case:
	Save 1 billion value in bitarray and every value take 2 bits. It will alloced about 238.42mb memory.  
	Every value will occupy 2-bit.   
For 2-bit width value have 4 choices.00 01 10 11
	
	mlen := uint32(1000000000)
	bm := NewBitArray(mlen, 2)  
	bm.SetB(4, 0x03) // 0000 0011
	x := bm.GetB(4) // x will be 0x03
	
View the test for more. 