package module

import (
	"context"
	"fmt"
	"os"

	"time"

	"github.com/MarlinaLubis/be_bap/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "bap_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPresensi(db *mongo.Database, col string, long float64, lat float64, lokasi string, phonenumber int, checkin string, biodata model.Mahasiswa) (InsertedID interface{}) {
	var presensi model.Presensi
	presensi.Latitude = long
	presensi.Longitude = lat
	presensi.Location = lokasi
	presensi.Phone_number = phonenumber
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc(db, col, presensi)
}

func GetPresensiFromCheckin(checkin string, db *mongo.Database, col string) (presensi model.Presensi) {
	data_presensi := db.Collection(col)
	filter := bson.M{"checkin": checkin}
	err := data_presensi.FindOne(context.TODO(), filter).Decode(&presensi)
	if err != nil {
		fmt.Printf("GetPresensiFromCheckin: %v\n", err)
	}
	return presensi
}

func InsertMahasiswa(db *mongo.Database, col string, nama string, phone_number int, email string, jurusan string, jam_sidang string, hari_sidang string) (InsertID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Phone_number = phone_number
	mahasiswa.Email = email
	mahasiswa.Jurusan = jurusan
	mahasiswa.Jam_sidang = jam_sidang
	mahasiswa.Hari_sidang = hari_sidang
	return InsertOneDoc(db, col, mahasiswa)
}

func GetMahasiswaFromNama(nama string, db *mongo.Database, col string) (mahasiswa model.Mahasiswa) {
	data_mahasiswa := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := data_mahasiswa.FindOne(context.TODO(), filter).Decode(&mahasiswa)
	if err != nil {
		fmt.Printf("GetMahasiswaFromNama: %v\n", err)
	}
	return mahasiswa
}

func InsertDosen(db *mongo.Database, col string, nama string, phone_number int, email string, jam_sidang string, hari_sidang string) (InsertID interface{}) {
	var dosen model.Dosen
	dosen.Nama = nama
	dosen.Phone_number = phone_number
	dosen.Email = email
	dosen.Jam_sidang = jam_sidang
	dosen.Hari_sidang = hari_sidang
	return InsertOneDoc(db, col, dosen)
}

func GetDosenFromNama(nama string, db *mongo.Database, col string) (dosen model.Dosen) {
	data_dosen := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := data_dosen.FindOne(context.TODO(), filter).Decode(&dosen)
	if err != nil {
		fmt.Printf("GetDosenFromNama: %v\n", err)
	}
	return dosen
}

func InsertJamSidang(db *mongo.Database, col string, durasi int, jam_masuk string, jam_keluar string, gmt int, hari string) (InsertedID interface{}) {
	var jamsidang model.JamSidang
	jamsidang.Durasi = durasi
	jamsidang.Jam_masuk = jam_masuk
	jamsidang.Jam_keluar = jam_keluar
	jamsidang.Gmt = gmt
	jamsidang.Hari = hari
	return InsertOneDoc(db, col, jamsidang)
}

func GetJamSidangFromDurasi(durasi int, db *mongo.Database, col string) (jamsidang model.JamSidang) {
	data_jamsidang := db.Collection(col)
	filter := bson.M{"durasi": durasi}
	err := data_jamsidang.FindOne(context.TODO(), filter).Decode(&jamsidang)
	if err != nil {
		fmt.Printf("GetJamsidangFromHari: %v\n", err)
	}
	return jamsidang
}

func InsertBap(db *mongo.Database, col string, judul string, nama_mahasiswa string, prodi string, hasil_revisi string, dosen string) (InsertedID interface{}) {
	var bap model.Bap
	bap.Judul = judul
	bap.Nama_Mahasiswa = nama_mahasiswa
	bap.Prodi = prodi
	bap.Hasil_revisi = hasil_revisi
	bap.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	bap.Dosen = dosen
	return InsertOneDoc(db, col, bap)
}

func GetBapFromJudul(judul string, db *mongo.Database, col string) (bap model.Bap) {
	data_bap := db.Collection(col)
	filter := bson.M{"judul": judul}
	err := data_bap.FindOne(context.TODO(), filter).Decode(&bap)
	if err != nil {
		fmt.Printf("GetBapFromJudul: %v\n", err)
	}
	return bap
}

func GetAllBapFromJudul(judul string, db *mongo.Database, col string) (bap []model.Bap) {
	data_bap := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_bap.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &bap)
	if err != nil {
		fmt.Println(err)
	}
	return bap
}

func GetAllBap(db *mongo.Database, col string) (data []model.Bap) {
	data_bap := db.Collection(col)
	filter := bson.M{}
	cursor, err := data_bap.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}


