package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseUtgid(c *gin.Context) (int64, error) {
	utgidStr := c.Param("utgid")

	utgid, err := strconv.ParseInt(utgidStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return utgid, nil
}
