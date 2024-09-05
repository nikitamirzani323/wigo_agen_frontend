package routers

import (
	"bitbucket.org/isbtotogroup/wigo_agen_frontend/controllers"
	"bitbucket.org/isbtotogroup/wigo_agen_frontend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("Content-Security-Policy", "frame-ancestors 'none'")
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/check/healthz", controllers.HealthCheck)
	app.Get("/check/dashboard", monitor.New())

	api := app.Group("/api", middleware.Gateway)
	api.Post("/login", controllers.CheckLogin)
	api.Post("/valid", controllers.Home)
	api.Post("/alladmin", controllers.Adminhome)
	api.Post("/saveadmin", controllers.AdminSave)
	api.Post("/alladminrule", controllers.Adminrulehome)
	api.Post("/saveadminrule", controllers.AdminruleSave)
	api.Post("/transaksi2d30s", controllers.Transaksi2d30shome)
	api.Post("/transaksi2d30ssummarydaily", controllers.Transaksi2d30ssummarydaily)
	api.Post("/transaksi2d30sinfo", controllers.Transaksi2d30sinfo)
	api.Post("/transaksi2d30sprediksi", controllers.Transaksi2d30sprediksi)
	api.Post("/transaksi2d30sdetail", controllers.Transaksi2d30sdetail)
	api.Post("/transaksi2d30ssave", controllers.Transaksi2d30sSave)
	api.Post("/conf", controllers.Agenconf)
	api.Post("/confsave", controllers.AgenconfSave)

	return app
}
