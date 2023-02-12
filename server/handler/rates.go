package handler

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getRates(c *gin.Context) {
	input := c.Query("pairs")

	if input == "" {
		newErrorResponse(c, http.StatusBadRequest, "Empty 'pairs' param")
		return
	}

	pairs := strings.Split(input, ",")

	if err := validatePairs(pairs); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	tickers, err := h.services.Rates.GetTickers(pairs)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickers)
}

type postRatesInput struct {
	Pairs []string `json:"pairs" binding:"required"`
}

func (h *Handler) postRates(c *gin.Context) {
	var input postRatesInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	pairs := input.Pairs

	if err := validatePairs(pairs); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	tickers, err := h.services.Rates.GetTickers(pairs)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickers)
}

func validatePairs(pairs []string) error {
	pairRegexp, _ := regexp.Compile("^[A-Z0-9]+-[A-Z0-9]+$")

	for _, pair := range pairs {
		match := pairRegexp.MatchString(pair)

		if match == false {
			msg := fmt.Sprintf("Pair symbol '%s' is invalid", pair)
			return errors.New(msg)
		}
	}

	return nil
}
