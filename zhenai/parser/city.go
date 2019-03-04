package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	log.Printf("%s",matches)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User:"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
	}
	return result
}