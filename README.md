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

`<C-S-o><A-i>igen<Down><Cr><C-w>`
