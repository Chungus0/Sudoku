package main
import (
	"os"
	"github.com/01-edu/z01"
)
var arr = make([][]int, 9)
var count = 0
func converter(arg []string) {
	for i := 0; i < 9; i++ {
		var p []int
		for j := 0; j < 9; j++ {
			for k := 49; k < 58; k++ {
				if rune(arg[i][j]) == rune(k) {
					p = append(p, k-48)
				}
			}
			if rune(arg[i][j]) == '.' {
				p = append(p, 0)
			}
		}
		arr[i] = p
	}
}
func check(row int, col int, num int) bool {
	for x := 0; x <= 8; x++ {
		if arr[row][x] == num {
			return false
		}
	}
	for y := 0; y <= 8; y++ {
		if arr[y][col] == num {
			return false
		}
	}
	var startRow int = row - row%3
	var startCol int = col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if arr[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}
func Solve(row int, col int) {
	if row == 8 && col == 9 {
		count++
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				z01.PrintRune(rune(arr[i][j]) + 48)
				z01.PrintRune(' ')
			}
			z01.PrintRune('\n')
		}
		return
	}
	if col == 9 {
		row++
		col = 0
	}
	if arr[row][col] > 0 {
		Solve(row, col+1)
		return
	}
	for num := 1; num <= 9; num++ {
		if check(row, col, num) == true {
			arr[row][col] = num
			Solve(row, col+1)
		}
		arr[row][col] = 0
	}
}
func main() {
	arg := os.Args[1:]
	str := "Error"
	if len(arg) != 9 {
		for _, letter := range str {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < len(arg); i++ {
			if len(arg[i]) != 9 {
				for _, letter := range str {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
				return
			}
		}
		converter(arg)
		Solve(0, 0)
		if count != 1 {
			for _, letter := range str {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}
	}
}