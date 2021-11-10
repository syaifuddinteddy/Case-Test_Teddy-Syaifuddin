package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/syaifuddin.teddy/test-case-majoo/api/responses"
)

func (server *Server) SudokuSolver(w http.ResponseWriter, r *http.Request) {
	// Get 9x9 Array Input'an Sudoku
	decoder := json.NewDecoder(r.Body)
	var t [9][9]int
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	// Console Input'an Sudoku
	formatPrint(t, "Input")

	if backtrackAlgoritm(&t) {
		// Console Hasil Akhir Sudoku
		formatPrint(t, "Solved")

		// Output
		responses.JSON(w, http.StatusOK, t)
	} else {
		// Jika Gagal
		fmt.Printf("The Sudoku can't be solved.")
		responses.ERROR(w, http.StatusUnauthorized, errors.New("The Sudoku can't be solved."))
	}
}

func formatPrint(input [9][9]int, IOtext string) {
	fmt.Println("Sudoku -", IOtext)
	fmt.Println("+-------------------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			fmt.Printf("%d ", input[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 8 {
			fmt.Println("\n+-------------------+")
		} else {
			fmt.Println()
		}
	}
}

func backtrackAlgoritm(sudoku *[9][9]int) bool {
	if !hasEmptyCell(sudoku) { // jika sudah tidak ada nilai 0 pada sudoku, maka dianggap selesai
		return true
	}

	// pemecahan masalah
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 { // jika terdapat nilai 0
				for candidate := 9; candidate >= 1; candidate-- { //mencoba mengisi nilai yg 0
					sudoku[i][j] = candidate
					if isSudokuValid(sudoku) { // apabila tidak ada yg kembar
						if backtrackAlgoritm(sudoku) { // lanjut ke nilai 0 berikutnya
							return true
						}
						sudoku[i][j] = 0
					} else {
						sudoku[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

// Check 0 value
func hasEmptyCell(sudoku *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isSudokuValid(board *[9][9]int) bool {
	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
