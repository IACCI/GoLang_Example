package main

import "fmt"

func main() {
	/*
		mymap := make(map[int]string)

		mymap[12] = "acc"
		mymap[13] = "sda"
		fmt.Println(mymap)
	*/

	states := make(map[string]string)
	states["IST"] = "Istanbul"
	states["ANK"] = "Ankara"
	states["DZC"] = "Duzce"
	states["GS"] = "Galatasaray"
	states["SKR"] = "Sakarya"
	fmt.Println(states)

	// istenilen sehiri elde etme
	Ankara := states["ANK"]
	fmt.Println("secili sehir : ", Ankara)

	// istenilen sehiri silme islemi
	delete(states, "ANK")
	fmt.Println(states)
	fmt.Println("il cikarildi")

	// sehir ekleme
	states["ERZ"] = "Erzurum"
	// map nesnesi adedince key olusturur
	keys := make([]string, len(states))
	i := 0

	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// key listesinde index degerlerine gore states nesnesinde bulunan sehirleri yazdir
	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
