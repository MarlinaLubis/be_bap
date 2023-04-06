package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number int                `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Jam_sidang   string             `bson:"jam_sidang,omitempty" json:"jam_sidang,omitempty"`
	Hari_sidang  string             `bson:"hari_sidang,omitempty" json:"hari_sidang,omitempty"`
}

type Dosen struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number int                `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Jam_sidang   string             `bson:"jam_sidang,omitempty" json:"jam_sidang,omitempty"`
	Hari_sidang  string             `bson:"hari_sidang,omitempty" json:"hari_sidang,omitempty"`
}

type JamSidang struct {
	Durasi     int    `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int    `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       string `bson:"hari,omitempty" json:"hari,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number int                `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Bap struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Judul          string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Nama_Mahasiswa string             `bson:"nama_mahasiswa,omitempty" json:"nama_mahasiswa,omitempty"`
	Prodi          string             `bson:"prodi,omitempty" json:"prodi,omitempty"`
	Hasil_revisi   string             `bson:"hasil_revisi,omitempty" json:"hasil_revisi,omitempty"`
	Datetime       primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Dosen          string             `bson:"dosen,omitempty" json:"dosen,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty" json:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
