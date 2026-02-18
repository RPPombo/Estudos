#pragma once
#include <iostream>
#include <vector>
#include <string>
#include <fstream>

// O namespace serve para criar uma espécie de "pasta virtual", para que funções fiquem mais organizadas dentro de projetos e resolvendo problemas de nomes redundantes
namespace Forca {
    std::vector<std::string> le_arquivo();

    void escrever_arquivo(std::vector <std::string> nova_lista);

    void adiciona_palavra();
}