package repository

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "backend/config"
    "backend/model"
)

func GetAllPemesanan(ctx context.Context) ([]model.Pemesanan, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    var data []model.Pemesanan
    if err := cursor.All(ctx, &data); err != nil {
        return nil, err
    }
    return data, nil
}

func GetPemesananByID(ctx context.Context, id string) (*model.Pemesanan, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    var pemesanan model.Pemesanan
    err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&pemesanan)
    if err != nil {
		return nil, err
    }
    return &pemesanan, nil
}

func InsertPemesanan(ctx context.Context, p model.Pemesanan) (interface{}, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
    result, err := collection.InsertOne(ctx, p)
    if err != nil {
        fmt.Printf("InsertPemesanan: %v\n", err)
        return nil, err
    }
    return result.InsertedID, nil
}

func UpdatePemesanan(ctx context.Context, id string, p model.Pemesanan) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
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

func DeletePemesanan(ctx context.Context, id string) (string, error) {
    collection := config.MongoConnect(config.DBName).Collection(config.PemesananCollection)
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