/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/9/12 1:29
*  @To:
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/models"
	"test/readexcel"
)

func main() {
	router := gin.Default()
	models.InitsqLite()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
		})

	})

	/*----------------普通路由-------------------*/
	router.POST("/api/test", readexcel.BridgeDeck)
	/*---------------------运行----------------*/
	router.StaticFS("/public", http.Dir("../asset"))
	err := router.Run(":8000") // 监听端口
	fmt.Println("启动成功")
	if err != nil {
		fmt.Println("启动失败") // 启动失败
		return
	}

}
