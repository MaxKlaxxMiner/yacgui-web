package goBit

import (
	"fmt"
	"math/rand"
)

////////////////////////////////////////////////////////
//////////////////////// HASH //////////////////////////
var randPcSq [12 * 64]uint64 // keyvalues for 'pc on sq'
var randEp [8]uint64         // keyvalues for 8 ep files
var randCastl [16]uint64     // keyvalues for castling states

// setup random generator with seed
var rnd = (*rand.Rand)(rand.New(rand.NewSource(1013))) //usage: rnd.Intn(n) NOTE: n > 0

// Rand64 creates one 64 bit random number
func rand64() uint64 {
	rand := uint64(0)

	for i := 0; i < 4; i++ {
		rand = uint64(int(rand<<16) | rnd.Intn(1<<16))
	}

	return rand
}

// initKeys computes random hash keyvalues for pc/sq, ep and castlings
func initKeys() {
	for i := 0; i < 12*64; i++ {
		randPcSq[i] = rand64()
	}
	for i := 0; i < 8; i++ {
		randEp[i] = rand64()
	}
	for i := 0; i < 16; i++ {
		randCastl[i] = rand64()
	}

	// check that all keys are different
	for pc := 0; pc < 12-1; pc++ {
		for sq := 0; sq < 64; sq++ {
			key1 := pcSqKey(pc, sq)
			for pc2 := pc + 1; pc2 < 12; pc2++ {
				for sq2 := 0; sq2 < 64; sq2++ {
					if key1 == pcSqKey(pc2, sq2) {
						fmt.Printf("pc=%v, sq=%v gives the same key as pc=%v, sq=%v \n", pc, sq, pc2, sq2)
					}
				}
			}
			for ep := A3; ep <= H3; ep++ {
				if key1 == epKey(ep) {
					fmt.Printf("pc=%v, sq=%v gives the same key as ep=%v \n", pc, sq, ep)
				}
			}
			for c := uint(0); c < 16; c++ {
				if key1 == castlKey(c) {
					fmt.Printf("pc=%v, sq=%v gives the same key as castl=%v \n", pc, sq, c)
				}
			}
		}
	}

	for ep := A3; ep < H3; ep++ {
		key1 := epKey(ep)
		for ep2 := ep + 1; ep2 <= H3; ep2++ {
			if key1 == epKey(ep2) {
				fmt.Printf("ep=%vgives the same key as ep=%v \n", ep, ep2)
			}
		}

		for c := uint(0); c < 16; c++ {
			if key1 == castlKey(c) {
				fmt.Printf("ep=%v is the same key as castl=%v \n", ep, c)
			}
		}
	}

	for c := uint(0); c < 16-1; c++ {
		key1 := castlKey(c)
		for c2 := c + 1; c2 < 16; c2++ {
			if key1 == castlKey(c2) {
				fmt.Printf("castl=%v is the same key as castl=%v \n", c, c2)
			}
		}

	}
}

// for color we just flip with XOR ffffffffffffffff
// hash key after change color
func flipSide(key uint64) uint64 {
	return ^key
}

// pcSqKey returns the keyvalue för piece on square
func pcSqKey(pc, sq int) uint64 {
	return randPcSq[pc*64+sq]
}

// epKey returns the keyvalue for the current ep state
func epKey(epSq int) uint64 {
	if epSq == 0 {
		return 0
	}
	return randEp[epSq%8]
}

// castlKey returns the keyvalue for the current castling state
func castlKey(castling uint) uint64 {
	return randCastl[castling]
}

func checkKey(b *BoardStruct) bool {
	key := uint64(0)
	for sq, pc := range b.sq {
		if pc == empty {
			continue
		}
		key ^= pcSqKey(pc, sq)
	}
	if b.stm == BLACK {
		key = ^key
	}

	if key != b.key {
		return false
	}

	return true
}

////////////////////////////////////////////////////////
//////////////////////// TRANS /////////////////////////
const entrySize = 128 / 8

type ttEntry struct {
	key       uint64 // the key, extra safety
	move      uint32 // the best move from the search
	score     int16  // the score from the search
	depth     int8   // the depth that the score is based on
	scoreType uint8  // the score has this score type
}

// clear one entry
func (e *ttEntry) clear() {
	e.key = 0
	e.move = uint32(noMove)
	e.score = 0
	e.depth = -1
	e.scoreType = 0
}

type transpStruct struct {
	entries uint // number of entries
	mask    uint // mask for the index
	tab     []ttEntry
}

var trans transpStruct

// allocate a new transposition table with the size from GUI
func (t *transpStruct) new(mB int) error {
	byteSize := mB << 20
	bits := sizeToBits(byteSize)

	t.entries = 1 << bits
	t.mask = t.entries - 1

	t.tab = make([]ttEntry, t.entries, t.entries)
	t.clear()
	fmt.Println(fmt.Sprintf("info string allocated %v MB to %v entries", len(t.tab)*entrySize/(1024*1024), t.entries))
	return nil
}

// returns how many bits the mask will need to cover the table size
func sizeToBits(size int) uint {
	bits := uint(0)
	for cntEntries := size / entrySize; cntEntries > 1; cntEntries /= 2 {
		bits++
	}

	return bits
}

// clear all entries, age and counters
func (t *transpStruct) clear() {
	var e ttEntry
	e.clear()

	for i := uint(0); i < t.entries; i++ {
		t.tab[i] = e
	}
}

// index uses the Key to compute an index into the table
func (t *transpStruct) index(fullKey uint64) int64 {
	return int64(fullKey)
}

func (board *BoardStruct) fullKey() uint64 {
	key := board.key ^ epKey(board.ep)
	key ^= castlKey(uint(board.castlings))
	return key
}

const (
	//no scoretype = 0
	scoreTypeLower   = 0x1                             // sc > alpha
	scoreTypeUpper   = 0x2                             // sc < beta
	scoreTypeBetween = scoreTypeLower | scoreTypeUpper // alpha < sc < beta
)
