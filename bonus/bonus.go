package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {

	torre := 0
	numero := 0

	for _, i := range barLengths {
		altura := barLengths[i]
		for i < len(barLengths) && barLengths[i] == altura {
			i++
		}
		torre = altura
		numero++
	}
	return torre, numero
}
