#include "funcoes_palavras.hpp"

// A definição extern faz com que o compilador entenda que a váriavel não vai ser definida nesse arquivo
extern std::string palavra_secreta;
extern std::map <char, bool> chutou;

void escolhe_palavra() {
    std::vector <std::string> palavras = le_arquivo();

    srand(time(NULL));
    int indice = rand() % palavras.size();

    palavra_secreta = palavras[indice];
}

bool acertou_palavra() {
    for (char letra : palavra_secreta) {
        if (!chutou[letra]) {
            return false;
        }
    }

    return true;
}

void mostra_palavra() {
    for (char letra : palavra_secreta) {
            if (chutou[letra]) {
                std::cout << letra << " ";
            } else {
                std::cout << "_ ";
            }
        }
    std::cout << std::endl;
}

bool letra_existe(char chute) {
    for (char letra : palavra_secreta) { // A partir do C++11 é possível iterar uma palavra da maneira como está no código ao invés de for (int i=0; i < palavra_secreta.size(); i++)
        if (chute == letra){
            return true;
        }
    }

    return false;
}