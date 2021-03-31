package main

// This code developed by Jose' Santos L. is a bot from the thousand-year-old board game Go. Used for the codingame minigame https://www.codingame.com/ide/puzzle/atari-go

//The name of the variables are assigned were assigned based on the vocabulary of the game.
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type cordinate struct {
	X int
	Y int
}

// (c cordinate) String() ...
func (c cordinate) String() string {
	return "(" + strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y) + ")"
}

// find: find the coordinate in the list: return the index of the cordinate or return -1 if the cordinate doesn't exist
// punto: the cordine
func find(punto cordinate, lista []cordinate) int {
	for i, v := range lista {
		if v.X == punto.X && v.Y == punto.Y {
			return i
		}
	}
	return -1

}

//findDestroy: directly destroy the ocupated element.
func findDestroy(punto cordinate, lista []cordinate) []cordinate {
	i := find(punto, lista)
	if i == -1 {
		return lista
	} else {
		return append(lista[:i], lista[i+1:]...)
	}

}

// suicide: try to confirm if this liberty could survive or not
// true: this stone gonna die
// false: can survive
// p: the cordinate to study
// mycolor: the color for the futere stone
// mappa: the map of the board
func suicide(p cordinate, mycolor string, mappa []string) bool {
	//making easiest the use of the cordinates
	x := p.X
	y := p.Y
	//define the color of my opponent.
	var color string
	if mycolor == "W" {
		color = "B"
	} else {
		color = "W"
	}
	if libertyCounter(x-1, y, color, mappa) <= 1 || libertyCounter(x+1, y, color, mappa) <= 1 || libertyCounter(x, y-1, color, mappa) <= 1 || libertyCounter(x, y+1, color, mappa) <= 1 {
		return false
	}

	//if doesn't find any other return, then the position is a suicide position.
	return true

}

// libertyCounter:iniciate the recursive search function libertyCounterR.
// if the first stone is not connected to a chain of stones of the same color, only return the number of of liberties that have that stone.
func libertyCounter(x int, y int, color string, mappa []string) int {

	if x <= 0 || x >= len(mappa) || y <= 0 || y >= len(mappa) || string(mappa[y][x]) != color {
		return 0
	} else {
		aleados := []cordinate{}  //buffer de piedras ya analizadas.
		espacios := []cordinate{} //buffer de libertades encontradas.
		_, espacios = libertyCounterR(x, y, color, mappa, aleados, espacios)
		return len(espacios)
	}

}

// libertyCounterR: recursive function that try to caunt the number of libertys of the stone o a chain of stones.
func libertyCounterR(x int, y int, color string, mappa []string, aleados []cordinate, espacios []cordinate) ([]cordinate, []cordinate) {
	if x <= 0 || x >= len(mappa) || y <= 0 || y >= len(mappa) || string(mappa[y][x]) != color {
		return aleados, espacios
	} else {
		aleados = append(aleados, cordinate{x, y})
		if x > 0 {
			switch string(mappa[y][x-1]) {
			case ".":
				if find(cordinate{x - 1, y}, espacios) == -1 {
					espacios = append(espacios, cordinate{x - 1, y})
				}
			case color:
				if find(cordinate{x - 1, y}, aleados) == -1 {
					aleados, espacios = libertyCounterR(x-1, y, color, mappa, aleados, espacios)
				}
			}
		}
		if x < len(mappa)-1 {
			switch string(mappa[y][x+1]) {
			case ".":
				if find(cordinate{x + 1, y}, espacios) == -1 {
					espacios = append(espacios, cordinate{x + 1, y})
				}
			case color:
				if find(cordinate{x + 1, y}, aleados) == -1 {
					aleados, espacios = libertyCounterR(x+1, y, color, mappa, aleados, espacios)
				}
			}
		}
		if y > 0 {
			switch string(mappa[y-1][x]) {
			case ".":
				if find(cordinate{x, y - 1}, espacios) == -1 {
					espacios = append(espacios, cordinate{x, y - 1})
				}
			case color:
				if find(cordinate{x, y - 1}, aleados) == -1 {
					aleados, espacios = libertyCounterR(x, y-1, color, mappa, aleados, espacios)
				}
			}
		}
		if y < len(mappa)-1 {
			switch string(mappa[y+1][x]) {
			case ".":
				if find(cordinate{x, y + 1}, espacios) == -1 {
					espacios = append(espacios, cordinate{x, y + 1})
				}
			case color:
				if find(cordinate{x, y + 1}, aleados) == -1 {
					aleados, espacios = libertyCounterR(x, y+1, color, mappa, aleados, espacios)
				}
			}

		}
		return aleados, espacios
	}

}

func main() {
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	myColor := scanner.Text()
	_ = myColor // to avoid unused error
	// boardSize: the size of the board (width and height)
	var boardSize int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &boardSize)

	liberta := []cordinate{}    //all the liberty cordinates of my oppontent
	ocupado := []cordinate{}    //all the cordinates that are ocuped by other stone in the board
	suicidecor := []cordinate{} //a list of liberty cordinates that can't be attacked because it's a suicide point at this time.
	mappa := make([]string, boardSize)
	for {
		//all this list are cleaned after each loop of the game
		ocupado = []cordinate{}
		suicidecor = []cordinate{}

		var opponentX, opponentY int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &opponentX, &opponentY)
		if opponentY != -1 {
			if opponentX > 0 && find(cordinate{opponentX - 1, opponentY}, liberta) == -1 {
				liberta = append(liberta, cordinate{opponentX - 1, opponentY})
			}
			if opponentX < boardSize-1 && find(cordinate{opponentX + 1, opponentY}, liberta) == -1 {
				liberta = append(liberta, cordinate{opponentX + 1, opponentY})
			}
			if opponentY > 0 && find(cordinate{opponentX, opponentY - 1}, liberta) == -1 {
				liberta = append(liberta, cordinate{opponentX, opponentY - 1})
			}
			if opponentY < boardSize-1 && find(cordinate{opponentX - 1, opponentY + 1}, liberta) == -1 {
				liberta = append(liberta, cordinate{opponentX, opponentY + 1})
			}

		}

		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)

		for i := 0; i < boardSize; i++ {
			scanner.Scan()
			line := scanner.Text()
			mappa[i] = line
			fmt.Fprintln(os.Stderr, line)
			for j := 0; j < boardSize; j++ {
				if line[j] != '.' {
					//destroy the fake liberties, I don't know why this happens, but this could be the solution.
					liberta = findDestroy(cordinate{j, i}, liberta)
					ocupado = append(ocupado, cordinate{j, i})
				}
			}

		}

		//now that I have the map to compare I can confirm the liberties, and doesnt comete suicie. This doesnt delete the cordinate from the list of liberties, only add to temporal list to
		for _, v := range liberta {
			if suicide(v, myColor, mappa) {
				suicidecor = append(suicidecor, v)
			}
		}

		//print all the possible liberties
		for _, v := range liberta {
			fmt.Fprint(os.Stderr, v, " ")
		}

		if opponentX == -1 && len(ocupado) == 0 {
			fmt.Println(int(boardSize/2), int(boardSize/2))
		} else {
			if len(liberta) != 0 {
				secure := false
				for i := 0; !secure; i++ {
					if find(liberta[i], suicidecor) == -1 && find(liberta[i], ocupado) == -1 {
						secure = true
						fmt.Println(liberta[i].X, liberta[i].Y)
						liberta = findDestroy(liberta[i], liberta)

					}

				}

			} else {
				fmt.Fprint(os.Stderr, "Random selection: ")
				pt := cordinate{r1.Intn(boardSize - 1), r1.Intn(boardSize - 1)}
				for find(pt, ocupado) != -1 && find(pt, suicidecor) != -1 {
					pt = cordinate{r1.Intn(boardSize), r1.Intn(boardSize)}
				}

				fmt.Println(pt.X, pt.Y)
				suicidecor = []cordinate{pt}

			}
		}

	}
}
