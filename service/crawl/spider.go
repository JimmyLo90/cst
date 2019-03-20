package crawl

import (
	"html"
	"regexp"
	"spider/sougou"
)

func SougouMpnewsUrl(title string) string {
	//查询搜索狗引擎，获取html内容供解析
	res := html.UnescapeString(string(sougou.Act(title)))
	//替换red_beg、red_end与em标签
	replace, err := regexp.Compile(`(\<\!\-\-red_beg\-\-\>)|(\<\!\-\-red_end\-\-\>)|(\<em\>)|(\<\/em\>)`)
	if err != nil {
		panic("replace red-class error")
	}
	res = replace.ReplaceAllString(res, "")
	//搜索title并记录url
	search, err := regexp.Compile(`\<a\s+target\=\"_blank\"\s+href\=\"(http[s]?\:\/\/[^\"\'\s]+)\"\s+id\=\"[a-zA-Z\d\_]+\"\s+uigs\=\"article_title\_0\"`)

	subMatch := search.FindStringSubmatch(res)
	if subMatch == nil {
		panic("未匹配到数据 ID=>")
	}

	return subMatch[1]
}
