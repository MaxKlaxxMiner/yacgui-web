package YacBoard

import (
	"fmt"
)

type Move struct {
	PromotionPiece Piece
	CapturePiece   Piece
	FromPos        byte
	ToPos          byte
}

func (m Move) IsValid() bool {
	return m.FromPos != m.ToPos
}

func (m Move) String() string {
	if !m.IsValid() {
		return "-"
	}

	result := fmt.Sprintf("%s-%s", PosToChars(int(m.FromPos)), PosToChars(int(m.ToPos)))

	if m.PromotionPiece != PieceNone {
		result += "->" + string(PieceToChar(m.PromotionPiece))
	}

	if m.CapturePiece != PieceNone {
		result += " (x" + string(PieceToChar(m.CapturePiece)) + ")"
	}

	return result
}

//public struct Move
//{
//
///// <summary>
///// sortiert die Züge in einem Array
///// </summary>
///// <param name="moves">Züge, welche sortiert werden sollen</param>
///// <param name="ofs">Startposition innerhalb des Arrays</param>
///// <param name="count">Anzahl der enthaltenen Züge</param>
//public static unsafe void Sort(Move[] moves, int ofs, int count)
//{
//if (count < 2) return;
//fixed (Move* ptr = &moves[ofs])
//{
//Sort((uint*)ptr, count);
//}
//}
//
///// <summary>
///// sortiert uint-Werte in einer Liste
///// </summary>
///// <param name="ptr">Pointer auf die Liste</param>
///// <param name="count">Anzahl der Elemente</param>
//static unsafe void Sort(uint* ptr, int count)
//{
//for (int start = 1; start < count; start++)
//{
//int i = start;
//uint tmp = ptr[start];
//for (; i > 0 && tmp < ptr[i - 1]; i--)
//{
//ptr[i] = ptr[i - 1];
//}
//ptr[i] = tmp;
//}
//}
//
///// <summary>
///// sortiert die Züge in einem Array rückwärts
///// </summary>
///// <param name="moves">Züge, welche sortiert werden sollen</param>
///// <param name="ofs">Startposition innerhalb des Arrays</param>
///// <param name="count">Anzahl der enthaltenen Züge</param>
//public static unsafe void SortBackward(Move[] moves, int ofs, int count)
//{
//if (count < 2) return;
//fixed (Move* ptr = &moves[ofs])
//{
//SortBackward((uint*)ptr, count);
//}
//}
//
///// <summary>
///// sortiert uint-Werte in einer Liste rückwärts
///// </summary>
///// <param name="ptr">Pointer auf die Liste</param>
///// <param name="count">Anzahl der Elemente</param>
//static unsafe void SortBackward(uint* ptr, int count)
//{
//for (int start = 1; start < count; start++)
//{
//int i = start;
//uint tmp = ptr[start];
//for (; i > 0 && tmp > ptr[i - 1]; i--)
//{
//ptr[i] = ptr[i - 1];
//}
//ptr[i] = tmp;
//}
//}
