package expense

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if err := validateExpression(expression); err != nil {
		return 0, errors.New("Expression is not valid")
	}
	var sl = []string{}
	var l string
	var curEl string
	for _, el := range expression {

		curEl = string(el)
		if curEl == "*" || curEl == "/" || curEl == "-" || curEl == "+" || curEl == "(" || curEl == ")" {
			if l != "" {
				sl = append(sl, l)

			}
			sl = append(sl, curEl)
			l = ""
		} else {
			l += curEl
		}
		// log.Println("Calc", sl)
	}
	if curEl == "*" || curEl == "/" || curEl == "-" || curEl == "+" || curEl == "(" || curEl == ")" {
		if l != "" {
			sl = append(sl, l)

		}

		sl = append(sl, curEl)
		l = ""
	} else {
		if l != "" {
			sl = append(sl, l)

		}

	}

	log.Println("ИСходник", strings.Join(sl, ","))
	for {
		var ok bool
		sl, ok = inBracket(sl)
		if !ok {
			break
		}
	}
	str, _ := mainCalc(sl)
	in1, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Printf("Ошибка конвертации 1 числа\n")
		return 0, nil
	}
	return in1, nil
}
func validateExpression(expression string) error {
	lastWasDigit := false
	parensCount := 0

	for i, char := range expression {
		if (char >= '0' && char <= '9') || char == '.' {
			lastWasDigit = true
		} else {
			if lastWasDigit == false {
				if char != '(' {
					return errors.New("некорректный символ перед: " + string(char))
				}
			}
			switch char {
			case '+', '-', '*', '/':
				if i == 0 || !lastWasDigit {
					return errors.New("некорректный оператор: " + string(char))
				}
				lastWasDigit = false

			case '(':
				parensCount++

			case ')':
				parensCount--
				if parensCount < 0 {
					return errors.New("несоответствующие скобки")
				}

			default:
				return errors.New("неизвестный символ: " + string(char))
			}
		}
	}

	if parensCount != 0 {
		return errors.New("несоответствующие скобки")
	}

	if !lastWasDigit {
		return errors.New("выражение должно заканчиваться цифрой")
	}

	return nil
}
func mainCalc(k []string) (string, bool) {
	var ok bool
	for {
		k, ok = priority(k)
		if !ok {
			return k[0], true
		}
	}
}
func priority(z []string) ([]string, bool) {
	log.Println(strings.Join(z, ","))
	for i, el := range z {
		if el == "*" || el == "/" {
			return run(z, i), true
		}
	}
	for i, el := range z {
		if el == "+" || el == "-" {
			return run(z, i), true
		}
	}
	return z, false
}
func run(z []string, i int) []string {
	res := calcularion(z[i-1], z[i], z[i+1])
	z[i-1] = res
	d := slices.Delete(z, i, i+2)
	return d
}
func calcularion(n1, sign, n2 string) string {
	in1, err := strconv.ParseFloat(n1, 64)
	if err != nil {
		log.Printf("Ошибка конвертации 1 числа\n")
		return ""
	}
	in2, err := strconv.ParseFloat(n2, 64)
	if err != nil {
		log.Printf("Ошибка конвертации 2 числа\n")
		return ""
	}
	switch sign {
	case "+":
		res := in1 + in2
		result := strconv.FormatFloat(res, 'f', 2, 64)
		return result
	case "-":
		res := in1 - in2
		result := strconv.FormatFloat(res, 'f', 2, 64)
		return result
	case "*":
		res := in1 * in2
		result := strconv.FormatFloat(res, 'f', 2, 64)
		return result
	case "/":
		if in2 == 0 {
			log.Println("Деление на ноль")
			return ""
		}
		res := in1 / in2
		result := strconv.FormatFloat(res, 'f', 2, 64)
		return result
	default:
		log.Println("Неизвестная операция")
		return ""
	}
}
func inBracket(sl []string) ([]string, bool) {
	var q []string
	var bracket1 int
	insideBrackets := false
	for i, el := range sl {
		if el == "(" {
			bracket1 = i
			insideBrackets = true
			continue
		}
		if el == ")" {
			str, _ := mainCalc(q)
			sl[bracket1] = str
			log.Println("Начало", strings.Join(sl, ","))
			sl = slices.Delete(sl, bracket1+1, i+1)
			log.Println("Конец", strings.Join(sl, ","))
			insideBrackets = false
			return sl, true
		}
		if insideBrackets {
			q = append(q, el)
			continue
		}
	}
	return sl, false
}
