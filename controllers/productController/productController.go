package productController

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tsaqifmaulana444/go-gin-restapi/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message:": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": "Couldn't Update Product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Successfully Updated"})
}

func Delete(c *gin.Context) {
    id := c.Param("id")

    var product models.Product
    if err := models.DB.First(&product, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }

    if err := models.DB.Delete(&product).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
