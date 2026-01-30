#include <iostream>
#include <string>
#include <cctype>
using namespace std;

const string PALAVRA_SECRETA = "MELANCIA";

bool letra_existe(char chute) {
    for (char letra : PALAVRA_SECRETA) { // A partir do C++11 é possível iterar uma palavra da maneira como está no código ao invés de for (int i=0; i < PALAVRA_SECRETA.size(); i++)
        if (chute == letra){
            return true;
        }
    }

    return false;
}

int main() {
    bool acertou = false;
    bool enforcou = false;

    while (!acertou && !enforcou <= 5) {
        char chute;

        cout << "Faça seu chute: ";
        cin >> chute;
        chute = toupper(chute);

        cout << "\nChute: " << chute << endl;

        if (letra_existe(chute)) {
            cout << "A letra existe na palavra" << endl;
        } else {
            cout << "A letra não existe na palavra" << endl;
        }
    }
}

// Para facilitar a compilação do arquivo, foi criado um arquivo Makefile.
/* Explicação para o arquivo Makefile:
    1ª parte: definição de variáveis.
    2ª parte: criação do comando de compilação (all:). Nessa parte é escrito que o arquivo a ser compilado é o TARGET, e que ele necessita do arquivo SRC para isso acontecer, além do próprio comando para compilar.
    3ª parte: criação do comando de rodar o arquivo (run:). Nessa parte é escrito como o arquivo TARGET é rodado.
    4ª parte: craição do comando para apagar o arquivo (clean:). Está escrito o comando para apagar o arquivo TARGET.
*/