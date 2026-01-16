import random

def jogar() -> None:

    print('***************************')
    print('Bem-vindo ao jogo da forca!')
    print('***************************')

    arquivo = open("Palavras.txt", "r")
    palavras = list()
    for linha in arquivo:
        linha = linha.strip()
        linha = linha.upper()
        palavras.append(linha)

    arquivo.close()

    numero = random.randrange(0,len(palavras))
    palavra_secreta = palavras[numero]

    enforcou = False
    acertou = False
    palavra_jogador = list('_' for letra_secreta in palavra_secreta)
    erros = 0

    while not enforcou and not acertou:
        print(f'Erros:{erros}/7')
        posicao = 0

        print(''.join(palavra_jogador))

        chute = input('Letra:')
        chute = chute.upper()
        
        if chute in palavra_secreta:
            posicao = 0
            for letra in palavra_secreta:
                if chute == letra:
                    palavra_jogador[posicao] = letra
                posicao += 1
        else:
            erros += 1
            desenha_forca(erros)
        
        if erros == 7:
            print('Você perdeu')
            enforcou = True

        if not '_' in palavra_jogador:
            print('Você ganhou')
            acertou = True

    print('Fim de Jogo!')

if __name__ == "__main__":
    jogar()

def desenha_forca(erros):
    print("  _______     ")
    print(" |/      |    ")

    if(erros == 1):
        print(" |      (_)   ")
        print(" |            ")
        print(" |            ")
        print(" |            ")

    if(erros == 2):
        print(" |      (_)   ")
        print(" |      \     ")
        print(" |            ")
        print(" |            ")

    if(erros == 3):
        print(" |      (_)   ")
        print(" |      \|    ")
        print(" |            ")
        print(" |            ")

    if(erros == 4):
        print(" |      (_)   ")
        print(" |      \|/   ")
        print(" |            ")
        print(" |            ")

    if(erros == 5):
        print(" |      (_)   ")
        print(" |      \|/   ")
        print(" |       |    ")
        print(" |            ")

    if(erros == 6):
        print(" |      (_)   ")
        print(" |      \|/   ")
        print(" |       |    ")
        print(" |      /     ")

    if (erros == 7):
        print(" |      (_)   ")
        print(" |      \|/   ")
        print(" |       |    ")
        print(" |      / \   ")

    print(" |            ")
    print("_|___         ")
    print()