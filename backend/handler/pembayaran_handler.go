package handler

import (
    "github.com/gofiber/fiber/v2"
    "backend/model"
    "backend/repository"
)

func GetAllPembayaran(c *fiber.Ctx) error {
    data, err := repository.GetAllPembayaran(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Gagal mengambil data pembayaran",
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data pembayaran berhasil diambil",
        "data":    data,
        "status":  fiber.StatusOK,
    })
}

func GetPembayaranByID(c *fiber.Ctx) error {
    id := c.Params("id")
    p, err := repository.GetPembayaranByID(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    if p == nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "Data tidak ditemukan",
            "status":  fiber.StatusNotFound,
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data pembayaran ditemukan",
        "data":    p,
        "status":  fiber.StatusOK,
    })
}

func InsertPembayaran(c *fiber.Ctx) error {
    var p model.Pembayaran
    if err := c.BodyParser(&p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    insertedID, err := repository.InsertPembayaran(c.Context(), p)
    if err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Pembayaran berhasil ditambahkan",
        "id":      insertedID,
        "status":  fiber.StatusCreated,
    })
}

func UpdatePembayaran(c *fiber.Ctx) error {
    id := c.Params("id")
    var p model.Pembayaran
    if err := c.BodyParser(&p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    updatedID, err := repository.UpdatePembayaran(c.Context(), id, p)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pembayaran berhasil diupdate",
        "id":      updatedID,
        "status":  fiber.StatusOK,
    })
}

func DeletePembayaran(c *fiber.Ctx) error {
    id := c.Params("id")
    deletedID, err := repository.DeletePembayaran(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pembayaran berhasil dihapus",
        "id":      deletedID,
        "status":  fiber.StatusOK,
    })
}