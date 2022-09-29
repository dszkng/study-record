package main

import "testing"

func TestSubtest(t *testing.T){
    t.Run("aa", func(t *testing.T){
        if res := subadd(1,2); res !=3 {
            t.Fatal("fatal aa")
        }
    })

    t.Run("bb", func(t *testing.T){
        if res := subadd(1,3); res !=5 {
            t.Fatal("fatal bb")
        }
    })
}
