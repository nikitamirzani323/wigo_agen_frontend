package controllers

import (
	"fmt"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/wigo_agen_frontend/configs"
	"bitbucket.org/isbtotogroup/wigo_agen_frontend/entities"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type responseinit struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}
type responselivescore struct {
	Record interface{} `json:"record"`
}
type responsedefault struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var PATH string = configs.GetPathAPI()

func CheckLogin(c *fiber.Ctx) error {
	hostname := c.Hostname()
	client := new(entities.Login)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseinit{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"username":        client.Username,
			"password":        client.Password,
			"ipaddress":       client.Ipaddress,
			"timezone":        client.Timezone,
		}).
		Post(PATH + "api/login")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*responseinit)
	return c.JSON(fiber.Map{
		"status": result.Status,
		"token":  result.Token,
		"time":   time.Since(render_page).String(),
	})
}
func Home(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	fmt.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/valid")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
