package hotel

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/dugiahuy/hotel-data-merge/src/usecase/hotel"
	"github.com/dugiahuy/hotel-data-merge/src/util/err_util"
)

type hotelHandler struct {
	usecase hotel.Usecase
}

func NewHandler(e *echo.Echo, us hotel.Usecase) {
	handler := &hotelHandler{
		usecase: us,
	}

	e.GET("/hotels", handler.Fetch)
	e.GET("/hotels/:id", handler.Get)
	e.GET("/hotels/destination/:id", handler.GetByDestination)
}

func (h *hotelHandler) Fetch(c echo.Context) error {
	hotels, err := h.usecase.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hotels)
}

func (h *hotelHandler) Get(c echo.Context) error {
	id := c.Param("id")
	hotel, err := h.usecase.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if hotel == nil {
		return c.JSON(http.StatusNotFound, err_util.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, hotel)
}

func (h *hotelHandler) GetByDestination(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := int64(idP)
	hotels, err := h.usecase.GetByDestination(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if hotels == nil {
		return c.JSON(http.StatusNotFound, err_util.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, hotels)
}
