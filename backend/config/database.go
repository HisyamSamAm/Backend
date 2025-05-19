package config

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "os"
)

var DBName = "db_booking_lapangan"
var LapanganCollection = "lapangan"
var JadwalCollection = "jadwal"
var PemesananCollection = "pemesanan"
var PembayaranCollection = "pembayaran"
var MongoString = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) *mongo.Database {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
    if err != nil {
        fmt.Printf("MongoConnect: %v\n", err)
    }
    return client.Database(dbname)
}