package matrixselector

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

// Coordinates holds the x and y positions in the matrix
type Coordinates struct {
	X, Y int
}

// drawMatrix draws the matrix based on the pivot point
func drawMatrix(matrix [][]int, pivotX, pivotY int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// Adjust the position by the pivot point
			posX := j*4 - pivotX*4
			posY := i - pivotY

			// Draw the cell value
			termbox.SetCell(posX, posY, rune('0'+matrix[i][j]/10), termbox.ColorWhite, termbox.ColorDefault)
			termbox.SetCell(posX+1, posY, rune('0'+matrix[i][j]%10), termbox.ColorWhite, termbox.ColorDefault)
			termbox.SetCell(posX+2, posY, ' ', termbox.ColorWhite, termbox.ColorDefault)
		}
	}
}

// highlightCell highlights a cell based on the pivot point
func highlightCell(x, y int, val int, pivotX, pivotY int) {
	// Convert the value to a string
	valStr := fmt.Sprintf("%d", val)

	// Ensure the value fits within the cell's width
	if len(valStr) > 3 {
		valStr = valStr[:3] // Truncate if it's too long
	}

	// Adjust the position by the pivot point
	posX := x*4 - pivotX*4
	posY := y - pivotY

	// Set the background color of the cell
	for i := 0; i < 3; i++ {
		termbox.SetCell(posX+i, posY, ' ', termbox.ColorBlack, termbox.ColorRed)
	}

	// Print the value within the highlighted cell
	for i, c := range valStr {
		termbox.SetCell(posX+i, posY, c, termbox.ColorWhite, termbox.ColorRed)
	}
}

// GetMatrixCoordinates allows a user to select a cell in the matrix and returns its coordinates
func GetMatrixCoordinates(matrix [][]int) (Coordinates, error) {
	var pivotX, pivotY int
	var x, y int

	// Initialize termbox
	if err := termbox.Init(); err != nil {
		return Coordinates{}, err
	}
	defer termbox.Close()

	drawMatrix(matrix, pivotX, pivotY)
	highlightCell(x, y, matrix[y][x], pivotX, pivotY)
	termbox.Flush()

	// Event loop
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				if y > 0 {
					y--
				}
			case termbox.KeyArrowDown:
				if y < len(matrix)-1 {
					y++
				}
			case termbox.KeyArrowLeft:
				if x > 0 {
					x--
				}
			case termbox.KeyArrowRight:
				if x < len(matrix[0])-1 {
					x++
				}
			case termbox.KeyEnter:
				return Coordinates{X: x, Y: y}, nil
			case termbox.KeyEsc:
				return Coordinates{}, nil
			default:
				// Handle `W`, `A`, `S`, `D` keys
				switch ev.Ch {
				case 'a':
					pivotX++
				case 'd':
					pivotX--
				case 'w':
					pivotY++
				case 's':
					pivotY--
				}
			}
			drawMatrix(matrix, pivotX, pivotY)
			highlightCell(x, y, matrix[y][x], pivotX, pivotY)
			termbox.Flush()
		case termbox.EventError:
			log.Fatal(ev.Err)
		}
	}
}
