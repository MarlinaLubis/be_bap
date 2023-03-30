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
	lokasi := "Universitas Logistik Bisnis dan International"
	phonenumber := 6287345699211
	checkin := "Masuk"
	biodata := model.Mahasiswa{
		Nama:         "Marlina Magdalena Lubis",
		Phone_number: 6281234678922,
		Email:        "lubismarlina06@gmail.com",
		Jurusan:      "D4 Teknik Informatika",
		Jam_sidang:   "08.00",
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
	nama := "Angelina"
	phone_number := 6281234678922
	email := "angelina18@gmail.com"
	jurusan := "D4 Teknik Informatika"
	jam_sidang := "10:00"
    
	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, phone_number, email, jurusan, jam_sidang,)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromNama(t *testing.T) {
	nama := "Angelina"
	biodata := module.GetMahasiswaFromNama(nama, module.MongoConn, "mahasiswa")
	fmt.Println(biodata)
}

func TestInsertDosen(t *testing.T) {
	nama := "Indra Riksa"
	phone_number := 6282556890062
	email := "indra@ulbi.ac"
	jam_sidang := "10.00"
	hari_sidang := "Selasa"

	hasil := module.InsertDosen(module.MongoConn, "dosen", nama, phone_number, email, jam_sidang, hari_sidang)
	fmt.Println(hasil)
}

func TestGetDosenFromNama(t *testing.T) {
	nama := "Indra Riksa"
	data := module.GetDosenFromNama(nama, module.MongoConn, "dosen")
	fmt.Println(data)
}

func TestInsertJamSidang(t *testing.T) {
	durasi := 7
	jam_masuk := "13:00"
	jam_keluar := "14:00"
	gmt := 7
	hari := "Selasa"

	hasil := module.InsertJamSidang(module.MongoConn, "jamsidang", durasi, jam_masuk, jam_keluar, gmt, hari)
	fmt.Println(hasil)
}

func TestGetJamSidangFromDurasi(t *testing.T) {
	durasi := 30
	data := module.GetJamSidangFromDurasi(durasi, module.MongoConn, "jamsidang")
	fmt.Println(data)
}

func TestInsertBap(t *testing.T) {
	judul := "Aplikasi Perpustakaan Berbasis Web"
	nama_mahasiswa := "Salsabila"
	prodi := "D4 Teknik Informatika"
	dosen := "Roni Habibi, S.Kom., M.T., SFPC"
	jam_sidang := "09:00"
    
	hasil := module.InsertBap(module.MongoConn, "bap", judul, nama_mahasiswa, prodi, dosen, jam_sidang)
	fmt.Println(hasil)
}

func TestGetBapFromJudul(t *testing.T) {
	judul := "Aplikasi Perpustakaan Berbasis Web"
	data := module.GetBapFromJudul(judul, module.MongoConn, "bap")
	fmt.Println(data)
}

func TestGetAllBapFromJudul(t *testing.T) {
	judul := "Aplikasi Perpustakaan Berbasis Web"
	data := module.GetAllBapFromJudul(judul, module.MongoConn, "bap")
	fmt.Println(data)
}
