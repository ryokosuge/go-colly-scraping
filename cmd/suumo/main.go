package main

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("went wrong: ", err)
	})

	c.OnHTML("div.property_unit ", func(e *colly.HTMLElement) {
		title := e.DOM.Find("h2.property_unit-title > a").Text()
		log.Println(title)
		lines := e.DOM.Find("div.property_unit-info div.dottable-line")
		propertyName := lines.Eq(0).Find("dl > dd").Text()
		log.Println("\t物件名: ", propertyName)
		price := lines.Eq(1).Find("dl > dd > span").Text()
		log.Println("\t販売価格: ", price)
		location := lines.Eq(2).Find("dl").Eq(0).Find("dd").Text()
		log.Println("\t所在地: ", location)
		nearestStation := lines.Eq(2).Find("dl").Eq(1).Find("dd").Text()
		log.Println("\t最寄り駅: ", nearestStation)
		area := lines.Eq(3).Find("table tr td").Eq(0).Find("dl > dd").Text()
		log.Println("\t専有面積: ", area)
		floorPlan := lines.Eq(3).Find("table tr td").Eq(1).Find("dl > dd").Text()
		log.Println("\t間取り: ", floorPlan)
		balcony := lines.Eq(4).Find("table tr td").Eq(0).Find("dl > dd").Text()
		log.Println("\tバルコニー: ", balcony)
		ageOfBuilding := lines.Eq(4).Find("table tr td").Eq(1).Find("dl > dd").Text()
		log.Println("\t築年数: ", ageOfBuilding)
	})

	c.Visit("https://suumo.jp/jj/bukken/ichiran/JJ012FC001/?ar=030&bs=011&cn=25&cnb=0&ekTjCd=&ekTjNm=&et=20&kb=1&kr=A&kt=8000&mb=70&md=3&mt=9999999&sc=13101&sc=13102&sc=13103&sc=13104&sc=13105&sc=13113&sc=13107&sc=13109&sc=13110&sc=13111&sc=13112&sc=13114&sc=13115&sc=13120&sc=13116&sc=13117&sc=13119&ta=13&tj=0&pc=100&po=1&pj=2")
}