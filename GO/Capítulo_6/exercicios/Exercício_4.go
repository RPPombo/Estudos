package exercicios

func Exercicio_4() {
	ano := 2005
	ano_atual := 2026

	for {
		println(ano)

		if ano == ano_atual {
			break
		} else {
			ano++
		}
	}
}
