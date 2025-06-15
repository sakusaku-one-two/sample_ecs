package handler

import (
	"context"
	"fmt"
	"module/infra"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	MAIN_HANDLER *Handler
)

func init() {
	MAIN_HANDLER = NewHandler()
}

func SetupRoutes(e *echo.Echo) {

	e.GET("/", MAIN_HANDLER.root)
	e.GET("/health", MAIN_HANDLER.healthCheck)
	e.GET("/all", MAIN_HANDLER.AllRedisData)

	e.POST("/add", MAIN_HANDLER.SetData)

}

type Handler struct {
	redis_clinet *infra.RedisClient
}

func NewHandler() *Handler {
	redis_client, _ := infra.NewRedisClient()
	return &Handler{
		redis_clinet: redis_client,
	}
}

// -------------GET----------------//
func (h *Handler) root(c echo.Context) error {
	fmt.Println("root called")
	self_ip := c.RealIP()

	return c.JSON(http.StatusOK, map[string]string{
		"self_ip": self_ip,
	})

}
func (h *Handler) healthCheck(c echo.Context) error {
	if h.redis_clinet.HealthCheck() {
		return c.String(http.StatusOK, "StatusOk")
	}
	return c.String(http.StatusInternalServerError, "Redis client not connection")
}

const LIST_KEY = "list_key"

func (h *Handler) AllRedisData(c echo.Context) error {
	fmt.Println("GET ALL REDIS DATA")
	// redis の情報を全て取得する
	client := h.redis_clinet.GetClient()
	all_data, err := client.LRange(context.Background(), LIST_KEY, 0, -1).Result()
	if err != nil {
		fmt.Println("FIALD GET ALL REDIS DATA")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string][]string{
		"result": all_data,
	})
}

//--------------POST----------------//

type Data struct {
	Value string `json:"value"`
}

func (h *Handler) SetData(c echo.Context) error {
	fmt.Println("SET DATA FOR REDIS")
	var set_data Data
	if err := c.Bind(&set_data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	client := h.redis_clinet.GetClient()
	ctx := context.Background()

	err := client.RPush(ctx, LIST_KEY, set_data.Value).Err()
	if err != nil {
		fmt.Println("FAILED REDIS ADD")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"error": "",
	})
}
