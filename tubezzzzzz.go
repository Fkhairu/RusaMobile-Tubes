package main

import (
	"fmt"
	"strings"
	"sort"  // Menambahkan import sort untuk fungsi sorting
)

// Struktur data untuk pasien
type Pasien struct {
	ID        int
	Nama      string
	Usia      int
	Tinggi    float64 // tinggi badan dalam meter
	Berat     float64 // berat badan dalam kg
	Kondisi   []string // daftar kondisi medis
}

// Struktur data untuk dokter
type Dokter struct {
	ID        int
	Nama      string
	Spesialisasi string
	Pasien    []Pasien // pasien yang ditangani oleh dokter
}

// Struktur data untuk rumah sakit
type RumahSakit struct {
	ID       int
	Nama     string
	Lokasi   string
	Dokter   []Dokter // daftar dokter yang bekerja di rumah sakit
}

func main() {
	var pasienList []Pasien
	var dokterList []Dokter
	var rumahSakitList []RumahSakit
	var pasienID, dokterID, rumahSakitID int

	// Menambahkan data pasien, dokter, dan rumah sakit
	for {
		var aksi string
		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Tambah Dokter")
		fmt.Println("3. Tambah Rumah Sakit")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Pencarian")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih aksi: ")
		fmt.Scanln(&aksi)

		switch aksi {
		case "1":
			// Menambah pasien
			pasienID++
			var nama, kondisi string
			var usia int
			var tinggi, berat float64
			fmt.Print("Masukkan Nama Pasien: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan Usia Pasien: ")
			fmt.Scanln(&usia)
			fmt.Print("Masukkan Tinggi Badan (dalam meter): ")
			fmt.Scanln(&tinggi)
			fmt.Print("Masukkan Berat Badan (dalam kg): ")
			fmt.Scanln(&berat)
			fmt.Print("Masukkan Kondisi Medis (pisahkan dengan koma jika lebih dari satu): ")
			fmt.Scanln(&kondisi)
			kondisiList := strings.Split(kondisi, ",")

			pasienList = append(pasienList, Pasien{
				ID:        pasienID,
				Nama:      nama,
				Usia:      usia,
				Tinggi:    tinggi,
				Berat:     berat,
				Kondisi:   kondisiList,
			})
			fmt.Println("Pasien berhasil ditambahkan!")

		case "2":
			// Menambah dokter
			dokterID++
			var nama, spesialisasi string
			var jumlahPasien int
			fmt.Print("Masukkan Nama Dokter: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan Spesialisasi Dokter: ")
			fmt.Scanln(&spesialisasi)

			// Menambahkan pasien untuk dokter
			fmt.Print("Berapa pasien yang ditangani oleh dokter ini? ")
			fmt.Scanln(&jumlahPasien)

			var pasienTerkait []Pasien
			for i := 0; i < jumlahPasien; i++ {
				var pasienNama string
				fmt.Print("Masukkan Nama Pasien yang Ditangani: ")
				fmt.Scanln(&pasienNama)

				// Mencari pasien yang sesuai
				for _, p := range pasienList {
					if p.Nama == pasienNama {
						pasienTerkait = append(pasienTerkait, p)
					}
				}
			}

			dokterList = append(dokterList, Dokter{
				ID:           dokterID,
				Nama:         nama,
				Spesialisasi: spesialisasi,
				Pasien:       pasienTerkait,
			})
			fmt.Println("Dokter berhasil ditambahkan!")

		case "3":
			// Menambah rumah sakit
			rumahSakitID++
			var nama, lokasi string
			fmt.Print("Masukkan Nama Rumah Sakit: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan Lokasi Rumah Sakit: ")
			fmt.Scanln(&lokasi)

			rumahSakitList = append(rumahSakitList, RumahSakit{
				ID:     rumahSakitID,
				Nama:   nama,
				Lokasi: lokasi,
			})
			fmt.Println("Rumah sakit berhasil ditambahkan!")

		case "4":
			// Sorting data pasien berdasarkan usia
			sort.Slice(pasienList, func(i, j int) bool {
				return pasienList[i].Usia < pasienList[j].Usia
			})

			// Menampilkan data pasien
			fmt.Println("\nDaftar Pasien:")
			for _, pasien := range pasienList {
				fmt.Printf("ID: %d, Nama: %s, Usia: %d, Tinggi: %.2f m, Berat: %.2f kg, Kondisi: %v\n", 
					pasien.ID, pasien.Nama, pasien.Usia, pasien.Tinggi, pasien.Berat, pasien.Kondisi)
			}

			// Menampilkan data dokter
			fmt.Println("\nDaftar Dokter:")
			for _, dokter := range dokterList {
				fmt.Printf("ID: %d, Nama: %s, Spesialisasi: %s\n", dokter.ID, dokter.Nama, dokter.Spesialisasi)
				for _, pasien := range dokter.Pasien {
					fmt.Printf("  Pasien: %s, Usia: %d\n", pasien.Nama, pasien.Usia)
				}
			}

			// Menampilkan data rumah sakit
			fmt.Println("\nDaftar Rumah Sakit:")
			for _, rumahSakit := range rumahSakitList {
				fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", rumahSakit.ID, rumahSakit.Nama, rumahSakit.Lokasi)
			}

		case "5":
			// Pencarian
			var jenisPencarian string
			fmt.Println("\nPilih jenis pencarian:")
			fmt.Println("1. Pencarian Pasien Berdasarkan Nama")
			fmt.Println("2. Pencarian Dokter Berdasarkan Spesialisasi")
			fmt.Println("3. Pencarian Rumah Sakit Berdasarkan Lokasi")
			fmt.Print("Pilih pencarian: ")
			fmt.Scanln(&jenisPencarian)

			switch jenisPencarian {
			case "1":
				var namaCari string
				fmt.Print("Masukkan Nama Pasien yang dicari: ")
				fmt.Scanln(&namaCari)
				for _, pasien := range pasienList {
					if pasien.Nama == namaCari {
						fmt.Printf("Pasien Ditemukan: ID: %d, Nama: %s, Usia: %d, Tinggi: %.2f m, Berat: %.2f kg, Kondisi: %v\n", 
							pasien.ID, pasien.Nama, pasien.Usia, pasien.Tinggi, pasien.Berat, pasien.Kondisi)
					}
				}

			case "2":
				var spesialisasiCari string
				fmt.Print("Masukkan Spesialisasi Dokter yang dicari: ")
				fmt.Scanln(&spesialisasiCari)
				for _, dokter := range dokterList {
					if dokter.Spesialisasi == spesialisasiCari {
						fmt.Printf("Dokter Ditemukan: ID: %d, Nama: %s, Spesialisasi: %s\n", dokter.ID, dokter.Nama, dokter.Spesialisasi)
					}
				}

			case "3":
				var lokasiCari string
				fmt.Print("Masukkan Lokasi Rumah Sakit yang dicari: ")
				fmt.Scanln(&lokasiCari)
				for _, rumahSakit := range rumahSakitList {
					if rumahSakit.Lokasi == lokasiCari {
						fmt.Printf("Rumah Sakit Ditemukan: ID: %d, Nama: %s, Lokasi: %s\n", rumahSakit.ID, rumahSakit.Nama, rumahSakit.Lokasi)
					}
				}
			}

		case "6":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Aksi tidak dikenal. Silakan pilih aksi yang valid.")
		}
	}
}
