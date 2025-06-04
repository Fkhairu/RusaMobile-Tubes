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

// Fungsi validasi nama hanya huruf
func isAlphabetic(s string) bool {
	for _, r := range s {
		if !(r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r == ' ') {
			return false
		}
	}
	return true
}

func main() {
	//bagian array
	var pasienList []Pasien
	var dokterList []Dokter
	var rumahSakitList []RumahSakit
	var pasienID, dokterID, rumahSakitID int

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Jerry",
		Spesialisasi: "Umum",
		Penyakit:     []string{"Flu", "Pilek", "Radang Tenggorokan", "Bronkitis"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Siti",
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
		Nama:   "RS Sehat",
		Lokasi: "Jakarta",
		Dokter: dokterList,
	})

	rumahSakitID++
	rumahSakitList = append(rumahSakitList, RumahSakit{
		ID:     rumahSakitID,
		Nama:   "RS Cinta Sehat",
		Lokasi: "Bandung",
		Dokter: dokterList,
	})

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

	// Contoh input pasien baru dengan validasi nama
	// (Letakkan sebelum for loop utama jika ingin input pasien baru)
	var namaPasien string
	for {
		fmt.Print("Masukkan nama pasien (huruf saja): ")
		fmt.Scanln(&namaPasien)
		if isAlphabetic(namaPasien) {
			break
		}
		fmt.Println("Nama hanya boleh berisi huruf dan spasi. Silakan coba lagi.")
	}
	// Lanjutkan proses input pasien lain seperti usia, tinggi, dsb.

	// Menampilkan Data Pasien, Dokter, Rumah Sakit
	for {
		var aksi string
		fmt.Println("\nPilih aksi:")
		fmt.Println("1. Lihat Daftar Pasien")
		fmt.Println("2. Lihat Daftar Dokter")
		fmt.Println("3. Lihat Daftar Rumah Sakit")
		fmt.Println("4. Cari Dokter Berdasarkan Penyakit")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan aksi (1-5): ")
		fmt.Scanln(&aksi)

		switch aksi {
		case "1":
			// Bagian sort Pasien
			fmt.Println("\nDaftar Pasien:")
			sort.Slice(pasienList, func(i, j int) bool {
				return pasienList[i].Usia < pasienList[j].Usia
			})
			for _, pasien := range pasienList {
				fmt.Printf("ID: %d, Nama: %s, Usia: %d, Tinggi: %.2f m, Berat: %.2f kg, Kondisi: %v\n",
					pasien.ID, pasien.Nama, pasien.Usia, pasien.Tinggi, pasien.Berat, pasien.Kondisi)
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
