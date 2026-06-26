package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


//MODEL - Struktur data dan variabel global


// Struktur data utama untuk menampung info akun
type PasporDigital struct {
	Platform     string // Nama layanan (misal: Netflix, Google)
	EmailUser    string
	SandiRahasia string
	WaktuUpdate  string
	DetikInput   int64 // Dipakai buat sorting berdasarkan waktu masuk
}

// Batas maksimal tampungan array sesuai ketentuan soal
const BatasMaksimal = 200 // Array statis buat nampung data akun, dengan kapasitas maksimal 200

var kumpulanAkun [BatasMaksimal]PasporDigital 
var totalAkunTersimpan int = 0


// CRUD - Tambah, Edit, Hapus, Lihat data
// Prosedur buat nambahin data akun baru ke array
func menuTambahAkun(input *bufio.Scanner) {
	if totalAkunTersimpan >= BatasMaksimal {
		fmt.Println("\n[Waduh, memori penyimpanan sudah penuh!]")
		return
	}

	fmt.Println("\n--- INPUT DATA AKUN BARU ---") // Proses input data akun baru
	fmt.Print("Nama Layanan/Platform : ")
	input.Scan()
	layanan := strings.TrimSpace(input.Text())

	fmt.Print("Alamat Email/Username : ") // Proses input email/username
	input.Scan()
	email := strings.TrimSpace(input.Text())

	fmt.Print("Kata Sandi/Password   : ") // Proses input password/sandi
	input.Scan()
	sandi := strings.TrimSpace(input.Text())

	// Ambil waktu realtime laptop buat timestamp
	waktuSekarang := time.Now()

	kumpulanAkun[totalAkunTersimpan] = PasporDigital{
		Platform:     layanan,
		EmailUser:    email,
		SandiRahasia: sandi,
		WaktuUpdate:  waktuSekarang.Format("02-01-2006 15:04:05"),
		DetikInput:   waktuSekarang.UnixNano(), // Simpan dalam detik nanosecond buat keperluan sorting
	}
	totalAkunTersimpan++
	fmt.Println(">> Sip! Akun berhasil disimpan ke sistem.")
}

// Prosedur buat ngedit data yang sudah ada
func menuEditAkun(input *bufio.Scanner) {
	if totalAkunTersimpan == 0 {
		fmt.Println("\n[Data masih kosong, apa yang mau diedit?]")
		return
	}

	fmt.Print("\nMasukkan Nama Layanan yang mau diubah: ") // Proses cari data yang mau diedit berdasarkan nama layanan
	input.Scan()
	target := strings.TrimSpace(input.Text())

	ketemuIdx := -1 // Cari indeks data yang sesuai dengan nama layanan yang dimasukkan
	for i := 0; i < totalAkunTersimpan; i++ {
		if strings.EqualFold(kumpulanAkun[i].Platform, target) {
			ketemuIdx = i
			break
		}
	}

	if ketemuIdx == -1 {
		fmt.Println(">> Maaf, layanan tersebut gak ketemu di database.")
		return
	}

	fmt.Printf("\nData ketemu! Silakan masukkan data baru untuk %s:\n", kumpulanAkun[ketemuIdx].Platform)
	fmt.Print("Email Baru: ")
	input.Scan()
	kumpulanAkun[ketemuIdx].EmailUser = strings.TrimSpace(input.Text())

	fmt.Print("Sandi Baru: ")
	input.Scan()
	kumpulanAkun[ketemuIdx].SandiRahasia = strings.TrimSpace(input.Text())

	// Perbarui tanggal modifikasinya
	kumpulanAkun[ketemuIdx].WaktuUpdate = time.Now().Format("02-01-2006 15:04:05")
	fmt.Println(">> Mantap! Data berhasil diperbarui.")
}

// Prosedur buat ngapus elemen array dengan cara digeser
func menuHapusAkun(input *bufio.Scanner) {
	if totalAkunTersimpan == 0 { // Cek dulu apakah ada data yang bisa dihapus
		fmt.Println("\n[Data kosong, tidak ada yang bisa dihapus]")
		return
	}

	fmt.Print("\nMasukkan Nama Layanan yang mau dihapus: ") 	// Proses cari data yang mau dihapus berdasarkan nama layanan
	input.Scan()
	target := strings.TrimSpace(input.Text())

	ketemuIdx := -1
	for i := 0; i < totalAkunTersimpan; i++ {
		if strings.EqualFold(kumpulanAkun[i].Platform, target) {
			ketemuIdx = i
			break
		}
	}

	if ketemuIdx == -1 {
		fmt.Println(">> Layanan tidak ditemukan.")
		return
	}

	// Proses menggeser elemen array biar gak kosong di tengah
	for i := ketemuIdx; i < totalAkunTersimpan-1; i++ {
		kumpulanAkun[i] = kumpulanAkun[i+1]
	}
	totalAkunTersimpan--
	fmt.Println(">> Sukses! Akun telah dihapus dari penyimpanan.") // Setelah data dihapus, totalAkunTersimpan dikurangi 1 untuk mencerminkan jumlah data yang tersimpan sekarang
}

// Fitur tambahan: Ngecek level keamanan password secara mandiri
func cekKadarSandi(sandi string) string {
	panjang := len(sandi)
	adaKapital := false
	adaAngka := false
// Cek karakter dalam sandi untuk menentukan apakah ada huruf kapital dan angka
	for _, karakter := range sandi {
		if karakter >= 'A' && karakter <= 'Z' { // Cek apakah  ada kapital
			adaKapital = true
		}
		if karakter >= '0' && karakter <= '9' { // Cek apakah ada angka
			adaAngka = true
		}
	}
	// Kriteria sederhana: minimal 8 karakter, ada huruf kapital, dan ada angka
	if panjang >= 8 && adaKapital && adaAngka {
		return "KUAT" // Kriteria kuat: panjang minimal 8, ada huruf kapital, dan ada angka
	} else if panjang >= 6 {
		return "SEDANG" // Kriteria sedang: panjang minimal 6, tapi mungkin kurang lengkap (misal: tidak ada kapital atau angka)
	}
	return "LEMAH" //Kriteria lemah: kurang dari 6 karakter atau tidak memenuhi kriteria lainnya
}

// Prosedur buat nampilin rangkuman data ke layar
func tampilkanRingkasanStatistik() {
	fmt.Println("\nRINGKASAN DATA & KEAMANAN")
	fmt.Printf("Total akun terdaftar: %d data\n", totalAkunTersimpan)

	if totalAkunTersimpan == 0 {
		return
	}
// Hitung jumlah akun berdasarkan kategori keamanan sandi
	k, s, l := 0, 0, 0
	for i := 0; i < totalAkunTersimpan; i++ {
		status := cekKadarSandi(kumpulanAkun[i].SandiRahasia)
		if strings.Contains(status, "KUAT") {
			k++
		} else if strings.Contains(status, "SEDANG") {
			s++
		} else {
			l++
		}
	}
	fmt.Printf("- Kategori Kuat   : %d\n", k) // Tampilkan jumlah akun dengan kategori kuat, sedang, dan lemah berdasarkan hasil cekKadarSandi
	fmt.Printf("- Kategori Sedang : %d\n", s) // Tampilkan jumlah akun dengan kategori sedang
	fmt.Printf("- Kategori Lemah  : %d\n", l) // Tampilkan jumlah akun dengan kategori lemah
}

func lihatSemuaData() {
	if totalAkunTersimpan == 0 {
		fmt.Println("\n[Belum ada data akun yang tersimpan]")
		return
	}
	fmt.Println("\n--- DAFTAR SELURUH DATA AKUN ---")
	for i := 0; i < totalAkunTersimpan; i++ {
		fmt.Printf("%d. Platform: %s | Email: %s | Sandi: %s | Diperbarui: %s\n",
			i+1, kumpulanAkun[i].Platform, kumpulanAkun[i].EmailUser, kumpulanAkun[i].SandiRahasia, kumpulanAkun[i].WaktuUpdate)
	}
}


// SEARCH & SORT - Pencarian dan Pengurutan data

func menuPencarianData(input *bufio.Scanner) {
	if totalAkunTersimpan == 0 {
		fmt.Println("\n[Data masih kosong, jalankan menu tambah data dulu!]")
		return
	}

	fmt.Print("\nKetik nama layanan yang dicari: ")
	input.Scan()
	target := strings.TrimSpace(input.Text())

	fmt.Println("Pilih Algoritma Pencarian:")
	fmt.Println("1. Sequential Search (Bisa langsung cari)")
	fmt.Println("2. Binary Search (Wajib urut alfabet dulu)")
	fmt.Print("Pilihan (1/2): ")

	input.Scan()
	pilihanCari := strings.TrimSpace(input.Text())
	var modelCari int
	fmt.Sscanf(pilihanCari, "%d", &modelCari)

	indexHasil := -1
	if modelCari == 1 {
		indexHasil = metodeSequential(target)
	} else if modelCari == 2 {
		indexHasil = metodeBinary(target)
	}

	if indexHasil != -1 {
		fmt.Println("\n[ DATA DITEMUKAN! ]")
		data := kumpulanAkun[indexHasil]
		fmt.Printf("Layanan    : %s\nEmail/User : %s\nKata Sandi : %s\nTgl Update : %s\nKadar Aman : %s\n",
			data.Platform, data.EmailUser, data.SandiRahasia, data.WaktuUpdate, cekKadarSandi(data.SandiRahasia))
	} else {
		fmt.Println("\n>> Waduh, data yang kamu cari gak ada di sistem.")
	}
}

func metodeSequential(target string) int {
	for i := 0; i < totalAkunTersimpan; i++ {
		if strings.EqualFold(kumpulanAkun[i].Platform, target) {
			return i
		}
	}
	return -1
}

func metodeBinary(target string) int {
	kiri := 0
	kanan := totalAkunTersimpan - 1 // Pastikan data sudah diurutkan alfabetis sebelum pakai binary search

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		bandingkan := strings.Compare(strings.ToLower(kumpulanAkun[tengah].Platform), strings.ToLower(target))

		if bandingkan == 0 { // Ketemu data yang dicari
			return tengah
		} else if bandingkan < 0 {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func menuPengurutanData(input *bufio.Scanner) {
	if totalAkunTersimpan == 0 {
		fmt.Println("\n[Gak ada data yang bisa diurutkan]")
		return
	}
// Tampilkan opsi pengurutan ke pengguna
	fmt.Println("\nMenu Pengurutan:")
	fmt.Println("1. Berdasarkan Nama Layanan (A-Z) -> Selection Sort")
	fmt.Println("2. Berdasarkan Waktu Input (Lama-Baru) -> Insertion Sort")
	fmt.Print("Pilih (1/2): ")

	input.Scan()
	pilihanSort := strings.TrimSpace(input.Text())
	var modelSort int
	fmt.Sscanf(pilihanSort, "%d", &modelSort)

// Selection Sort (A-Z berdasarkan nama layanan)
	if modelSort == 1 {
		jalankanSelectionSort()
		fmt.Println(">> Urutan berhasil dirapikan alfabetis (Selection Sort)!")
	} else if modelSort == 2 {
		jalankanInsertionSort()  //berdasarkan waktu input (lama ke baru)
		fmt.Println(">> Urutan berhasil dirapikan berdasarkan kronologi (Insertion Sort)!")
	} else {
		fmt.Println("Opsi salah pilih.")
	}
}

func jalankanSelectionSort() {
	for i := 0; i < totalAkunTersimpan-1; i++ {
		indeksTerkecil := i
		for j := i + 1; j < totalAkunTersimpan; j++ {
			if strings.ToLower(kumpulanAkun[j].Platform) < strings.ToLower(kumpulanAkun[indeksTerkecil].Platform) {
				indeksTerkecil = j
			}
		}
		kumpulanAkun[i], kumpulanAkun[indeksTerkecil] = kumpulanAkun[indeksTerkecil], kumpulanAkun[i]
	}
}

func jalankanInsertionSort() {
	for i := 1; i < totalAkunTersimpan; i++ {
		kunciData := kumpulanAkun[i]
		j := i - 1
		for j >= 0 && kumpulanAkun[j].DetikInput > kunciData.DetikInput {
			kumpulanAkun[j+1] = kumpulanAkun[j]
			j--
		}
		kumpulanAkun[j+1] = kunciData
	}
}


//MAIN - Entry point aplikasi

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
			tampilkanRingkasanStatistik()
		case 7:
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