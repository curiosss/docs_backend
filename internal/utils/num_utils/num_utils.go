package numutils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetUintParam parses a path or query parameter into uint
func GetUintParam(c echo.Context, name string) (uint, error) {
	param := c.Param(name)
	fmt.Println(param)
	if param == "" {
		param = c.QueryParam(name)
	}
	id64, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, errors.New("invalid parameter: " + name)
	}
	return uint(id64), nil
}

// GetIntParam parses a path or query parameter into int
func GetIntParam(c echo.Context, name string) (int, error) {
	param := c.Param(name)
	if param == "" {
		param = c.QueryParam(name)
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, errors.New("invalid parameter: " + name)
	}
	return id, nil
}

// GetFloatParam parses a path or query parameter into float64
func GetFloatParam(c echo.Context, name string) (float64, error) {
	param := c.Param(name)
	if param == "" {
		param = c.QueryParam(name)
	}
	val, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, errors.New("invalid parameter: " + name)
	}
	return val, nil
}

// GetBoolParam parses a path or query parameter into bool
func GetBoolParam(c echo.Context, name string) (bool, error) {
	param := c.Param(name)
	if param == "" {
		param = c.QueryParam(name)
	}
	val, err := strconv.ParseBool(param)
	if err != nil {
		return false, errors.New("invalid parameter: " + name)
	}
	return val, nil
}
