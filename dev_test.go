package NPM

import (
	"fmt"
	"testing"

	model "github.com/MarlinaLubis/be_bap/model"
	module "github.com/MarlinaLubis/be_bap/module"
)

func TestInsertPresensi(t *testing.T) {
	long := 110.828316
	lat :=   -7.550676
	lokasi := "	Solo"
	phonenumber := 6281262487651
	checkin := "Masuk"
	biodata := model.Mahasiswa{    
		Nama:         "Novi Febriani",
		Phone_number:   6281262487651,
		Email:        "novi@gmail.com",
		Jurusan:      "D4 Teknik Informatika",
		Jam_sidang:   "10.00",
		Hari_sidang: "Senin", 
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
	nama := "Novi Febriani"
	phone_number := 6281262487651
	email := "novifebri@gmail.com"
	jurusan := "D4 Teknik Informatika"
	jam_sidang := "10:00"
    harisidang := "Senin"

	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, phone_number, email, jurusan, jam_sidang, harisidang)
	fmt.Println(hasil)
}
func TestGetMahasiswaFromNama(t *testing.T) {
	nama := "Novi Febriani"
	biodata := module.GetMahasiswaFromNama(nama, module.MongoConn, "mahasiswa")
	fmt.Println(biodata)
}

func TestInsertDosen(t *testing.T) {
	nama := " Roni Habibi, S.Kom., M.T., SFPC"
	phone_number := 6281522176359
	email := "roni.habibi@ulbi.ac.id"
	jam_sidang := "10.00"
	hari_sidang := "Senin"

	hasil := module.InsertDosen(module.MongoConn, "dosen", nama, phone_number, email, jam_sidang, hari_sidang)
	fmt.Println(hasil)
}

func TestGetDosenFromNama(t *testing.T) {
	nama := " Roni Habibi, S.Kom., M.T., SFPC"
	data := module.GetDosenFromNama(nama, module.MongoConn, "dosen")
	fmt.Println(data)
}

func TestInsertJamSidang(t *testing.T) {
	durasi := 60
	jam_masuk := "10:00"
	jam_keluar := "11:00"
	gmt := 7
	hari := "Senin"

	hasil := module.InsertJamSidang(module.MongoConn, "jamsidang", durasi, jam_masuk, jam_keluar, gmt, hari)
	fmt.Println(hasil)
}

func TestGetJamSidangFromDurasi(t *testing.T) {
	durasi := 60
	data := module.GetJamSidangFromDurasi(durasi, module.MongoConn, "jamsidang")
	fmt.Println(data)
}

func TestInsertBap(t *testing.T) {
	judul := "Aplikasi Raport Berbasis Web"
	nama_mahasiswa := "Novi Febriani"
	prodi := "D4 Teknik Informatika"
	hasil_revisi := "Memperbaiki database dan dashboard"
	dosen := " Roni Habibi, S.Kom., M.T., SFPC"

	hasil := module.InsertBap(module.MongoConn, "bap", judul, nama_mahasiswa, prodi, hasil_revisi, dosen )
	fmt.Println(hasil)
}

func TestGetBapFromJudul(t *testing.T) {
	judul := "Aplikasi Raport Berbasis Web"
	data := module.GetBapFromJudul(judul, module.MongoConn, "bap")
	fmt.Println(data)
}

// func TestGetAllBapFromJudul(t *testing.T) {
// 	judul := "Aplikasi Rental Mobil Berbasis Web"
// 	data := module.GetAllBapFromJudul(judul, module.MongoConn, "bap")
// 	fmt.Println(data)
// }

// func TestGetAllPresensi(t *testing.T) {
// 	data := module.GetAllPresensi(module.MongoConn, "presensi")
// 	fmt.Println(data)
// }


func TestGetAll(t *testing.T){
	data := module.GetAllBap(module.MongoConn, "bap")
	fmt.Println(data)
}