package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pemesanan struct {
    ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
    NamaPelanggan    string             `bson:"nama_pelanggan" json:"nama_pelanggan"`
    KontakPelanggan  string             `bson:"kontak_pelanggan" json:"kontak_pelanggan"`
    LapanganID       string             `bson:"lapangan_id" json:"lapangan_id"`
    TanggalMain      string             `bson:"tanggal_main" json:"tanggal_main"`
    JamMulai         string             `bson:"jam_mulai" json:"jam_mulai"`
    JamSelesai       string             `bson:"jam_selesai" json:"jam_selesai"`
    DurasiJam        int                `bson:"durasi_jam" json:"durasi_jam"`
    TotalBiaya       int                `bson:"total_biaya" json:"total_biaya"`
    StatusPemesanan  string             `bson:"status_pemesanan" json:"status_pemesanan"`
    StatusPembayaran string             `bson:"status_pembayaran" json:"status_pembayaran"`
    CatatanAdmin     string             `bson:"catatan_admin,omitempty" json:"catatan_admin,omitempty"`
    AdminPencatat    string             `bson:"admin_pencatat,omitempty" json:"admin_pencatat,omitempty"`
}