package httpservice

import (
	"architect/modules/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
	uc usecase.UseCase
}

func NewUserService(uc usecase.UseCase) service {
	return service{uc: uc}
}

func (s service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordRegisterDTO

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.uc.Register(c.Request.Context(), dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func (s service) handleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordLoginDTO

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := s.uc.Login(c.Request.Context(), dto)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": resp})
	}
}

func (s service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/login", s.handleLogin())
}
