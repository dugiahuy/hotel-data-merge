package updater

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/dugiahuy/hotel-data-merge/src/usecase/updater"
)

type updaterHandler struct {
	usecase updater.Usecase
}

func NewHandler(e *echo.Echo, us updater.Usecase) {
	handler := &updaterHandler{
		usecase: us,
	}

	e.GET("/updater", handler.UpdateData)
}

func (h *updaterHandler) UpdateData(c echo.Context) error {
	if err := h.usecase.CollectData(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": "OK",
	})
}
