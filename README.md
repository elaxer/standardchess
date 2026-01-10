# Documentation
## Board encoding/decoding
### FEN

Forsyth–Edwards Notation (FEN) is a standard notation for describing a particular board position of a chess game.
It consists of piece placement and other information about the board

Let's see how can you encode your board to FEN string:

Then encode your board to FEN string:

```go
import (
    "github.com/elaxer/chess"
    "github.com/elaxer/standardchess"
    "github.com/elaxer/standardchess/encoding/fen"
 )

func main() {
    var board chess.Board = standardchess.NewBoard()

    fenStr := fen.Encode(board)
    fenStr == "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
}

```

... or encode only piece placement:
```go
fenStr = fen.EncodePiecePlacement(board)
fenStr == "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
```

Decode a FEN string to your board:
```go
board, err := fen.Decode("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
```

### PGN

Portable Game Notation (PGN) is a standard plain text format for recording chess games (both the moves and related data).

Here is examples of encoding a board to the PGN strings:

```go
import (
    "github.com/elaxer/chess"
    "github.com/elaxer/standardchess"
    "github.com/elaxer/standardchess/encoding/pgn"
 )

func main() {
    var board chess.Board = standardchess.NewBoard()
    headers := []pgn.Header{
        pgn.NewHeader("Event", "F/S Return Match"),
        pgn.NewHeader("Site", "Belgrade, Serbia JUG"),
        pgn.NewHeader("Date", time.Now().Format("2006.01.02")),
    }

    result = "*"

    pgnStr := pgn.Encode(headers, pgn.MoveHistoryToStrings(board.MoveHistory()), result)

    // Or decode only headers:
    pgnHeadersStr := pgn.EncodeHeaders(headers)

    // ... or only moves:
    pgnMovesStr := pgn.EncodeMoves(board.MoveHistory())
}

```

Let's try to decode a PGN string into headers with a board:
```go
pgnStr := `
[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]

1.e4 e5 2.Nf3 Nc6 *
`

header, moves, err := pgn.Decode(pgnStr)
```