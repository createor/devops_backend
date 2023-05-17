package test

import (
	"fmt"
	"server/utils"
	"testing"
)

func TestSaveArticle(t *testing.T) {
	testStr := `<script>alert("XSS");</script><p>This is a <a href="javascript:alert('XSS')">dangerous</a> link.</p>`
	filename := utils.GetFileNameByRandom()
	want, err := utils.SaveArticle(utils.FilterArticle(testStr), filename)
	if err != nil || want != true {
		t.Errorf("保存失败:%v", err)
	}
}

func TestShowArtcle(t *testing.T) {
	filename := "839e7732959e0aeca078eb0f482e9a15d5d8e693511b076b4a886265c3b2"
	want := utils.ShowArtcle(filename)
	if want == "" {
		t.Error("读取文章失败")
	} else {
		fmt.Println(want)
	}
}
