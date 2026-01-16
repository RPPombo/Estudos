#include <stdio.h>
#include <stdlib.h>
#include "mapa.h"

FILE* le_mapa(){
    FILE* f;
    f = fopen("mapa.txt", "r");
    if (f == 0) {
        printf("Arquivo de mapa não encontrado!\n");
        exit(1);
    } else{
        return f;
    }
}

void atribui_mapa(MAPA* mapa) {
    FILE* f;
    f = le_mapa();

    fscanf(f, "%d %d %d", &mapa->linhas, &mapa->colunas, &mapa->fantasmas); // Há duas formas de acessarmos o conteúdo de uma struct: (*mapa).colunas ou mapa->colunas

    (*mapa).matriz = malloc(sizeof(char*) * (*mapa).linhas);
    for (int i=0; i<(*mapa).linhas; i++){
        (*mapa).matriz[i] = malloc(sizeof(char) * ((*mapa).colunas+1));
    }

    for (int i=0; i<(*mapa).linhas; i++){
        fscanf(f, "%s", (*mapa).matriz[i]);
    }

    fclose(f);
}

void libera_mapa(MAPA* mapa){
    for (int i=0; i< (*mapa).linhas; i++){
        free((*mapa).matriz[i]);
    }

    free(mapa);
}



int encontra_mapa(MAPA* mapa, POSICAO* p, char c, int ocorrencia){
    int contador = 0;

    for (int i=0; i<mapa->linhas; i++){
        for (int j=0; j<mapa->colunas; j++){
            if (mapa->matriz[i][j] == c){
                contador++;
                if (contador == ocorrencia){
                    p->x = j;
                    p->y = i;
                    return 1;
                }  
            }
        }
    }

    return 0;
}

// Outra forma de externalizar uma váriavel, além de enviar o ponteiro dela é utilizando o comando extern nela, quando ela for declarada