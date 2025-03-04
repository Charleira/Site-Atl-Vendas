package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole") // Obtém o papel do usuário do token JWT
		if !exists || userRole != "admin" {
			// Log para registrar tentativas de acesso não autorizadas
			log.Printf("Tentativa de acesso não autorizado por usuário com papel: %v", userRole)

			c.JSON(http.StatusForbidden, gin.H{
				"error": "Acesso negado: usuário sem permissão de administrador",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
