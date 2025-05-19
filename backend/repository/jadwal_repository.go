package repository

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "backend/config"
    "backend/model"
)

func GetJadwalByLapanganTanggal(ctx context.Context, lapanganID, tanggal string) ([]model.Jadwal, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.JadwalCollection)
    filter := bson.M{}
    if lapanganID != "" {
        filter["lapangan_id"] = lapanganID
    }
    if tanggal != "" {
        filter["tanggal"] = tanggal
    }
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    var data []model.Jadwal
    if err := cursor.All(ctx, &data); err != nil {
        return nil, err
    }
    return data, nil
}

func InsertJadwal(ctx context.Context, jadwal model.Jadwal) (interface{}, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.JadwalCollection)
    result, err := collection.InsertOne(ctx, jadwal)
    if err != nil {
        fmt.Printf("InsertJadwal: %v\n", err)
        return nil, err
    }
    return result.InsertedID, nil
}

func UpdateJadwal(ctx context.Context, id string, jadwal model.Jadwal) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.JadwalCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return "", err
    }
    update := bson.M{"$set": jadwal}
    result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    if err != nil {
        return "", err
    }
    if result.ModifiedCount == 0 {
        return "", fmt.Errorf("tidak ada data yang diupdate untuk id %v", id)
    }
    return id, nil
}

func DeleteJadwal(ctx context.Context, id string) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.JadwalCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return "", err
    }
    result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        return "", err
    }
    if result.DeletedCount == 0 {
        return "", fmt.Errorf("tidak ada data yang dihapus untuk id %v", id)
    }
    return id, nil
}