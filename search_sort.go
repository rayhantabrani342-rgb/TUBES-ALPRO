package main

import (
	"bufio"
	"fmt"
	"strings"
)

// BAGIAN ALGORITMA PENCARIAN (SEARCHING)

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
	kanan := totalAkunTersimpan - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		bandingkan := strings.Compare(strings.ToLower(kumpulanAkun[tengah].Platform), strings.ToLower(target))

		if bandingkan == 0 {
			return tengah
		} else if bandingkan < 0 {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

// BAGIAN ALGORITMA PENGURUTAN (SORTING)

func menuPengurutanData(input *bufio.Scanner) {
	if totalAkunTersimpan == 0 {
		fmt.Println("\n[Gak ada data yang bisa diurutkan]")
		return
	}

	fmt.Println("\nMenu Pengurutan:")
	fmt.Println("1. Berdasarkan Nama Layanan (A-Z) -> Selection Sort")
	fmt.Println("2. Berdasarkan Waktu Input (Lama-Baru) -> Insertion Sort")
	fmt.Print("Pilih (1/2): ")

	input.Scan()
	pilihanSort := strings.TrimSpace(input.Text())
	var modelSort int
	fmt.Sscanf(pilihanSort, "%d", &modelSort)

	if modelSort == 1 {
		jalankanSelectionSort()
		fmt.Println(">> Urutan berhasil dirapikan alfabetis (Selection Sort)!")
	} else if modelSort == 2 {
		jalankanInsertionSort()
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