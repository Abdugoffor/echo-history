package history

import (
	"net"

	"github.com/labstack/echo/v4"
)

func RequestContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		if ip == "" {
			ip = getLocalIP()
		}
		c.Set("ip", ip)

		// auth middleware qo‘shilgach shu yerga c.Set("user_id", int64(1)) yoziladi

		return next(c)
	}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "unknown"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "unknown"
}
