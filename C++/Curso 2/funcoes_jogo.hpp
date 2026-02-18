#pragma once
#include <iostream>
#include <string>
#include <vector>
#include <map>
#include "funcoes_palavras.hpp"

// O namespace serve para criar uma espécie de "pasta virtual", para que funções fiquem mais organizadas dentro de projetos e resolvendo problemas de nomes redundantes
namespace Forca {
    // Uma função inline serve para que o compilador decida se é melhor deixá-la em um espaço de memória separado do arquivo que está chamando-a ou copiá-la diretamente no mesmo espaço de memória do arquivo
    // Isso faz com que o programa seja otimizado
    inline void imprime_cabecalho(){
        std::cout << "*******************************" << std::endl;
        std::cout << "**Bem-vindo ao Jogo da Forca!**" << std::endl;
        std::cout << "*******************************" << std::endl;
    }
    inline bool enforcou(std::vector <char> chutes_errados) {
        return chutes_errados.size() >= 5;
    }

    void mostra_erros(const std::vector <char>& chutes_errados);
    void chuta(std::map <char, bool>& chutou, std::vector <char>& chutes_errados, const std::string& palavra_secreta);
}