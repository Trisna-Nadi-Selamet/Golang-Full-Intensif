package main

import "fmt"

func main(){
	
	//penulisan variabel pada golang :

	//cara penulisan pertama
	var namadepan string = "Trisna"
	
	//cara penulisan kedua
    var nama_belakang = "Selamet"

	//cara penuisan ketiga
	nama_tengah := "Nadi"

    //merubah nilai pada variabel
	nama_belakang = "Fadlah"

	//output
	fmt.Printf("Hallo %s %s %s",namadepan,nama_tengah,nama_belakang)
}