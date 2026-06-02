package valueobject

import "regexp"

func OnlyDigits(value string) string {
	//regex que busca por qualquer caractere que não seja um dígito
	regex := regexp.MustCompile(`\D`)
	return regex.ReplaceAllString(value, "")
}
