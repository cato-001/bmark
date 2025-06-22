package main

import (
	"maps"
	"os"

	"github.com/cato-001/junkbox/out"

	arg "github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		LuaFile []string `arg:"positional" help:"the lua file to generate the bookmarks from"`
	}

	parser, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		out.Ejsonln(err)
		return
	}
	parser.MustParse(os.Args[1:])

	bookmarks := make(Bookmarks, 0)
	for _, file := range args.LuaFile {
		fileBookmarks, err := LoadBookmarks(file)
		if err != nil {
			out.Ejsonln(err)
			return
		}
		maps.Copy(bookmarks, fileBookmarks)
	}

	output := "~/downloads"
	if IsWslActive() {
		output = "/mnt/c/Users/lukas/Downloads"
	}

	output += "/generated-bookmarks.html"

	content := GenerateFirefoxBookmarks(bookmarks)
	err = os.WriteFile(output, []byte(content), 0o777)
	if err != nil {
		out.Ejsonln(err)
		return
	}

	out.Jsonln(map[string]any {
		"output": output,
	})
}

func IsWslActive() bool {
	_, ok := os.LookupEnv("WSLENV")
	return ok
}
