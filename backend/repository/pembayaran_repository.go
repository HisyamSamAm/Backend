package repository

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "backend/config"
    "backend/model"
)

func GetAllPembayaran(ctx context.Context) ([]model.Pembayaran, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PembayaranCollection)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    var data []model.Pembayaran
    if err := cursor.All(ctx, &data); err != nil {
        return nil, err
    }
    return data, nil
}

func GetPembayaranByID(ctx context.Context, id string) (*model.Pembayaran, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PembayaranCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var pembayaran model.Pembayaran
    err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&pembayaran)
    if err != nil {
        return nil, err
    }
    return &pembayaran, nil
}

func InsertPembayaran(ctx context.Context, p model.Pembayaran) (interface{}, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PembayaranCollection)
    result, err := collection.InsertOne(ctx, p)
    if err != nil {
        fmt.Printf("InsertPembayaran: %v\n", err)
        return nil, err
    }
    return result.InsertedID, nil
}

func UpdatePembayaran(ctx context.Context, id string, p model.Pembayaran) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PembayaranCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return "", err
    }
    update := bson.M{"$set": p}
    result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    if err != nil {
        return "", err
    }
    if result.ModifiedCount == 0 {
        return "", fmt.Errorf("tidak ada data yang diupdate untuk id %v", id)
    }
    return id, nil
}

func DeletePembayaran(ctx context.Context, id string) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PembayaranCollection)
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