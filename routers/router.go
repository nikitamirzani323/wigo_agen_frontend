package routers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/WIGO_AGEN_FRONTEND/controllers"
	"github.com/valyala/fasthttp"
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
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", controllers.Home)
	app.Post("/api/alladmin", controllers.Adminhome)
	app.Post("/api/saveadmin", controllers.AdminSave)
	app.Post("/api/alladminrule", controllers.Adminrulehome)
	app.Post("/api/saveadminrule", controllers.AdminruleSave)
	app.Post("/api/transaksi2d30s", controllers.Transaksi2d30shome)
	app.Post("/api/transaksi2d30sinfo", controllers.Transaksi2d30sinfo)
	app.Post("/api/transaksi2d30sprediksi", controllers.Transaksi2d30sprediksi)
	app.Post("/api/transaksi2d30sdetail", controllers.Transaksi2d30sdetail)
	app.Post("/api/transaksi2d30ssave", controllers.Transaksi2d30sSave)
	app.Post("/api/conf", controllers.Agenconf)
	app.Post("/api/confsave", controllers.AgenconfSave)

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")

			dbHost := os.Getenv("DB_REDIS_HOST") + ":" + os.Getenv("DB_REDIS_PORT")
			dbPass := os.Getenv("DB_REDIS_PASSWORD")
			dbName, _ := strconv.Atoi(os.Getenv("DB_REDIS_NAME"))

			rdb := redis.NewClient(&redis.Options{
				Addr:     dbHost,
				Password: dbPass,
				DB:       dbName,
			})

			resultredis := rdb.Subscribe("", "payload_agen_nuke")

			for {
				msg, err := resultredis.ReceiveMessage()
				if err != nil {
					panic(err)
				}

				// fmt.Println("Received message from " + msg.Payload + " channel.")
				// data_pubsub := strings.Split(msg.Payload, ":")

				msg_sse := msg.Payload

				fmt.Fprintf(w, "data: %s\n\n", msg_sse)
				// fmt.Println(msg_sse)
				err_sse := w.Flush()
				if err_sse != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
				// time.Sleep(1 * time.Second)

			}

		}))

		return nil
	})
	return app
}
