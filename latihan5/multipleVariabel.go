//? Multiple Variabel
/*
	- Di go-lang kita bisa membuat variabel secara sekaligus banayak
	- Code yang dibuat akan lebih bagus dan mudah di baca
*/

package main
import "fmt"

func main(){
	var (
		firstName = "Muh";
		lastName = "Bintang";
	);

	fmt.Println(firstName + " " +  lastName);
}