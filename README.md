# BMark

Autogenerate bookmarks for firefox using lua.

To generate bookmarks just create a lua file:
```lua
return {
  -- shortcut = "address"
  g = "github.com/search",
  gu = "github.com/search?type=users",
}
```

After that run:
`bmark file.lua`

And import the bookmarks from you download folder into firefox:
`Ctrl+Shift+O` > `Alt+I` > `I` > `gen<Down><Enter>` > `Ctrl+W`
