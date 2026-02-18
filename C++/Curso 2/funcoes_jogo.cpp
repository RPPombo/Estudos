#include "funcoes_jogo.hpp"

/* A definição extern faz com que o compilador entenda que a váriavel não vai ser definida nesse arquivo
extern std::map <char,bool> chutou;
extern std::vector <char> chutes_errados;
*/

// O namespace serve para criar uma espécie de "pasta virtual", para que funções fiquem mais organizadas dentro de projetos e resolvendo problemas de nomes redundantes
namespace Forca {
    // Para melhorar o gerenciamento de memória é possível enviar por meio de uma referência o valor de uma variável sem precisar derefenciá-la ao enviar um pointeiro
    void mostra_erros(const std::vector <char>& chutes_errados) {
        std::cout << "Chutes errados: ";
            for (char letra : chutes_errados) {
                std::cout << letra << " ";
            }
        std::cout << std::endl;
    }

    void chuta(std::map <char, bool>& chutou, std::vector <char>& chutes_errados, const std::string& palavra_secreta) {
        char chute;
        std::cout << "Faça seu chute: ";
        std::cin >> chute;

        chutou[chute] = true;

        std::cout << "\nChute: " << chute << std::endl;

        if (letra_existe(palavra_secreta ,chute)) {
            std::cout << "A letra existe na palavra" << std::endl;
        } else {
            std::cout << "A letra não existe na palavra" << std::endl;
            chutes_errados.push_back(chute);
        }
        std::cout << std::endl;
    }
}