package goBit

import (
	"fmt"
	"strconv"
	"strings"
)

// various consts
const (
	nPc      = 12
	nPt      = 6
	WHITE    = color(0)
	BLACK    = color(1)
	startpos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - "
	row1     = bitBoard(0x00000000000000FF)
	row2     = bitBoard(0x000000000000FF00)
	row3     = bitBoard(0x0000000000FF0000)
	row4     = bitBoard(0x00000000FF000000)
	row5     = bitBoard(0x000000FF00000000)
	row6     = bitBoard(0x0000FF0000000000)
	row7     = bitBoard(0x00FF000000000000)
	row8     = bitBoard(0xFF00000000000000)
	fileA    = bitBoard(0x0101010101010101)
	fileB    = bitBoard(0x0202020202020202)
	fileG    = bitBoard(0x4040404040404040)
	fileH    = bitBoard(0x8080808080808080)
)

var atksKnights [64]bitBoard
var atksKings [64]bitBoard

// initialize all possible knight attacks
func initAtksKnights() {
	for fr := A1; fr <= H8; fr++ {
		toBB := bitBoard(0)
		rk := fr / 8
		fl := fr % 8
		// NNE  2,1
		if rk+2 < 8 && fl+1 < 8 {
			to := (rk+2)*8 + fl + 1
			toBB.set(to)
		}

		// ENE  1,2
		if rk+1 < 8 && fl+2 < 8 {
			to := (rk+1)*8 + fl + 2
			toBB.set(to)
		}

		// ESE  -1,2
		if rk-1 >= 0 && fl+2 < 8 {
			to := (rk-1)*8 + fl + 2
			toBB.set(to)
		}

		// SSE  -2,+1
		if rk-2 >= 0 && fl+1 < 8 {
			to := (rk-2)*8 + fl + 1
			toBB.set(to)
		}

		// NNW  2,-1
		if rk+2 < 8 && fl-1 >= 0 {
			to := (rk+2)*8 + fl - 1
			toBB.set(to)
		}

		// WNW  1,-2
		if rk+1 < 8 && fl-2 >= 0 {
			to := (rk+1)*8 + fl - 2
			toBB.set(to)
		}

		// WSW  -1,-2
		if rk-1 >= 0 && fl-2 >= 0 {
			to := (rk-1)*8 + fl - 2
			toBB.set(to)
		}

		// SSW  -2,-1
		if rk-2 >= 0 && fl-1 >= 0 {
			to := (rk-2)*8 + fl - 1
			toBB.set(to)
		}
		atksKnights[fr] = toBB
	}
}

// initialize all possible King attacks
func initAtksKings() {

	for fr := A1; fr <= H8; fr++ {
		toBB := bitBoard(0)
		rk := fr / 8
		fl := fr % 8
		//N 1,0
		if rk+1 < 8 {
			to := (rk+1)*8 + fl
			toBB.set(to)
		}

		//NE 1,1
		if rk+1 < 8 && fl+1 < 8 {
			to := (rk+1)*8 + fl + 1
			toBB.set(to)
		}

		//E   0,1
		if fl+1 < 8 {
			to := (rk)*8 + fl + 1
			toBB.set(to)
		}

		//SE -1,1
		if rk-1 >= 0 && fl+1 < 8 {
			to := (rk-1)*8 + fl + 1
			toBB.set(to)
		}

		//S  -1,0
		if rk-1 >= 0 {
			to := (rk-1)*8 + fl
			toBB.set(to)
		}

		//SW -1,-1
		if rk-1 >= 0 && fl-1 >= 0 {
			to := (rk-1)*8 + fl - 1
			toBB.set(to)
		}

		//W   0,-1
		if fl-1 >= 0 {
			to := (rk)*8 + fl - 1
			toBB.set(to)
		}

		//NW  1,-1
		if rk+1 < 8 && fl-1 >= 0 {
			to := (rk+1)*8 + fl - 1
			toBB.set(to)
		}
		atksKings[fr] = toBB
	}
}

type BoardStruct struct {
	key     uint64
	sq      [64]int
	wbBB    [2]bitBoard
	pieceBB [nPt]bitBoard
	King    [2]int
	ep      int
	castlings
	stm    color
	count  [12]int
	rule50 int //set to 0 if a pawn or capt move otherwise increment
}
type color int

func (c color) opp() color {
	return c ^ 0x1
}
func (c color) String() string {
	if c == WHITE {
		return "W"
	}
	return "B"
}

func (board *BoardStruct) allBB() bitBoard {
	return board.wbBB[0] | board.wbBB[1]
}

var initialized = false

// clear the Board, flags, bitboards etc
func (board *BoardStruct) clear() {
	if !initialized {
		initFen2Sq()
		initMagic()
		initKeys()
		initAtksKings()
		initAtksKnights()
		initCastlings()
		initialized = true
	}

	board.stm = WHITE
	board.rule50 = 0
	board.sq = [64]int{}
	board.King = [2]int{}
	board.ep = 0
	board.castlings = 0

	for ix := A1; ix <= H8; ix++ {
		board.sq[ix] = empty
	}

	for ix := wP; ix < nPc; ix++ {
		board.count[ix] = 0
	}

	// bitBoards
	board.wbBB[WHITE], board.wbBB[BLACK] = 0, 0
	for ix := 0; ix < nPt; ix++ {
		board.pieceBB[ix] = 0
	}
	board.key = 0
}

// make a move
func (board *BoardStruct) Move(mv move) bool {
	newEp := 0
	// we assume that the move is legally correct (except for inCheck())
	fr := mv.fr()
	to := mv.to()
	pr := mv.pr()
	pc := board.sq[fr]
	switch {
	case pc == wK:
		board.castlings.off(shortW | longW)
		if abs(int(to)-int(fr)) == 2 {
			if to == G1 {
				board.setSq(wR, F1)
				board.setSq(empty, H1)
			} else {
				board.setSq(wR, D1)
				board.setSq(empty, A1)
			}
		}
	case pc == bK:
		board.castlings.off(shortB | longB)
		if abs(int(to)-int(fr)) == 2 {
			if to == G8 {
				board.setSq(bR, F8)
				board.setSq(empty, H8)
			} else {
				board.setSq(bR, D8)
				board.setSq(empty, A8)
			}
		}
	case pc == wR:
		if fr == A1 {
			board.off(longW)
		} else if fr == H1 {
			board.off(shortW)
		}
	case pc == bR:
		if fr == A8 {
			board.off(longB)
		} else if fr == H8 {
			board.off(shortB)
		}

	case pc == wP && board.sq[to] == empty: // ep move or set ep
		if to-fr == 16 {
			newEp = fr + 8
		} else if to-fr == 7 { // must be ep
			board.setSq(empty, to-8)
		} else if to-fr == 9 { // must be ep
			board.setSq(empty, to-8)
		}
	case pc == bP && board.sq[to] == empty: //  ep move or set ep
		if fr-to == 16 {
			newEp = to + 8
		} else if fr-to == 7 { // must be ep
			board.setSq(empty, to+8)
		} else if fr-to == 9 { // must be ep
			board.setSq(empty, to+8)
		}
	}
	board.ep = newEp
	board.setSq(empty, fr)

	if pr != empty {
		board.setSq(pr, to)
	} else {
		board.setSq(pc, to)
	}

	board.key = ^board.key
	board.stm = board.stm ^ 0x1
	if board.isAttacked(board.King[board.stm^0x1], board.stm) {
		board.UnMove(mv)
		return false
	}

	return true
}

func (board *BoardStruct) UnMove(mv move) {
	board.ep = mv.ep(board.stm.opp())
	board.castlings = mv.castl()
	pc := int(mv.pc())
	fr := int(mv.fr())
	to := int(mv.to())

	board.setSq(mv.cp(), to)
	board.setSq(pc, fr)

	if pc2pt(pc) == Pawn {
		if to == board.ep && board.ep != 0 { // ep move
			board.setSq(empty, to)
			switch to - fr {
			case NW, NE:
				board.setSq(bP, to-N)
			case SW, SE:
				board.setSq(wP, to-S)
			}
		}
	} else if pc2pt(pc) == King {
		sd := pcColor(pc)
		if fr-to == 2 { // long castling
			board.setSq(castl[sd].rook, int(castl[sd].rookL))
			board.setSq(empty, fr-1)
		} else if fr-to == -2 { // short castling
			board.setSq(castl[sd].rook, int(castl[sd].rookSh))
			board.setSq(empty, fr+1)
		}
	}
	board.key = ^board.key
	board.stm = board.stm ^ 0x1
}

// make Null move
func (board *BoardStruct) moveNull() move {
	mv := noMove
	mv.packMove(0, 0, empty, empty, empty, board.ep, board.castlings)

	board.ep = 0
	board.key = ^board.key
	board.stm = board.stm ^ 0x1
	return mv
}

func (board *BoardStruct) undoNull(mv move) {
	board.key = ^board.key
	board.stm = board.stm ^ 0x1

	board.ep = mv.ep(board.stm)
	// b.castlings = mv.castl()     // no need!
}

// is the move legal (except from inCheck)
func (board *BoardStruct) isLegal(mv move) bool {
	fr := mv.fr()
	pc := mv.pc()
	if board.sq[fr] != pc || pc == empty {
		return false
	}
	if board.stm != pcColor(pc) {
		return false
	}

	to := mv.to()
	cp := mv.cp()
	if !((pc == wP || pc == bP) && to == board.ep && board.ep != 0) {
		if board.sq[to] != cp {
			return false
		}
		if cp != empty && pcColor(cp) == pcColor(pc) {
			return false
		}
	}

	switch {
	case pc == wP:
		if to-fr == 8 { // wP one step
			if board.sq[to] == empty {
				return true
			}
		} else if to-fr == 16 {
			if board.sq[fr+8] == empty && board.sq[fr+16] == empty { // wP two step
				return true
			}
		} else if board.ep == mv.ep(board.stm) && board.sq[to-8] == bP { // wP ep
			return true
		} else if to-fr == 7 && cp != empty { // wP capture left
			return true
		} else if to-fr == 9 && cp != empty { // wp capture right
			return true
		}

		return false
	case pc == bP:
		if fr-to == 8 { // bP one step
			if board.sq[to] == empty {
				return true
			}
		} else if fr-to == 16 {
			if board.sq[fr-8] == empty && board.sq[fr-16] == empty { // bP two step
				return true
			}
		} else if board.ep == mv.ep(board.stm) && board.sq[to+8] == wP { // bP ep
			return true
		} else if fr-to == 7 && cp != empty { // bP capture right
			return true
		} else if fr-to == 9 && cp != empty { // bp capture left
			return true
		}

		return false
	case pc == wB, pc == bB:
		toBB := bitBoard(1) << uint(to)
		if mBishopTab[fr].atks(board.allBB())&toBB != 0 {
			return true
		}
		return false
	case pc == wR, pc == bR:
		toBB := bitBoard(1) << uint(to)
		if mRookTab[fr].atks(board.allBB())&toBB != 0 {
			return true
		}
		return false
	case pc == wQ, pc == bQ:
		toBB := bitBoard(1) << uint(to)
		if mBishopTab[fr].atks(board.allBB())&toBB != 0 {
			return true
		}
		if mRookTab[fr].atks(board.allBB())&toBB != 0 {
			return true
		}
		return false
	case pc == wK:
		if abs(int(to)-int(fr)) == 2 { //castlings
			if to == G1 {
				if board.sq[H1] != wR || board.sq[E1] != wK {
					return false
				}

				if board.sq[F1] != empty || board.sq[G1] != empty {
					return false
				}

				if !board.isShortOk(board.stm) {
					return false
				}
			} else {
				if board.sq[A1] != wR || board.sq[E1] != wK {
					return false
				}
				if to != C1 {
					return false
				}
				if board.sq[B1] != empty || board.sq[C1] != empty || board.sq[D1] != empty {
					return false
				}
				if !board.isLongOk(board.stm) {
					return false
				}
			}
		}
		return true
	case pc == bK:
		if abs(int(to)-int(fr)) == 2 { //castlings
			if to == G8 {
				if board.sq[H8] != bR || board.sq[E8] != bK {
					return false
				}
				if board.sq[F8] != empty || board.sq[G8] != empty {
					return false
				}
				if !board.isShortOk(board.stm) {
					return false
				}
			} else {
				if board.sq[A8] != bR || board.sq[E8] != bK {
					return false
				}
				if to != C8 {
					return false
				}
				if board.sq[B8] != empty || board.sq[C8] != empty || board.sq[D8] != empty {
					return false
				}
				if !board.isLongOk(board.stm) {
					return false
				}
			}
		}
		return true
	}

	return true
}

func (board *BoardStruct) setSq(pc, sq int) {
	pt := pc2pt(pc)
	sd := pcColor(pc)

	if board.sq[sq] != empty { // capture
		cp := board.sq[sq]
		board.count[cp]--
		board.wbBB[sd^0x1].clr(sq)
		board.pieceBB[pc2pt(cp)].clr(sq)
		board.key ^= pcSqKey(cp, sq)
	}
	board.sq[sq] = pc

	if pc == empty {
		board.wbBB[WHITE].clr(sq)
		board.wbBB[BLACK].clr(sq)
		for p := 0; p < nPt; p++ {
			board.pieceBB[p].clr(sq)
		}
		return
	}

	board.key ^= pcSqKey(pc, sq)

	board.count[pc]++

	if pt == King {
		board.King[sd] = sq
	}

	board.wbBB[sd].set(sq)
	board.pieceBB[pt].set(sq)
}

func (board *BoardStruct) NewGame() {
	board.ParseFEN(startpos)
}

// check if short castlings is legal
func (board *BoardStruct) isShortOk(sd color) bool {
	if !board.shortFlag(sd) {
		return false
	}

	opp := sd ^ 0x1
	if castl[sd].pawnsSh&board.pieceBB[Pawn]&board.wbBB[opp] != 0 { // stopped by pawns?
		return false
	}
	if castl[sd].pawnsSh&board.pieceBB[King]&board.wbBB[opp] != 0 { // stopped by king?
		return false
	}
	if castl[sd].knightsSh&board.pieceBB[Knight]&board.wbBB[opp] != 0 { // stopped by Knights?
		return false
	}

	// sliding to e1/e8	//NOTE: Maybe not needed during search because we know if we are in check
	sq := board.King[sd]
	if (mBishopTab[sq].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	if (mRookTab[sq].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}

	// slidings to f1/f8
	if (mBishopTab[sq+1].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	if (mRookTab[sq+1].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}

	// slidings to g1/g8		//NOTE: Maybe not needed because we always make isAttacked() after a move
	if (mBishopTab[sq+2].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	if (mRookTab[sq+2].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	return true
}

// check if long castlings is legal
func (board *BoardStruct) isLongOk(sd color) bool {
	if !board.longFlag(sd) {
		return false
	}

	opp := sd ^ 0x1
	if castl[sd].pawnsL&board.pieceBB[Pawn]&board.wbBB[opp] != 0 {
		return false
	}
	if castl[sd].pawnsL&board.pieceBB[King]&board.wbBB[opp] != 0 {
		return false
	}
	if castl[sd].knightsL&board.pieceBB[Knight]&board.wbBB[opp] != 0 {
		return false
	}

	// sliding e1/e8
	sq := board.King[sd]
	if (mBishopTab[sq].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	if (mRookTab[sq].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}

	// sliding d1/d8
	if (mBishopTab[sq-1].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	if (mRookTab[sq-1].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}

	// sliding c1/c8	//NOTE: Maybe not needed because we always make inCheck() before a move
	if (mBishopTab[sq-2].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	if (mRookTab[sq-2].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[opp]) != 0 {
		return false
	}
	return true
}

func (board *BoardStruct) genRookMoves(ml *MoveList, targetBB bitBoard) {
	sd := board.stm
	allRBB := board.pieceBB[Rook] & board.wbBB[sd]
	pc := pt2pc(Rook, color(sd))
	var mv move
	for fr := allRBB.firstOne(); fr != 64; fr = allRBB.firstOne() {
		toBB := mRookTab[fr].atks(board.allBB()) & targetBB
		for to := toBB.firstOne(); to != 64; to = toBB.firstOne() {
			mv.packMove(fr, to, pc, board.sq[to], empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}
}

func (board *BoardStruct) genBishopMoves(ml *MoveList, targetBB bitBoard) {
	sd := board.stm
	allBBB := board.pieceBB[Bishop] & board.wbBB[sd]
	pc := pt2pc(Bishop, color(sd))
	ep := board.ep
	castlings := board.castlings
	var mv move

	for fr := allBBB.firstOne(); fr != 64; fr = allBBB.firstOne() {
		toBB := mBishopTab[fr].atks(board.allBB()) & targetBB
		for to := toBB.lastOne(); to != 64; to = toBB.lastOne() {
			mv.packMove(fr, to, pc, board.sq[to], empty, ep, castlings)
			ml.Add(mv)
		}
	}
}

func (board *BoardStruct) genQueenMoves(mlq *MoveList, targetBB bitBoard) {
	sd := board.stm
	allQBB := board.pieceBB[Queen] & board.wbBB[sd]
	pc := pt2pc(Queen, color(sd))
	ep := board.ep
	castlings := board.castlings
	var mv move

	for fr := allQBB.firstOne(); fr != 64; fr = allQBB.firstOne() {
		toBB := mBishopTab[fr].atks(board.allBB()) & targetBB
		toBB |= mRookTab[fr].atks(board.allBB()) & targetBB
		for to := toBB.firstOne(); to != 64; to = toBB.firstOne() {
			mv.packMove(fr, to, pc, board.sq[to], empty, ep, castlings)
			mlq.Add(mv)
		}
	}
}

func (board *BoardStruct) genKnightMoves(ml *MoveList, targetBB bitBoard) {
	sd := board.stm
	allNBB := board.pieceBB[Knight] & board.wbBB[sd]
	pc := pt2pc(Knight, color(sd))
	ep := board.ep
	castlings := board.castlings
	var mv move
	for fr := allNBB.firstOne(); fr != 64; fr = allNBB.firstOne() {
		toBB := atksKnights[fr] & targetBB
		for to := toBB.firstOne(); to != 64; to = toBB.firstOne() {
			mv.packMove(fr, to, pc, board.sq[to], empty, ep, castlings)
			ml.Add(mv)
		}
	}
}

func (board *BoardStruct) genKingMoves(ml *MoveList, targetBB bitBoard) {
	sd := board.stm
	// 'normal' moves
	pc := pt2pc(King, color(sd))
	ep := board.ep
	castlings := board.castlings
	var mv move

	toBB := atksKings[board.King[sd]] & targetBB
	for to := toBB.firstOne(); to != 64; to = toBB.firstOne() {
		mv.packMove(board.King[sd], to, pc, board.sq[to], empty, ep, castlings)
		ml.Add(mv)
	}

	// castlings
	if board.King[sd] == castl[sd].kingPos { // NOTE: Maybe not needed. We should know that the king is there if the flags are ok
		if targetBB.test(board.King[sd] + 2) {
			// short castling
			if board.sq[castl[sd].rookSh] == castl[sd].rook && // NOTE: Maybe not needed. We should know that the rook is there if the flags are ok
				(castl[sd].betweenSh&board.allBB()) == 0 {
				if board.isShortOk(sd) {
					mv.packMove(board.King[sd], board.King[sd]+2, board.sq[board.King[sd]], empty, empty, board.ep, board.castlings)
					ml.Add(mv)
				}
			}
		}

		if targetBB.test(board.King[sd] - 2) {
			// long castling
			if board.sq[castl[sd].rookL] == castl[sd].rook && // NOTE: Maybe not needed. We should know that the rook is there if the flags are ok
				(castl[sd].betweenL&board.allBB()) == 0 {
				if board.isLongOk(sd) {
					mv.packMove(board.King[sd], board.King[sd]-2, board.sq[board.King[sd]], empty, empty, board.ep, board.castlings)
					ml.Add(mv)
				}
			}
		}
	}
}

var genPawns = [2]func(*BoardStruct, *MoveList){(*BoardStruct).genWPawnMoves, (*BoardStruct).genBPawnMoves}
var genPawnCapt = [2]func(*BoardStruct, *MoveList){(*BoardStruct).genWPawnCapt, (*BoardStruct).genBPawnCapt}
var genPawnNonCapt = [2]func(*BoardStruct, *MoveList){(*BoardStruct).genWPawnNonCapt, (*BoardStruct).genBPawnNonCapt}

func (board *BoardStruct) genPawnMoves(ml *MoveList) {
	genPawns[board.stm](board, ml)
}
func (board *BoardStruct) genPawnCapt(ml *MoveList) {
	genPawnCapt[board.stm](board, ml)
}
func (board *BoardStruct) genPawnNonCapt(ml *MoveList) {
	genPawnNonCapt[board.stm](board, ml)
}
func (board *BoardStruct) genWPawnMoves(ml *MoveList) {
	wPawns := board.pieceBB[Pawn] & board.wbBB[WHITE]

	// one step
	to1Step := (wPawns << N) & ^board.allBB()
	// two steps,
	to2Step := ((to1Step & row3) << N) & ^board.allBB()
	// captures
	toCapL := ((wPawns & ^fileA) << NW) & board.wbBB[BLACK]
	toCapR := ((wPawns & ^fileH) << NE) & board.wbBB[BLACK]

	mv := noMove

	// prom
	prom := (to1Step | toCapL | toCapR) & row8
	if prom != 0 {
		for to := prom.firstOne(); to != 64; to = prom.firstOne() {
			cp := board.sq[to]
			frTab := make([]int, 0, 3)
			if board.sq[to] == empty {
				frTab = append(frTab, to-N) // not capture
			} else {
				if toCapL.test(to) { // capture left
					frTab = append(frTab, to-NW)
				}
				if toCapR.test(to) { // capture right
					frTab = append(frTab, to-NE)
				}
			}

			for _, fr := range frTab {
				mv.packMove(fr, to, wP, cp, wQ, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, wP, cp, wR, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, wP, cp, wN, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, wP, cp, wB, board.ep, board.castlings)
				ml.Add(mv)
			}
		}
		to1Step &= ^row8
		toCapL &= ^row8
		toCapR &= ^row8
	}

	// ep move
	if board.ep != 0 {
		epBB := bitBoard(1) << uint(board.ep)
		// ep left
		epToL := ((wPawns & ^fileA) << NW) & epBB
		if epToL != 0 {
			mv.packMove(board.ep-NW, board.ep, wP, bP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
		epToR := ((wPawns & ^fileH) << NE) & epBB
		if epToR != 0 {
			mv.packMove(board.ep-NE, board.ep, wP, bP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}
	// Add one step forward
	for to := to1Step.firstOne(); to != 64; to = to1Step.firstOne() {
		mv.packMove(to-N, to, wP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}
	// Add two steps forward
	for to := to2Step.firstOne(); to != 64; to = to2Step.firstOne() {
		mv.packMove(to-2*N, to, wP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}

	// add Captures left
	for to := toCapL.firstOne(); to != 64; to = toCapL.firstOne() {
		mv.packMove(to-NW, to, wP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}

	// add Captures right
	for to := toCapR.firstOne(); to != 64; to = toCapR.firstOne() {
		mv.packMove(to-NE, to, wP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}
}

func (board *BoardStruct) genBPawnMoves(ml *MoveList) {
	bPawns := board.pieceBB[Pawn] & board.wbBB[BLACK]

	// one step
	to1Step := (bPawns >> (-S)) & ^board.allBB()
	// two steps,
	to2Step := ((to1Step & row6) >> (-S)) & ^board.allBB()
	// captures
	toCapL := ((bPawns & ^fileA) >> (-SW)) & board.wbBB[WHITE]
	toCapR := ((bPawns & ^fileH) >> (-SE)) & board.wbBB[WHITE]

	var mv move

	// prom
	prom := (to1Step | toCapL | toCapR) & row1
	if prom != 0 {
		for to := prom.firstOne(); to != 64; to = prom.firstOne() {
			cp := board.sq[to]
			frTab := make([]int, 0, 3)
			if board.sq[to] == empty {
				frTab = append(frTab, to-S) // not capture
			} else {
				if toCapL.test(to) { // capture left
					frTab = append(frTab, to-SW)
				}
				if toCapR.test(to) { // capture right
					frTab = append(frTab, to-SE)
				}
			}

			for _, fr := range frTab {
				mv.packMove(fr, to, bP, cp, bQ, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, bP, cp, bR, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, bP, cp, bN, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, bP, cp, bB, board.ep, board.castlings)
				ml.Add(mv)
			}
		}
		to1Step &= ^row1
		toCapL &= ^row1
		toCapR &= ^row1
	}
	// ep move
	if board.ep != 0 {
		epBB := bitBoard(1) << uint(board.ep)
		// ep left
		epToL := ((bPawns & ^fileA) >> (-SW)) & epBB
		if epToL != 0 {
			mv.packMove(board.ep-SW, board.ep, bP, wP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
		epToR := ((bPawns & ^fileH) >> (-SE)) & epBB
		if epToR != 0 {
			mv.packMove(board.ep-SE, board.ep, bP, wP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}
	// Add one step forward
	for to := to1Step.firstOne(); to != 64; to = to1Step.firstOne() {
		mv.packMove(to-S, to, bP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}
	// Add two steps forward
	for to := to2Step.firstOne(); to != 64; to = to2Step.firstOne() {
		mv.packMove(to-2*S, to, bP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}

	// add Captures left
	for to := toCapL.firstOne(); to != 64; to = toCapL.firstOne() {
		mv.packMove(to-SW, to, bP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}

	// add Captures right
	for to := toCapR.firstOne(); to != 64; to = toCapR.firstOne() {
		mv.packMove(to-SE, to, bP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}
}

// W pawns  captures or promotions alt 2
func (board *BoardStruct) genWPawnCapt(ml *MoveList) {
	wPawns := board.pieceBB[Pawn] & board.wbBB[WHITE]

	// captures
	toCapL := ((wPawns & ^fileA) << NW) & board.wbBB[BLACK]
	toCapR := ((wPawns & ^fileH) << NE) & board.wbBB[BLACK]
	// prom
	prom := row8 & ((toCapL | toCapR) | ((wPawns << N) & ^board.allBB()))

	var mv move
	if prom != 0 {
		for to := prom.firstOne(); to != 64; to = prom.firstOne() {
			cp := board.sq[to]
			frTab := make([]int, 0, 3)
			if board.sq[to] == empty {
				frTab = append(frTab, to-N) // not capture
			} else {
				if toCapL.test(to) { // capture left
					frTab = append(frTab, to-NW)
				}
				if toCapR.test(to) { // capture right
					frTab = append(frTab, to-NE)
				}
			}
			for _, fr := range frTab {
				mv.packMove(fr, to, wP, cp, wQ, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, wP, cp, wR, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, wP, cp, wN, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, wP, cp, wB, board.ep, board.castlings)
				ml.Add(mv)
			}
		}
		toCapL &= ^row8
		toCapR &= ^row8
	}
	// ep move
	if board.ep != 0 {
		epBB := bitBoard(1) << uint(board.ep)
		// ep left
		epToL := ((wPawns & ^fileA) << NW) & epBB
		if epToL != 0 {
			mv.packMove(board.ep-NW, board.ep, wP, bP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
		epToR := ((wPawns & ^fileH) << NE) & epBB
		if epToR != 0 {
			mv.packMove(board.ep-NE, board.ep, wP, bP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}

	// add Captures left
	for to := toCapL.firstOne(); to != 64; to = toCapL.firstOne() {
		mv.packMove(to-NW, to, wP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}

	// add Captures right
	for to := toCapR.firstOne(); to != 64; to = toCapR.firstOne() {
		mv.packMove(to-NE, to, wP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}
}

// B pawn captures or promotions alternativ 2
func (board *BoardStruct) genBPawnCapt(ml *MoveList) {
	bPawns := board.pieceBB[Pawn] & board.wbBB[BLACK]

	// captures
	toCapL := ((bPawns & ^fileA) >> (-SW)) & board.wbBB[WHITE]
	toCapR := ((bPawns & ^fileH) >> (-SE)) & board.wbBB[WHITE]

	var mv move

	// prom
	prom := row1 & ((toCapL | toCapR) | ((bPawns >> (-S)) & ^board.allBB()))
	if prom != 0 {
		for to := prom.firstOne(); to != 64; to = prom.firstOne() {
			cp := board.sq[to]
			frTab := make([]int, 0, 3)
			if board.sq[to] == empty {
				frTab = append(frTab, to-S) // not capture
			} else {
				if toCapL.test(to) { // capture left
					frTab = append(frTab, to-SW)
				}
				if toCapR.test(to) { // capture right
					frTab = append(frTab, to-SE)
				}
			}

			for _, fr := range frTab {
				mv.packMove(fr, to, bP, cp, bQ, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, bP, cp, bR, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, bP, cp, bN, board.ep, board.castlings)
				ml.Add(mv)
				mv.packMove(fr, to, bP, cp, bB, board.ep, board.castlings)
				ml.Add(mv)
			}
		}
		toCapL &= ^row1
		toCapR &= ^row1
	}
	// ep move
	if board.ep != 0 {
		epBB := bitBoard(1) << uint(board.ep)
		// ep left
		epToL := ((bPawns & ^fileA) >> (-SW)) & epBB
		if epToL != 0 {
			mv.packMove(board.ep-SW, board.ep, bP, wP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
		epToR := ((bPawns & ^fileH) >> (-SE)) & epBB
		if epToR != 0 {
			mv.packMove(board.ep-SE, board.ep, bP, wP, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}

	// add Captures left
	for to := toCapL.firstOne(); to != 64; to = toCapL.firstOne() {
		mv.packMove(to-SW, to, bP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}

	// add Captures right
	for to := toCapR.firstOne(); to != 64; to = toCapR.firstOne() {
		mv.packMove(to-SE, to, bP, board.sq[to], empty, board.ep, board.castlings)
		ml.Add(mv)
	}
}

// W pawns moves that doesn't capture aand not promotions
func (board *BoardStruct) genWPawnNonCapt(ml *MoveList) {
	var mv move
	wPawns := board.pieceBB[Pawn] & board.wbBB[WHITE]

	// one step
	to1Step := (wPawns << N) & ^board.allBB()
	// two steps,
	to2Step := ((to1Step & row3) << N) & ^board.allBB()
	to1Step &= ^row8

	// Add one step forward
	for to := to1Step.firstOne(); to != 64; to = to1Step.firstOne() {
		mv.packMove(to-N, to, wP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}
	// Add two steps forward
	for to := to2Step.firstOne(); to != 64; to = to2Step.firstOne() {
		mv.packMove(to-2*N, to, wP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}
}

//B pawns moves that doesn't capture aand not promotions
func (board *BoardStruct) genBPawnNonCapt(ml *MoveList) {
	var mv move
	bPawns := board.pieceBB[Pawn] & board.wbBB[BLACK]

	// one step
	to1Step := (bPawns >> (-S)) & ^board.allBB()
	// two steps,
	to2Step := ((to1Step & row6) >> (-S)) & ^board.allBB()
	to1Step &= ^row1

	// Add one step forward
	for to := to1Step.firstOne(); to != 64; to = to1Step.firstOne() {
		mv.packMove(to-S, to, bP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}
	// Add two steps forward
	for to := to2Step.firstOne(); to != 64; to = to2Step.firstOne() {
		mv.packMove(to-2*S, to, bP, empty, empty, board.ep, board.castlings)
		ml.Add(mv)
	}
}

// generates all pseudomoves
func (board *BoardStruct) GenAllMoves(ml *MoveList) {
	board.genPawnMoves(ml)
	board.genKnightMoves(ml, ^board.wbBB[board.stm])
	board.genBishopMoves(ml, ^board.wbBB[board.stm])
	board.genRookMoves(ml, ^board.wbBB[board.stm])
	board.genQueenMoves(ml, ^board.wbBB[board.stm])
	board.genKingMoves(ml, ^board.wbBB[board.stm])
}

func (board *BoardStruct) genAllCaptures(ml *MoveList) {
	oppBB := board.wbBB[board.stm.opp()]
	board.genPawnCapt(ml)
	board.genKnightMoves(ml, oppBB)
	board.genBishopMoves(ml, oppBB)
	board.genRookMoves(ml, oppBB)
	board.genQueenMoves(ml, oppBB)
	board.genKingMoves(ml, oppBB)
}

// Create a list of captures from pawns to Kings (including promotions) - alternative
func (board *BoardStruct) genAllCapturesy(ml *MoveList) {
	us := board.stm
	them := us.opp()
	usBB := board.wbBB[us]
	themBB := board.wbBB[them]
	allBB := board.allBB()
	var atkBB, frBB bitBoard
	var mv move

	// Pawns (including ep and promotions)
	board.genPawnCapt(ml)

	// Knights
	pc := pt2pc(Knight, us)
	frBB = board.pieceBB[Knight] & usBB
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB = atksKnights[fr] & themBB
		for to := atkBB.firstOne(); to != 64; to = atkBB.firstOne() {
			cp := board.sq[to]
			mv.packMove(fr, to, pc, cp, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}

	// Bishops
	pc = pt2pc(Bishop, us)
	frBB = board.pieceBB[Bishop] & usBB
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB = mBishopTab[fr].atks(allBB) & themBB
		for to := atkBB.firstOne(); to != 64; to = atkBB.firstOne() {
			cp := board.sq[to]
			mv.packMove(fr, to, pc, cp, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}

	// Rooks
	pc = pt2pc(Rook, us)
	frBB = board.pieceBB[Rook] & usBB
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB = mRookTab[fr].atks(allBB) & themBB
		for to := atkBB.firstOne(); to != 64; to = atkBB.firstOne() {
			cp := board.sq[to]
			mv.packMove(fr, to, pc, cp, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}

	// Queens
	pc = pt2pc(Queen, us)
	frBB = board.pieceBB[Queen] & usBB
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB = mBishopTab[fr].atks(allBB) & themBB
		atkBB |= mRookTab[fr].atks(allBB) & themBB
		for to := atkBB.firstOne(); to != 64; to = atkBB.firstOne() {
			cp := board.sq[to]
			mv.packMove(fr, to, pc, cp, empty, board.ep, board.castlings)
			ml.Add(mv)
		}
	}

	// King
	pc = pt2pc(King, us)
	fr := board.King[us]
	atkBB = atksKings[fr] & themBB
	for to := atkBB.firstOne(); to != 64; to = atkBB.firstOne() {
		cp := board.sq[to]
		mv.packMove(fr, to, pc, cp, empty, board.ep, board.castlings)
		ml.Add(mv)
	}

}
func (board *BoardStruct) genAllNonCaptures(ml *MoveList) {
	emptyBB := ^board.allBB()
	board.genPawnNonCapt(ml)
	board.genKnightMoves(ml, emptyBB)
	board.genBishopMoves(ml, emptyBB)
	board.genRookMoves(ml, emptyBB)
	board.genQueenMoves(ml, emptyBB)
	board.genKingMoves(ml, emptyBB)
}

// generates all legal moves
func (board *BoardStruct) GenAllLegals(ml *MoveList) {
	board.GenAllMoves(ml)
	board.FilterLegals(ml)
}

// generate all legal moves
func (board *BoardStruct) FilterLegals(ml *MoveList) {
	for ix := len(*ml) - 1; ix >= 0; ix-- {
		mov := (*ml)[ix]
		if board.Move(mov) {
			board.UnMove(mov)
		} else {
			ml.Remove(ix)
		}
	}
}

// is sq attacked by the sd color side
func (board *BoardStruct) isAttacked(to int, sd color) bool {
	if isPawnAtkingSq[sd](board, to) {
		return true
	}

	if atksKnights[to]&board.pieceBB[Knight]&board.wbBB[sd] != 0 {
		return true
	}
	if atksKings[to]&board.pieceBB[King]&board.wbBB[sd] != 0 {
		return true
	}
	if (mBishopTab[to].atks(board.allBB()) & (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[sd]) != 0 {
		return true
	}
	if (mRookTab[to].atks(board.allBB()) & (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[sd]) != 0 {
		return true
	}

	return false
}

// allAttacks from color to any square, empty or not
func (board *BoardStruct) attacksBB(us color) bitBoard {
	allSq := ^bitBoard(0) // all squares

	atkBB := atksKings[board.King[us]]

	atkBB |= allPawnAtksBB[us](board)

	frBB := board.pieceBB[Knight] & board.wbBB[us]
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB |= atksKnights[fr]
	}

	frBB = (board.pieceBB[Bishop] | board.pieceBB[Queen]) & board.wbBB[us]
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB |= mBishopTab[fr].atks(allSq)
	}

	frBB = (board.pieceBB[Rook] | board.pieceBB[Queen]) & board.wbBB[us]
	for fr := frBB.firstOne(); fr != 64; fr = frBB.firstOne() {
		atkBB |= mRookTab[fr].atks(allSq)
	}

	return atkBB
}

var isPawnAtkingSq = [2]func(*BoardStruct, int) bool{(*BoardStruct).iswPawnAtkingSq, (*BoardStruct).isbPawnAtkingSq}
var allPawnAtksBB = [2]func(*BoardStruct) bitBoard{(*BoardStruct).wPawnAtksBB, (*BoardStruct).bPawnAtksBB}
var pawnAtksFr = [2]func(*BoardStruct, int) bitBoard{(*BoardStruct).wPawnAtksFr, (*BoardStruct).bPawnAtksFr}
var pawnAtkers = [2]func(*BoardStruct) bitBoard{(*BoardStruct).wPawnAtkers, (*BoardStruct).bPawnAtkers}

// Returns true or false if to-sq is attacked by white pawn
func (board *BoardStruct) iswPawnAtkingSq(to int) bool {
	sqBB := bitBoard(1) << uint(to)

	wPawns := board.pieceBB[Pawn] & board.wbBB[WHITE]

	// Attacks left and right
	toCap := ((wPawns & ^fileA) << NW) & board.wbBB[BLACK]
	toCap |= ((wPawns & ^fileH) << NE) & board.wbBB[BLACK]
	return (toCap & sqBB) != 0
}

// Returns true or false if to-sq is attacked by white pawn
func (board *BoardStruct) isbPawnAtkingSq(to int) bool {
	sqBB := bitBoard(1) << uint(to)

	bPawns := board.pieceBB[Pawn] & board.wbBB[BLACK]

	// Attacks left and right
	toCap := ((bPawns & ^fileA) >> (-SW)) & board.wbBB[WHITE]
	toCap |= ((bPawns & ^fileH) >> (-SE)) & board.wbBB[WHITE]

	return (toCap & sqBB) != 0
}

// returns all w pawns that attacka black pieces
func (board *BoardStruct) wPawnAtkers() bitBoard {

	BB := board.wbBB[BLACK] // all their pieces
	// pretend that all their pieces are pawns
	// Get pawn Attacks left and right from their pieces into our pawns that now are all our pwan attackers
	ourPawnAttackers := ((BB & ^fileA) >> (-SW)) & board.wbBB[WHITE] & board.pieceBB[Pawn]
	ourPawnAttackers |= ((BB & ^fileH) >> (-SE)) & board.wbBB[WHITE] & board.pieceBB[Pawn]

	return ourPawnAttackers
}

// returns all bl pawns that attacks white pieces
func (board *BoardStruct) bPawnAtkers() bitBoard {

	BB := board.wbBB[WHITE] // all their pieces
	// pretend that all their pieces are pawns
	// Get pawn Attacks left and right from their pieces into our pawns that now are all our pwan attackers
	ourPawnAttackers := ((BB & ^fileA) << NW) & board.wbBB[BLACK] & board.pieceBB[Pawn]
	ourPawnAttackers |= ((BB & ^fileH) << NE) & board.wbBB[BLACK] & board.pieceBB[Pawn]

	return ourPawnAttackers
}

// returns captures from fr-sq
func (board *BoardStruct) wPawnAtksFr(fr int) bitBoard {
	frBB := bitBoard(1) << uint(fr)

	// Attacks left and right
	toCap := ((frBB & ^fileA) << NW) & board.wbBB[BLACK]
	toCap |= ((frBB & ^fileH) << NE) & board.wbBB[BLACK]
	return toCap
}

// returns captures from fr-sq
func (board *BoardStruct) bPawnAtksFr(fr int) bitBoard {
	frBB := bitBoard(1) << uint(fr)

	// Attacks left and right
	toCap := ((frBB & ^fileA) >> (-SW)) & board.wbBB[WHITE]
	toCap |= ((frBB & ^fileH) >> (-SE)) & board.wbBB[WHITE]

	return toCap
}

// returns bitBoard with all attacks, empty or not, from all white Pawns
func (board *BoardStruct) wPawnAtksBB() bitBoard {
	frBB := board.pieceBB[Pawn] & board.wbBB[WHITE]

	// Attacks left and right
	toCap := ((frBB & ^fileA) << NW)
	toCap |= ((frBB & ^fileH) << NE)
	return toCap
}

// returns bitBoard with all attacks, empty or not, from all white Pawns
func (board *BoardStruct) bPawnAtksBB() bitBoard {
	frBB := board.pieceBB[Pawn] & board.wbBB[BLACK]

	// Attacks left and right
	toCap := ((frBB & ^fileA) << NW)
	toCap |= ((frBB & ^fileH) << NE)
	return toCap
}

//////////////////////////////////// my own commands - NOT UCI /////////////////////////////////////

// print all legal moves
func (board *BoardStruct) printAllLegals() {
	var ml MoveList
	ml.Clear()
	board.GenAllLegals(&ml)
	fmt.Println(len(ml), "moves:", ml.String())
}

func (board *BoardStruct) Print() {
	fmt.Println()
	txtStm := "BLACK"
	if board.stm == WHITE {
		txtStm = "WHITE"
	}
	txtEp := "-"
	if board.ep != 0 {
		txtEp = sq2Fen[board.ep]
	}
	key, fullKey := board.key, board.fullKey()
	index := fullKey & uint64(trans.mask)
	fmt.Printf("%v to move; ep: %v  castling:%v fullKey=%x key=%x index=%x\n", txtStm, txtEp, board.castlings.String(), fullKey, key, index)

	fmt.Println("  +------+------+------+------+------+------+------+------+")
	for lines := 8; lines > 0; lines-- {
		fmt.Println("  |      |      |      |      |      |      |      |      |")
		fmt.Printf("%v |", lines)
		for ix := (lines - 1) * 8; ix < lines*8; ix++ {
			if board.sq[ix] == bP {
				fmt.Printf("   o  |")
			} else {
				fmt.Printf("   %v  |", pc2Fen(board.sq[ix]))
			}
		}
		fmt.Println()
		fmt.Println("  |      |      |      |      |      |      |      |      |")
		fmt.Println("  +------+------+------+------+------+------+------+------+")
	}

	fmt.Printf("       A      B      C      D      E      F      G      H\n")
}

func (board *BoardStruct) printAllBB() {
	txtStm := "BLACK"
	if board.stm == WHITE {
		txtStm = "WHITE"
	}
	txtEp := "-"
	if board.ep != 0 {
		txtEp = sq2Fen[board.ep]
	}
	fmt.Printf("%v to move; ep: %v   castling:%v\n", txtStm, txtEp, board.castlings.String())

	fmt.Println("white pieces")
	fmt.Println(board.wbBB[WHITE].Stringln())
	fmt.Println("black pieces")
	fmt.Println(board.wbBB[BLACK].Stringln())

	fmt.Println("wP")
	fmt.Println((board.pieceBB[Pawn] & board.wbBB[WHITE]).Stringln())
	fmt.Println("wN")
	fmt.Println((board.pieceBB[Knight] & board.wbBB[WHITE]).Stringln())
	fmt.Println("wB")
	fmt.Println((board.pieceBB[Bishop] & board.wbBB[WHITE]).Stringln())
	fmt.Println("wR")
	fmt.Println((board.pieceBB[Rook] & board.wbBB[WHITE]).Stringln())
	fmt.Println("wQ")
	fmt.Println((board.pieceBB[Queen] & board.wbBB[WHITE]).Stringln())
	fmt.Println("wK")
	fmt.Println((board.pieceBB[King] & board.wbBB[WHITE]).Stringln())

	fmt.Println("bP")
	fmt.Println((board.pieceBB[Pawn] & board.wbBB[BLACK]).Stringln())
	fmt.Println("bN")
	fmt.Println((board.pieceBB[Knight] & board.wbBB[BLACK]).Stringln())
	fmt.Println("bB")
	fmt.Println((board.pieceBB[Bishop] & board.wbBB[BLACK]).Stringln())
	fmt.Println("bR")
	fmt.Println((board.pieceBB[Rook] & board.wbBB[BLACK]).Stringln())
	fmt.Println("bQ")
	fmt.Println((board.pieceBB[Queen] & board.wbBB[BLACK]).Stringln())
	fmt.Println("bK")
	fmt.Println((board.pieceBB[King] & board.wbBB[BLACK]).Stringln())
}

// parse a FEN string and setup that position
func (board *BoardStruct) ParseFEN(FEN string) {
	board.clear()
	fenIx := 0
	sq := 0
	for row := 7; row >= 0; row-- {
		for sq = row * 8; sq < row*8+8; {

			char := string(FEN[fenIx])
			fenIx++
			if char == "/" {
				continue
			}

			if i, err := strconv.Atoi(char); err == nil { //numeriskt
				for j := 0; j < i; j++ {
					board.setSq(empty, sq)
					sq++
				}
				continue
			}

			if strings.IndexAny(pcFen, char) == -1 {
				fmt.Println("info string invalid piece ", char, " try next one")
				continue
			}

			board.setSq(fen2pc(char), sq)

			sq++
		}
	}

	remaining := strings.Split(strings.TrimSpace(FEN[fenIx:]), " ")

	// stm
	if len(remaining) > 0 {
		if remaining[0] == "w" {
			board.stm = WHITE
		} else if remaining[0] == "b" {
			board.stm = BLACK
		} else {
			r := fmt.Sprintf("%v; sq=%v;  fenIx=%v", strings.Join(remaining, " "), sq, fenIx)

			fmt.Println("info string remaining=", r, ";")
			fmt.Println("info string ", remaining[0], " invalid stm color")
			board.stm = WHITE
		}
	}
	if board.stm == BLACK {
		board.key = ^board.key
	}

	// castling
	board.castlings = 0
	if len(remaining) > 1 {
		board.castlings = parseCastlings(remaining[1])
	}

	// ep square
	board.ep = 0
	if len(remaining) > 2 {
		if remaining[2] != "-" {
			board.ep = fen2Sq[remaining[2]]
		}
	}

	// 50-move
	board.rule50 = 0
	if len(remaining) > 3 {
		board.rule50 = parse50(remaining[3])
	}
}

// parse 50 move rue in fenstring
func parse50(fen50 string) int {
	r50, err := strconv.Atoi(fen50)
	if err != nil || r50 < 0 {
		fmt.Println("info string 50 move rule in fenstring ", fen50, " is not a valid number >= 0 ")
		return 0
	}
	return r50
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// fen2pc convert pieceString to pc int
func fen2pc(c string) int {
	for p, x := range pcFen {
		if string(x) == c {
			return p
		}
	}
	return empty
}

// pc2Fen convert pc to fenString
func pc2Fen(pc int) string {
	if pc == empty {
		return " "
	}
	return string(pcFen[pc])
}

// pc2pt returns the pt from pc
func pc2pt(pc int) int {
	return pc >> 1
}

// pcColor returns the color of a pc form
func pcColor(pc int) color {
	return color(pc & 0x1)
}

// pt2pc returns pc from pt and sd
func pt2pc(pt int, sd color) int {
	return (pt << 1) | int(sd)
}

// map fen-sq to int
var fen2Sq = make(map[string]int)

// map int-sq to fen
var sq2Fen = make(map[int]string)

// init the square map from string to int and int to string
func initFen2Sq() {
	fen2Sq["a1"] = A1
	fen2Sq["a2"] = A2
	fen2Sq["a3"] = A3
	fen2Sq["a4"] = A4
	fen2Sq["a5"] = A5
	fen2Sq["a6"] = A6
	fen2Sq["a7"] = A7
	fen2Sq["a8"] = A8

	fen2Sq["b1"] = B1
	fen2Sq["b2"] = B2
	fen2Sq["b3"] = B3
	fen2Sq["b4"] = B4
	fen2Sq["b5"] = B5
	fen2Sq["b6"] = B6
	fen2Sq["b7"] = B7
	fen2Sq["b8"] = B8

	fen2Sq["c1"] = C1
	fen2Sq["c2"] = C2
	fen2Sq["c3"] = C3
	fen2Sq["c4"] = C4
	fen2Sq["c5"] = C5
	fen2Sq["c6"] = C6
	fen2Sq["c7"] = C7
	fen2Sq["c8"] = C8

	fen2Sq["d1"] = D1
	fen2Sq["d2"] = D2
	fen2Sq["d3"] = D3
	fen2Sq["d4"] = D4
	fen2Sq["d5"] = D5
	fen2Sq["d6"] = D6
	fen2Sq["d7"] = D7
	fen2Sq["d8"] = D8

	fen2Sq["e1"] = E1
	fen2Sq["e2"] = E2
	fen2Sq["e3"] = E3
	fen2Sq["e4"] = E4
	fen2Sq["e5"] = E5
	fen2Sq["e6"] = E6
	fen2Sq["e7"] = E7
	fen2Sq["e8"] = E8

	fen2Sq["f1"] = F1
	fen2Sq["f2"] = F2
	fen2Sq["f3"] = F3
	fen2Sq["f4"] = F4
	fen2Sq["f5"] = F5
	fen2Sq["f6"] = F6
	fen2Sq["f7"] = F7
	fen2Sq["f8"] = F8

	fen2Sq["g1"] = G1
	fen2Sq["g2"] = G2
	fen2Sq["g3"] = G3
	fen2Sq["g4"] = G4
	fen2Sq["g5"] = G5
	fen2Sq["g6"] = G6
	fen2Sq["g7"] = G7
	fen2Sq["g8"] = G8

	fen2Sq["h1"] = H1
	fen2Sq["h2"] = H2
	fen2Sq["h3"] = H3
	fen2Sq["h4"] = H4
	fen2Sq["h5"] = H5
	fen2Sq["h6"] = H6
	fen2Sq["h7"] = H7
	fen2Sq["h8"] = H8

	// -------------- sq2Fen
	sq2Fen[A1] = "a1"
	sq2Fen[A2] = "a2"
	sq2Fen[A3] = "a3"
	sq2Fen[A4] = "a4"
	sq2Fen[A5] = "a5"
	sq2Fen[A6] = "a6"
	sq2Fen[A7] = "a7"
	sq2Fen[A8] = "a8"

	sq2Fen[B1] = "b1"
	sq2Fen[B2] = "b2"
	sq2Fen[B3] = "b3"
	sq2Fen[B4] = "b4"
	sq2Fen[B5] = "b5"
	sq2Fen[B6] = "b6"
	sq2Fen[B7] = "b7"
	sq2Fen[B8] = "b8"

	sq2Fen[C1] = "c1"
	sq2Fen[C2] = "c2"
	sq2Fen[C3] = "c3"
	sq2Fen[C4] = "c4"
	sq2Fen[C5] = "c5"
	sq2Fen[C6] = "c6"
	sq2Fen[C7] = "c7"
	sq2Fen[C8] = "c8"

	sq2Fen[D1] = "d1"
	sq2Fen[D2] = "d2"
	sq2Fen[D3] = "d3"
	sq2Fen[D4] = "d4"
	sq2Fen[D5] = "d5"
	sq2Fen[D6] = "d6"
	sq2Fen[D7] = "d7"
	sq2Fen[D8] = "d8"

	sq2Fen[E1] = "e1"
	sq2Fen[E2] = "e2"
	sq2Fen[E3] = "e3"
	sq2Fen[E4] = "e4"
	sq2Fen[E5] = "e5"
	sq2Fen[E6] = "e6"
	sq2Fen[E7] = "e7"
	sq2Fen[E8] = "e8"

	sq2Fen[F1] = "f1"
	sq2Fen[F2] = "f2"
	sq2Fen[F3] = "f3"
	sq2Fen[F4] = "f4"
	sq2Fen[F5] = "f5"
	sq2Fen[F6] = "f6"
	sq2Fen[F7] = "f7"
	sq2Fen[F8] = "f8"

	sq2Fen[G1] = "g1"
	sq2Fen[G2] = "g2"
	sq2Fen[G3] = "g3"
	sq2Fen[G4] = "g4"
	sq2Fen[G5] = "g5"
	sq2Fen[G6] = "g6"
	sq2Fen[G7] = "g7"
	sq2Fen[G8] = "g8"

	sq2Fen[H1] = "h1"
	sq2Fen[H2] = "h2"
	sq2Fen[H3] = "h3"
	sq2Fen[H4] = "h4"
	sq2Fen[H5] = "h5"
	sq2Fen[H6] = "h6"
	sq2Fen[H7] = "h7"
	sq2Fen[H8] = "h8"
}

// 6 piece types - no color (P)
const (
	Pawn int = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// 12 pieces with color plus empty
const (
	wP = iota
	bP
	wN
	bN
	wB
	bB
	wR
	bR
	wQ
	bQ
	wK
	bK
	empty = 15
)

// piece char definitions
const (
	pcFen = "PpNnBbRrQqKk     "
	ptFen = "PNBRQK?"
)

// square names
const (
	A1 = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1

	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2

	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3

	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4

	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5

	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6

	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7

	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

//////////////////////////////// TODO: remove this after benchmarking ////////////////////////////////////////
func (board *BoardStruct) genSimpleRookMoves(ml *MoveList, sd color) {
	allRBB := board.pieceBB[Rook] & board.wbBB[sd]
	pc := pt2pc(Rook, color(sd))
	ep := board.ep
	castlings := board.castlings
	var mv move
	for fr := allRBB.firstOne(); fr != 64; fr = allRBB.firstOne() {
		rk := fr / 8
		fl := fr % 8
		//N
		for r := rk + 1; r < 8; r++ {
			to := r*8 + fl
			cp := board.sq[to]
			if cp != empty && pcColor(int(cp)) == sd {
				break
			}
			mv.packMove(fr, to, pc, cp, empty, ep, castlings)
			ml.Add(mv)
			if cp != empty {
				break
			}
		}
		//S
		for r := rk - 1; r >= 0; r-- {
			to := r*8 + fl
			cp := board.sq[to]
			if cp != empty && pcColor(int(cp)) == sd {
				break
			}
			mv.packMove(fr, to, pc, cp, empty, ep, castlings)
			ml.Add(mv)
			if cp != empty {
				break
			}
		}
		//E
		for f := fl + 1; f < 8; f++ {
			to := rk*8 + f
			cp := board.sq[to]
			if cp != empty && pcColor(int(cp)) == sd {
				break
			}
			mv.packMove(fr, to, pc, cp, empty, ep, castlings)
			ml.Add(mv)
			if cp != empty {
				break
			}
		}
		//W
		for f := fl - 1; f >= 0; f-- {
			to := rk*8 + f
			cp := board.sq[to]
			if cp != empty && pcColor(int(cp)) == sd {
				break
			}
			mv.packMove(fr, to, pc, cp, empty, ep, castlings)
			ml.Add(mv)
			if cp != empty {
				break
			}
		}
	}
}
