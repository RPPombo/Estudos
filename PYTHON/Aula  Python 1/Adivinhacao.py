import random

def jogar() -> None:

    print('*********************************')
    print('Bem vindo ao jogo de adivinhação!')
    print('*********************************')

    numero_secreto = random.randrange(1, 101)
    tentativas = 20
    rodada = 1
    pontos = 1000

    print('Qual nível de dificuldade você quer?')
    print('''
    (1) Fácil
    (2) Médio
    (3) Dificil
    ''')
    dificuldade = input("Dificuldade:")

    if dificuldade == '1':
        tentativas = 20
    elif dificuldade == '2':
        tentativas = 10
    elif dificuldade == '3':
        tentativas = 5
    else:
        print("A dificuldade está na fácil")

    total_de_rodadas = tentativas

    while tentativas > 0:
        print(f"RODADA {rodada} DE {total_de_rodadas}")
        chute = int(input("Digite seu número de 1 a 100:"))

        acertou = chute == numero_secreto
        menor = chute < numero_secreto
        maior = chute > numero_secreto

        print("você digitou ",chute)
        if not (1 < chute > 100):
            print("Digite um número de 1 a 100")

        if acertou:
            print("Você acertou!")
            print(f"O seu total de pontos foi {pontos} pontos")
            break
        elif menor:
            print("Seu chute foi menor que o número")
        elif maior:
            print("Seu chute foi maior que o número")
        
        pontos_perdidos = abs(numero_secreto - chute)
        pontos = pontos - pontos_perdidos
        tentativas -= 1
        rodada += 1
        
        print()
        print(f"você tem {tentativas} tentativas")
        print()
        if tentativas == 0:
            print("Você perdeu!")
            print(f"A sua pontuação foi {pontos} pontos")

    print('Fim de jogo!')

if __name__ == "__main__":
    jogar()