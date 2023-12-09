package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	x int
	y int
}

func main() {
	var py, px int
	fmt.Println("Inserisci piega y:")
	fmt.Scan(&py)
	fmt.Println("Inserisci piega x:")
	fmt.Scan(&px)

	var v []Punto

	//leggo i punti da standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), ",")
		var punto Punto
		punto.x, _ = strconv.Atoi(p[0])
		punto.y, _ = strconv.Atoi(p[1])
		v = append(v, punto)
	}

	var mappaPunti map[Punto]int = make(map[Punto]int)

	//quando piego il foglio, è come se i punti sotto/a destra della linea di piega si spostassero sopra/a sinistra
	//di conseguenza modifico le coordinate di questi punti e tengo conto di quanti punti con le stesse coordinate ho, usando una mappa da Punto a int
	//il numero di punti visibili sarà pari alla lunghezza della mia mappa
	for _, punto := range v {
		if punto.x > px && punto.y > py {
			punto.x -= px
			punto.y -= py
		} else if punto.y > py {
			punto.y -= py
		} else if punto.x > px {
			punto.x -= px
		}
		mappaPunti[punto]++
	}

	//fmt.Println(v)

	fmt.Println("Punti visibili:", len(mappaPunti))
}
