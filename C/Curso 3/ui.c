#include <stdio.h>
#include <stdlib.h>
#include "ui.h"

char desenhofantasma[4][7] = {
    {" .-.  " },
    {"| OO| " },
    {"|   | " },
    {"'^^^' " }
};

char desenhoparede[4][7] = {
    {"......" },
    {"......" },
    {"......" },
    {"......" }
};

char desenhoheroi[4][7] = {
    {" .--. "  },
    {"/ _.-'"  },
    {"\\  '-." },
    {" '--' "  }
};

char desenhopilula[4][7] = {
    {"      "},
    {" .-.  "},
    {" '-'  "},
    {"      "}
};

char desenhovazio[4][7] = {
    {"      "},
    {"      "},
    {"      "},
    {"      "}
};


void imprime_parte(char desenho[4][7], int parte){
    printf("%s", desenho[parte]);
}

void imprime_mapa(MAPA* mapa){
    for (int i=0; i<(*mapa).linhas; i++){
        for (int parte = 0; parte<4; parte++){
            for (int j=0; j<mapa->colunas; j++){
                switch (mapa->matriz[i][j])
                {
                case FANTASMA:
                    imprime_parte(desenhofantasma, parte);
                    break;
                
                case HEROI:
                    imprime_parte(desenhoheroi, parte);
                    break;

                case PAREDE_H:
                case PAREDE_V:
                    imprime_parte(desenhoparede, parte);
                    break;

                case PILULA:
                    imprime_parte(desenhopilula, parte);
                    break;
                
                case VAZIO:
                    imprime_parte(desenhovazio, parte);
                    break;
                } 
            }
            printf("\n");
        }
    }
}