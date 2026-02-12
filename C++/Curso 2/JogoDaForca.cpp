#include "JogoDaForca.hpp"

std::string palavra_secreta;
std::map <char,bool> chutou;
std:: vector <char> chutes_errados;

int main() {
    std::cout << "Bem-vindo ao jogo da forca!" << std::endl;

    escolhe_palavra();

    while (!acertou_palavra() && !enforcou()) {
        mostra_erros();

        mostra_palavra();

        chuta();
    }

    std::cout << "Palavra secreta: " << palavra_secreta << std::endl << std::endl;

    if (enforcou()) {
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