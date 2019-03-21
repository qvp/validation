package validation

//todo тут обертки для вызова валидаторов ф функциональном стиле
// валидаторы с параметрами переименовать с маленькой буквы а обертки сделать доступными

func Min(params ...interface{}) Rule {
	return Rule{Name: "min", Params: params}
}
