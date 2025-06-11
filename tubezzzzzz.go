package main

import (
	"fmt"
	"strings"
	"sort"
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
	Penyakit     []string
}

type RumahSakit struct {
	ID     int
	Nama   string
	Lokasi string
	Dokter []Dokter
}

func main() {
	var pasienList []Pasien
	var dokterList []Dokter
	var rumahSakitList []RumahSakit
	var pasienID, dokterID, rumahSakitID int

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
		Spesialisasi: "Umum",
		Penyakit:     []string{"Sakit Gigi", "Radang Gusi", "Infeksi Gigi"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Jerry",
		Spesialisasi: "Ortopedi",
		Penyakit:     []string{"Cedera ACL", "Patah Tulang", "Fraktur Ligamen", "Cedera Sendi"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Anisa",
		Spesialisasi: "Ortopedi",
		Penyakit:     []string{"Cedera Paha", "Dislokasi Sendi", "Fraktur Tibia"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Lani",
		Spesialisasi: "Gigi",
		Penyakit:     []string{"Gigi Berlubang", "Radang Gusi", "Infeksi Gigi"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Rizal",
		Spesialisasi: "Gigi",
		Penyakit:     []string{"Gigi Sensitif", "Gigi Rontok", "Karies Gigi"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Budi",
		Spesialisasi: "Penyakit Dalam",
		Penyakit:     []string{"Diabetes", "Hipertensi", "Penyakit Jantung"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Mira",
		Spesialisasi: "Penyakit Dalam",
		Penyakit:     []string{"Kolesterol Tinggi", "Gagal Ginjal", "Stroke"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Agus",
		Spesialisasi: "Anak",
		Penyakit:     []string{"Demam", "Batuk", "Flu"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Fitri",
		Spesialisasi: "Anak",
		Penyakit:     []string{"Infeksi Saluran Pernafasan", "Pneumonia", "Cacar Air"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Taufik",
		Spesialisasi: "Kandungan",
		Penyakit:     []string{"Kehamilan", "Endometriosis", "Fibroid Uterus"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Nadia",
		Spesialisasi: "Kandungan",
		Penyakit:     []string{"Infertilitas", "Kanker Serviks", "Menstruasi Tidak Teratur"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Kevin",
		Spesialisasi: "Kulit & Kelamin",
		Penyakit:     []string{"Jerawat", "Psoriasis", "Eksim"},
	})

	dokterID++
	dokterList = append(dokterList, Dokter{
		ID:           dokterID,
		Nama:         "Dr. Dina",
		Spesialisasi: "Kulit & Kelamin",
		Penyakit:     []string{"Vitiligo", "Kanker Kulit", "Infeksi Jamur"},
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
		Lokasi: "Surabaya",
		Dokter: dokterList,
	})

	rumahSakitID++
	rumahSakitList = append(rumahSakitList, RumahSakit{
		ID:     rumahSakitID,
		Nama:   "RS Telkom",
		Lokasi: "Jakarta",
		Dokter: dokterList,
	})

	rumahSakitID++
	rumahSakitList = append(rumahSakitList, RumahSakit{
		ID:     rumahSakitID,
		Nama:   "RS Telkom",
		Lokasi: "Purwokerto",
		Dokter: dokterList,
	})

	for {
	var aksi string
	fmt.Println("\nPilih menu:")
	fmt.Println("1. Daftar Pasien")
	fmt.Println("2. Lihat Daftar Dokter")
	fmt.Println("3. Lihat Daftar Rumah Sakit")
	fmt.Println("4. Cari Dokter berdasarkan Penyakit")
	fmt.Println("5. Cari Dokter berdasarkan Spesialisasi")
	fmt.Println("6. Keluar")
	fmt.Print("Masukkan menu (1-6): ")
	fmt.Scanln(&aksi)

	switch aksi {
	case "1":
		var pilihan string
		fmt.Println("\n1. Lihat Daftar Pasien")
		fmt.Println("2. Tambah Pasien")
		fmt.Print("Pilih menu (1-2): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			if len(pasienList) == 0 {
				fmt.Println("Tidak ada pasien terdaftar.")
			} else {
				fmt.Println("\nDaftar Pasien:")
				for _, pasien := range pasienList {
					fmt.Printf("ID: %d, Nama: %s, Usia: %d, Tinggi: %.2f m, Berat: %.2f kg, Kondisi: %v\n",
						pasien.ID, pasien.Nama, pasien.Usia, pasien.Tinggi, pasien.Berat, pasien.Kondisi)
				}
			}

		case "2":
			var nama, kondisiInput string
			var usia int
			var tinggi, berat float64
			var kondisi []string

			fmt.Print("\nMasukkan Nama Pasien: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan Usia Pasien: ")
			fmt.Scanln(&usia)
			fmt.Print("Masukkan Tinggi Pasien (cm): ")
			fmt.Scanln(&tinggi)
			fmt.Print("Masukkan Berat Pasien (kg): ")
			fmt.Scanln(&berat)
			fmt.Print("Masukkan Kondisi Pasien (pisahkan dengan koma): ")
			fmt.Scanf("%s\n", &kondisiInput)

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

			fmt.Println("Pasien berhasil ditambahkan!")
			fmt.Println("Rekomendasi Dokter berdasarkan kondisi pasien:")

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
		sort.Slice(dokterList, func(i, j int) bool {
			return dokterList[i].Nama < dokterList[j].Nama
		})

		fmt.Println("\nDaftar Dokter:")
		for _, dokter := range dokterList {
			fmt.Printf("ID: %d, Nama: %s, Spesialisasi: %s, Penyakit yang Ditangani: %v\n",
				dokter.ID, dokter.Nama, dokter.Spesialisasi, dokter.Penyakit)
		}

	case "3":
		sort.Slice(rumahSakitList, func(i, j int) bool {
			return rumahSakitList[i].Nama < rumahSakitList[j].Nama
		})

		fmt.Println("\nDaftar Rumah Sakit:")
		for _, rumahSakit := range rumahSakitList {
			fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", rumahSakit.ID, rumahSakit.Nama, rumahSakit.Lokasi)
		}

	case "4":
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

	case "5":
		var spesialisasiCari string
		fmt.Print("\nMasukkan Spesialisasi Dokter yang ingin dicari: ")
		fmt.Scanln(&spesialisasiCari)

		fmt.Println("\nDokter dengan spesialisasi ini:")
		found := false
		for _, dokter := range dokterList {
			if strings.Contains(strings.ToLower(dokter.Spesialisasi), strings.ToLower(spesialisasiCari)) {
				fmt.Printf("ID: %d, Nama: %s, Spesialisasi: %s\n", dokter.ID, dokter.Nama, dokter.Spesialisasi)
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada dokter dengan spesialisasi ini.")
		}

	case "6":
		fmt.Println("Terima kasih!")
		return

	default:
		fmt.Println("Aksi tidak valid. Silakan coba lagi.")
	}
	}
}