package main

import (
	"fmt"
	"math/rand" // paket
	"strings"   // paket
)

func main() {
	// random sayi atama
	fmt.Println("benim favori sayim", rand.Intn(100))
	// 2. string deger 1. string degerde varsa true yoksa false cevirir.
	fmt.Println(strings.Contains("test", "es"))
	// 2. string deger 1. string degerde kac tane var?
	fmt.Println(strings.Count("test", "t"))
	// on ek dogru mu kontrol eder.
	fmt.Println(strings.HasPrefix("unit_test", "unit"))
	// son ek dogru mu kontrol eder.
	fmt.Println(strings.HasSuffix("dosya.rar", "unit"))
	// 2.stringdeki deger 1. stringde kacinci index
	fmt.Println(strings.Index("test", "s"))

}
