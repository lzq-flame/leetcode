/**
 * @Author: liuzhiqi
 * @Data: 2021/11/18 10:34
 */

package main

//medium46 子母异位词分组
func groupAnagrams(strs []string) [][]string {
	hash := map[[26]int][]string{}
	for _, s := range strs {
		cnt := [26]int{}
		for _, a := range s {
			cnt[a-'a']++
		}
		hash[cnt] = append(hash[cnt], s)
	}
	ans := make([][]string, 0, len(hash))
	for _, value := range hash {
		ans = append(ans, value)
	}
	return ans
}
