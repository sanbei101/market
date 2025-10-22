package handle

import (
	"market/handle/product"
	"market/model"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *model.Queries) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	categoryHandle := product.NewCategoryHandle(db)
	productHandle := product.NewProductHandle(db)

	productGroup := r.Group("/product")
	{
		// Category routes
		category := productGroup.Group("/category")
		{
			category.GET("/list", categoryHandle.ListCategory)
			category.GET("/detail", categoryHandle.GetCategory)
			category.POST("/create", categoryHandle.CreateCategory)
			category.PUT("/update", categoryHandle.UpdateCategory)
			category.DELETE("/delete", categoryHandle.DeleteCategory)
			category.GET("/children", categoryHandle.ListCategoriesByParent)
		}

		// SPU routes
		spu := productGroup.Group("/spu")
		{
			spu.GET("/list", productHandle.ListProductSPU)
			spu.GET("/detail", productHandle.GetProductSPU)
			spu.GET("/with-skus", productHandle.GetSPUWithSKUs)
			spu.GET("/list-by-category", productHandle.ListProductSPUsByCategory)
			spu.POST("/create", productHandle.CreateProductSPU)
			spu.PUT("/update", productHandle.UpdateProductSPU)
			spu.DELETE("/delete", productHandle.DeleteProductSPU)
		}

		// SKU routes
		sku := productGroup.Group("/sku")
		{
			sku.GET("/detail", productHandle.GetProductSKU)
			sku.GET("/list-by-spu", productHandle.ListProductSKUsBySPU)
			sku.POST("/create", productHandle.CreateProductSKU)
			sku.PUT("/update", productHandle.UpdateProductSKU)
			sku.DELETE("/delete", productHandle.DeleteProductSKU)
			sku.DELETE("/delete-by-spu", productHandle.DeleteProductSKUsBySPU)
		}

		// Specialized routes
		productGroup.GET("/category-spec-values", productHandle.GetCategorySpecValues)
		productGroup.POST("/search-by-specs", productHandle.SearchProductsBySpecs)
	}
	return r
}
