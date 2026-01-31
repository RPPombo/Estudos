#include <iostream>
#include <string>
#include <map>
#include <vector>
using namespace std;

const string PALAVRA_SECRETA = "MELANCIA";
map <char, bool> chutou; // O map é uma estrutura de dicionário, funcionando com chave/valor
vector <char> chutes_errados; // O vector é um vetor dinâmico

bool letra_existe(char chute) {
    for (char letra : PALAVRA_SECRETA) { // A partir do C++11 é possível iterar uma palavra da maneira como está no código ao invés de for (int i=0; i < PALAVRA_SECRETA.size(); i++)
        if (chute == letra){
            return true;
        }
    }

    return false;
}

bool acertou_palavra() {
    for (char letra : PALAVRA_SECRETA) {
        if (!chutou[letra]) {
            return false;
        }
    }

    return true;
}

bool enforcou() {
    return chutes_errados.size() >= 5;
}

void mostra_erros() {
    cout << "Chutes errados: ";
        for (char letra : chutes_errados) {
            cout << letra << " ";
        }
    cout << endl;
}

void mostra_palavra() {
    for (char letra : PALAVRA_SECRETA) {
            if (chutou[letra]) {
                cout << letra << " ";
            } else {
                cout << "_ ";
            }
        }
    cout << endl;
}

void chuta() {
    char chute;
    cout << "Faça seu chute: ";
    cin >> chute;

    chutou[chute] = true;

    cout << "\nChute: " << chute << endl;

    if (letra_existe(chute)) {
        cout << "A letra existe na palavra" << endl;
    } else {
        cout << "A letra não existe na palavra" << endl;
        chutes_errados.push_back(chute);
    }
    cout << endl;
}

int main() {
    cout << "Bem-vindo ao jogo da forca!" << endl;

    while (!acertou_palavra() && !enforcou()) {
        mostra_erros();

        mostra_palavra();

        chuta();
    }
    if (enforcou()) {
        cout << "Que pena! Tente novamente!" << endl;
    } else {
        cout << "Parabéns! Você acertou a palavra secreta!" << endl;
    }

    cout << "Palavra secreta: " << PALAVRA_SECRETA << endl << endl;

    cout<< "Fim de Jogo!" << endl;
}

// Para facilitar a compilação do arquivo, foi criado um arquivo Makefile.
/* Explicação para o arquivo Makefile:
    1ª parte: definição de variáveis.
    2ª parte: criação do comando de compilação (all:). Nessa parte é escrito que o arquivo a ser compilado é o TARGET, e que ele necessita do arquivo SRC para isso acontecer, além do próprio comando para compilar.
    3ª parte: criação do comando de rodar o arquivo (run:). Nessa parte é escrito como o arquivo TARGET é rodado.
    4ª parte: craição do comando para apagar o arquivo (clean:). Está escrito o comando para apagar o arquivo TARGET.
*/