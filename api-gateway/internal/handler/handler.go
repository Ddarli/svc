package handler

import (
	"bytes"
	"fmt"
	"github.com/Ddarli/svc/gateway/conf"
	"github.com/Ddarli/svc/gateway/internal/domain"
	"github.com/Ddarli/svc/gateway/internal/service"
	"github.com/gofiber/fiber/v2"
	"io"
)

type router struct {
	app     *fiber.App
	cfg     *conf.Conf
	service *service.Service
}

func New(app *fiber.App, cfg *conf.Conf, service *service.Service) *router {
	return &router{
		app:     app,
		cfg:     cfg,
		service: service,
	}
}

func (r *router) SetupRoutes() {
	r.app.Post("/auth/authenticate", r.handleAuth)
	r.app.Post("/auth/register", r.handleRegistration)
	r.app.Get("api/v1/file/download", r.jwtProtected(), r.cookieAuthMiddleware, r.handleDownloadFile)
	r.app.Get("api/v1/files", r.jwtProtected(), r.cookieAuthMiddleware, r.handleListFile)
	r.app.Post("api/v1/file/upload", r.jwtProtected(), r.cookieAuthMiddleware, r.handleUploadFile)
}

func (r *router) handleAuth(ctx *fiber.Ctx) error {
	var loginRequest domain.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	resp := r.service.Auth(ctx.Context(), loginRequest)
	if resp == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Token": resp,
	})
}

func (r *router) handleRegistration(ctx *fiber.Ctx) error {
	var registerRequest domain.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	resp := r.service.Register(ctx.Context(), registerRequest)
	if resp == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Token": resp,
	})
}

func (r *router) handleDownloadFile(ctx *fiber.Ctx) error {
	resp, err := r.service.DownloadFile(ctx.Context(), domain.DownloadFileRequest{
		UserID: ctx.Locals("user_id").(string),
		FileID: ctx.Query("file_id"),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to download file",
		})
	}

	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", resp.FileName))
	ctx.Set("Content-Type", resp.MimeType)

	return ctx.SendStream(bytes.NewReader(resp.FileData))
}

func (r *router) handleListFile(ctx *fiber.Ctx) error {
	resp, err := r.service.ListFile(ctx.Context(), domain.ListUserFileRequest{
		UserID: ctx.Locals("user_id").(string),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get files",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Files": resp,
	})
}

func (r *router) handleUploadFile(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("file is required")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("failed to open uploaded file")
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("failed to read file content")
	}

	description := ctx.FormValue("description")
	mimeType := fileHeader.Header.Get("Content-Type")

	res, err := r.service.UploadFile(ctx.Context(), domain.UploadFileRequest{
		UserID:      ctx.Locals("user_id").(string),
		FileName:    fileHeader.Filename,
		MimeType:    mimeType,
		FileData:    fileBytes,
		Description: description,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("failed to upload file")
	}

	return ctx.JSON(fiber.Map{
		"message":  res.Message,
		"filename": fileHeader.Filename,
	})
}
