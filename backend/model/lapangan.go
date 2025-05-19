package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type JenisOlahragaDetail struct {
    Nama             string `bson:"nama" json:"nama"`
    KodeJenis        string `bson:"kode_jenis" json:"kode_jenis"`
    DeskripsiSingkat string `bson:"deskripsi_singkat" json:"deskripsi_singkat"`
}

type JamOperasional struct {
    SeninJumat  string `bson:"senin_jumat" json:"senin_jumat"`
    SabtuMinggu string `bson:"sabtu_minggu" json:"sabtu_minggu"`
}

type Lapangan struct {
    ID                  primitive.ObjectID   `bson:"_id,omitempty" json:"_id"`
    NamaLapangan        string               `bson:"nama_lapangan" json:"nama_lapangan"`
    JenisOlahraga       string               `bson:"jenis_olahraga" json:"jenis_olahraga"`
    JenisOlahragaDetail JenisOlahragaDetail  `bson:"jenis_olahraga_detail" json:"jenis_olahraga_detail"`
    DeskripsiLapangan   string               `bson:"deskripsi_lapangan" json:"deskripsi_lapangan"`
    HargaSewaPerJam     int                  `bson:"harga_sewa_per_jam" json:"harga_sewa_per_jam"`
    FotoLapanganUrls    []string             `bson:"foto_lapangan_urls" json:"foto_lapangan_urls"`
    StatusKetersediaan  string               `bson:"status_ketersediaan" json:"status_ketersediaan"`
    Fasilitas           []string             `bson:"fasilitas" json:"fasilitas"`
    JamOperasional      JamOperasional       `bson:"jam_operasional" json:"jam_operasional"`
}