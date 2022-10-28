package styles

import (
	"github.com/alecthomas/chroma/v2"
)

// LaunchCode style.
var LaunchCode = Register(chroma.MustNewStyle("launchcode", chroma.StyleEntries{
	chroma.Background:        "#354a5f",
	chroma.GenericOutput:     "#ffffff bg:#354a5f",
	chroma.Keyword:           "#f1A640 bold",
	chroma.KeywordPseudo:     "nobold",
	chroma.LiteralNumber:     "#DF552D bold",
	chroma.NameTag:           "#FF98CE bold",
	chroma.NameVariable:      "#FF98CE",
	chroma.Comment:           "#E67F74 bg:#354a5f italic",
	chroma.NameAttribute:     "#FF98CE bold",
	chroma.LiteralString:     "#61C08D",
	chroma.NameFunction:      "#FF98CE bold",
	chroma.KeywordType:       "#f1A640 bold",
	chroma.NameConstant:      "#FF98CE",
	chroma.CommentPreproc:    "#E67F74 bold"
}))
