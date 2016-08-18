package parser

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func ParseURL(input *http.Response) string {
	b, err := ioutil.ReadAll(input.Body)
	defer input.Body.Close()
	if err != nil {
		return ""
	}
	html := string(b[:])
	re := regexp.MustCompile(`http\:\/\/stackoverflow\.com.*?\"`)
	matches := re.FindString(html)
	return matches[:len(matches)-1]
}

func ParseAnswer(input *http.Response) string {
	b, err := ioutil.ReadAll(input.Body)
	defer input.Body.Close()
	if err != nil {
		return ""
	}
	html := string(b[:])

	reQuestion := regexp.MustCompile(`class=\"question-hyperlink\"\>(.*?)\<\/a\>`)
	question := reQuestion.FindStringSubmatch(html)[1]

	reAnswer := regexp.MustCompile(`(?s)class=\"answercell\"\>.*?itemprop=\"text\"\>(.*?)\<\/div\>`)
	answer := reAnswer.FindStringSubmatch(html)[1]

	reAnswer = regexp.MustCompile(`(?s)\<code\>(.*?)\<\/code\>`)
	answer = reAnswer.ReplaceAllString(answer, "<yellow>$1</yellow>")

	reAnswer = regexp.MustCompile(`\<strong\>(.*?)\<\/strong\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$1</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<em\>(.*?)\<\/em\>`)
	answer = reAnswer.ReplaceAllString(answer, "<cyan>$1</cyan>")

	reAnswer = regexp.MustCompile(`(?s)\<blockquote\>(.*?)\<\/blockquote\>`)
	answer = reAnswer.ReplaceAllString(answer, "<green>$1</green>")

	reAnswer = regexp.MustCompile(`(?s)\<p\>(.*?)\<\/p\>`)
	answer = reAnswer.ReplaceAllString(answer, "$1")

	reAnswer = regexp.MustCompile(`(?s)\<pre\>(.*?)\<\/pre\>`)
	answer = reAnswer.ReplaceAllString(answer, "\n$1\n")

	reAnswer = regexp.MustCompile(`\<a href=\"(.*?)\"\>.*?\<\/a\>`)
	answer = reAnswer.ReplaceAllString(answer, "<blue>$1</blue>")

	output := "<green><u>Question:</u></green> " + question + "\n\n"
	output += "<green><u>Answer:</u></green>\n" + answer

	return output
}
