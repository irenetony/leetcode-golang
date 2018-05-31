
func lengthOfLongestSubstring(s string) int {
    res := 0
    
    runeMap := make(map[rune]bool)
    
    runeArr := []rune(s)
    
    for i:=0;i<len(runeArr);i++{
        cur := 0
        for j:=i;j<len(runeArr);j++ {
            if _,ok := runeMap[runeArr[j]];!ok{
                cur++
                runeMap[runeArr[j]] = true
                if cur> res {
                    res =cur
                }
            }else{
                if cur> res{
                    res = cur
                }
                runeMap= make(map[rune]bool)
                break
            }
        }
    }
    
    return res
}
