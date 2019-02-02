package functional

import (
	"regexp"
)

func RegexMatch(regex string, s string) bool {
	match, _ := regexp.MatchString(regex, s)
	return match
}

func RegexFindAll(regex string, s string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(s, n)
}

func RegexFind(regex string, s string) string {
	r := regexp.MustCompile(regex)
	return r.FindString(s)
}

func RegexReplaceAll(regex string, s string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, repl)
}

func RegExReplaceAllLiteral(regex string, s string, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllLiteralString(s, repl)
}

func RegExSplit(regex string, s string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.Split(s, n)
}
