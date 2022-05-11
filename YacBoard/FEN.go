package YacBoard

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const FenMaxLength = (Width+1)*Height + 20

func (board *YacBoard) GetFEN() string {
	result := make([]byte, 0, 64)

	for y := 0; y < Height; y++ {
		result = append(result, '/')
		for x := 0; x < Width; x++ {
			c := PieceToChar(board.GetFieldXY(x, y))
			if c == '.' {
				if unicode.IsDigit(rune(result[len(result)-1])) {
					result[len(result)-1]++
				} else {
					result = append(result, '1')
				}
			} else {
				result = append(result, c)
			}
		}
	}

	if board.WhiteMove {
		result = append(result, ' ', 'w', ' ')
	} else {
		result = append(result, ' ', 'b', ' ')
	}

	if board.WhiteCanCastleKingside {
		result = append(result, 'K')
	}
	if board.WhiteCanCastleQueenside {
		result = append(result, 'Q')
	}
	if board.BlackCanCastleKingside {
		result = append(result, 'k')
	}
	if board.BlackCanCastleQueenside {
		result = append(result, 'q')
	}
	if !board.WhiteCanCastleKingside && !board.WhiteCanCastleQueenside && !board.BlackCanCastleKingside && !board.BlackCanCastleQueenside {
		result = append(result, '-')
	}

	result = append(result, fmt.Sprintf(" %s %d %d", PosToChars(board.EnPassantPos), board.HalfmoveClock, board.MoveNumber)...)

	return string(result[1:])
}

func (board *YacBoard) SetFEN(fen string) error {
	board.Clear()

	if len(fen) > FenMaxLength {
		return errors.New("invalid FEN: too long")
	}

	splits := strings.Split(strings.TrimSpace(fen), " ")
	if len(splits) != 6 {
		if len(splits) < 6 {
			return errors.New("invalid FEN: too short")
		} else {
			return errors.New("invalid FEN: too long")
		}
	}
	lines := strings.Split(splits[0], "/")
	if len(lines) != Height {
		return errors.New(fmt.Sprintf("invalid FEN: ranks: %d, expected: %d", len(lines), Height))
	}

	// --- 1 / 6 - read pieces ---

	for y := 0; y < len(lines); y++ {
		x := 0
		for _, c := range lines[y] {
			piece := PieceFromChar(c)
			if piece == Blocked {
				return errors.New(fmt.Sprintf("invalid FEN: unknown char: %v", strconv.QuoteRune(c)))
			}
			if piece == PieceNone {
				x += int(c - '0')
				continue
			}
			board.SetField(x+y*Width, piece)
			x++
		}

		if x != Width {
			return errors.New(fmt.Sprintf("invalid FEN: at rank %d: files: %d, expected: %d", Height-y, x, Width))
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
	board.EnPassantPos = PosFrom2Chars(splits[3])
	if board.EnPassantPos == -1 && splits[3] != "-" {
		return errors.New(fmt.Sprintf("invalid FEN: unknown en passant value: \"%v\"", splits[3]))
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
