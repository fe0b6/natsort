package natsort

import (
	"sort"
	"strconv"
	"strings"
)

// Sort - Сортируем массив
func Sort(arr []string) {
	_sort(arr, true, false)
}

// Reverse - Сортируем массив в обратном порядке
func Reverse(arr []string) {
	_sort(arr, false, true)
}

func _sort(arr []string, tb, fb bool) {
	sort.Slice(arr, func(i, j int) bool {
		return less(arr[i], arr[j], tb, fb)
	})
}

// Проверка число это или нет
func isNum(s string) bool {
	if s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" {
		return true
	}

	return false
}

// Less - определяем позицию элемента
func Less(str1, str2 string) bool {
	return less(str1, str2, true, false)
}

func less(str1, str2 string, tb, fb bool) (ok bool) {
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")

	var k1, k2 int
	for {
		if len(arr1) < k1+1 {
			break
		}

		if len(arr2) < k2+1 {
			return fb
		}

		v1 := arr1[k1]
		v2 := arr2[k2]
		var isNum1, isNum2 bool
		if isNum(v1) {
			isNum1 = tb
		}
		if isNum(v2) {
			isNum2 = tb
		}

		// Число всегда вперед
		if isNum1 && !isNum2 {
			break
		}

		// Если это не числа
		if !isNum1 {
			if v1 < v2 {
				break
			} else if v1 > v2 {
				return fb
			}

			k1++
			k2++
			continue
		}

		nstr1 := []string{}
		for k := k1; k < len(arr1); k++ {
			if isNum(arr1[k]) {
				nstr1 = append(nstr1, arr1[k])
				continue
				k1 = k
			}
			k1 = k - 1
			break
		}

		nstr2 := []string{}
		for k := k2; k < len(arr2); k++ {
			if isNum(arr2[k]) {
				nstr2 = append(nstr2, arr2[k])
				k2 = k
				continue
			}
			k2 = k - 1
			break
		}

		num1, _ := strconv.ParseInt(strings.Join(nstr1, ""), 10, 64)
		num2, _ := strconv.ParseInt(strings.Join(nstr2, ""), 10, 64)

		if num1 < num2 {
			break
		} else if num1 > num2 {
			return fb
		}

		k1++
		k2++
	}

	return tb
}
