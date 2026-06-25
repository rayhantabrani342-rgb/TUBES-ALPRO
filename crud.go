package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

// Prosedur buat nambahin data akun baru ke array 
func menuTambahAkun(input *bufio.Scanner) {
	if totalAkunTersimpan >= BatasMaksimal {
		fmt.Println("\n[Waduh, memori penyimpanan sudah penuh!]")
		return
	}

	fmt.Println("\n--- INPUT DATA AKUN BARU ---")
	fmt.Print("Nama Layanan/Platform : ")
	input.Scan()
	layanan := strings.TrimSpace(input.Text())

	fmt.Print("Alamat Email/Username : ")
	input.Scan()
	email := strings.TrimSpace(input.Text())

	fmt.Print("Kata Sandi/Password   : ")
	input.Scan()
	sandi := strings.TrimSpace(input.Text())

	// Ambil waktu realtime laptop buat timestamp
	waktuSekarang := time.Now()

	//
	kumpulanAkun[totalAkunTersimpan] = PasporDigital{
		Platform:     layanan,
		EmailUser:    email,
		SandiRahasia: sandi,
		WaktuUpdate:  waktuSekarang.Format("02-01-2006 15:04:05"),
		DetikInput:   waktuSekarang.UnixNano(),
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

	fmt.Print("\nMasukkan Nama Layanan yang mau diubah: ")
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

	//perbarui tanggal modifikasinya
	kumpulanAkun[ketemuIdx].WaktuUpdate = time.Now().Format("02-01-2006 15:04:05")
	fmt.Println(">> Mantap! Data berhasil diperbarui.")
}

//prosedur buat ngapus elemen array dengan cara digeser 
func menuHapusAkun(input *bufio.Scanner) {
	if totalAkunTersimpan == 0 {
		fmt.Println("\n[Data kosong, tidak ada yang bisa dihapus]")
		return
	}

	fmt.Print("\nMasukkan Nama Layanan yang mau dihapus: ")
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

	//proses menggeser elemen array biar gak kosong di tengah
	for i := ketemuIdx; i < totalAkunTersimpan-1; i++ {
		kumpulanAkun[i] = kumpulanAkun[i+1]
	}
	totalAkunTersimpan--
	fmt.Println(">> Sukses! Akun telah dihapus dari penyimpanan.")
}

//fitur tambahan: Ngecek level keamanan password secara mandiri
func cekKadarSandi(sandi string) string {
	panjang := len(sandi)
	adaKapital := false
	adaAngka := false

	for _, karakter := range sandi {
		if karakter >= 'A' && karakter <= 'Z' {
			adaKapital = true
		}
		if karakter >= '0' && karakter <= '9' {
			adaAngka = true
		}
	}
	// kriteria sederhana: minimal 8 karakter, ada huruf kapital, dan ada angka
	if panjang >= 8 && adaKapital && adaAngka { //kalau ada kapital dan angka itu kuat
		return "KUAT"
	} else if panjang >= 6 {
		return "SEDANG"
	}
	return "LEMAH"
}

// Prosedur buat nampilin rangkuman data ke layar 
func tampilkanRingkasanStatistik() {
	fmt.Println("\n=== RINGKASAN DATA & KEAMANAN ===")
	fmt.Printf("Total akun terdaftar: %d data\n", totalAkunTersimpan)

	if totalAkunTersimpan == 0 {
		return
	}

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
	fmt.Printf("- Kategori Kuat   : %d\n", k)
	fmt.Printf("- Kategori Sedang : %d\n", s)
	fmt.Printf("- Kategori Lemah  : %d\n", l)
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