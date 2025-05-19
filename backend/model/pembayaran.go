package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pembayaran struct {
    ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
    PemesananID       string             `bson:"pemesanan_id" json:"pemesanan_id"`
    NamaPelanggan     string             `bson:"nama_pelanggan" json:"nama_pelanggan"`
    LapanganID        string             `bson:"lapangan_id" json:"lapangan_id"`
    TanggalMain       string             `bson:"tanggal_main" json:"tanggal_main"`
    JumlahBayar       int                `bson:"jumlah_bayar" json:"jumlah_bayar"`
    MetodePembayaran  string             `bson:"metode_pembayaran" json:"metode_pembayaran"`
    TanggalPembayaran string             `bson:"tanggal_pembayaran" json:"tanggal_pembayaran"`
    BuktiPembayaranURL string            `bson:"bukti_pembayaran_url" json:"bukti_pembayaran_url"`
    StatusPembayaran  string             `bson:"status_pembayaran" json:"status_pembayaran"`
    AdminVerifikator  string             `bson:"admin_verifikator,omitempty" json:"admin_verifikator,omitempty"`
    Catatan           string             `bson:"catatan,omitempty" json:"catatan,omitempty"`
}