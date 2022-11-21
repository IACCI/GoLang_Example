package main

import "fmt"

func main() {

	//    BOS     TANIMLAYICI   ( _ )

	// " _ " alt tre bos olarak tanimlanir fakat bir deger atanmasi gerekir. bu atanan degeri ise yazdiramayiz.
	// ornekte gorundugu gibi
	// sadece x olasiligi disinda baska bir olasiliga da olasilik verir ve alt tre koyulur.

	var _, x int = 0, 999
	fmt.Println(x)

}
