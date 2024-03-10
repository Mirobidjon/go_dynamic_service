package helper

import "github.com/gofiber/fiber/v2"

func GetIP(c *fiber.Ctx, signIp string) string {
	if signIp != "" {
		return signIp
	}

	ip := c.Get("X-Real-IP")
	if ip == "" {
		ips := c.IPs()
		if len(ips) > 0 {
			return ips[0]
		}
		return c.IP()
	}
	return ip
}
