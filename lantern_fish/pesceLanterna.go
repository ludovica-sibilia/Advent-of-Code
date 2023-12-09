package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//inserisco gli input in una mappa che ne conti le occorrenze
	var mappa map[string]int = make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ",")
		for _, s := range riga {
			mappa[s]++
		}
	}

	var giorni int = 256
	/*fmt.Println("Inserisci il numero di giorni:")
	fmt.Scan(&giorni)*/

	/*
		Ogni giorno, il numero di pesci con un clock interno di n diventa pari al numero di pesci che il giorno precedente avevano clock n+1.
		Le uniche eccezioni sono 6 e 8: nel primo caso il numero di pesci sarà pari alla somma di quelli che il giorno precedente avevano clock
		0 e 7 (i 7 diminuiscono di 1 e diventano 6, ogni 0 aggiunge un 6); nel secondo caso il numero sarà uguale al numero di pesci che il
		giorno precedente avevano clock 0.
	*/
	for i := 0; i < giorni; i++ {
		mappa["6"], mappa["0"], mappa["1"], mappa["2"], mappa["3"], mappa["4"], mappa["5"], mappa["7"], mappa["8"] = mappa["0"]+mappa["7"], mappa["1"], mappa["2"], mappa["3"], mappa["4"], mappa["5"], mappa["6"], mappa["8"], mappa["0"]
	}

	var numeroPesci int = mappa["0"] + mappa["1"] + mappa["2"] + mappa["3"] + mappa["4"] + mappa["5"] + mappa["6"] + mappa["7"] + mappa["8"]
	fmt.Println("Dopo", giorni, "giorni ci sono", numeroPesci, "pesci")
}
