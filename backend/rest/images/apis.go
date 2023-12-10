package images

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

func MatchRegex(regex string, path string) bool {
	pathRegex, err := regexp.Compile(regex)
	// fmt.Println("pathRegex", pathRegex)
	if err != nil {
		fmt.Println("Invalid regex pattern:", err)
		return false
	}
	// 提取路径参数
	match := pathRegex.FindStringSubmatch(path)
	// fmt.Println("url match result -> ", match)
	if len(match) > 0 {
		for _, v := range match {
			if v == path {
				fmt.Println("match regex success", regex, path)
				return true
			}
		}
	}
	return false
}

func GetImageHandler(c *gin.Context) {
	// 获取路径参数
	filePath := c.Param("filepath")
	// 构造完整的文件路径，这里假设图片文件都放在 "images" 目录下
	imagePath := "images/" + filePath
	fmt.Println("get image ->", filePath)
	regex := "/burning/api/v1/image/(?P<file_path>.*)"
	path := "/burning/api/v1/image/actions/yingla.png"
	fmt.Println("match result -> ", MatchRegex(regex, path))
	if _, err := os.Stat(imagePath); err != nil {
		c.JSON(404, gin.H{
			"message": "image not found",
		})
		return
	}
	// 返回图片文件
	c.File(imagePath)
}
