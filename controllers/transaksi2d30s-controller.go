package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_transaksi2d30s struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	Record         interface{} `json:"record"`
	Periode        string      `json:"periode"`
	Perpage        int         `json:"perpage"`
	Totalrecord    int         `json:"totalrecord"`
	Totalbet       int         `json:"totalbet"`
	Totalwin       int         `json:"totalwin"`
	Winlose_agen   int         `json:"winlose_agen"`
	Winlose_member int         `json:"winlose_member"`
	Time           string      `json:"time"`
}
type response_transaksi2d30sinfo struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Summary interface{} `json:"summary"`
	Time    string      `json:"time"`
}
type response_transaksi2d30sprediksi struct {
	Status         int         `json:"status"`
	Message        string      `json:"message"`
	Record         interface{} `json:"record"`
	Totalmember    int         `json:"totalmember"`
	Totalbet       int         `json:"totalbet"`
	Totalwin       int         `json:"totalwin"`
	Winlose_agen   int         `json:"winlose_agen"`
	Winlose_member int         `json:"winlose_member"`
	Time           string      `json:"time"`
}

func Transaksi2d30shome(c *fiber.Ctx) error {
	type payload_transaksi2d30shome struct {
		Transaksi2D30s_search  string `json:"transaksi2D30s_search"`
		Transaksi2D30s_page    int    `json:"transaksi2D30s_page"`
		Transaksi2D30s_invoice string `json:"transaksi2D30s_invoice"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30shome)
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
		SetResult(response_transaksi2d30s{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":        hostname,
			"transaksi2D30s_search":  client.Transaksi2D30s_search,
			"transaksi2D30s_page":    client.Transaksi2D30s_page,
			"transaksi2D30s_invoice": client.Transaksi2D30s_invoice,
		}).
		Post(PATH + "api/transaksi2d30s")
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
	result := resp.Result().(*response_transaksi2d30s)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":         result.Status,
			"message":        result.Message,
			"record":         result.Record,
			"perpage":        result.Perpage,
			"totalrecord":    result.Totalrecord,
			"totalbet":       result.Totalbet,
			"totalwin":       result.Totalwin,
			"winlose_agen":   result.Winlose_agen,
			"winlose_member": result.Winlose_member,
			"periode":        result.Periode,
			"time":           time.Since(render_page).String(),
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
func Transaksi2d30ssummarydaily(c *fiber.Ctx) error {
	type payload_transaksi2d30ssummarydaily struct {
		Transaksi2D30ssummarydaily_search string `json:"transaksi2D30ssummarydaily_search"`
		Transaksi2D30ssummarydaily_page   int    `json:"transaksi2D30ssummarydaily_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30ssummarydaily)
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
		SetResult(response_transaksi2d30s{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":                   hostname,
			"transaksi2D30ssummarydaily_search": client.Transaksi2D30ssummarydaily_search,
			"transaksi2D30ssummarydaily_page":   client.Transaksi2D30ssummarydaily_page,
		}).
		Post(PATH + "api/transaksi2d30ssummarydaily")
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
	result := resp.Result().(*response_transaksi2d30s)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":         result.Status,
			"message":        result.Message,
			"record":         result.Record,
			"perpage":        result.Perpage,
			"totalrecord":    result.Totalrecord,
			"totalbet":       result.Totalbet,
			"totalwin":       result.Totalwin,
			"winlose_agen":   result.Winlose_agen,
			"winlose_member": result.Winlose_member,
			"periode":        result.Periode,
			"time":           time.Since(render_page).String(),
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
func Transaksi2d30sinfo(c *fiber.Ctx) error {
	type payload_transaksi2d30sinfo struct {
		Transaksi2D30s_invoice string `json:"transaksi2D30s_invoice"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30sinfo)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_transaksi2d30sinfo{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":        hostname,
			"transaksi2D30s_invoice": client.Transaksi2D30s_invoice,
		}).
		Post(PATH + "api/transaksi2d30sinfo")
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
	result := resp.Result().(*response_transaksi2d30sinfo)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"summary": result.Summary,
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
func Transaksi2d30sprediksi(c *fiber.Ctx) error {
	type payload_transaksi2d30sprediksi struct {
		Transaksi2D30sprediksi_invoice string `json:"transaksi2D30sprediksi_invoice"`
		Transaksi2D30sprediksi_result  string `json:"transaksi2D30sprediksi_result"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30sprediksi)
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
		SetResult(response_transaksi2d30sprediksi{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":                hostname,
			"transaksi2D30sprediksi_invoice": client.Transaksi2D30sprediksi_invoice,
			"transaksi2D30sprediksi_result":  client.Transaksi2D30sprediksi_result,
		}).
		Post(PATH + "api/transaksi2d30sprediksi")
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
	result := resp.Result().(*response_transaksi2d30sprediksi)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":         result.Status,
			"message":        result.Message,
			"record":         result.Record,
			"totalmember":    result.Totalmember,
			"totalbet":       result.Totalbet,
			"totalwin":       result.Totalwin,
			"winlose_agen":   result.Winlose_agen,
			"winlose_member": result.Winlose_member,
			"time":           time.Since(render_page).String(),
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
func Transaksi2d30sdetail(c *fiber.Ctx) error {
	type payload_transaksi2d30sdetail struct {
		Transaksidetail2D30s_invoice string `json:"transaksidetail2D30s_invoice"`
		Transaksidetail2D30s_status  string `json:"transaksidetail2D30s_status"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30sdetail)
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
		SetResult(response_transaksi2d30s{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":              hostname,
			"transaksidetail2D30s_invoice": client.Transaksidetail2D30s_invoice,
			"transaksidetail2D30s_status":  client.Transaksidetail2D30s_status,
		}).
		Post(PATH + "api/transaksi2d30sdetail")
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
	result := resp.Result().(*response_transaksi2d30s)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":         result.Status,
			"message":        result.Message,
			"record":         result.Record,
			"perpage":        result.Perpage,
			"totalrecord":    result.Totalrecord,
			"totalbet":       result.Totalbet,
			"totalwin":       result.Totalwin,
			"winlose_agen":   result.Winlose_agen,
			"winlose_member": result.Winlose_member,
			"periode":        result.Periode,
			"time":           time.Since(render_page).String(),
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
func Agenconf(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_transaksi2d30sprediksi{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
		}).
		Post(PATH + "api/conf")
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
	result := resp.Result().(*response_transaksi2d30sprediksi)
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
func Transaksi2d30sSave(c *fiber.Ctx) error {
	type payload_transaksi2d30ssave struct {
		Page                   string `json:"page"`
		Transaksi2D30s_invoice string `json:"transaksi2D30s_invoice" `
		Transaksi2D30s_result  string `json:"transaksi2D30s_result" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30ssave)
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
			"client_hostname":        hostname,
			"page":                   client.Page,
			"transaksi2D30s_invoice": client.Transaksi2D30s_invoice,
			"transaksi2D30s_result":  client.Transaksi2D30s_result,
		}).
		Post(PATH + "api/transaksi2d30ssave")
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
func AgenconfSave(c *fiber.Ctx) error {
	type payload_transaksi2d30ssave struct {
		Page                        string `json:"page"`
		Agenconf_2digit_30_operator string `json:"agenconf_2digit_30_operator" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transaksi2d30ssave)
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
			"client_hostname":             hostname,
			"page":                        client.Page,
			"agenconf_2digit_30_operator": client.Agenconf_2digit_30_operator,
		}).
		Post(PATH + "api/confsave")
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
