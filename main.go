package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bacaInputan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n")
		fmt.Println("      SECUREPASS - UTILITY KATA SANDI    ")
		fmt.Println("")
		fmt.Println("1. Tambah Akun Baru")
		fmt.Println("2. Edit Info Akun")
		fmt.Println("3. Hapus Akun")
		fmt.Println("4. Cari Akun Layanan")
		fmt.Println("5. Urutkan Record Database")
		fmt.Println("6. Lihat Statistik & Kadar Keamanan")
		fmt.Println("7. Cetak Semua Record")
		fmt.Println("8. Keluar Aplikasi")
		fmt.Print("Masukkan pilihan menu (1-8): ")

		bacaInputan.Scan()
		pilihan := strings.TrimSpace(bacaInputan.Text())

		var tombolPilihan int
		fmt.Sscanf(pilihan, "%d", &tombolPilihan)

		switch tombolPilihan {
		case 1:
			menuTambahAkun(bacaInputan)
		case 2:
			menuEditAkun(bacaInputan)
		case 3:
			menuHapusAkun(bacaInputan)
		case 4:
			menuPencarianData(bacaInputan)
		case 5:
			menuPengurutanData(bacaInputan)
		case 6:
			fmt.Println("DEBUG: masuk case 6")
			tampilkanRingkasanStatistik()
		case 7:
			fmt.Println("DEBUG: masuk case 7")
			lihatSemuaData()
		case 8:
			fmt.Println("\nAplikasi SecurePass ditutup. Sampai jumpa!")
			return
		default:
			fmt.Println("\nPilihan menu tidak ada di daftar, coba lagi ya.")
		}

		fmt.Print("\n[Tekan Enter untuk kembali ke menu...]")
		bacaInputan.Scan()
	}
}