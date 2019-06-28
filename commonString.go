package main 
import ("fmt")
// func longestCommonPrefix(strs []string) string {
//     // 1. find minx 
//     return "okay"
// }

func chectStringInString(strs []string, str string) {

}
func getMinLength(strs []string) string {
		min := len(strs[0]) 
		index := 0
    for  i, str := range strs {
			if len(str) < min {
				min = len(str)
				index = i 
			}
    }
    return strs[index]
}

func main() {
	strs := []string{"gddddd", "hgggg", "ixx"}
	result := getMinLength(strs)
	fmt.Println(result)
} 