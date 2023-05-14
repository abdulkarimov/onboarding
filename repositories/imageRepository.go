package repositories

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := path.Ext(filename)
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image upload error: ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	imageUrl := fmt.Sprintf("http://127.0.0.1:3000/images/%s", image)
	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}

func DeleteImage(c *fiber.Ctx) error {
	imageName := c.Params("imageName")

	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}
