package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// implement Interface interface
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	/*tmp := hs[i]
	hs[i] = hs[j]
	hs[j] = tmp*/
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("Hero~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	for _, v := range heros {
		fmt.Println(v)
	}

	sort.Sort(heros)

	for _, v := range heros {
		fmt.Println(v)
	}

}
