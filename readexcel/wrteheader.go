/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/9/12 1:31
*  @To:
 */

package readexcel

import (
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"test/models"
)

type resdata struct {
	List []string  `json:"list"`
	BA   models.BA `json:"ElementA"`
	BB   models.BB `json:"ElementB"`
	BC   models.BC `json:"ElementC"`
	BD   models.BD `json:"ElementD"`
	BE   models.BE `json:"ElementE"`
	BF   models.BF `json:"ElementF"`
}

func BridgeDeck(c *gin.Context) {
	var sqlite = models.GetSqlitedb()
	var res resdata
	c.ShouldBindJSON(&res)
	var deck models.BridgeDeck
	for _, item := range res.List {
		switch item {
		case "BA":
			res.BA.MDP()
			var ba models.BA
			ba = res.BA
			sqlite.Create(&ba)
			deck.BAId = uint(ba.ID)
		case "BB":
			res.BB.MDP()
			var bb models.BB
			sqlite.Create(&bb)
			deck.BBId = uint(bb.ID)
		case "BC":
			res.BC.MDP()
			var bc models.BC
			bc = res.BC
			sqlite.Create(&bc)
			deck.BCId = uint(bc.ID)
		case "BD":
			res.BD.MDP()
			var bd models.BD
			sqlite.Create(&bd)
			deck.BDId = uint(bd.ID)
		case "BE":
			res.BE.MDP()
			var be models.BE
			be = res.BE
			sqlite.Create(&be)
			deck.BEId = uint(be.ID)
		case "BF":
			res.BF.MDP()
			var bf models.BF
			bf = res.BF
			sqlite.Create(&bf)
			deck.BFId = uint(bf.ID)
		}
	}
	sqlite.Create(&deck)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "桥面信息保存成功",
		"data": deck,
	})
}
func Style(f *excelize.File, aglin string) int {
	style, _ := f.NewStyle(&excelize.Style{
		Border: nil,
		Fill:   excelize.Fill{},
		Font:   nil,
		Alignment: &excelize.Alignment{
			Horizontal: aglin,    //水平居中
			Vertical:   "center", //垂直居中
		},
	})

	return style

}

func Config() {

	f, _ := excelize.OpenFile("./file/template/template.xlsx")
	// 创建一个工作表
	sheet := "结果输出"
	f.SetCellValue(sheet, "B3", 93.00)
	f.SetCellValue(sheet, "B4", 95.38)
	f.SetCellValue(sheet, "B5", 94.76)
	f.SetCellFormula(sheet, "D3", "=B3*C3+B4*C4+B5*C5")
	// 根据指定路径保存文件
	if err := f.SaveAs("./media/Book1.xlsx"); err != nil {
		println(err.Error())
	}
}
