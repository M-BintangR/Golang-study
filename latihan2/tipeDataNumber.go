// dalam bahasa go ada dua tipe data numbr
/*
	1. Integer (bilangan bulat)
	2. Floating Point (bilangan desimal)
*/

//? Tipe data integer
/*
	di golang menggunakan key sensitive
	- di bahawah adalah nilai integer yang mempunyai
	nilai minimum mines sampai positif
	1. int8
	2. int16
	3. int32
	4. int64
	- di bawah adalah nilai integer yang tidak mempunyai
	nilai minimum mines hanya di mulai dari 0
	1. uint8
	2. uint16
	3. uint32
	4. uint64
*/ 
//? Floating Point
/*
	1. float32
	2. float64
	3. complex64
	3. complex128
*/
//? Alias (nama lain dari tipe data)
/*
	1. byte = uint8
	2. rune = int32
	3. int = Minimal int32
	4. uint = Minimal uint32
*/

//* Contoh program
package main
import "fmt"

func main(){
	fmt.Println("satu = ", 1);
	fmt.Println("dua = ", 2);
	fmt.Println("tiga koma lima = ", 3.5);
}

