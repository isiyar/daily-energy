package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseStartFinish(c *gin.Context) (int64, int64, error) {
	start := c.Query("start_at")
	finish := c.Query("finish_at")

	startInt, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		return 0, 0, errors.New("invalid start")
	}

	finishInt, err := strconv.ParseInt(finish, 10, 64)
	if err != nil {
		return 0, 0, errors.New("invalid end")
	}

	return startInt, finishInt, nil
}
