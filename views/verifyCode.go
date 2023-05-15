package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	captcha "github.com/mojocn/base64Captcha"
)

/*
验证码模块
*/

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverString {
	driver := new(captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 0
	driver.ShowLineOptions = captcha.OptionShowSlimeLine
	driver.Length = 6
	driver.Source = "1234567890qwertyuioplkjhgfdsazxcvbnm"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

// 生成图片验证码
func GenerateCaptchaHandler(c *gin.Context) {
	driver := NewDriver().ConvertFonts()
	capt := captcha.NewCaptcha(driver, store)
	id, content, answer := capt.Driver.GenerateIdQuestionAnswer()
	item, err := capt.Driver.DrawCaptcha(content)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "请求失败",
		})
		return
	}
	fmt.Println(id)
	fmt.Println(answer)
	capt.Store.Set(id, answer)
	// 将id返回到cookies中
	c.SetCookie("CODEID", id, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   item.EncodeB64string(),
		"msg":    "请求成功",
	})
}

// 验证图片验证码
func VerifyCaptchaHandler(id, code string) bool {
	return store.Verify(id, code, true)
}
