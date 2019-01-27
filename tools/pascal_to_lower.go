/**
*  Version:1.0
 * Copyright (C), 2019
 * @author Luobiken
 * @email:26015696@qq.com
 * @Tel/Weixin:13078095006
 * @Date  2019-01 11:37
 * Description：
*/
package tools

import (
	"unicode"
)

//把驼峰命名规则改成全部小写的下横杠的命名规则
func PascalToLower(structName string) string {
	//log.Println("转换前:",structName)
	runes := []rune(structName)
	if len(runes) <= 0 {
		return structName
	}
	newRunes := make([]rune, 0)
	newRunes = append(newRunes, unicode.ToLower(runes[0]))
	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			newRunes = append(newRunes, rune('_'))
			newRunes = append(newRunes, unicode.ToLower(runes[i]))
		} else {
			newRunes = append(newRunes, runes[i])
		}
	}
	//log.Println("转换后:",string(newRunes))
	return string(newRunes)
}
