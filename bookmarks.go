package main

import (
	"fmt"
	"maps"

	lua "github.com/yuin/gopher-lua"
)

type (
	Bookmarks map[string]string
)

func LoadBookmarks(file string) (Bookmarks, error) {
	l := lua.NewState()
	defer l.Close()

	err := l.DoFile(file)
	if err != nil {
		return nil, err
	}

	luaBookmarks := l.ToTable(-1)
	if luaBookmarks == nil {
		return nil, fmt.Errorf("the value returned from the bookmarks file has to be a table")
	}

	bookmarks := ResolveStringMapRecursive(luaBookmarks)
	return bookmarks, nil
}

func ResolveStringMapRecursive(table *lua.LTable) map[string]string {
	result := make(map[string]string)
	table.ForEach(func(lvKey, lvValue lua.LValue) {
		lKey, ok := lvKey.(lua.LString)
		if !ok {
			return
		}
		switch lValue := lvValue.(type) {
		case lua.LString:
			result[lKey.String()] = lValue.String()
		case *lua.LTable:
			childResult := ResolveStringMapRecursive(lValue)
			maps.Insert(result, maps.All(childResult))
		}
	})
	return result
}
