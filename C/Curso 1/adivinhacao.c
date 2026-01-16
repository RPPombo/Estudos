#include <stdio.h> // o include é a maneira utilizada para importar bibliotecas
#include <stdlib.h>
#include <time.h>

// caso haja a necessidade de criar uma constante utiliza-se #define, o nome da constante é colocado tudo em maiúsculo

int numerosecreto;
int chute;
unsigned int acertou;
double pontos = 1000;
int tentativas;
int nivel;


int main(){
    printf("Bem vindo ao jogo de adivinhação!\n");
    
    printf("Qual o nível de dificuldade desejada?\n");
    printf("1: Fácil\n");
    printf("2: Médio\n");
    printf("3: Difícil\n");
    printf("Escolha:");
    scanf("%d", &nivel);
    printf("\n");

    switch (nivel) {
        case 1:
            tentativas = 20;
            break;
        case 2:
            tentativas = 15;
            break;

        default:
            tentativas = 6;
            break; 
    }

    int seed = time(0); // a função time() serve para retornar o número de segundos desde janeiro de 1970(EPOCH) e como esse tempo muda constantemente utilizaremos ele como seed para o número secreto
    srand(seed);

    numerosecreto = rand() % 100; // para o número secreto ser sempre entre 0 e 99, pegaremos o resto da divisão entre o número aleatório e o número 100

    while (tentativas> 0){
        printf("Tentativas: %d\n", tentativas);
    
        printf("Faça o seu chute: ");
        scanf("%d", &chute);
        printf("Seu chute foi: %d\n", chute);

        acertou = chute == numerosecreto;

        if (acertou){
            printf("Parabéns! Você acertou!\n");
            break;
        } else {
            printf("Você errou! Tente novamente!\n");
            pontos -= abs((double)((chute - numerosecreto)/2)); // esse (double) é utilizado para fazer o casting das variáveis inteiras para elas se comportarem como um double
            tentativas -= 1;

            printf("Dica: ");
            if (chute > numerosecreto){
                printf("Seu chute foi maior que o número secreto.\n");
            } else if (chute < numerosecreto) {
                printf("Seu chute foi menor que o número secreto.\n");
            }
        }
    }

    printf("Fim de jogo!\n");
    printf("O número secreto era: %d\n", numerosecreto);
    printf("Você fez %.1f pontos!\n", pontos);

    return 0;
}