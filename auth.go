package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	PartnerKey = "your_partner_key_here"
)

func validateShopeeSignature(c *fiber.Ctx) error {
	partnerID := c.Query("partner_id")
	timestamp := c.Query("timestamp")
	accessToken := c.Query("access_token")
	shopID := c.Query("shop_id")
	sign := c.Query("sign")

	if partnerID == "" || timestamp == "" || sign == "" {
		return c.Status(401).JSON(fiber.Map{
			"error":   "unauthorized",
			"message": "Missing required authentication parameters",
		})
	}

	path := c.Path()
	baseString := fmt.Sprintf("%s%s%s%s", partnerID, path, timestamp, accessToken)
	
	if shopID != "" {
		baseString += shopID
	}

	expectedSign := generateHMACSHA256(baseString, PartnerKey)
	
	if !hmac.Equal([]byte(sign), []byte(expectedSign)) {
		return c.Status(401).JSON(fiber.Map{
			"error":   "invalid_signature",
			"message": "Invalid signature",
		})
	}

	return c.Next()
}

func generateHMACSHA256(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func validateTimestamp(c *fiber.Ctx) error {
	timestampStr := c.Query("timestamp")
	if timestampStr == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   "missing_timestamp",
			"message": "Timestamp is required",
		})
	}

	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "invalid_timestamp",
			"message": "Invalid timestamp format",
		})
	}

	c.Locals("timestamp", timestamp)
	return c.Next()
}