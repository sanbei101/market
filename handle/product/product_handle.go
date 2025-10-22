package product

import (
	"market/handle/response"
	"market/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

type ProductHandle struct {
	db *model.Queries
}

func NewProductHandle(db *model.Queries) *ProductHandle {
	return &ProductHandle{db: db}
}

// SPU handlers
func (h *ProductHandle) ListProductSPU(c *gin.Context) {
	spulist, err := h.db.ListProductSPUs(c)
	if err != nil {
		log.Error().Err(err).Msg("获取商品SPU列表失败")
		response.Error(c, err, "获取商品SPU列表失败")
		return
	}
	log.Info().Msgf("获取商品SPU列表成功")
	response.Success(c, spulist, "获取商品SPU列表成功")
}

func (h *ProductHandle) GetProductSPU(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	spu, err := h.db.GetProductSPU(c, id)
	if err != nil {
		log.Error().Err(err).Msg("获取商品SPU详情失败")
		response.Error(c, err, "获取商品SPU详情失败")
		return
	}
	log.Info().Msgf("获取商品SPU详情成功")
	response.Success(c, spu, "获取商品SPU详情成功")
}

func (h *ProductHandle) CreateProductSPU(c *gin.Context) {
	var req model.CreateProductSPUParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}

	spu, err := h.db.CreateProductSPU(c, req)
	if err != nil {
		log.Error().Err(err).Msg("创建商品SPU失败")
		response.Error(c, err, "创建商品SPU失败")
		return
	}
	log.Info().Msgf("创建商品SPU成功")
	response.Success(c, spu, "创建商品SPU成功")
}

func (h *ProductHandle) UpdateProductSPU(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	var req model.UpdateProductSPUParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}
	req.ID = id

	spu, err := h.db.UpdateProductSPU(c, req)
	if err != nil {
		log.Error().Err(err).Msg("更新商品SPU失败")
		response.Error(c, err, "更新商品SPU失败")
		return
	}
	log.Info().Msgf("更新商品SPU成功")
	response.Success(c, spu, "更新商品SPU成功")
}

func (h *ProductHandle) DeleteProductSPU(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	err = h.db.DeleteProductSPU(c, id)
	if err != nil {
		log.Error().Err(err).Msg("删除商品SPU失败")
		response.Error(c, err, "删除商品SPU失败")
		return
	}
	log.Info().Msgf("删除商品SPU成功")
	response.Success(c, true, "删除商品SPU成功")
}

func (h *ProductHandle) ListProductSPUsByCategory(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析分类ID失败")
		response.Error(c, err, "无效的分类ID参数")
		return
	}

	spulist, err := h.db.ListProductSPUsByCategory(c, categoryID)
	if err != nil {
		log.Error().Err(err).Msg("获取分类下商品SPU列表失败")
		response.Error(c, err, "获取分类下商品SPU列表失败")
		return
	}
	log.Info().Msgf("获取分类下商品SPU列表成功")
	response.Success(c, spulist, "获取分类下商品SPU列表成功")
}

// SKU handlers
func (h *ProductHandle) GetProductSKU(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	sku, err := h.db.GetProductSKU(c, id)
	if err != nil {
		log.Error().Err(err).Msg("获取商品SKU详情失败")
		response.Error(c, err, "获取商品SKU详情失败")
		return
	}
	log.Info().Msgf("获取商品SKU详情成功")
	response.Success(c, sku, "获取商品SKU详情成功")
}

func (h *ProductHandle) CreateProductSKU(c *gin.Context) {
	var req model.CreateProductSKUParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}

	sku, err := h.db.CreateProductSKU(c, req)
	if err != nil {
		log.Error().Err(err).Msg("创建商品SKU失败")
		response.Error(c, err, "创建商品SKU失败")
		return
	}
	log.Info().Msgf("创建商品SKU成功")
	response.Success(c, sku, "创建商品SKU成功")
}

func (h *ProductHandle) UpdateProductSKU(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	var req model.UpdateProductSKUParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}
	req.ID = id

	sku, err := h.db.UpdateProductSKU(c, req)
	if err != nil {
		log.Error().Err(err).Msg("更新商品SKU失败")
		response.Error(c, err, "更新商品SKU失败")
		return
	}
	log.Info().Msgf("更新商品SKU成功")
	response.Success(c, sku, "更新商品SKU成功")
}

func (h *ProductHandle) DeleteProductSKU(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	err = h.db.DeleteProductSKU(c, id)
	if err != nil {
		log.Error().Err(err).Msg("删除商品SKU失败")
		response.Error(c, err, "删除商品SKU失败")
		return
	}
	log.Info().Msgf("删除商品SKU成功")
	response.Success(c, true, "删除商品SKU成功")
}

func (h *ProductHandle) ListProductSKUsBySPU(c *gin.Context) {
	spuIDStr := c.Query("spu_id")
	spuID, err := strconv.ParseInt(spuIDStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析SPU ID失败")
		response.Error(c, err, "无效的SPU ID参数")
		return
	}

	skuList, err := h.db.ListProductSKUsBySPU(c, spuID)
	if err != nil {
		log.Error().Err(err).Msg("获取SPU下的SKU列表失败")
		response.Error(c, err, "获取SPU下的SKU列表失败")
		return
	}
	log.Info().Msgf("获取SPU下的SKU列表成功")
	response.Success(c, skuList, "获取SPU下的SKU列表成功")
}

// Specialized handlers
func (h *ProductHandle) GetSPUWithSKUs(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	result, err := h.db.GetSPUWithSKUs(c, id)
	if err != nil {
		log.Error().Err(err).Msg("获取SPU及所有SKU失败")
		response.Error(c, err, "获取SPU及所有SKU失败")
		return
	}
	log.Info().Msgf("获取SPU及所有SKU成功")
	response.Success(c, result, "获取SPU及所有SKU成功")
}

func (h *ProductHandle) GetCategorySpecValues(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析分类ID失败")
		response.Error(c, err, "无效的分类ID参数")
		return
	}

	specValues, err := h.db.GetCategorySpecValues(c, categoryID)
	if err != nil {
		log.Error().Err(err).Msg("获取分类规格值失败")
		response.Error(c, err, "获取分类规格值失败")
		return
	}
	log.Info().Msgf("获取分类规格值成功")
	response.Success(c, specValues, "获取分类规格值成功")
}

func (h *ProductHandle) SearchProductsBySpecs(c *gin.Context) {
	var req model.SearchProductsBySpecsParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}

	products, err := h.db.SearchProductsBySpecs(c, req)
	if err != nil {
		log.Error().Err(err).Msg("根据规格搜索商品失败")
		response.Error(c, err, "根据规格搜索商品失败")
		return
	}
	log.Info().Msgf("根据规格搜索商品成功")
	response.Success(c, products, "根据规格搜索商品成功")
}

func (h *ProductHandle) DeleteProductSKUsBySPU(c *gin.Context) {
	spuIDStr := c.Query("spu_id")
	spuID, err := strconv.ParseInt(spuIDStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析SPU ID失败")
		response.Error(c, err, "无效的SPU ID参数")
		return
	}

	err = h.db.DeleteProductSKUsBySPU(c, spuID)
	if err != nil {
		log.Error().Err(err).Msg("删除SPU下所有SKU失败")
		response.Error(c, err, "删除SPU下所有SKU失败")
		return
	}
	log.Info().Msgf("删除SPU下所有SKU成功")
	response.Success(c, true, "删除SPU下所有SKU成功")
}
