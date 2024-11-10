package main

import (
	"fmt"

	"github.com/shi-yunsheng/driver"
)

func main() {
	d := driver.New()
	d.Connect("jjx8496hc6nntcqs")

	d.StartApp("com.ss.android.ugc.aweme")

	searchBtn, _ := d.WaitElement(driver.By{Selector: driver.ContentDesc, Value: "搜索"})
	if searchBtn != nil {
		searchBtn.Tap()
	}

	dy := "ismeSYS"
	msg := "哈喽，老石"

	search, _ := d.WaitElement(driver.By{Selector: driver.ResourceID, Value: "com.ss.android.ugc.aweme:id/et_search_kw"})
	if search != nil {
		search.Input(dy)
		search.Search()
	}

	user, _ := d.WaitElement(driver.By{Selector: driver.EndsWithText, Value: fmt.Sprintf("抖音号：%s，按钮", dy)})
	if user != nil {
		user.Tap()
	}

	more, _ := d.WaitElement(driver.By{Selector: driver.ContentDesc, Value: "更多"})
	if more != nil {
		more.Tap()
	}

	sendBtn, _ := d.WaitElement(driver.By{Selector: driver.Text, Value: "发私信"})
	if sendBtn != nil {
		sendBtn.Tap()
	}

	message, _ := d.WaitElement(driver.By{Selector: driver.Text, Value: "发送消息"})
	if message != nil {
		message.Input(msg)
		message.Send()

		for i := 0; i < 5; i++ {
			d.Back()
		}
	}

	defer d.Cleanup()
}
