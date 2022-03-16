package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/m-yosefpor/httpmon/internal/http/request"
	"github.com/m-yosefpor/httpmon/internal/model"
	"github.com/m-yosefpor/httpmon/internal/store"
)

type User struct {
	Store store.User
}

// nolint: wrapcheck
func (u User) CreateUser(c *fiber.Ctx) error {
	req := new(request.User)

	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load user data %s", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Printf("cannot validate user data %s", err)

		return fiber.ErrBadRequest
	}

	user := model.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := u.Store.CreateUser(c.Context(), user); err != nil {
		if errors.Is(err, store.ErrUserDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "user already exists")
		}

		log.Printf("cannot save user %s", err)

		return fiber.ErrInternalServerError
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":         req.ID,
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"token": t})
}

// nolint: wrapcheck
func (u User) CreateEndpoint(c *fiber.Ctx) error {
	req := new(request.Endpoint)

	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load endpoint data %s", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Printf("cannot validate endpoint data %s", err)

		return fiber.ErrBadRequest
	}

	ep := model.Endpoint{
		URL:       req.URL,
		Interval:  req.Interval,
		Threshold: req.Threshold,
	}

	id := jwtGetId(c)
	log.Println(ep)
	if err := u.Store.CreateEndpoint(c.Context(), id, ep); err != nil {
		if errors.Is(err, store.ErrUserDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "endpoint already exists")
		}

		log.Printf("cannot update endpoint %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(ep)
}

// nolint: wrapcheck
func (u User) ListEndpoints(c *fiber.Ctx) error {
	id := jwtGetId(c)

	eps, err := u.Store.ListEndpoints(c.Context(), id)
	if err != nil {
		log.Printf("cannot load user %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(eps)
}

// nolint: wrapcheck
func (u User) ListAlerts(c *fiber.Ctx) error {
	id := jwtGetId(c)

	eps, err := u.Store.ListEndpoints(c.Context(), id)
	if err != nil {
		log.Printf("cannot list alerts %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(eps)
}

// nolint: wrapcheck
func (u User) StatEndpoint(c *fiber.Ctx) error {
	id := jwtGetId(c)
	url := c.Params("url", "")
	ep, err := u.Store.StatEndpoint(c.Context(), id, url)
	if err != nil {
		log.Printf("cannot stat endpoint %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(ep)
}

func (u User) Register(g fiber.Router) {
	// g.Post("/create_user", u.CreateUser)
	g.Post("/create_endpoint", u.CreateEndpoint)
	g.Get("/list_endpoint", u.ListEndpoints)
	g.Get("/list_alerts", u.ListAlerts)
	g.Get("/stat_endpoint/:url", u.StatEndpoint)
}
