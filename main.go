package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var b Board
	ifname := "input.txt"
	ofname := "output.txt"
	fbytes, err := ioutil.ReadFile(ifname)
	filetext := string(fbytes)
	if err != nil {
		log.Panic("File not find", ifname)
	}

	if berr := b.LoadBoard(filetext); berr != nil {
		fmt.Print(berr)
	} else {
		log.Println("Game Board  Loaded successfully from ", ifname)
		// pretty.Print("Board b=", b)
		// fmt.Printf("%v", b.Grid[0][0])
	}

	if err := ParseMissileActions(filetext); err != nil {
		log.Print(err)
	}

	for m := 0; m < TotalMissiles; m++ {
		/// Launch Player 0
		for p := 0; p < NPlayers; p++ {
			hit := b.LaunchMissile(p, Missles[p][m])
			log.Printf("Player %d Launching Missile %d : HIT ? %v", p, m, hit)
		}

	}

	// Uncomment below line to Visualize board !
	// pretty.Print("Board b=\n", b)

	// fmt.Printf("%v", Missles)
	wd, ferr := os.Create(ofname)
	if ferr != nil {
		log.Panic(err)
	}
	defer wd.Close()
	for p := 0; p < NPlayers; p++ {
		fmt.Fprintf(wd, "Player %d\n", p+1)
		for i := 0; i < b.GridSize; i++ {
			for j := 0; j < b.GridSize; j++ {
				fmt.Fprintf(wd, "%s ", b.Grid[p][i][j])
			}
			fmt.Fprintf(wd, "\n")
		}
		fmt.Fprintf(wd, "\n")
	}
	fmt.Fprintf(wd, "P1:%d\n", b.Hits[0])
	fmt.Fprintf(wd, "P2:%d\n", b.Hits[1])
	fmt.Fprintf(wd, b.Result())
	log.Println("File Created ", ofname)

}
