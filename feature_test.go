package rabc

import (
	"testing"
	"fmt"
)

var (
	FRemoveMsgTail = 0x01
	FWebHook       = 0x02
	FWebStats      = 0x04
	//12个功能点 二级制 12个1
	FeatureMask = 0xfff
)

func Test_Featue(t *testing.T) {
	userFeatue := Feature{}
	fmt.Printf("xx %d \r\n", FRemoveMsgTail|FWebHook)
	fmt.Printf("user has %b \r\n", userFeatue.Has(FRemoveMsgTail|FWebHook, FRemoveMsgTail))
	all := userFeatue.Grant(FRemoveMsgTail|FWebHook, FWebStats)
	fmt.Printf("user grant %d \r\n", all)
	fmt.Printf("user grant %d \r\n", userFeatue.Revoke(all, FWebStats|FRemoveMsgTail|FWebHook))
}
