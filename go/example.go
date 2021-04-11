package main

import(
    my "github.com/badrelmers/mysnips/go"
    "encoding/json"
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


    my.Rm("met2.t")
    my.Append("met2.t",text)
    my.Append("met2.t",intt)
    my.Append("met2.t","\n")
    my.Append("met2.t",byte)
    my.Append("met2.t",mapp)
    my.Append("met2.t","\n")
    my.Append("met2.t",jsonn)
    my.Append("met2.t","\n")
    my.Append("met2.t",slicee)
    my.Append("met2.t",slicee2)

    // ##################################
    my.Cp("met2.t", "met2copy.t")
    my.Mkdir("mkdirtest/2")
    my.Mv("met2copy.t", "mkdirtest/2/met2copyMoved.t")
    
    my.Cp("mkdirtest", "mkdirtest2")

}
