//? Constant
/**
	- menggunakan katakunci const
	- saat pembuatannya kita wajib langsung mengisi / menginisialisasikan datanya
*/

package main
import "fmt"

func main(){
	const firstName = "Bitang";
	const lastName string = "Kocak";
	const value = 1000;

	// kita tidak bisa menggunakan := kayak di variabel

	fmt.Println(firstName);
	fmt.Println(lastName);
	fmt.Println(value);

}
