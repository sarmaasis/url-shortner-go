package routes

import(
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type request struct {
	URL				string				`json:"url"`
	CustomShort		string				`json:"short"`
	Expiry			time.Duration		`json:"expiry"`
}

type response struct {
	URL					string			`json:"url"`
	CustomShort			string			`json:"short"`
	Expiry				time.Duration	`json:"expiry"`
	XRateRemaining		int				`json:"rate_limit"`
	XRateLimitReset		time.Duration	`json:"rate_limit_reset"`
}

func ShortenURL(c * fiber.Ctx) error{

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	/*
		Rate limiting
	*/



	/*
		Check if the input is valid URL
	*/
	
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}


	/*
		Check for domain error
	*/

	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": ""})
	}


	/*
		enforce https, ssl

	*/

	body.URL = helpers.EnforceHTTP(bodu.URL)
	
}