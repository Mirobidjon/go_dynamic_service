package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func WriteUz(n string) string {
	if len(n) == 0 {
		return ""
	}
	var (
		t  int
		st string
	)
	sp := strings.Split(fmt.Sprint(n), ".")
	nu, _ := strconv.Atoi(sp[0])
	if len(sp) > 1 {
		if sp[1] != "00" {
			if len(sp[1]) == 1 && sp[1] != "0" {
				sp[1] += "0"
			}
			if sp[1] != "0" {
				t, _ = strconv.Atoi(sp[1])
				st = writeTyUz(t, "")
			}
		} else {
			st = "nol tiyin"
		}
	}

	return WriteNumUz(nu, "") + st
}

func WriteNumUz(n int, str string) string {
	if n < 1 {
		return ""
	}

	if n < 10 {
		return str + " " + uz9[n-1] + " so'm "
	}

	if n < 100 {
		st := uz90[n/10-1]
		if n%10 != 0 {
			st := WriteNumUz(n%10, st)
			return str + " " + st
		}

		return str + " " + st + " so'm "
	}

	if n < 1000 {
		st := uz9[n/100-1]
		if n%100 != 0 {
			st := WriteNumUz(n%100, st+" "+uz900[0])
			return str + " " + st
		}
		return str + " " + st + " " + uz900[0] + " so'm "

	}

	if n < 10000 {
		str := uz9[int(n/1000)-1]
		if int(n)%1000 != 0 {
			st := WriteNumUz(n%1000, str+" "+uz900[1])
			return st
		}
		return str + " " + uz900[1] + " so'm "
	}

	if n < 1000000 {
		st1 := WriteNumUz(n/1000, "")
		st1 = strings.ReplaceAll(st1, "so'm ", "")
		st2 := WriteNumUz(n%1000, "")
		return st1 + " " + uz900[1] + st2
	}

	if n < 1000000000 {
		var st string
		st1 := WriteNumUz(n/1000000, "")
		st1 = strings.ReplaceAll(st1, "so'm ", "")
		st2 := WriteNumUz(n%1000000, "")
		st2 = strings.ReplaceAll(st2, "so'm ", "")
		st = st1 + " " + uz900[2]
		if st2 != "" {
			st += " " + st2
		}

		return st + " so'm "
	}

	if n < 1000000000000 {
		var st string
		st1 := WriteNumUz(n/1000000000, "")
		st1 = strings.ReplaceAll(st1, "so'm ", "")
		st2 := WriteNumUz(n%1000000000, "")
		st2 = strings.ReplaceAll(st2, "so'm ", "")
		st = st1 + " " + uz900[3]
		if st2 != "" {
			st += " " + st2
		}

		return st + " so'm "
	}

	return ""
}

func writeTyUz(n int, str string) string {
	if n < 10 {
		return str + " " + uz9[n-1] + " tiyin "
	}

	if n < 100 {
		st := uz90[n/10-1]
		if n%10 != 0 {
			st := writeTyUz(n%10, st)
			return str + " " + st
		}

		return str + " " + st + " tiyin "
	}

	return ""
}

func WriteRu(n string) string {
	if len(n) == 0 {
		return ""
	}
	var (
		t  int
		st string
	)
	sp := strings.Split(fmt.Sprint(n), ".")
	nu, _ := strconv.Atoi(sp[0])
	if len(sp) > 1 {
		if sp[1] != "00" {
			if len(sp[1]) == 1 && sp[1] != "0" {
				sp[1] += "0"
			}
			if sp[1] != "0" {
				t, _ = strconv.Atoi(sp[1])
				st = writeTyRu(t, "")
			}
		} else {
			st = "нол тийин"
		}
	}

	return WriteNumRu(nu, "") + st
}

func WriteNumRu(n int, str string) string {
	if n < 1 {
		return ""
	}

	if n < 20 {
		return str + " " + ru19[n-1] + " сум "
	}

	if n < 100 {
		st := ru90[n/10-2]
		if n%10 != 0 {
			st := WriteNumRu(n%10, st)
			return str + " " + st
		}

		return str + " " + st + " сум "
	}

	if n < 1000 {
		st := ru900[n/100-1]
		if n%100 != 0 {
			st := WriteNumRu(n%100, st)
			return str + " " + st
		}

		return str + " " + st + " сум "
	}

	if n < 1000000 {
		st1 := WriteNumRu(n/1000, "")
		st1 = strings.ReplaceAll(st1, "сум ", "")
		st2 := WriteNumRu(n%1000, "")
		st2 = strings.ReplaceAll(st2, "сум ", "")
		return st1 + " " + ru0[0] + st2 + " сум "
	}

	if n < 1000000000 {
		st1 := WriteNumRu(n/1000000, "")
		st1 = strings.ReplaceAll(st1, "сум ", "")
		st2 := WriteNumRu(n%1000000, "")
		st2 = strings.ReplaceAll(st2, "сум ", "")
		return st1 + " " + ru0[1] + st2 + " сум "
	}

	if n < 1000000000000 {
		st1 := WriteNumRu(n/1000000000, "")
		st1 = strings.ReplaceAll(st1, "сум ", "")
		st2 := WriteNumRu(n%1000000000, "")
		st2 = strings.ReplaceAll(st2, "сум ", "")
		return st1 + " " + ru0[2] + st2 + " сум "
	}

	return ""
}

func writeTyRu(n int, str string) string {
	if n < 20 {
		return str + " " + ru19[n-1] + " тийин"
	}

	if n < 100 {
		st := ru90[n/10-2]
		if n%10 != 0 {
			st := writeTyRu(n%10, st)
			return str + " " + st
		}

		return str + " " + st + " тийин"
	}

	return ""
}

func WriteOz(n string) string {
	if len(n) == 0 {
		return ""
	}
	var (
		t  int
		st string
	)
	sp := strings.Split(fmt.Sprint(n), ".")
	nu, _ := strconv.Atoi(sp[0])
	if len(sp) > 1 {
		if sp[1] != "00" {
			if len(sp[1]) == 1 && sp[1] != "0" {
				sp[1] += "0"
			}
			if sp[1] != "0" {
				t, _ = strconv.Atoi(sp[1])
				st = writeTyOz(t)
			}
		} else {
			st = "нол тийин"
		}
	}
	return WriteNumOz(nu, "") + st
}

func writeTyOz(n int) string {

	if n < 10 && n != 0 {
		return oz9[n-1] + " тийин"
	}

	if n < 100 {
		st := oz90[n/10-1] + " "
		if n%10 != 0 {
			st += oz9[(n%10)-1]
		}
		return st + " тийин"
	}

	return ""
}

func WriteNumOz(n int, str string) string {
	if n < 1 {
		return ""
	}

	if n < 10 {
		return str + " " + oz9[n-1] + " сўм "
	}

	if n < 100 {
		st := oz90[n/10-1]
		if n%10 != 0 {
			st := WriteNumOz(n%10, st)
			return str + " " + st
		}

		return str + " " + st + " сўм "
	}

	if n < 1000 {
		st := oz9[n/100-1]
		if n%100 != 0 {
			st := WriteNumOz(n%100, st+" "+oz900[0])
			return str + " " + st
		}
		return str + " " + st + " " + oz900[0] + " сўм "

	}

	if n < 10000 {
		str := oz9[int(n/1000)-1]
		if int(n)%1000 != 0 {
			st := WriteNumOz(n%1000, str+" "+oz900[1])
			return st
		}
		return str + " " + oz900[1] + " сўм "
	}

	if n < 1000000 {
		st1 := WriteNumOz(n/1000, "")
		st1 = strings.ReplaceAll(st1, "сўм ", "")
		st2 := WriteNumOz(n%1000, "")
		return st1 + " " + oz900[1] + st2
	}

	if n < 1000000000 {
		var st string
		st1 := WriteNumOz(n/1000000, "")
		st1 = strings.ReplaceAll(st1, "сўм ", "")
		st2 := WriteNumOz(n%1000000, "")
		st2 = strings.ReplaceAll(st2, "сўм ", "")
		st = st1 + " " + oz900[2]
		if st2 != "" {
			st += " " + st2
		}

		return st + " сўм "
	}

	if n < 1000000000000 {
		var st string
		st1 := WriteNumOz(n/1000000000, "")
		st1 = strings.ReplaceAll(st1, "сўм ", "")
		st2 := WriteNumOz(n%1000000000, "")
		st2 = strings.ReplaceAll(st2, "сўм ", "")
		st = st1 + " " + oz900[3]
		if st2 != "" {
			st += " " + st2
		}

		return st + " сўм "
	}

	return ""
}

var uz9 = []string{"bir", "ikki", "uch", "to'rt", "besh", "olti", "yetti", "sakkiz", "to'qqiz"}

var uz90 = []string{"o'n", "yigirma", "o'ttiz", "qirq", "ellik", "oltmish", "yetmish", "sakson", "to'qson"}

var uz900 = []string{"yuz", "ming", "million", "milliard"}

var oz9 = []string{"бир", "икки", "уч", "тўрт", "беш", "олти", "етти", "саккиз", "тўққиз"}

var oz90 = []string{"ўн", "йигирма", "ўттиз", "қирқ", "эллик", "олтмиш", "етмиш", "саксон", "тўқсон"}

var oz900 = []string{"юз", "минг", "миллион", "миллиард"}

var ru19 = []string{"один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять",
	"одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать",
	"восемнадцать", "девятнадцать"}

var ru90 = []string{"двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}

var ru900 = []string{"сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}

var ru0 = []string{"тысяча", "миллион", "миллиард"}
