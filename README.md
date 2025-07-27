# BMark

Autogenerate bookmarks for firefox using lua.

To generate bookmarks just create a lua file:
```lua
return {
  -- shortcut = "address"
  g = "github.com/search",
  gu = "github.com/search?type=users",

  -- sub tables are only used for organisation
  -- they are flattend in the actual output, and do not change it
  sub_table = {
    mail = "mail.google.com",
    maps = "maps.google.com",
    tl = "translate.google.com",
  },
}
```

After that run:
`bmark file.lua`

And import the bookmarks from you download folder into firefox:

`<C-S-o><A-i>igen<Down><Cr><C-w>`
