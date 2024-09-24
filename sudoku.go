package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	if checkErrors(arr) {
		printError()
		return
	}
	sudoku := optimizeSudoku(arr)
	if !checkDuplicate(sudoku) && solveSudoku(&sudoku, len(sudoku)) {
		printSudoku(sudoku)
	} else {
		printError()
	}
}

// checkErrors, girdi hatalarını kontrol eder.
// Hatalar varsa true, aksi takdirde false döner.
func checkErrors(arr []string) bool {
	for _, str := range arr {
		if len(str) != 9 {
			return true
		}
		for _, ch := range str {
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
		}
		if containsLetter(str) {
			return true
		}
	}
	return false
}

// containsLetter, bir stringin içinde harf olup olmadığını kontrol eder.
func containsLetter(str string) bool {
	for _, ch := range str {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			return true
		}
	}
	return false
}

// solveSudoku, Sudoku bulmacasını backtracking kullanarak çözer.
// Bir çözüm bulunursa true, aksi takdirde false döner.
func solveSudoku(arr *[][]int, length int) bool {
	isEmpty := true
	row := -1
	column := -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*arr)[i][j] == 0 {
				row = i
				column = j
				isEmpty = false
				break
			}
		}
	}
	if isEmpty {
		return true
	}
	for number := 1; number <= 9; number++ {
		if isCorrect(*arr, row, column, number) {
			(*arr)[row][column] = number
			if solveSudoku(arr, length) {
				return true
			}
			(*arr)[row][column] = 0
		}
	}
	return false
}

// optimizeSudoku, girdi string dizisini 2D bir tamsayı dizisine dönüştürür.
func optimizeSudoku(arr []string) [][]int {
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	for i, str := range arr {
		for j, ch := range str {
			sudoku[i][j] = runeToInt(ch)
		}
	}
	return sudoku
}

// printSudoku, Sudoku gridini yazdırır.
func printSudoku(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			z01.PrintRune(rune(arr[i][j] + '0'))
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

// isCorrect, bir sayıyı belirli bir pozisyona yerleştirmenin doğru olup olmadığını kontrol eder.
func isCorrect(arr [][]int, row int, column int, num int) bool {
	return !checkRow(arr, row, num) && !checkColumn(arr, column, num) && !checkSubSudoku(arr, row-(row%3), column-(column%3), num)
}

// checkRow, belirtilen satırda bir sayının var olup olmadığını kontrol eder.
func checkRow(arr [][]int, row int, num int) bool {
	for column := 0; column < len(arr[row]); column++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

// checkColumn, belirtilen sütunda bir sayının var olup olmadığını kontrol eder.
func checkColumn(arr [][]int, column int, num int) bool {
	for row := 0; row < len(arr); row++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

// checkSubSudoku, belirtilen 3x3 alt ızgarada bir sayının var olup olmadığını kontrol eder.
func checkSubSudoku(arr [][]int, rowIndex int, columnIndex int, num int) bool {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if arr[rowIndex+row][columnIndex+column] == num {
				return true
			}
		}
	}
	return false
}

// runeToInt, bir rune'ı karşılık gelen tamsayı değerine dönüştürür.
func runeToInt(number rune) int {
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}

// printError, bir hata mesajı yazdırır.
func printError() {
	errors := "Error"
	for _, char := range errors {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// checkDuplicate, satırlarda ve sütunlarda yinelenen değerleri kontrol eder.
// Eğer yinelenen değerler bulunursa true, aksi takdirde false döner.
func checkDuplicate(arr [][]int) bool {
	// Satır kontrolü
	for row := 0; row < len(arr); row++ {
		for col1 := 0; col1 < len(arr[row])-1; col1++ {
			for col2 := col1 + 1; col2 < len(arr[row]); col2++ {
				if arr[row][col1] != 0 && arr[row][col1] == arr[row][col2] {
					return true
				}
			}
		}
	}
	// Sütun kontrolü
	for col := 0; col < len(arr[0]); col++ {
		for row1 := 0; row1 < len(arr)-1; row1++ {
			for row2 := row1 + 1; row2 < len(arr); row2++ {
				if arr[row1][col] != 0 && arr[row1][col] == arr[row2][col] {
					return true
				}
			}
		}
	}
	return false
}
