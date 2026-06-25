package main

// Tempat ngumpulnya cetakan data dan variabel global

// Struktur data utama untuk menampung info akun 
type PasporDigital struct {
	Platform    string // Nama layanan (misal: Netflix, Google)
	EmailUser   string
	SandiRahasia string
	WaktuUpdate string
	DetikInput  int64  // Dipakai buat sorting berdasarkan waktu masuk
}

// Batas maksimal tampungan array sesuai ketentuan soal 
const BatasMaksimal = 200

var kumpulanAkun [BatasMaksimal]PasporDigital
var totalAkunTersimpan int = 0