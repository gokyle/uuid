package uuid

import (
	"fmt"
	"regexp"
	"testing"
)

const hexOctet = "[A-F0-9]"

var uuidV4RegexStr = fmt.Sprintf("^%s{8}-%s{4}-4%s{3}-[89AB]%s{3}-%s{12}$",
	hexOctet, hexOctet, hexOctet, hexOctet, hexOctet)
var uuidVersion4Regex = regexp.MustCompile(uuidV4RegexStr)

func TestGenerateV4(t *testing.T) {
	fmt.Println("[+] testing GenerateV4()")
	var u []byte
	var err error
	u, err = GenerateV4()
	if err != nil {
		fmt.Printf("[!] %s\n", err.Error())
		t.FailNow()
	} else if len(u) != 36 {
		fmt.Printf("[!] UUID length mismatch: %s\n", u)
		t.FailNow()
	} else if !uuidVersion4Regex.Match(u) {
                fmt.Printf("[!] invalid UUID: %s\n", u)
                t.FailNow()
        }
	fmt.Printf("\t[*] generate UUID: %s\n", string(u))
}

func TestGenerateV4String(t *testing.T) {
	fmt.Println("[+] testing GenerateV4String()")
	var u string
	var err error
	u, err = GenerateV4String()
	if err != nil {
		fmt.Printf("[!] %s\n", err.Error())
		t.FailNow()
	} else if len(u) != 36 {
		fmt.Printf("[!] UUID length mismatch: %s\n", u)
		t.FailNow()
	} else if !uuidVersion4Regex.MatchString(u) {
                fmt.Printf("[!] invalid UUID: %s\n", u)
                t.FailNow()
        }
	fmt.Printf("\t[*] generate UUID: %s\n", u)
}
