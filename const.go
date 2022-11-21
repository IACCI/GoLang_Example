package main

import "fmt"

// const kullanimi bu ornekte enum veri tipi gibi kullanilmistir.

// trakzer adinda string yapisinda bir type olusturulur.
type trakzer string

// const degerleri yazilir.
const (
	thy  trakzer = "example hava yollari"
	acc  trakzer = "example"
	snk  trakzer = "example2"
	glng trakzer = "GOLANG"
)

// olusturulan degerleri yazdirmak icin printtrakzer adinda bir func olusturulur.
func printtrakzer(Trakzer trakzer) {

	fmt.Println(Trakzer)

}

// main func icinde olusturulan stringleri yazdirma islemi printtrakzer ile.
func main() {

	printtrakzer(acc)
	printtrakzer(snk)
	printtrakzer(thy)
	printtrakzer(glng)

}
