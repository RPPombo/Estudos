#include "funcoes_palavras.hpp"

/* A definição extern faz com que o compilador entenda que a váriavel não vai ser definida nesse arquivo
extern std::string palavra_secreta;
extern std::map <char, bool> chutou;
*/
// O namespace serve para criar uma espécie de "pasta virtual", para que funções fiquem mais organizadas dentro de projetos e resolvendo problemas de nomes redundantes
namespace Forca {
    // Para melhorar o gerenciamento de memória é possível enviar por meio de uma referência o valor de uma variável sem precisar derefenciá-la ao enviar um pointeiro
    std::string escolhe_palavra() {
        std::vector <std::string> palavras = Forca::le_arquivo();

        srand(time(NULL));
        int indice = rand() % palavras.size();

        return palavras[indice];
    }

    bool acertou_palavra(const std::string& palavra_secreta, const std::map<char, bool>& chutou) {
        for (char letra : palavra_secreta) {
            if (chutou.find(letra) == chutou.end() || !(chutou.at(letra))) {
                return false;
            }
        }

        return true;
    }

    void mostra_palavra(const std::string& palavra_secreta, const std::map <char,bool>& chutou) {
        for (char letra : palavra_secreta) {
                if (!(chutou.find(letra) == chutou.end()) && chutou.at(letra)) {
                    std::cout << letra << " ";
                } else {
                    std::cout << "_ ";
                }
            }
        std::cout << std::endl;
    }

    bool letra_existe(const std::string& palavra_secreta ,char chute) {
        for (char letra : palavra_secreta) { // A partir do C++11 é possível iterar uma palavra da maneira como está no código ao invés de for (int i=0; i < palavra_secreta.size(); i++)
            if (chute == letra){
                return true;
            }
        }

        return false;
    }
}