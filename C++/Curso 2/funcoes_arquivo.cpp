#include "funcoes_arquivo.hpp"

std::vector<std::string> le_arquivo() {
    std::ifstream arquivo;

    arquivo.open("palavras.txt"); // Abertura de aqrquivo

    if (!arquivo.is_open()){ // Verficando se o arquivo está aberto
        std::cout << "Erro na leitura do arquivo de palavras." << std::endl;
        exit(1);
    }

    int quantidade_palavras;
    arquivo >> quantidade_palavras;

    std::vector<std::string> lista;
    std::string palavra;

    for (int linha=0; linha<quantidade_palavras; linha++){
        arquivo >> palavra;
        lista.push_back(palavra);
    }

    arquivo.close();

    return lista;
}

void escrever_arquivo(std::vector <std::string> nova_lista) {
    std::ofstream arquivo;
    arquivo.open("palavras.txt");

    if (arquivo.is_open()){
        arquivo << nova_lista.size() << std::endl;

        for (std::string palavra : nova_lista){
            arquivo << palavra << std::endl;
        }
    } else {
        std::cout << "Erro na abertura do arquivo de palavras." << std::endl;
    }

    arquivo.close();
}

void adiciona_palavra() {
    std::cout << "Digite a nova palavra em letras maiúsculas:";
    std::string palavra_nova;
    std::cin >> palavra_nova;

    std::vector <std::string> lista_palavras = le_arquivo();

    lista_palavras.push_back(palavra_nova);

    escrever_arquivo(lista_palavras);
}