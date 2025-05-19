package router

import (
    "github.com/gofiber/fiber/v2"
    "backend/handler"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    api.Get("/", handler.Homepage)
    api.Get("/lapangan", handler.GetAllLapangan)
    api.Get("/lapangan/:id", handler.GetLapanganByID)
    api.Post("/lapangan", handler.InsertLapangan)
    api.Put("/lapangan/:id", handler.UpdateLapangan)
    api.Delete("/lapangan/:id", handler.DeleteLapangan)

    api.Get("/jadwal", handler.GetJadwalByLapanganTanggal)
    api.Post("/jadwal", handler.InsertJadwal)
    api.Put("/jadwal/:id", handler.UpdateJadwal)
    api.Delete("/jadwal/:id", handler.DeleteJadwal)

    api.Get("/pemesanan", handler.GetAllPemesanan)
    api.Get("/pemesanan/:id", handler.GetPemesananByID)
    api.Post("/pemesanan", handler.InsertPemesanan)
    api.Put("/pemesanan/:id", handler.UpdatePemesanan)
    api.Delete("/pemesanan/:id", handler.DeletePemesanan)

    api.Get("/pembayaran", handler.GetAllPembayaran)
    api.Get("/pembayaran/:id", handler.GetPembayaranByID)
    api.Post("/pembayaran", handler.InsertPembayaran)
    api.Put("/pembayaran/:id", handler.UpdatePembayaran)
    api.Delete("/pembayaran/:id", handler.DeletePembayaran)
}