//? Deklarasi Multiple constat
/*
	- sama seperti pada var. Constant juga bisa membuat multiple constant
*/

package main
import "fmt"

func main(){
	const (
		firstName = "Muhammad";
		lastName = "Bintang";
	);

	fmt.Println(firstName + " " + lastName);
}