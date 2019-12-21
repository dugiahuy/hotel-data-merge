package updater

import (
	"net/http"

	"github.com/dugiahuy/hotel-data-merge/src/usecase/updater"
	"github.com/labstack/echo"
)

type updaterHandler struct {
	usecase updater.Usecase
}

func NewUpdaterHandler(e *echo.Echo, us updater.Usecase) {
	handler := &updaterHandler{
		usecase: us,
	}
	e.POST("/updater", handler.UpdateData)
}

func (h *updaterHandler) UpdateData(c echo.Context) error {
	hotels, err := h.usecase.CollectData()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, hotels)
}
