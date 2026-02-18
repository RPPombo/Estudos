#include "JogoDaForca.hpp"

// Essa parte serve para que o compilador procure funções que ele desconhece por não estarem em arquivos do projeto, em namespaces específicos
// É possível usar namespace a = std
using namespace Forca;

// O static nesse contexto faz com que as variáveis globais existam apenas na unidade de tradução que elas estão, ou seja ela não "existem" para outros arquivos
static std::string palavra_secreta;
static std::map <char,bool> chutou;
static std:: vector <char> chutes_errados;

int main() {
    imprime_cabecalho();

    palavra_secreta = escolhe_palavra();

    while (!acertou_palavra(palavra_secreta, chutou) && !enforcou(chutes_errados)) {
        mostra_erros(chutes_errados);

        mostra_palavra(palavra_secreta, chutou);

        chuta(chutou, chutes_errados, palavra_secreta);
    }

    std::cout << "Palavra secreta: " << palavra_secreta << std::endl << std::endl;

    if (enforcou(chutes_errados)) {
        std::cout << "Que pena! Tente novamente!" << std::endl;
    } else {
        std::cout << "Parabéns! Você acertou a palavra secreta!" << std::endl;
        std::cout << "Deseja adicionar uma nova palavra? (S/N)" << std::endl;
        char resposta;
        std::cin >> resposta;

        if (resposta == 'S') {
            adiciona_palavra();
        }
    }

    std::cout<< "Fim de Jogo!" << std::endl;
}

// Para facilitar a compilação do arquivo, foi criado um arquivo Makefile.
/* Explicação para o arquivo Makefile:
    1ª parte: definição de variáveis.
    2ª parte: criação do comando de compilação (all:). Nessa parte é escrito que o arquivo a ser compilado é o TARGET, e que ele necessita do arquivo SRC para isso acontecer, além do próprio comando para compilar.
    3ª parte: criação do comando de rodar o arquivo (run:). Nessa parte é escrito como o arquivo TARGET é rodado.
    4ª parte: craição do comando para apagar o arquivo (clean:). Está escrito o comando para apagar o arquivo TARGET.
*/