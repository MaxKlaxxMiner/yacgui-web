package yacboard

import (
	"errors"
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/boardsize"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/piece"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard/pos"
	"strconv"
	"strings"
)

func (board *YacBoard) SetFEN(fen string) error {
	board.Clear()

	if len(fen) > boardsize.FenMaxBytes {
		return errors.New("invalid FEN: too long")
	}

	splits := strings.Split(strings.TrimSpace(fen), " ")
	if len(splits) != 6 {
		if len(splits) < 6 {
			return errors.New("invalid FEN: too few elements (" + strconv.Itoa(len(splits)) + " < 6)")
		} else {
			return errors.New("invalid FEN: too many elements (" + strconv.Itoa(len(splits)) + " > 6)")
		}
	}
	lines := strings.Split(splits[0], "/")
	if len(lines) != boardsize.Height {
		return errors.New(fmt.Sprintf("invalid FEN: ranks: %d, expected: %d", len(lines), boardsize.Height))
	}

	// --- 1 / 6 - read pieces ---

	for y := 0; y < len(lines); y++ {
		x := 0
		for _, c := range lines[y] {
			p := piece.FromChar(byte(c))
			if p == piece.Blocked {
				return errors.New(fmt.Sprintf("invalid FEN: unknown char: %v", strconv.QuoteRune(c)))
			}
			if p == piece.None {
				x += int(c - '0')
				continue
			}
			if x < boardsize.Width {
				board.SetField(pos.FromXY(x, y), p)
			}
			x++
		}

		if x != boardsize.Width {
			return errors.New(fmt.Sprintf("invalid FEN: at rank %d: files: %d, expected: %d", boardsize.Height-y, x, boardsize.Width))
		}
	}

	// --- 2 / 6 - side ---
	switch splits[1] {
	case "w":
		board.WhiteMove = true
	case "b":
		board.WhiteMove = false
	default:
		return errors.New(fmt.Sprintf("invalid FEN: unknown move color: \"%v\"", splits[1]))
	}

	// --- 3 / 6 - castling opportunities ---
	switch splits[2] {
	case "-":
	case "q":
		board.BlackCanCastleQueenside = true
	case "k":
		board.BlackCanCastleKingside = true
	case "kq":
		board.BlackCanCastleKingside = true
		board.BlackCanCastleQueenside = true
	case "Q":
		board.WhiteCanCastleQueenside = true
	case "Qq":
		board.WhiteCanCastleQueenside = true
		board.BlackCanCastleQueenside = true
	case "Qk":
		board.WhiteCanCastleQueenside = true
		board.BlackCanCastleKingside = true
	case "Qkq":
		board.WhiteCanCastleQueenside = true
		board.BlackCanCastleKingside = true
		board.BlackCanCastleQueenside = true
	case "K":
		board.WhiteCanCastleKingside = true
	case "Kq":
		board.WhiteCanCastleKingside = true
		board.BlackCanCastleQueenside = true
	case "Kk":
		board.WhiteCanCastleKingside = true
		board.BlackCanCastleKingside = true
	case "Kkq":
		board.WhiteCanCastleKingside = true
		board.BlackCanCastleKingside = true
		board.BlackCanCastleQueenside = true
	case "KQ":
		board.WhiteCanCastleKingside = true
		board.WhiteCanCastleQueenside = true
	case "KQq":
		board.WhiteCanCastleKingside = true
		board.WhiteCanCastleQueenside = true
		board.BlackCanCastleQueenside = true
	case "KQk":
		board.WhiteCanCastleKingside = true
		board.WhiteCanCastleQueenside = true
		board.BlackCanCastleKingside = true
	case "KQkq":
		board.WhiteCanCastleKingside = true
		board.WhiteCanCastleQueenside = true
		board.BlackCanCastleKingside = true
		board.BlackCanCastleQueenside = true
	default:
		return errors.New(fmt.Sprintf("invalid FEN: unknown castling: \"%v\"", splits[2]))
	}

	// --- 4 / 6 - "en passant" ---
	board.EnPassantPos = pos.FromChars(splits[3])
	if board.EnPassantPos > 0 {
		if board.WhiteMove {
			if board.EnPassantPos < pos.FromChars("a6") || board.EnPassantPos > pos.FromChars("h6") {
				board.EnPassantPos = -1
			}
			if board.EnPassantPos > 0 && board.Fields[board.EnPassantPos+boardsize.Width] != piece.BlackPawn {
				board.EnPassantPos = -1
			}
		} else {
			if board.EnPassantPos < pos.FromChars("a3") || board.EnPassantPos > pos.FromChars("h3") {
				board.EnPassantPos = -1
			}
			if board.EnPassantPos > 0 && board.Fields[board.EnPassantPos-boardsize.Width] != piece.WhitePawn {
				board.EnPassantPos = -1
			}
		}
	}

	if board.EnPassantPos == -1 && splits[3] != "-" {
		return errors.New(fmt.Sprintf("invalid FEN: invalid en passant value: \"%v\"", splits[3]))
	}

	// --- 5 / 6 - read halfmove clock ---
	tmp, err := strconv.Atoi(splits[4])
	if err != nil || tmp < 0 || tmp > 9999 {
		return errors.New(fmt.Sprintf("invalid FEN: halfmove-counter: \"%v\"", splits[4]))
	}
	board.HalfmoveClock = tmp

	// --- 6 / 6 - read move number ---
	tmp, err = strconv.Atoi(splits[5])
	if err != nil || tmp < 1 || tmp > 9999 {
		return errors.New(fmt.Sprintf("invalid FEN: movenumber: \"%v\"", splits[4]))
	}
	board.MoveNumber = tmp

	return nil
}
