package handler

import (
    "github.com/gofiber/fiber/v2"
    "backend/model"
    "backend/repository"
)

func Homepage(c *fiber.Ctx) error {
    return c.SendString("API Lapangan Fiber sudah jalan")
}

func GetAllLapangan(c *fiber.Ctx) error {
    data, err := repository.GetAllLapangan(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Gagal mengambil data dari database",
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data berhasil diambil",
        "data":    data,
        "status":  fiber.StatusOK,
    })
}

func GetLapanganByID(c *fiber.Ctx) error {
    id := c.Params("id")
    lap, err := repository.GetLapanganByID(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    if lap == nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Data tidak ditemukan",
            "status":  fiber.StatusNotFound,
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data lapangan ditemukan",
        "data":    lap,
        "status":  fiber.StatusOK,
    })
}

func InsertLapangan(c *fiber.Ctx) error {
    var lap model.Lapangan
    if err := c.BodyParser(&lap); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    insertedID, err := repository.InsertLapangan(c.Context(), lap)
    if err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Lapangan berhasil ditambahkan",
        "id":      insertedID,
        "status":  fiber.StatusCreated,
    })
}

func UpdateLapangan(c *fiber.Ctx) error {
    id := c.Params("id")
    var lap model.Lapangan
    if err := c.BodyParser(&lap); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    updatedID, err := repository.UpdateLapangan(c.Context(), id, lap)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data lapangan berhasil diupdate",
        "id":      updatedID,
        "status":  fiber.StatusOK,
    })
}

func DeleteLapangan(c *fiber.Ctx) error {
    id := c.Params("id")
    deletedID, err := repository.DeleteLapangan(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Lapangan berhasil dihapus",
        "id":      deletedID,
        "status":  fiber.StatusOK,
    })
}