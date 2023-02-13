//? Variabel
/*
	- Di-Go lang variabel hanya dapat menyimpan tipe data yang sama, jika kita
		ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa 
		variabel
	- Untuk membuat variabel, kita bisa menggunakan kata kunci var, lalu di ikuti dengan
		nama variabel data tipe data nya

	- Jika variabel tidak di pakai maka akan terjadi error
*/


package main
import "fmt"

func main(){
	var nama string;

	nama = "Muhammad Bintang";
	fmt.Println(nama);

	nama = "Zahaka Hidayat";
	fmt.Println(nama);

	var friendName string;
	friendName = "budi";
	fmt.Println(friendName);

	var age = 18; // akan otomatis int => int32
	fmt.Println("umur = ", age);

}



