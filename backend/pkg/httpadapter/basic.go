package httpadapter

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFuncWithContext func(ctx context.Context, c *gin.Context) (interface{}, error)
type HandlerFuncWithoutContext func(c *gin.Context) (interface{}, error)

func WrapWithAdditionalContext(fn HandlerFuncWithContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Вызываем бизнес‑логику
		result, err := fn(c.Request.Context(), c)
		if err != nil {
			// Вариант простого error handling’а
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Отдаём ответ 200 + результат
		c.JSON(http.StatusOK, result)
	}
}

func WrapWithoutAdditionalContext(fn HandlerFuncWithoutContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Вызываем бизнес‑логику
		result, err := fn(c)
		if err != nil {
			// Вариант простого error handling’а
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Отдаём ответ 200 + результат
		c.JSON(http.StatusOK, result)
	}
}
