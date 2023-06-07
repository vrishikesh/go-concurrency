package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	printSomething("This is the first line", wg)

	words := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "theta", "iota", "kappa", "lambda", "mu", "nu", "xi", "omicron", "pi", "rho", "sigma", "tau", "upsilon", "phi", "chi", "psi", "omega"}

	wg.Add(len(words))
	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), wg)
	}

	wg.Add(1)
	printSomething("This is the last line", wg)

	wg.Wait()
}
