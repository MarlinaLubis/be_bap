package NPM

import (
	"fmt"
	"testing"

	model "github.com/MarlinaLubis/be_bap/model"
	module "github.com/MarlinaLubis/be_bap/module"
)

func TestInsertPresensi(t *testing.T) {
	long := 114.590111
	lat :=  -3.316694
	lokasi := "	Kalimantan "
	phonenumber := 6281221089054
	checkin := "Masuk"
	biodata := model.Mahasiswa{    
		Nama:         "Cahaya Satriani",
		Phone_number: 6281221089054,
		Email:        "cahaya@gmail.com",
		Jurusan:      "D4 Teknik Informatika",
		Jam_sidang:   "09.00",
		Hari_sidang: "Rabu", 
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
	nama := "Cahaya Sastriani"
	phone_number :=  6281221089054
	email := "cahaya@gmail.com"
	jurusan := "D4 Teknik Informatika"
	jam_sidang := "09:00"
    harisidang := "Rabu"

	hasil := module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, phone_number, email, jurusan, jam_sidang, harisidang)
	fmt.Println(hasil)
}
func TestGetMahasiswaFromNama(t *testing.T) {
	nama := "Cahaya Sastriani"
	biodata := module.GetMahasiswaFromNama(nama, module.MongoConn, "mahasiswa")
	fmt.Println(biodata)
}

func TestInsertDosen(t *testing.T) {
	nama := "Woro Isti Rahayu, S.T., M.T., SFPC"
	phone_number := 6281267885432
	email := "woroisti@ulbi.ac.id"
	jam_sidang := "09.00"
	hari_sidang := "Rabu"

	hasil := module.InsertDosen(module.MongoConn, "dosen", nama, phone_number, email, jam_sidang, hari_sidang)
	fmt.Println(hasil)
}

func TestGetDosenFromNama(t *testing.T) {
	nama := "Woro Isti Rahayu, S.T., M.T., SFPC"
	data := module.GetDosenFromNama(nama, module.MongoConn, "dosen")
	fmt.Println(data)
}

func TestInsertJamSidang(t *testing.T) {
	durasi := 60
	jam_masuk := "09:00"
	jam_keluar := "10:00"
	gmt := 7
	hari := "Rabu"

	hasil := module.InsertJamSidang(module.MongoConn, "jamsidang", durasi, jam_masuk, jam_keluar, gmt, hari)
	fmt.Println(hasil)
}

func TestGetJamSidangFromDurasi(t *testing.T) {
	durasi := 60
	data := module.GetJamSidangFromDurasi(durasi, module.MongoConn, "jamsidang")
	fmt.Println(data)
}

func TestInsertBap(t *testing.T) {
	judul := "Aplikasi Penjualan Berbasis Web"
	nama_mahasiswa := "Cahaya Sastriani"
	prodi := "D4 Teknik Informatika"
	hasil_revisi := "Memperbaiki bagian usecase"
	dosen := "Woro Isti Rahayu, S.T., M.T., SFPC"

	hasil := module.InsertBap(module.MongoConn, "bap", judul, nama_mahasiswa, prodi, hasil_revisi, dosen )
	fmt.Println(hasil)
}

func TestGetBapFromJudul(t *testing.T) {
	judul := "Aplikasi Penjualan Berbasis Web"
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