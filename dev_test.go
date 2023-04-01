package NPM

import (
	"fmt"
	"testing"

	model "github.com/MarlinaLubis/be_bap/model"
	module "github.com/MarlinaLubis/be_bap/module"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Surabaya"
	phonenumber := 6281256892345
	checkin := "Masuk"
	biodata := []model.Mahasiswa{
		{
			Nama:         "Amar Tana",
			Phone_number:  6281256892345,
			Email:        "amartana@gmail.com",
			Jurusan:      "D4 Teknik Informatika",
			Jam_sidang:   "10.00",
		},
	}

	hasil := module.InsertPresensi(module.MongoConn, "presensi", long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetPresensiFromCheckin(t *testing.T) {
	checkin := "Masuk"
	data := module.GetPresensiFromCheckin(checkin, module.MongoConn, "presensi")
	fmt.Println(data)
}

func TestInsertMahasiswa(t *testing.T) {
	nama := "Amar Tana"
	phone_number := 6281211345090
	email := "amartana@gmail.com"
	jurusan := "D4 Teknik Informatika"
	jam_sidang := "10:00"

	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, phone_number, email, jurusan, jam_sidang)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromNama(t *testing.T) {
	nama := "Amar Tana"
	biodata := module.GetMahasiswaFromNama(nama, module.MongoConn, "mahasiswa")
	fmt.Println(biodata)
}

func TestInsertDosen(t *testing.T) {
	nama := "Noviana Riza, S.Si., M.T., SFPC"
	phone_number := 6281377908321
	email := "novianarizau@ulbi.ac"
	jam_sidang := "10.00"
	hari_sidang := "Senin"

	hasil := module.InsertDosen(module.MongoConn, "dosen", nama, phone_number, email, jam_sidang, hari_sidang)
	fmt.Println(hasil)
}

func TestGetDosenFromNama(t *testing.T) {
	nama := "Noviana Riza, S.Si., M.T., SFPC"
	data := module.GetDosenFromNama(nama, module.MongoConn, "dosen")
	fmt.Println(data)
}

func TestInsertJamSidang(t *testing.T) {
	durasi := 15
	jam_masuk := "10:00"
	jam_keluar := "11:00"
	gmt := 7
	hari := "Senin"

	hasil := module.InsertJamSidang(module.MongoConn, "jamsidang", durasi, jam_masuk, jam_keluar, gmt, hari)
	fmt.Println(hasil)
}

func TestGetJamSidangFromDurasi(t *testing.T) {
	durasi := 15
	data := module.GetJamSidangFromDurasi(durasi, module.MongoConn, "jamsidang")
	fmt.Println(data)
}

func TestInsertBap(t *testing.T) {
	judul := "Aplikasi Penjualan Baju Berbasis Web"
	nama_mahasiswa := "Amar Tana"
	prodi := "D4 Teknik Informatika"
	dosen := "Noviana Riza, S.Si., M.T., SFPC"
	jam_sidang := "10:00"

	hasil := module.InsertBap(module.MongoConn, "bap", judul, nama_mahasiswa, prodi, dosen, jam_sidang)
	fmt.Println(hasil)
}

func TestGetBapFromJudul(t *testing.T) {
	judul := "Aplikasi Penjualan Baju Berbasis Web"
	data := module.GetBapFromJudul(judul, module.MongoConn, "bap")
	fmt.Println(data)
}

func TestGetAllBapFromJudul(t *testing.T) {
	judul := "Aplikasi Penjualan Baju Berbasis Web"
	data := module.GetAllBapFromJudul(judul, module.MongoConn, "bap")
	fmt.Println(data)
}

func TestGetAllPresensi(t *testing.T) {
	data := module.GetAllPresensi(module.MongoConn, "presensi")
	fmt.Println(data)
}
