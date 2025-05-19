package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Jadwal struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
    LapanganID   string             `bson:"lapangan_id" json:"lapangan_id"`
    Tanggal      string             `bson:"tanggal" json:"tanggal"`
    JamMulai     string             `bson:"jam_mulai" json:"jam_mulai"`
    JamSelesai   string             `bson:"jam_selesai" json:"jam_selesai"`
    StatusSlot   string             `bson:"status_slot" json:"status_slot"`
    PemesananID  string             `bson:"pemesanan_id,omitempty" json:"pemesanan_id,omitempty"`
    NamaPelanggan string            `bson:"nama_pelanggan,omitempty" json:"nama_pelanggan,omitempty"`
}