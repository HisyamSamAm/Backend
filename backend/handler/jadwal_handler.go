package handler

import (
    "github.com/gofiber/fiber/v2"
    "backend/model"
    "backend/repository"
)

func GetJadwalByLapanganTanggal(c *fiber.Ctx) error {
    lapanganID := c.Query("lapangan_id")
    tanggal := c.Query("tanggal")
    data, err := repository.GetJadwalByLapanganTanggal(c.Context(), lapanganID, tanggal)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Gagal mengambil data jadwal",
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data jadwal berhasil diambil",
        "data":    data,
        "status":  fiber.StatusOK,
    })
}

func InsertJadwal(c *fiber.Ctx) error {
    var jadwal model.Jadwal
    if err := c.BodyParser(&jadwal); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    insertedID, err := repository.InsertJadwal(c.Context(), jadwal)
    if err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Jadwal berhasil ditambahkan",
        "id":      insertedID,
        "status":  fiber.StatusCreated,
    })
}

func UpdateJadwal(c *fiber.Ctx) error {
    id := c.Params("id")
    var jadwal model.Jadwal
    if err := c.BodyParser(&jadwal); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    updatedID, err := repository.UpdateJadwal(c.Context(), id, jadwal)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Jadwal berhasil diupdate",
        "id":      updatedID,
        "status":  fiber.StatusOK,
    })
}

func DeleteJadwal(c *fiber.Ctx) error {
    id := c.Params("id")
    deletedID, err := repository.DeleteJadwal(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Jadwal berhasil dihapus",
        "id":      deletedID,
        "status":  fiber.StatusOK,
    })
}