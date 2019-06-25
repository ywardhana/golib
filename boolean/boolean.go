package boolean

func ToInt(param bool) (res int) {
	res = 0
	if param {
		res = 1
	}
	return
}

func ToString(param bool) (res string) {
	res = "false"
	if param {
		res = "true"
	}
	return
}
