package main

import "fmt"

func stringMatching(words []string) []string {
    result := make([]string, 0)
    isDuplicated := make(map[string]bool)
    for i:=0; i<len(words)-1; i++ {
        if _, ok := isDuplicated[words[i]]; ok {
            continue
        }
        for j:=i+1; j<len(words); j++ {
            target := isSubString(words[i], words[j])
            if target != "" {
                if _, ok :=  isDuplicated[target]; ok {
                    continue
                }
                isDuplicated[target] = true
                result = append(result, target)
            }
        }
    }
    return result
}

func isSubString(a, b string) string {
    x := []rune(a)
    lenX := len(x)
    y := []rune(b)
    lenY := len(y)
    if lenX > lenY {
        for i:=0; i<= lenX - lenY; i++ {
            if x[i] != y[0] {
                continue
            }
            temp := string(x[i:i+lenY])
            if temp == b {
                return b
            }
        }
        return ""
    } else if lenX < lenY {
        for j:=0; j<= lenY-lenX; j++ {
            if y[j] != x[0] {
                continue
            }
            temp := string(y[j:j+lenX])
            if temp == a {
                return a
            }
        }
        return ""
    } else {
        return ""
    }
}

func main() {
	array := []string{"ehvj","ehvjw","mmgeue","ehvjwy","hdqtlhtwhx","hdqtlhtwhxo","wkkxjsns","kohdqtlhtwhxfw","acftf","kohdqtlhtwhxodp","cwpeiowwms","ehvjk","mkdmsmxb","mehvjks","mxdoz","xnacftf","mgksgencwhk","hacftf","jdofko","mcwpeiowwms","x","pommgeuefd","kenptsaoyr","bmgksgencwhk","pcmgvojskh","xnacftfx","fpnpzvmckle","pjdofkone","ssheyptxddttxjm","xxnacftfxp"}
	fmt.Println(stringMatching(array))
}