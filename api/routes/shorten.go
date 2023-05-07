package routes

import(
	"time"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sarmaasis/url-shortner-go/database"
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

	r2 := database.CreateClient(1)
	defer r2.Close()

	value, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil{
		_ = r2.Set(database.Ctx, c.IP, os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		value, _ := r2.Get(database.Ctx, c.IP()).Result()
		valueInt, _ = strconv.Atoi(value)

		if valueInt < = 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "Rate limit exceeded",
				"rate_limit_reset": limit/time.Nanosecond/time.Minute,
			})
		}
	}


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

	r2.Decr(database.Ctx, c.IP())
	
}