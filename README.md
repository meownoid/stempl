# stempl – Simple Templates

`stempl` implements simple and fast named formatting.

### Installation
```bash
go get github.com/meownoid/stempl
```

### Examples
```go
result, err = stempl.Format(
    "The {{quick}} {brown} {fox} jumps over the lazy {dog}",
    map[string]string{
        "brown": "gray",
        "fox": "cat",
        "dog": "owl",
    },
)
if err != nil {
    panic(err)
}
// The {quick} gray cat jumps over the lazy owl
fmt.Println(result)
```
