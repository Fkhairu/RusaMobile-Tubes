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
		Nama:         "Dr. Andi",
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

<<<<<<< Updated upstream
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
=======
	// Menampilkan Data Pasien, Dokter, Rumah Sakit
	for {
		var aksi string
		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Input Data Pasien")
		fmt.Println("2. Lihat Daftar Dokter")
		fmt.Println("3. Lihat Daftar Rumah Sakit")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan aksi (1-4): ")
>>>>>>> Stashed changes
		fmt.Scanln(&aksi)

		switch aksi {
		case "1":
			// Input Data Pasien
			pasienID++
			var nama string
			var usia int
			var tinggi, berat float64
			var kondisiInput string

			fmt.Print("Masukkan Nama Pasien: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan Usia Pasien: ")
			fmt.Scanln(&usia)
			fmt.Print("Masukkan Tinggi Pasien (dalam meter): ")
			fmt.Scanln(&tinggi)
			fmt.Print("Masukkan Berat Pasien (dalam kg): ")
			fmt.Scanln(&berat)
			fmt.Print("Masukkan Kondisi Penyakit Pasien (pisahkan dengan koma jika lebih dari satu): ")
			fmt.Scanln(&kondisiInput)

			kondisi := strings.Split(kondisiInput, ",")
			for i := range kondisi {
				kondisi[i] = strings.TrimSpace(kondisi[i]) // Menghapus spasi
			}

			pasienList = append(pasienList, Pasien{
				ID:      pasienID,
				Nama:    nama,
				Usia:    usia,
				Tinggi:  tinggi,
				Berat:   berat,
				Kondisi: kondisi,
			})

			// Mencari Dokter yang cocok
			fmt.Println("\nDokter yang dapat menangani penyakit ini:")
			found := false
			for _, pasien := range pasienList {
				for _, penyakit := range pasien.Kondisi {
					for _, dokter := range dokterList {
						for _, dPenyakit := range dokter.Penyakit {
							if strings.Contains(strings.ToLower(dPenyakit), strings.ToLower(penyakit)) {
								fmt.Printf("ID: %d, Nama: %s, Spesialisasi: %s\n", dokter.ID, dokter.Nama, dokter.Spesialisasi)
								found = true
							}
						}
					}
				}
			}
			if !found {
				fmt.Println("Dokter ini tidak ada.")
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
			// Keluar dari program
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Aksi tidak valid. Silakan coba lagi.")
		}
	}
}
