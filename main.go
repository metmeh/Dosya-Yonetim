package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var secim int
	fmt.Println("Dosya Yönetim Uygulaması")
	fmt.Println("-------------------------")
	for {
		fmt.Println("Yapmak istediğiniz işlemi seçiniz: \n1.Dosya Oluştur \n2.Dosya Sil \n3.Dosya Taşı \n4.Çıkış")
		fmt.Scan(&secim)

		switch secim {
		case 1:
			dosyaOlustur()
		case 2:
			dosyaSil()
		case 3:
			dosyaTasi()
		case 4:
			fmt.Println("Çıkıs Yapılıyor.....")
			return
		default:
			fmt.Println("Geçersiz seçim! lütfen tekrar deneyiniz.")
		}
	}
}

func dosyaOlustur() {

	var dosyaAdı string

	fmt.Print("Oluşturulacak Dosya Adını Giriniz: ")
	fmt.Scanln(&dosyaAdı)

	os.Create(dosyaAdı)
	fmt.Println("Dosya Oluşturuldu.")
}

func dosyaSil() {
	var dosyaAdı string
	fmt.Print("Silinecek dosyanın adnını giriniz: ")
	fmt.Scanln(&dosyaAdı)

	err := os.Remove(dosyaAdı)
	if err != nil {
		fmt.Print("Dosya Silinemedi ", err)
		return
	}
	fmt.Println("Dosya Silindi.")
}

func dosyaTasi() {

	var dosyaAdı, tasi string

	fmt.Println("Taşıncak dosyanın adı: ")
	fmt.Scanln(&dosyaAdı)

	fmt.Println("Taşınacak konum adresi: ")
	fmt.Scanln(&tasi)

	os.Rename(dosyaAdı, filepath.Join(tasi, filepath.Base(dosyaAdı)))

	fmt.Println("Dosya Taşındı.")
}
