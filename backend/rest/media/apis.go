package media

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/burning/common"
	"github.com/weridolin/burning/models"
)

func GetMusicHandler(c *gin.Context) {
	// 获取查询参数中的分页参数
	page := common.Str2Int(c.DefaultQuery("page", "1"))
	pageSize := common.Str2Int(c.DefaultQuery("size", "10"))
	fmt.Println("get music list params -> ", page, pageSize)

	// 查询音乐列表
	musics, total, err := models.QueryMusicList(page, pageSize, map[string]interface{}{"enable": true}, common.DB)
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// pre := c.Request.URL.Path + "?page=" + common.Int2Str(page-1) + "&size=" + common.Int2Str(pageSize)
	// next := ""
	common.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"list":  musics,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
