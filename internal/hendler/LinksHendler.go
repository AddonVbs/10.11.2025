package hendler

import (
	"modul/internal/ServiceLinks"
	"modul/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ServiceLinskHendler struct {
	service ServiceLinks.LinksServiceInterface
}

func NewLinksServiHendler(s ServiceLinks.LinksServiceInterface) *ServiceLinskHendler {
	return &ServiceLinskHendler{service: s}
}

func (h *ServiceLinskHendler) GetLinkByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	links, err := h.service.GetStatusLinksServiceById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, links)
}

func (h *ServiceLinskHendler) GetLink(c echo.Context) error {
	var input models.Links

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	links, err := h.service.GetStatusLinksService(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, links)
}
