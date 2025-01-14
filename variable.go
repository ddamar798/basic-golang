package main

import "fmt"

func main() {
    // Deklarasi variabel dengan tipe data
    var nama string = "Damar"
    var umur int = 16
    var tinggi float64 = 180.5
    var isMember bool = true

    // Deklarasi variabel tanpa tipe data (tipe akan ditentukan secara otomatis)
    var hobi = "Ngoding"
    var skor = 95.5

    // Deklarasi variabel menggunakan pendeklarasian singkat (hanya dapat digunakan di dalam fungsi)
    kota := "Sukoharjo"
    negara := "Adoh Kono"

    // Menampilkan nilai variabel
    fmt.Println("Nama:", nama)
    fmt.Println("Umur:", umur)
    fmt.Println("Tinggi:", tinggi)
    fmt.Println("Anggota:", isMember)
    fmt.Println("Hobi:", hobi)
    fmt.Println("Skor:", skor)
    fmt.Println("Kota:", kota)
    fmt.Println("Negara:", negara)

    // Deklarasi banyak variabel sekaligus
    var (
        pekerjaan = "Programmer"
		posisi = "Back-end"
        gaji      = 15000000
    )

    fmt.Println("Pekerjaan:", pekerjaan)
	fmt.Println("posisi", posisi)
    fmt.Println("Gaji:", gaji)

    // Konstanta (nilai tetap yang tidak dapat diubah)
    const pi = 3.14
    fmt.Println("Nilai Pi:", pi)
}
