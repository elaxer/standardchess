# standardchess - golang chess engine

[![godoc](https://godoc.org/github.com/elaxer/standardchess?status.svg)](https://godoc.org/github.com/elaxer/standardchess)
[![go report card](https://goreportcard.com/badge/elaxer/standardchess)](https://goreportcard.com/report/elaxer/standardchess)

The `standardchess` library offers an API for executing chess moves, working with boards and pieces, and encoding boards into FEN/PGN formats and vice versa. The library is based on the [github.com/elaxer/chess](https://github.com/elaxer/chess) library.
For a better understanding of how the engine works, I recommend reading the documentation for that library.

## Documentation

Note: this library is based on the [github.com/elaxer/chess](https://github.com/elaxer/chess) library, so you should import the `github.com/elaxer/chess` package for the `chess` namespace.

### Board creation

You can create a board in several ways:
```go
// Create a new board with the position "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1":
board := standardchess.NewBoard()
```

```go
// Create a new empty board 8x8 with White's turn:
board, err := standardchess.NewBoardEmpty(chess.SideWhite, nil, standardchess.EdgePosition)
```

```go
// ... or create a board with the position of the pieces specified by you:
piecesPosition := map[chess.Position]chess.Piece{
    chess.PositionFromString("d3"): standardchess.NewPawn(chess.SideWhite),
    chess.PositionFromString("c1"): standardchess.NewKing(chess.SideWhite),

    chess.PositionFromString("a2"): standardchess.NewRook(chess.SideBlack),
    chess.PositionFromString("h8"): standardchess.NewKing(chess.SideBlack),
}
board := standardchess.NewBoardEmpty(chess.SideBlack, piecesPosition, standardchess.EdgePosition)
```

```go
// Create a new board from moves list:
board, err := standardchess.NewBoardFromMoves([]chess.Move{
    chess.StringMove("e4"),
    chess.StringMove("e5"),
    // ...
})
```

### Making and undoing of moves

After you create a filled one, you can perform and cancel moves:
```go
moveResult, err := board.MakeMove(chess.StringMove("Nc3"))
if err != nil {
    // ...
}

// Undo the last "Nc3" move:
poppedMoveResult, err := board.UndoLastMove()
if err != nil {
    // ...
}
```

Also make/undo castling moves and pawn promotions:
```go
moveResult, err := board.MakeMove(chess.StringMove("0-0"))
if err != nil {
    // ...
}

moveResult, err = board.MakeMove(chess.StringMove("c8=Q"))
if err != nil {
    // ...
}
```

### Checking the board state

Each move can change the state of the board. You can get state of the board using method `State`:

```go
state := board.State()
```

Now let's check the value:

```go
switch state {
case standardchess.StateClear:
    fmt.Println("Nothing special on the board")
case standardchess.StateCheck:
    fmt.Println("Check on the board")
case standardchess.StateCheckmate:
    fmt.Println("Checkmate on the board, no new moves can be made")
case standardchess.StateStalemate:
    fmt.Println("Stalemate on the board, no new moves can be made")
case standardchess.StateFiftyMoves:
    fmt.Println("Case of the fifty moves rule")
case standardchess.StateThreefoldRepetition:
    fmt.Println("Case of the threefold repetition rule")
case standardchess.StateInsufficientMaterial:
    fmt.Println("Insufficient material on the board, no new moves can be made")
}
```

... or check the type of the state:
```go
switch {
case state.IsTerminal():
    fmt.Print("Checkmate, stalemate, fifty moves rule draw, ")
    fmt.Println("threefold repetition draw or insufficient material")
case state.IsThreat():
    fmt.Println("Check")
case state.IsClear():
    fmt.Println("Nothing special on the board")
}
```

### Pieces

You can create any chess piece of any color:

```go
wKing := standardchess.NewKing(chess.SideWhite)
bKing := standardchess.NewKing(chess.SideBlack)

pawn := standardchess.NewPawn(chess.SideWhite)
rook := standardchess.NewRook(chess.SideBlack)
knight := standardchess.NewKnight(chess.SideWhite)
bishop := standardchess.NewBishop(chess.SideBlack)
queen := standardchess.NewQueen(chess.SideWhite)
```

... or create via the universal function:
```go
pawn, err := standardchess.NewPiece(standardchess.NotationPawn, chess.SideWhite)
if err != nil {
    // ...
}

rook, err := standardchess.NewPiece(standardchess.NotationRook, chess.SideBlack)
if err != nil {
    // ...
}

// etc.
```

### Working with pieces and board squares. More basic work with the board

Lower-level stuff such as placing/moving/removing pieces from the board,
iterations over the board, and so on are described in the documentation
[github.com/elaxer/chess](https://github.com/elaxer/chess)

### Board player

Use `*BoardPlayer` to rewind the board position to see past board positions:

```go
player := standardchess.NewBoardPlayer(board)

// Use this methods for navigation:
player.Reset() // Sets the cursor at the init board position
ok := player.Prev() // Moves the cursor back one position
ok = player.GoTo(4) // Sets the cursor to a specific position
ok = player.Next() // Moves the cursor forward one position
player.End() // Sets the cursor at the last board position

// Get a snapshot of the board after moving the cursor:
boardSnapshot := player.Board() 
```

### Board encoding/decoding
#### FEN

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
    board := standardchess.NewBoard()

    f := fen.Encode(board)
    
    f.Placement() == "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
    f.Turn() == chess.ColorWhite
    short, long := f.Castlings(chess.ColorBlack)
    f.EnPassantSquare() == chess.NewPositionEmpty()
    f.HalfmoveClock() == 0
    f.MoveNumber() == 1

    f.String() == "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
}

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
    board, err := standardchess.NewBoardFromMoves(
        // ...
    )
    headers := []pgn.Header{
        pgn.NewHeader("Event", "F/S Return Match"),
        pgn.NewHeader("Site", "Belgrade, Serbia JUG"),
        pgn.NewHeader("Date", time.Now().Format("2006.01.02")),
    }

    result = pgn.ResultFromBoard(board)

    pgnStr := pgn.Encode(headers, board, result)
}

```

Let's try to create a PGN from a string:
```go
const pgnStr = `
[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]

1.e4 e5 2.Nf3 Nc6 *
`

p, err := pgn.FromString(pgnStr)
p.Headers()
p.Moves()
p.Result()
p.String()
```

Now let's parse several PGNs from a reader:
```go
f, err := os.Open("games.pgn")
if err != nil {
    panic(err)
}
defer f.Close()

pgns, err := pgn.Parse(f)
if err != nil {
    // ...
}
```

#### JSON

Marshal your board into json format:

```go
board := standardchess.NewBoard()
b, err := json.Marshal(board)
if err != nil {
    // ...
}
```

It encodes your board into the format:

```json
{
  "turn": false,
  "state": { "name": "string", "type": "string" },
  "castlings": { "O-O": false, "O-O-O": false },
  "captured_pieces": [
    { "color": false, "notation": "string", "is_moved": false }
  ],
  "move_history": [
    {
      "move": "string",
      "side": false,
      "captured_piece": { "side": false, "notation": "string", "is_moved": false },
      "board_new_state": { "name": "string", "type": "string" },
      "str": "string"
    },
    {
      "move": "string",
      "side": false,
      "captured_piece": null,
      "board_new_state": { "name": "string", "type": "string" },
      "str": "string"
    }
  ],
  "placement": [
    {
      "piece": { "side": false, "notation": "string", "is_moved": false },
      "position": { "file": "string", "rank": 0 },
      "legal_moves": [
        { "file": "string", "rank": 0 },
      ]
    }
  ]
}
```

## Contributing

Bug reports and contributions are welcome. Please open issues or pull requests against this repository. Keep changes small and add tests for new behavior.

## License

The GNU General Public License