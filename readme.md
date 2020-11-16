# GenID
GenID is a small command line utility used to generate Version 4 UUIDs. I use it
with `xclip` so the generated UUID can be sent to the clipboard.

## Usage
```shell script
$ genid | xclip -selection c
```

This would put a UUID similar to `a1a1be12-9c64-440a-91a8-086229a1fc0d` into the
clipboard. If you want a UUID without the dashes between the segments, use the
`-plain` argument.

```shell script
$ genid -plain | xclip -selection c
```

This would put a UUID similar to `80346b0873a747b99777903359215165` into the
clipboard.
