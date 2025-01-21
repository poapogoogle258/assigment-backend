package main

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service := NewPieFireDire(GetMeetText)

	r.GET("/beef/summary", service.PieFireDire)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type IPieFireDire interface {
	PieFireDire(c *gin.Context)
}

type PieFireDire struct {
	GetMeetText func() string
}

func NewPieFireDire(GetMeetText func() string) IPieFireDire {
	return &PieFireDire{
		GetMeetText: GetMeetText,
	}
}

func (p *PieFireDire) PieFireDire(c *gin.Context) {

	text := p.GetMeetText()
	text = strings.ToLower(text)

	beefsPattern := `\b(?:alcatra|andouille|bacon|beef|belly|biltong|boudin|bresaola|brisket|buffalo|capicola|chicken|chislic|chop|chuck|corned|cow|cupim|drumstick|fatback|filet|flank|frankfurter|ground|ham|hamburger|hock|jerky|jowl|kielbasa|landjaeger|leberkas|loin|meatball|meatloaf|mignon|pancetta|pastrami|picanha|pig|porchetta|pork|prosciutto|ribeye|ribs|round|rump|salami|sausage|shank|short|shoulder|sirloin|spare|steak|strip|swine|t-bone|tail|tenderloin|tip|tongue|tri-tip|turducken|turkey|venison)\b`
	var beef = regexp.MustCompile(beefsPattern)

	list := beef.FindAllString(text, -1)

	var mapSummary = make(map[string]int)

	for i := 0; i < len(list); i++ {
		if value, has := mapSummary[list[i]]; has {
			mapSummary[list[i]] = value + 1
		} else {
			mapSummary[list[i]] = 1

		}
	}

	c.JSON(http.StatusOK, mapSummary)
}

func GetMeetText() string {

	ScrapingEndPoint := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text" // should change querystring to be "?type=all-meat&format=text" is better.

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(ScrapingEndPoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
