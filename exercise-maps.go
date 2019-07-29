package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    fields := strings.Fields(s)
    for i:=range(fields){
        word := fields[i]
        count:=0
        for j:=0; j<len(fields);j++{
            if fields[j]==word{
                count=count+1
            }
        }
        m[fields[i]] = count
    }
    return m
}

func main() {
    wc.Test(WordCount)
}


