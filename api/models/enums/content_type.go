package enums

import (
//    "fmt"
)

type ContentType int

const (
    Clip ContentType = iota 
    Twit                 
    Story               
)

func (s ContentType) String() string {
    return [...]string{"Clip", "Twit", "Story"}[s]
}

