package enums

import (
//    "fmt"
)

type PlatformType int

const (
    YouTube PlatformType = iota 
    TikTok                 
    Instagram               
)

func (s PlatformType) String() string {
    return [...]string{"YouTube", "TikTok", "Instagram"}[s]
}

