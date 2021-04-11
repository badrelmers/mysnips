# update package
after editing the package it takes about 20min to be able to download it : https://stackoverflow.com/questions/58364988/go-module-not-downloading-latest-minor-version

If you are using the default public proxy (proxy.golang.org), it has a cache on the mapping from latest to a specific version.

temporarily set GONOPROXY to bypass the cache:
```
GONOPROXY=github.com/badrelmers/mysnips go get -u github.com/badrelmers/mysnips
```

# test
```
export GONOPROXY=github.com/badrelmers/mysnips
go mod init mysnipsexample ; go mod tidy ; go build 
```

cat main.go 
```go
package main

import(
    my "github.com/badrelmers/mysnips/go"
    "encoding/json"
    "fmt"
)


func main() {
    text := "text %s text\n"
    intt := 1234
    
    byte := []byte("byte\n")

    mapp := map[int]string{}
    mapp[1] = "map 1 %s test"
    mapp[2] = "map 2\n"

    jsonn,_ := json.Marshal(mapp)

    slicee := []string{"one", "two", "three"}
    slicee2 := []int{10, 200, 3000}


    my.Rm("testAppend.t")
    my.Append("testAppend.t",text)
    my.Append("testAppend.t",intt)
    my.Append("testAppend.t","\n")
    my.Append("testAppend.t",byte)
    my.Append("testAppend.t",mapp)
    my.Append("testAppend.t","\n")
    my.Append("testAppend.t",jsonn)
    my.Append("testAppend.t","\n")
    my.Append("testAppend.t",slicee)
    my.Append("testAppend.t",slicee2)

    // ##################################
    my.Cp("testAppend.t", "testAppend_copy.t")
    my.Mkdir("mkdirtest/2")
    my.Mv("testAppend_copy.t", "mkdirtest/2/testAppend_copy_Moved.t")
    
    // ##################################
    fmt.Println(my.INFOCB, "INFOC", my.ENDC)
    fmt.Println(my.WARNCB, "WARNC", my.ENDC)
    fmt.Println(my.ERRORCB, "ERRORC", my.ENDC)
    fmt.Println(my.HIDECB, "HIDEC", my.ENDC)
    fmt.Println(my.INFO2CB, "INFO2C", my.ENDC)
    fmt.Println(my.INFO3CB, "INFO3C", my.ENDC)
    fmt.Println(my.INFO4CB, "INFO4C", my.ENDC)

    fmt.Println(my.INFOC, "INFOC", my.ENDC)
    fmt.Println(my.WARNC, "WARNC", my.ENDC)
    fmt.Println(my.ERRORC, "ERRORC", my.ENDC)
    fmt.Println(my.HIDEC, "HIDEC", my.ENDC)
    fmt.Println(my.INFO2C, "INFO2C", my.ENDC)
    fmt.Println(my.INFO3C, "INFO3C", my.ENDC)
    fmt.Println(my.INFO4C, "INFO4C", my.ENDC)
}

```
