#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "fogefoge.h"

MAPA mapa1;
POSICAO heroi;
POSICAO** fantasmas;
int pilula = 0;

int validacao_mover(int x, int y){
    if (x < 0 || x >= mapa1.colunas || y < 0 || y >= mapa1.linhas){ // Limite de mapa
        return 0;
    } else if (mapa1.matriz[y][x] == PILULA){ // Verificação de pílula
        pilula ++;
        return 1;
    } else if (mapa1.matriz[y][x] != VAZIO && mapa1.matriz[y][x] != HEROI){ // Local possível de andar
        return 0; 
    } else {
        return 1;
    }
}

int acabou(){
    int fantasmas_mortos = 0;
    for (int i=0; i<mapa1.fantasmas; i++){
        if (fantasmas[i]->x == -1 && fantasmas[i]->y == -1){
            fantasmas_mortos ++;
        }
    }

    int heroivivo = encontra_mapa(&mapa1, &heroi, HEROI, 1);

    if(fantasmas_mortos == mapa1.fantasmas) {
        printf("Você ganhou!\n");
    } else if (!heroivivo) {
        printf("Você perdeu!\n");
    }

    return (!heroivivo || fantasmas_mortos == mapa1.fantasmas);
}

void mover_heroi(char direcao){
    int x = heroi.x;
    int y = heroi.y;

    switch (direcao){
        case ESQUERDA:
            x--;
            break;
        case DIREITA:
            x++;
            break;
        case BAIXO:
            y++;
            break;
        case CIMA:
            y--;
            break;
        default:
            return;
    }
    
    if (!validacao_mover(x, y)){
        return;
    } else {
        mapa1.matriz[y][x] = HEROI;
        mapa1.matriz[heroi.y][heroi.x] = VAZIO;
        heroi.x = x;
        heroi.y = y;
    }

}

void libera_fantasmas(){
    for (int i=0; i<mapa1.fantasmas; i++){
        free(fantasmas[i]);
    }

    free(fantasmas);
}

void atribui_fantasmas() {
    fantasmas = malloc(sizeof(POSICAO*) * mapa1.fantasmas);

    for (int i = 0; i < mapa1.fantasmas; i++) {
        fantasmas[i] = malloc(sizeof(POSICAO));
        encontra_mapa(&mapa1, fantasmas[i], FANTASMA, i + 1);
    }
}

void decide_direcao(POSICAO* fantasma){
    int xatual = fantasma->x;
    int yatual = fantasma->y;

    int opcoes[4][2] = {
        {xatual, yatual+1},
        {xatual, yatual-1},
        {xatual+1, yatual},
        {xatual-1, yatual}
    };

    
    for (int i=0; i<10; i++){
        int opcao = rand() % 4;

        if (validacao_mover(opcoes[opcao][0], opcoes[opcao][1])){
            fantasma->x = opcoes[opcao][0];
            fantasma->y = opcoes[opcao][1];
            break;
        }
    }

}

int mover_fantasmas(){
    int matou =0;
    for (int i=0; i<mapa1.fantasmas; i++){
        if (fantasmas[i]->x != -1 && fantasmas[i]->y != -1){
            int xatual = fantasmas[i]->x;
            int yatual = fantasmas[i]->y;

            decide_direcao(fantasmas[i]);

            int xnovo = fantasmas[i]->x;
            int ynovo = fantasmas[i]->y;

            if (xatual != xnovo || yatual != ynovo){
                if (mapa1.matriz[ynovo][xnovo] == HEROI){
                    matou = 1;
                }
                mapa1.matriz[yatual][xatual] = VAZIO;
                mapa1.matriz[ynovo][xnovo] = FANTASMA;
            }
        }
    }
    return matou;
}

void verifica_fantasmas(int x, int y){
    for (int i=0; i<mapa1.fantasmas; i++){
        if (x == fantasmas[i]->x && y == fantasmas[i]->y){
            fantasmas[i]->x = -1;
            fantasmas[i]->y = -1;
            mapa1.matriz[y][x] = VAZIO;
        }
    }
}

int pode_explodir(int x, int y){
    return !(x < 0 || x >= mapa1.colunas || y < 0 || y >= mapa1.linhas) && (mapa1.matriz[y][x] != PAREDE_H && mapa1.matriz[y][x] != PAREDE_V);
}

void usar_pilula(){
    int x = heroi.x;
    int y = heroi.y;

    for (int dir = 0; dir < 4; dir++){
        int continuar = 1;

        for (int dist = 1; dist < 4 && continuar; dist++){
            int nx = x, ny = y;

            switch (dir){
                case 0: nx = x + dist; break; // direita
                case 1: nx = x - dist; break; // esquerda
                case 2: ny = y - dist; break; // cima
                case 3: ny = y + dist; break; // baixo
            }

            if (!pode_explodir(nx, ny)){
                continuar = 0;
                continue;
            }

            verifica_fantasmas(nx, ny);
            mapa1.matriz[ny][nx] = VAZIO;
        }
    }

    pilula--;
}


int main() {

    int matou = 0;

    srand(time(0));

    atribui_mapa(&mapa1);
    encontra_mapa(&mapa1, &heroi, HEROI, 1);
    atribui_fantasmas();

    do{
        matou = mover_fantasmas();

        printf("Pílulas: %d\n", pilula);
        imprime_mapa(&mapa1);

        if(matou) break;

        char comando;
        scanf(" %c", &comando);

        mover_heroi(comando);

        if (comando == BOMBA && pilula>0) usar_pilula(); // Para condicionais com ações simples pode-se usar sem chaves

    } while (!acabou());
    
    printf("Liberando memória\n");
    libera_fantasmas();
    libera_mapa(&mapa1);
}

// Para a alocação dinâmica de memória utiliza-se a função malloc() para a alocar e a função free() para desalocar
// A função sizeof() é utilizada em conjunto com as duas, pois o tamanho de cada tipo de variável pode mudar de plataforma para plataforma