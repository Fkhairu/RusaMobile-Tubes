package main

import (
	"fmt"
	"sort"
	"strings"
)

// bagian structure
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
	Penyakit     []string
}

type RumahSakit struct {
	ID     int
	Nama   string
	Lokasi string
	Dokter []Dokter
}

func main() {
	// bagian array
	var pasienList []Pasien
	var dokterList []Dokter
	var rumahSakitList []RumahSakit
	var pasienID, dokterID, rumahSakitID int

	// Menambahkan dokter dan rumah sakit secara manual
	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Ieam",
		Spesialisasi: "Umum",
		Penyakit:     []string{"Flu", "Pilek", "Radang Tenggorokan", "Bronkitis"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Panus",
		Spesialisasi: "Gigi",
		Penyakit:     []string{"Sakit Gigi", "Radang Gusi", "Infeksi Gigi"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Jerry",
		Spesialisasi: "Ortopedi",
		Penyakit:     []string{"Cedera ACL", "Patah Tulang", "Fraktur Ligamen", "Cedera Sendi"},
	})

	rumahSakitID++
	rumahSakitList = append(rumahSakitList, RumahSakit{
		ID:     rumahSakitID,
		Nama:   "RS Telkom",
		Lokasi: "Bandung",
		Dokter: dokterList,
	})

	rumahSakitID++
	rumahSakitList = append(rumahSakitList, RumahSakit{
		ID:     rumahSakitID,
		Nama:   "RS Telkom",
		Lokasi: "Jakarta",
		Dokter: dokterList,
	})

	// Menggunakan pasien yang telah ada
	pasienID++
	pasienList = append(pasienList, Pasien{
		ID:      pasienID,
		Nama:    "Budi",
		Usia:    25,
		Tinggi:  1.75,
		Berat:   70,
		Kondisi: []string{"Flu", "Pilek"},
	})

	pasienID++
	pasienList = append(pasienList, Pasien{
		ID:      pasienID,
		Nama:    "Dewi",
		Usia:    30,
		Tinggi:  1.60,
		Berat:   60,
		Kondisi: []string{"Sakit Gigi", "Radang Gusi"},
	})

	// Menu utama
	for {
		var aksi string
		fmt.Println("\nPilih menu:")
		fmt.Println("1. Daftar Pasien")
		fmt.Println("2. Lihat Daftar Dokter")
		fmt.Println("3. Lihat Daftar Rumah Sakit")
		fmt.Println("4. Cari Dokter Berdasarkan Penyakit")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan menu (1-5): ")
		fmt.Scanln(&aksi)

		switch aksi {
		case "1":
			// Daftar Pasien: Menampilkan pasien dan menambah pasien
			var pilihan string
			fmt.Println("\n1. Lihat Daftar Pasien")
			fmt.Println("2. Tambah Pasien")
			fmt.Print("Pilih menu (1-2): ")
			fmt.Scanln(&pilihan)

			switch pilihan {
			case "1":
				// Menampilkan daftar pasien
				fmt.Println("\nDaftar Pasien:")
				sort.Slice(pasienList, func(i, j int) bool {
					return pasienList[i].Usia < pasienList[j].Usia
				})
				for _, pasien := range pasienList {
					fmt.Printf("ID: %d, Nama: %s, Usia: %d, Tinggi: %.2f m, Berat: %.2f kg, Kondisi: %v\n",
						pasien.ID, pasien.Nama, pasien.Usia, pasien.Tinggi, pasien.Berat, pasien.Kondisi)
				}

			case "2":
				// Menambah pasien baru
				var nama, kondisiInput string
				var usia int
				var tinggi, berat float64
				var kondisi []string

				fmt.Print("\nMasukkan Nama Pasien: ")
				fmt.Scanln(&nama)
				fmt.Print("Masukkan Usia Pasien: ")
				fmt.Scanln(&usia)
				fmt.Print("Masukkan Tinggi Pasien (dalam meter): ")
				fmt.Scanln(&tinggi)
				fmt.Print("Masukkan Berat Pasien (dalam kg): ")
				fmt.Scanln(&berat)
				fmt.Print("Masukkan Kondisi Pasien (pisahkan dengan koma): ")
				fmt.Scanln(&kondisiInput)

				// Memasukkan kondisi ke dalam array
				kondisi = strings.Split(kondisiInput, ",")
				pasienID++
				pasienList = append(pasienList, Pasien{
					ID:      pasienID,
					Nama:    nama,
					Usia:    usia,
					Tinggi:  tinggi,
					Berat:   berat,
					Kondisi: kondisi,
				})

				// Rekomendasi dokter berdasarkan kondisi pasien
				fmt.Println("Pasien berhasil ditambahkan!")
				fmt.Println("Rekomendasi Dokter berdasarkan kondisi pasien:")

				// Mencocokkan kondisi pasien dengan dokter yang relevan
				for _, kondisiPenyakit := range kondisi {
					for _, dokter := range dokterList {
						for _, penyakit := range dokter.Penyakit {
							if strings.Contains(strings.ToLower(penyakit), strings.ToLower(kondisiPenyakit)) {
								fmt.Printf("Dokter: %s, Spesialisasi: %s\n", dokter.Nama, dokter.Spesialisasi)
							}
						}
					}
				}

			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}

		case "2":
			// Bagian sort Dokter
			fmt.Println("\nDaftar Dokter:")
			sort.Slice(dokterList, func(i, j int) bool {
				return dokterList[i].Nama < dokterList[j].Nama
			})
			for _, dokter := range dokterList {
				fmt.Printf("ID: %d, Nama: %s, Spesialisasi: %s, Penyakit yang Ditangani: %v\n",
					dokter.ID, dokter.Nama, dokter.Spesialisasi, dokter.Penyakit)
			}

		case "3":
			// Bagian sort Rumah Sakit
			fmt.Println("\nDaftar Rumah Sakit:")
			sort.Slice(rumahSakitList, func(i, j int) bool {
				return rumahSakitList[i].Lokasi < rumahSakitList[j].Lokasi
			})
			for _, rumahSakit := range rumahSakitList {
				fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", rumahSakit.ID, rumahSakit.Nama, rumahSakit.Lokasi)
			}

		case "4":
			// Bagian search Penyakit
			var penyakitCari string
			fmt.Print("\nMasukkan Penyakit yang ingin dicari: ")
			fmt.Scanln(&penyakitCari)

			fmt.Println("\nDokter yang dapat menangani penyakit ini:")
			found := false
			for _, dokter := range dokterList {
				for _, p := range dokter.Penyakit {
					if strings.Contains(strings.ToLower(p), strings.ToLower(penyakitCari)) {
						fmt.Printf("ID: %d, Nama: %s, Spesialisasi: %s\n", dokter.ID, dokter.Nama, dokter.Spesialisasi)
						found = true
					}
				}
			}
			if !found {
				fmt.Println("Tidak ada dokter yang menangani penyakit ini.")
			}

			fmt.Println("\nRumah Sakit yang dapat menangani penyakit ini:")
			for _, rumahSakit := range rumahSakitList {
				fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", rumahSakit.ID, rumahSakit.Nama, rumahSakit.Lokasi)
			}

		case "5":
			// Keluar dari program
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Aksi tidak valid. Silakan coba lagi.")
		}
	}
}
