package handler

import (
    "github.com/gofiber/fiber/v2"
    "backend/model"
    "backend/repository"
)

func GetAllPemesanan(c *fiber.Ctx) error {
    data, err := repository.GetAllPemesanan(c.Context())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Gagal mengambil data pemesanan",
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Data pemesanan berhasil diambil",
        "data":    data,
        "status":  fiber.StatusOK,
    })
}

func GetPemesananByID(c *fiber.Ctx) error {
    id := c.Params("id")
    p, err := repository.GetPemesananByID(c.Context(), id)
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
        "message": "Data pemesanan ditemukan",
        "data":    p,
        "status":  fiber.StatusOK,
    })
}

func InsertPemesanan(c *fiber.Ctx) error {
    var p model.Pemesanan
    if err := c.BodyParser(&p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    insertedID, err := repository.InsertPemesanan(c.Context(), p)
    if err != nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Pemesanan berhasil ditambahkan",
        "id":      insertedID,
        "status":  fiber.StatusCreated,
    })
}

func UpdatePemesanan(c *fiber.Ctx) error {
    id := c.Params("id")
    var p model.Pemesanan
    if err := c.BodyParser(&p); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request data",
        })
    }
    updatedID, err := repository.UpdatePemesanan(c.Context(), id, p)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pemesanan berhasil diupdate",
        "id":      updatedID,
        "status":  fiber.StatusOK,
    })
}

func DeletePemesanan(c *fiber.Ctx) error {
    id := c.Params("id")
    deletedID, err := repository.DeletePemesanan(c.Context(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Pemesanan berhasil dihapus",
        "id":      deletedID,
        "status":  fiber.StatusOK,
    })
}