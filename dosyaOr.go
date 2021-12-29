package main

/*   //DOSYA OLUŞTURMA
var (
	yeniDosya *os.File
	err       error
)

func main() {

	yeniDosya, err := os.Create("main.txt")
	if err != nil {
		log.Fatal(err)
		yeniDosya.Close()
	}

}
*/

/*
var (
	DosyBilgisi *os.FileInfo
	err         error
)

func main() {
	DosyBilgisi, err := os.Stat("main.txt")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Dosya isim: ", DosyBilgisi.Name())
	fmt.Println("Dosya Boyutu: ", DosyBilgisi.Size())
	fmt.Println("izinler: ", DosyBilgisi.Mode())
	fmt.Println("Güncelleme Tarihi: ", DosyBilgisi.ModTime())

}
*/
/* dosya kopyalama
func main() {
	orjinal, err := os.Open("main.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer orjinal.Close()

	yenidos, _ := os.Create("main2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer yenidos.Close()

	icerik, err := io.Copy(yenidos, orjinal)

	log.Printf("%d byte kopyalandı", icerik)

	err = yenidos.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
*/
