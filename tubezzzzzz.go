package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pasien struct {
	ID      int
	Nama    string
	Usia    int
	Tinggi  float64
	Berat   float64
	Kondisi []string
}

type Dokter struct {
	ID           int
	Nama         string
	Spesialisasi string
	Pasien       []Pasien
}

type RumahSakit struct {
	ID     int
	Nama   string
	Lokasi string
	Dokter []Dokter
}

func main() {
	var pasienID int

	// Data awal dokter dan rumah sakit
	dokterList := []Dokter{
		{ID: 1, Nama: "Dr. Andi", Spesialisasi: "Penyakit Dalam"},
		{ID: 2, Nama: "Dr. Budi", Spesialisasi: "Anak"},
		{ID: 3, Nama: "Dr. Citra", Spesialisasi: "Bedah Umum"},
		{ID: 4, Nama: "Dr. Dedi", Spesialisasi: "Kandungan"},
		{ID: 5, Nama: "Dr. Evi", Spesialisasi: "THT"},
		{ID: 6, Nama: "Dr. Fajar", Spesialisasi: "Mata"},
		{ID: 7, Nama: "Dr. Gita", Spesialisasi: "Gigi"},
	}
	rumahSakitList := []RumahSakit{
		{ID: 1, Nama: "RS Wico Selatan", Lokasi: "Jakarta", Dokter: []Dokter{dokterList[0]}},
		{ID: 2, Nama: "RS Wico Utara", Lokasi: "Bandung", Dokter: []Dokter{dokterList[1]}},
		{ID: 3, Nama: "RS Wico Barat", Lokasi: "Surabaya", Dokter: []Dokter{dokterList[2]}},
		{ID: 4, Nama: "RS Wico Timur", Lokasi: "Yogyakarta", Dokter: []Dokter{dokterList[3]}},
		{ID: 5, Nama: "RS SWK Ketintang", Lokasi: "Medan", Dokter: []Dokter{dokterList[4]}},
		{ID: 6, Nama: "RS SWK Gayungan", Lokasi: "Makassar", Dokter: []Dokter{dokterList[5]}},
		{ID: 7, Nama: "RS Swk kebon", Lokasi: "Palembang", Dokter: []Dokter{dokterList[6]}},
	}

	for {
		var aksi string
		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Pencarian")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih aksi: ")
		fmt.Scanln(&aksi)

		switch aksi {
		case "1":
			pasienID++
			var nama, kondisi string
			var usia int
			var tinggi, berat float64

			fmt.Print("Masukkan Nama Pasien: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan Usia Pasien: ")
			fmt.Scanln(&usia)
			fmt.Print("Masukkan Tinggi Badan (m): ")
			fmt.Scanln(&tinggi)
			fmt.Print("Masukkan Berat Badan (kg): ")
			fmt.Scanln(&berat)
			fmt.Print("Masukkan Kondisi Medis (pisahkan dengan koma): ")
			fmt.Scanln(&kondisi)
			kondisiList := strings.Split(kondisi, ",")

			// Pilih Rumah Sakit
			fmt.Println("Pilih Rumah Sakit:")
			for _, rs := range rumahSakitList {
				fmt.Printf("%d. %s (%s)\n", rs.ID, rs.Nama, rs.Lokasi)
			}
			var pilihRS int
			fmt.Print("Masukkan ID Rumah Sakit: ")
			fmt.Scanln(&pilihRS)

			// Validasi Rumah Sakit
			var rumahSakit *RumahSakit
			for i := range rumahSakitList {
				if rumahSakitList[i].ID == pilihRS {
					rumahSakit = &rumahSakitList[i]
					break
				}
			}
			if rumahSakit == nil {
				fmt.Println("Rumah sakit tidak ditemukan.")
				continue
			}

			// Pilih Dokter di Rumah Sakit tersebut
			fmt.Println("Pilih Dokter:")
			for _, d := range rumahSakit.Dokter {
				fmt.Printf("%d. %s (%s)\n", d.ID, d.Nama, d.Spesialisasi)
			}
			var pilihDokter int
			fmt.Print("Masukkan ID Dokter: ")
			fmt.Scanln(&pilihDokter)

			// Validasi Dokter
			var dokter *Dokter
			for i := range rumahSakit.Dokter {
				if rumahSakit.Dokter[i].ID == pilihDokter {
					dokter = &rumahSakit.Dokter[i]
					break
				}
			}
			if dokter == nil {
				fmt.Println("Dokter tidak ditemukan.")
				continue
			}

			// Tambah pasien ke dokter
			pasienBaru := Pasien{
				ID:      pasienID,
				Nama:    nama,
				Usia:    usia,
				Tinggi:  tinggi,
				Berat:   berat,
				Kondisi: kondisiList,
			}
			dokter.Pasien = append(dokter.Pasien, pasienBaru)
			fmt.Println("Pasien berhasil ditambahkan!")

		case "2":
			for _, rs := range rumahSakitList {
				fmt.Printf("RS: %s (%s)\n", rs.Nama, rs.Lokasi)

				// 🔽 SORTING DOKTER BERDASARKAN NAMA
				sort.Slice(rs.Dokter, func(i, j int) bool {
					return rs.Dokter[i].Nama < rs.Dokter[j].Nama
				})

				for _, d := range rs.Dokter {
					fmt.Printf("  Dokter: %s (%s)\n", d.Nama, d.Spesialisasi)

					// 🔽 SORTING PASIEN BERDASARKAN USIA
					sort.Slice(d.Pasien, func(i, j int) bool {
						return d.Pasien[i].Usia < d.Pasien[j].Usia
					})

					for _, p := range d.Pasien {
						fmt.Printf("    Pasien: %s, Usia: %d, Tinggi: %.2f, Berat: %.2f, Kondisi: %v\n",
							p.Nama, p.Usia, p.Tinggi, p.Berat, p.Kondisi)
					}
				}
			}

		case "3":
			var cariNama string
			fmt.Print("Masukkan nama pasien yang dicari: ")
			fmt.Scanln(&cariNama)
			ditemukan := false
			for _, rs := range rumahSakitList {
				for _, d := range rs.Dokter {
					for _, p := range d.Pasien {
						if strings.EqualFold(p.Nama, cariNama) {
							fmt.Printf("Pasien ditemukan di %s, Dokter: %s\n", rs.Nama, d.Nama)
							fmt.Printf("  Usia: %d, Tinggi: %.2f, Berat: %.2f, Kondisi: %v\n", p.Usia, p.Tinggi, p.Berat, p.Kondisi)
							ditemukan = true
						}
					}
				}
			}
			if !ditemukan {
				fmt.Println("Pasien tidak ditemukan.")
			}

		case "4":
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak dikenal.")
		}
	}
}
