#ifndef _FOGEFOGE_H_
#define _FOGEFOGE_H_

#include "mapa.h"
#include "ui.h"

#define CIMA 'w'
#define BAIXO 's'
#define ESQUERDA 'a'
#define DIREITA 'd'

#define BOMBA 'b'

void atribui_fantasmas();
void decide_direcao(POSICAO* fantasma);
void mover_fantasma();
void verifica_fantasmas(int x, int y);
int pode_explodir(int x, int y);
void usar_pilula();
int acabou();
void mover_heroi(char direcao);
int validacao_mover(int x, int y);
#endif