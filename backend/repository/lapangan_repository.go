package repository

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "backend/config"
    "backend/model"
)

func InsertLapangan(ctx context.Context, lap model.Lapangan) (interface{}, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.LapanganCollection)
    result, err := collection.InsertOne(ctx, lap)
    if err != nil {
        fmt.Printf("InsertLapangan: %v\n", err)
        return nil, err
    }
    return result.InsertedID, nil
}

func GetLapanganByID(ctx context.Context, id string) (*model.Lapangan, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.LapanganCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var lap model.Lapangan
    err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&lap)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &lap, nil
}

func GetAllLapangan(ctx context.Context) ([]model.Lapangan, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.LapanganCollection)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    var data []model.Lapangan
    if err := cursor.All(ctx, &data); err != nil {
        return nil, err
    }
    return data, nil
}

func UpdateLapangan(ctx context.Context, id string, lap model.Lapangan) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.LapanganCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return "", err
    }
    update := bson.M{"$set": lap}
    result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    if err != nil {
        return "", err
    }
    if result.ModifiedCount == 0 {
        return "", fmt.Errorf("tidak ada data yang diupdate untuk id %v", id)
    }
    return id, nil
}

func DeleteLapangan(ctx context.Context, id string) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.LapanganCollection)
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