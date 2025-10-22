package product

import (
	"market/handle/response"
	"market/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/phuslu/log"
)

type CategoryHandle struct {
	db *model.Queries
}

func NewCategoryHandle(db *model.Queries) *CategoryHandle {
	return &CategoryHandle{db: db}
}

func (h *CategoryHandle) ListCategory(c *gin.Context) {
	categorys, err := h.db.ListCategories(c)
	if err != nil {
		log.Error().Err(err).Msg("获取分类失败")
		response.Error(c, err, "获取分类失败")
		return
	}
	log.Info().Msgf("获取分类成功")
	response.Success(c, categorys, "获取分类成功")
}

func (h *CategoryHandle) GetCategory(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	category, err := h.db.GetCategory(c, id)
	if err != nil {
		log.Error().Err(err).Msg("获取分类详情失败")
		response.Error(c, err, "获取分类详情失败")
		return
	}
	log.Info().Msgf("获取分类详情成功")
	response.Success(c, category, "获取分类详情成功")
}

func (h *CategoryHandle) CreateCategory(c *gin.Context) {
	var req model.CreateCategoryParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}

	category, err := h.db.CreateCategory(c, req)
	if err != nil {
		log.Error().Err(err).Msg("创建分类失败")
		response.Error(c, err, "创建分类失败")
		return
	}
	log.Info().Msgf("创建分类成功")
	response.Success(c, category, "创建分类成功")
}

func (h *CategoryHandle) UpdateCategory(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	var req model.UpdateCategoryParams
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("参数绑定失败")
		response.Error(c, err, "参数格式错误")
		return
	}
	req.ID = id

	category, err := h.db.UpdateCategory(c, req)
	if err != nil {
		log.Error().Err(err).Msg("更新分类失败")
		response.Error(c, err, "更新分类失败")
		return
	}
	log.Info().Msgf("更新分类成功")
	response.Success(c, category, "更新分类成功")
}

func (h *CategoryHandle) DeleteCategory(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("解析ID失败")
		response.Error(c, err, "无效的ID参数")
		return
	}

	err = h.db.DeleteCategory(c, id)
	if err != nil {
		log.Error().Err(err).Msg("删除分类失败")
		response.Error(c, err, "删除分类失败")
		return
	}
	log.Info().Msgf("删除分类成功")
	response.Success(c, true, "删除分类成功")
}

func (h *CategoryHandle) ListCategoriesByParent(c *gin.Context) {
	parentIDStr := c.Query("parent_id")
	var parentID pgtype.Int8

	if parentIDStr == "0" || parentIDStr == "" {
		parentID.Valid = false
	} else {
		id, err := strconv.ParseInt(parentIDStr, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("解析父级ID失败")
			response.Error(c, err, "无效的父级ID参数")
			return
		}
		parentID.Int64 = id
		parentID.Valid = true
	}

	categories, err := h.db.ListCategoriesByParent(c, parentID)
	if err != nil {
		log.Error().Err(err).Msg("获取子分类失败")
		response.Error(c, err, "获取子分类失败")
		return
	}
	log.Info().Msgf("获取子分类成功")
	response.Success(c, categories, "获取子分类成功")
}
