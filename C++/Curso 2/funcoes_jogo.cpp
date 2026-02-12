#include "funcoes_jogo.hpp"

// A definição extern faz com que o compilador entenda que a váriavel não vai ser definida nesse arquivo
extern std::map <char,bool> chutou;
extern std::vector <char> chutes_errados;

bool enforcou() {
    return chutes_errados.size() >= 5;
}

void mostra_erros() {
    std::cout << "Chutes errados: ";
        for (char letra : chutes_errados) {
            std::cout << letra << " ";
        }
    std::cout << std::endl;
}

void chuta() {
    char chute;
    std::cout << "Faça seu chute: ";
    std::cin >> chute;

    chutou[chute] = true;

    std::cout << "\nChute: " << chute << std::endl;

    if (letra_existe(chute)) {
        std::cout << "A letra existe na palavra" << std::endl;
    } else {
        std::cout << "A letra não existe na palavra" << std::endl;
        chutes_errados.push_back(chute);
    }
    std::cout << std::endl;
}