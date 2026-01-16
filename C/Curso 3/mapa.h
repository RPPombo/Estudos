#ifndef _MAPA_H_
#define _MAPA_H_

#define HEROI '@'
#define VAZIO '.'
#define FANTASMA 'F'
#define PILULA 'P'
#define PAREDE_H '-'
#define PAREDE_V '|'

// A struct serve para criar um conjunto de variáveis, para ser mais fácil de organizar
struct mapa {
    // Para declarar uma matriz basta colocar a quantidade de linhas e colunas
    char** matriz;
    int linhas;
    int colunas;
    int fantasmas;
};

// O typedef serve para criar um "apelido" para uma struct
typedef struct mapa MAPA;

struct posicao {
    int x;
    int y;
};

typedef struct posicao POSICAO;

void libera_mapa(MAPA* mapa);
FILE* le_mapa();
void atribui_mapa(MAPA* mapa);
int encontra_mapa(MAPA* mapa, POSICAO* p, char c, int ocorrencia);
#endif