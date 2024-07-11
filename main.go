package main

import "fmt"

const (
	emptySymbol  = 'S'
	bookedSymbol = 'B'
	priceFront   = 10
	priceBack    = 8
)

func main() {
	rows := getRows()
	seats := getSeats()
	cinema := initCinema(rows, seats)

	menuOption := -1
	for menuOption != 0 {
		menuOption = getMenuOption()
		switch menuOption {
		case 1:
			fmt.Println(toStringCinema(cinema))
		case 2:
			buyTicket(cinema, rows, seats)
		case 3:
			statistics(cinema, rows, seats)
		}
	}
}

func statistics(cinema [][]rune, rows int, seats int) {
	var fullSeats, bookedSeats, totalIncome, currentIncome int
	fullSeats = rows * seats
	for i := range cinema {
		price := calculatePrice(rows, seats, i+1)
		totalIncome += price * seats
		for j := range cinema[i] {
			if cinema[i][j] == bookedSymbol {
				bookedSeats++
				currentIncome += price
			}
		}
	}
	percentage := float64(bookedSeats) * 100.0 / float64(fullSeats)

	fmt.Printf("Number of purchased tickets: %d\n", bookedSeats)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Current income: $%d\n", currentIncome)
	fmt.Printf("Total income: $%d\n", totalIncome)
}

func buyTicket(cinema [][]rune, rows int, seats int) {
	seatRow, seat := askSeat(cinema)
	cinema[seatRow-1][seat-1] = bookedSymbol
	price := calculatePrice(rows, seats, seatRow)
	fmt.Printf("Ticket price:\n$%d\n", price)
}

func askSeat(cinema [][]rune) (int, int) {
	var seatRow, seat int
	for {
		seatRow, seat = getSeat()
		if (seatRow < 1 || seatRow > len(cinema)) || (seat < 1 || seat > len(cinema[0])) {
			fmt.Println("Wrong input!")
			continue
		}
		if cinema[seatRow-1][seat-1] == bookedSymbol {
			fmt.Println("That ticket has already been purchased!")
			continue
		}
		break
	}
	return seatRow, seat
}

func getMenuOption() int {
	fmt.Println()
	fmt.Println("1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("3. Statistics")
	fmt.Println("0. Exit")
	var option int
	fmt.Scanln(&option)
	return option
}

func getRows() int {
	var rows int
	fmt.Println("Enter the number of rows:")
	fmt.Scanln(&rows)
	return rows
}

func getSeats() int {
	var seats int
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scanln(&seats)
	return seats
}

func initCinema(rows, seats int) [][]rune {
	cinema := make([][]rune, rows)
	for i := range cinema {
		cinema[i] = make([]rune, seats)
	}
	for i := range cinema {
		for j := range cinema[i] {
			cinema[i][j] = emptySymbol
		}
	}
	return cinema
}

func toStringCinema(cinema [][]rune) string {
	result := "Cinema:\n "

	for i := 0; i < len(cinema[0]); i++ {
		result += fmt.Sprintf(" %d", i+1)
	}
	for i := range cinema {
		result += fmt.Sprintf("\n%d", i+1)
		for _, value := range cinema[i] {
			result += fmt.Sprintf(" %c", value)
		}
	}
	return result + "\n"
}

func getSeat() (int, int) {
	var row, seat int
	fmt.Println("Enter a row number:")
	fmt.Scanln(&row)
	fmt.Println("Enter a seat number in that row:")
	fmt.Scanln(&seat)
	return row, seat
}

func calculatePrice(rows, seats, seatRow int) int {
	total := rows * seats
	price := 0

	if total <= 60 {
		price = priceFront
	} else {
		frontRows := rows / 2
		if seatRow <= frontRows {
			price = priceFront
		} else {
			price = priceBack
		}
	}
	return price
}
