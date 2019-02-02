package fmap

import (
	"errors"
	"github.com/gofunct/functional"
	"html/template"
	"os"
	"path"
	"strconv"
	"strings"
	ttemplate "text/template"
	"time"

	util "github.com/aokoli/goutils"
	"github.com/huandu/xstrings"
)

func FuncMap() template.FuncMap {
	return HtmlFuncMap()
}

// HermeticTextFuncMap returns a 'text/template'.FuncMap with only repeatable functions.
func HermeticTxtFuncMap() ttemplate.FuncMap {
	r := TxtFuncMap()
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// HermeticHtmlFuncMap returns an 'Html/template'.Funcmap with only repeatable functions.
func HermeticHtmlFuncMap() template.FuncMap {
	r := HtmlFuncMap()
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// TextFuncMap returns a 'text/template'.FuncMap
func TxtFuncMap() ttemplate.FuncMap {
	return ttemplate.FuncMap(GenericFuncMap())
}

// HtmlFuncMap returns an 'Html/template'.Funcmap
func HtmlFuncMap() template.FuncMap {
	return template.FuncMap(GenericFuncMap())
}

// GenericFuncMap returns a copy of the basic function map as a map[string]interface{}.
func GenericFuncMap() map[string]interface{} {
	gfm := make(map[string]interface{}, len(genericMap))
	for k, v := range genericMap {
		gfm[k] = v
	}
	return gfm
}

// These functions are not guaranteed to evaluate to the same result for given input, because they
// refer to the environemnt or global state.
var nonhermeticFunctions = []string{
	// Date functions
	"Date",
	"Date_in_zone",
	"Date_modify",
	"Now",
	"HtmlDate",
	"HtmlDateInZone",
	"DateInZone",
	"DateModify",

	// Strings
	"RandAlphaNum",
	"RandAlpha",
	"RandAscii",
	"RandNumeric",
	"Uuidv4",

	// OS
	"Env",
	"Expandenv",
}

var genericMap = map[string]interface{}{
	"Hello": func() string { return "Hello!" },

	// Date functions
	"Date":           functional.Date,
	"Date_in_zone":   functional.DateInZone,
	"Date_modify":    functional.DateModify,
	"Now":            func() time.Time { return time.Now() },
	"HtmlDate":       functional.HtmlDate,
	"HtmlDateInZone": functional.HtmlDateInZone,
	"DateInZone":     functional.DateInZone,
	"DateModify":     functional.DateModify,
	"Ago":            functional.DateAgo,
	"ToDate":         functional.ToDate,

	// Strings
	"Abbrev":     functional.Abbrev,
	"Abbrevboth": functional.Abbrevboth,
	"Trunc":      functional.Trunc,
	"Trim":       strings.TrimSpace,
	"Upper":      strings.ToUpper,
	"Lower":      strings.ToLower,
	"Title":      strings.Title,
	"Untitle":    functional.Untitle,
	"Substr":     functional.Substring,
	// Switch order so that "foo" | repeat 5
	"Repeat": func(count int, str string) string { return strings.Repeat(str, count) },
	// Deprecated: Use trimAll.
	"Trimall": func(a, b string) string { return strings.Trim(b, a) },
	// Switch order so that "$foo" | trimall "$"
	"TrimAll":      func(a, b string) string { return strings.Trim(b, a) },
	"TrimSuffix":   func(a, b string) string { return strings.TrimSuffix(b, a) },
	"TrimPrefix":   func(a, b string) string { return strings.TrimPrefix(b, a) },
	"Nospace":      util.DeleteWhiteSpace,
	"Initials":     functional.Initials,
	"RandAlphaNum": functional.RandAlphaNumeric,
	"RandAlpha":    functional.RandAlpha,
	"RandAscii":    functional.RandAscii,
	"RandNumeric":  functional.RandNumeric,
	"Swapcase":     util.SwapCase,
	"Shuffle":      xstrings.Shuffle,
	"Snakecase":    xstrings.ToSnakeCase,
	"Camelcase":    xstrings.ToCamelCase,
	"Kebabcase":    xstrings.ToKebabCase,
	"Wrap":         func(l int, s string) string { return util.Wrap(s, l) },
	"WrapWith":     func(l int, sep, str string) string { return util.WrapCustom(str, l, sep, true) },
	// Switch order so that "foobar" | contains "foo"
	"contains":   func(substr string, str string) bool { return strings.Contains(str, substr) },
	"HasPrefix":  func(substr string, str string) bool { return strings.HasPrefix(str, substr) },
	"HasSuffix":  func(substr string, str string) bool { return strings.HasSuffix(str, substr) },
	"Quote":      functional.Quote,
	"SQuote":     functional.SQuote,
	"cat":        functional.Cat,
	"Indent":     functional.Indent,
	"NIndent":    functional.NIndent,
	"Replace":    functional.Replace,
	"Plural":     functional.Plural,
	"Sha1sum":    functional.Sha1sum,
	"Sha256sum":  functional.Sha256sum,
	"Adler32sum": functional.Adler32sum,
	"ToString":   functional.StrVal,

	// Wrap Atoi to stop errors.
	"Atoi":    func(a string) int { i, _ := strconv.Atoi(a); return i },
	"Int64":   functional.ToInt64,
	"Int":     functional.ToInt,
	"Float64": functional.ToFloat64,

	// Split "/" foo/bar returns map[int]string{0: foo, 1: bar}
	"Split":     functional.Split,
	"SplitList": func(sep, orig string) []string { return strings.Split(orig, sep) },
	// Splitn "/" foo/bar/fuu returns map[int]string{0: foo, 1: bar/fuu}
	"Splitn":    functional.Splitn,
	"toStrings": functional.StrSlice,

	"Until":     functional.Until,
	"UntilStep": functional.UntilStep,

	// VERY basic arithmetic.
	"Add1": func(i interface{}) int64 { return functional.ToInt64(i) + 1 },
	"Add": func(i ...interface{}) int64 {
		var a int64 = 0
		for _, b := range i {
			a += functional.ToInt64(b)
		}
		return a
	},
	"Sub": func(a, b interface{}) int64 { return functional.ToInt64(a) - functional.ToInt64(b) },
	"Div": func(a, b interface{}) int64 { return functional.ToInt64(a) / functional.ToInt64(b) },
	"Mod": func(a, b interface{}) int64 { return functional.ToInt64(a) % functional.ToInt64(b) },
	"Mul": func(a interface{}, v ...interface{}) int64 {
		val := functional.ToInt64(a)
		for _, b := range v {
			val = val * functional.ToInt64(b)
		}
		return val
	},
	"Biggest": functional.Max,
	"Max":     functional.Max,
	"Min":     functional.Min,
	"Ceil":    functional.Ceil,
	"Floor":   functional.Floor,
	"Round":   functional.Round,

	"Join":      functional.Join,
	"SortAlpha": functional.SortAlpha,

	// Defaults
	"Default":      functional.Dfault,
	"IsEmpty":      functional.IsEmpty,
	"Coalesce":     functional.Coalesce,
	"Compact":      functional.Compact,
	"ToJson":       functional.ToJson,
	"ToPrettyJson": functional.ToPrettyJson,
	"Ternary":      functional.Ternary,

	// Reflection
	"ReflectTypeOf":     functional.TypeOf,
	"ReflecTypeIs":      functional.TypeIs,
	"ReflectTypeIsLike": functional.TypeIsLike,
	"ReflectKindOf":     functional.KindOf,
	"ReflectKindIs":     functional.KindIs,

	// OS:
	"Env":       func(s string) string { return os.Getenv(s) },
	"Expandenv": func(s string) string { return os.ExpandEnv(s) },

	// File Paths:
	"BasePath":  path.Base,
	"BaseDir":   path.Dir,
	"PathClean": path.Clean,
	"PathExt":   path.Ext,
	"PathIsAbs": path.IsAbs,

	// Encoding:
	"Base64enc": functional.Base64encode,
	"Base64dec": functional.Base64decode,
	"Base32enc": functional.Base32encode,
	"Base32dec": functional.Base32decode,

	// Data Structures:
	"List":   functional.List,
	"Dict":   functional.Dict,
	"Set":    functional.Set,
	"Unset":  functional.Unset,
	"HasKey": functional.HasKey,
	"Pluck":  functional.Pluck,
	"Keys":   functional.Keys,
	"Pick":   functional.Pick,
	"Omit":   functional.Omit,
	"Merge":  functional.Merge,
	"Values": functional.Values,

	"Append": functional.Push, "push": functional.Push,
	"Prepend": functional.Prepend,
	"First":   functional.First,
	"Rest":    functional.Rest,
	"Last":    functional.Last,
	"Initial": functional.Initial,
	"Reverse": functional.Reverse,
	"Uniq":    functional.Uniq,
	"Without": functional.Without,
	"Has":     functional.Has,
	"Slice":   functional.Slice,

	// Crypto:
	"GenPrivateKey":     functional.GeneratePrivateKey,
	"DerivePassword":    functional.DerivePassword,
	"BuildCustomCert":   functional.BuildCustomCertificate,
	"GenCA":             functional.GenerateCertificateAuthority,
	"GenSelfSignedCert": functional.GenerateSelfSignedCertificate,
	"GenSignedCert":     functional.GenerateSignedCertificate,

	// UUIDs:
	"uuidv4": functional.Uuidv4,

	// SemVer:
	"SemVer":        functional.SemVer,
	"SemVerCompare": functional.SemVerCompare,

	// Flow Control:
	"Fail": func(msg string) (string, error) { return "", errors.New(msg) },

	// Regex
	"RegexMatch":             functional.RegexMatch,
	"RegexFindAll":           functional.RegexFindAll,
	"RegexFind":              functional.RegexFind,
	"RegexReplaceAll":        functional.RegexReplaceAll,
	"RegExReplaceAllLiteral": functional.RegExReplaceAllLiteral,
	"RegExSplit":             functional.RegExSplit,
}
