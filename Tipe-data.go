package main

import "fmt"


func main(){

	//tipe data golang terdiri dari berikut :
	/*
	#number positif :
	uint8 = 0 - 255
	unit16 = 0 - 65535
    unit32 = 0 - 4294967295
	unit64 = 0 - 18446744073709551615

	#number negative :
    int8 = -128 - 127
	(kecuali value nilai diberikan minus)
	*/
	var number_positif uint8 = 10 // 0 - 225
    var desimal = 2.55 //perlu tanda titik untuk penulisan desimal karena sebagai pemisah angka
    var check bool = true //nilai boolean meiliki dua nilai yaitu true or false
    var name string = "Trisna" //untuk string penggunaan sama pada umum nya menggunakan tanda "" or tanda ''
    //nilai nil merupakan bukan tipe data melainkan sebuah nilai,variabel yang isinya nill berarti meniliki kosong, 
	//biasanya berada pada pngembalian nilai error.
	
	/* ada beberapa nilai nill yang bisa di set sama nill 
    1. pointer,
	2. tipe data fungsi,
	3. / sleas
	4. map,
	5. cenel,
	6. inteface kosong,
	7. interface
	*/

	fmt.Println(number_positif)
	fmt.Println(desimal)
	fmt.Println(check)
	fmt.Println(name)

}