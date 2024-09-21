package interfaces

import (
	"github.com/gin-gonic/gin"
)

type CompanyController interface {
	RegisterCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	GetCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}