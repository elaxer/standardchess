# Documentation
## Board encoding/decoding
### FEN

Forsyth–Edwards Notation (FEN) is a standard notation for describing a particular board position of a chess game.
It consists of piece placement and other information about the board

Let's see how can you build your own FEN encoder:

```go
fenEnc := &fen.Encoder{
    MetricFuncs: []metric.MetricFunc{
        func (board chess.Board) metric.Metric {
            turn := "w"
            if board.Turn().IsBlack() {
                turn = "w"
            }

            return metric.New("Turn", turn)
        },
        metric.HalfmoveCounter,
    },
}
```

Then encode your board to FEN string:

```go
var board chess.Board

fenStr := fenEnc.Encode(board)
```

... or encode only piece placement:
```go
fenStr = fenEnc.EncodePiecePlacement(board)
```

Decode a FEN string to your board:
```go
// Here should be your implementation
var (
    boardFactory chess.BoardFactory
    pieceFactory chess.PieceFactory
)

// Create your FEN decoder:
fenDec := fen.NewDecoder(boardFactory, pieceFactory)
board, err := fenDec.Decode("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
```

### PGN

Portable Game Notation (PGN) is a standard plain text format for recording chess games (both the moves and related data).

Here is examples of encoding a board to the PGN strings:

```go
var board chess.Board
headers := []pgn.Header{
    pgn.NewHeader("Event", "F/S Return Match"),
    pgn.NewHeader("Site", "Belgrade, Serbia JUG"),
    pgn.NewHeader("Date", time.Now().Format("2006.01.02")),
}

pgnStr := pgn.Encode(headers, board)

// Or decode only headers:
pgnHeadersStr := pgn.EncodeHeaders(headers)

// ... or only moves:
pgnMovesStr := pgn.EncodeMoves(board.MoveHistory())
```

Let's try to decode a PGN string into headers with a board. At first create the decoder:
```go
// A regular expression for matching with the moves string representation.
// It is usually used for the standard chess game variant.
var regexpMoves = regexp.MustCompile(
	"([NBKRQ]?[a-h]?[1-8]?x?[a-h][1-8](?:=[NBRQ])?)|([0Oo]-[0Oo](-[0Oo])?)",
)

pgnDec := pgn.NewDecoder(regexpMoves)
```

Then decode your string:
```go
pgnStr := `
[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]

1.e4 e5 2.Nf3 Nc6
`

header, board, err := pgnDec.Decode(pgnStr)
```