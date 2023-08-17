package transcript

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ygunayer/restbench/internal/response"
)

func handleGetV1(c *fiber.Ctx) error {
	id := c.Params("id")
	transcriptV2, err := GetTranscript(id)

	if err != nil {
		return err
	}

	transcriptV1 := SplitIntoWords(transcriptV2)

	return response.Ok(c, transcriptV1)
}

func handleGetV2(c *fiber.Ctx) error {
	id := c.Params("id")
	transcript, err := GetTranscript(id)

	if err != nil {
		return err
	}

	return response.Ok(c, transcript)
}

func BindV2Routes(r fiber.Router) {
	r.Get("/transcript/:id/full", handleGetV2)
}

func BindV1Routes(r fiber.Router) {
	r.Get("/transcript/:id", handleGetV1)
}
