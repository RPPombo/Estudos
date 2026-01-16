#include <stdio.h>
#include <string.h>
#include "forca.h" // Include com aspas servem para importarmos header files que estão presentes no computador e não no compilador
#include <time.h>
#include <stdlib.h>

void abertura() {
    printf("\n-----//-----\n");
    printf("Jogo de Forca\n");
    printf("-----//-----\n\n");
}

void fazer_chute(char chutes[26], int* chutes_dados){ // Para fazer com que a função receba um endereço de memória ao invés de uma variável, deve-se colocar o * depois do tipo da variável
    char chute;
    printf("Faça seu chute: ");
    scanf(" %c", &chute); // Por conta do ENTER poder ser considerado um CHAR, para ignorá-lo o  certo é colocar espaço %c
    printf("\n");

    chutes[*chutes_dados] = chute; // O * antes da variável faz com que acesse o conteúdo do endereço de memória
    (*chutes_dados)++;
}

int ja_chutou(char chutes[26], int chutes_dados, char letra){
    int achou = 0;

    for (int j=0; j< chutes_dados; j++){
        if (chutes[j] == letra){
            achou = 1;
            break;
        }
    }

    return achou;
}

void desenha_forca(char palavrasecreta[20], char chutes[26], int chutes_dados){
    int erros = chutes_errados(chutes_dados, palavrasecreta, chutes);

    printf("  _______       \n");
    printf(" |/      |      \n");
    printf(" |      %c%c%c  \n", (erros >=1 ? '(': ' '),(erros >=1 ? '_' : ' '),(erros >=1 ? ')' : ' ')); // Esse jeito de escrever o if é chamado de if ternário (True : False)
    printf(" |      %c%c%c  \n", (erros >=3 ? '/' : ' '),(erros >=2 ? '|' : ' '),(erros >=4 ? '\\' : ' '));
    printf(" |       %c     \n", (erros >=2 ? '|' : ' '));
    printf(" |      %c %c   \n", (erros >=5 ? '/' : ' '), (erros >=5 ? '\\' : ' '));
    printf(" |              \n");
    printf("_|___           \n");

    printf("\n\n");

    for (int i=0; i < strlen(palavrasecreta); i++){
        int achou = ja_chutou(chutes, chutes_dados, palavrasecreta[i]);

        if (!achou){
            printf("_ ");
        } else {
            printf("%c ", palavrasecreta[i]);
        }
    }
    printf("\n");
}

int chutes_errados(int chutes_dados, char palavrasecreta[20], char chutes[26]){
    int erros = 0;
    for (int i=0; i< chutes_dados; i++){
        int existe = 0;
        for (int j=0;j<strlen(palavrasecreta);j++) {
            if (palavrasecreta[j] == chutes[i]){
                existe = 1;
                break;
            }
        }
        if (!existe){
            erros++;
        }
    }
    return erros;
}

int ja_enforcou(int chutes_dados, char palavrasecreta[20], char chutes[26]) {
    int erros;
    erros = chutes_errados(chutes_dados, palavrasecreta, chutes);
    return erros >= 5;
}

int ja_ganhou(char palavrasecreta[20], int chutes_dados, char chutes[26]) {
    for (int i=0; i<strlen(palavrasecreta); i++){
        if (!ja_chutou(chutes, chutes_dados, palavrasecreta[i])) {
            return 0;
        }
    }
    return 1;
}

void escolhe_palavra(char palavrasecreta[20]) {
    FILE* file = fopen("palavras.txt", "r");
    if (!file) {
        printf("Arquivo não disponível\n\n");
        exit(1);
    }

    int quantidade_palavras;
    fscanf(file, "%d", &quantidade_palavras);

    srand(time(0));
    int numero_random = rand() % quantidade_palavras;

    for (int i = 0; i <= numero_random; i++) {
        fscanf(file, "%s", palavrasecreta);
    }

    fclose(file);
}


int main() {
    char palavrasecreta[20]; // Como não existe o tipo string em C, utilizamos um array de char para isso

    escolhe_palavra(palavrasecreta);

    int acertou = 0; 
    int enforcou = 0;
    char chutes[26]; // O array já é um ponteiro
    int chutes_dados = 0;

    abertura();
    
    do{
        desenha_forca(palavrasecreta, chutes, chutes_dados);

        fazer_chute(chutes, &chutes_dados); // Para utilizar o endereço de memória da variável, coloca-se o & antes dela

        enforcou = ja_enforcou(chutes_dados, palavrasecreta, chutes);

        acertou = ja_ganhou(palavrasecreta, chutes_dados, chutes);

    } while (!acertou && !enforcou);

    desenha_forca(palavrasecreta, chutes, chutes_dados);

    if (acertou) {
        printf("Parabéns, você ganhou!\n");
    } else {
        printf("Que pena, a palvra secreta era %s.\nTente novamente!\n", palavrasecreta);
    }

}

// Esse código está com os locais das variáveis dentro da função main() para o aprendizado sobre ponteiros poder ficar demonstrado, ao invés de apenas as utilizarmos como variáveis globais
// Para escrever em arquivos, utilizar a função fprintf(), e para posicionar o ponteiro dentro do arquivo utilizar fseek()
// outra função útil para a leitura de arquivos é a feof(), ela retorna se o ponteiro está no final do arquivo