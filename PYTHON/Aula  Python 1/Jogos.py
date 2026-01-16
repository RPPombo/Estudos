import Adivinhacao
import Forca

while True:
    print("""
    Escolha o jogo que você quer jogar
    (1) Forca  (2) Adivinhação
        
    (0)Sair
    """)

    escolha = int(input("Escolha:"))

    match escolha:
        case 1:
            print('Jogando Forca')
            Forca.jogar()
        case 2:
            print('Jogando Adivinhação')
            Adivinhacao.jogar()
        case 0:
            break
        case _:
            print('Escolha uma opção válida')
            continue