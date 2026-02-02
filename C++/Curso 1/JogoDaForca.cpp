#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <fstream>
#include <cstdlib>
#include <ctime>
using namespace std;

string palavra_secreta;
map <char, bool> chutou; // O map é uma estrutura de dicionário, funcionando com chave/valor
vector <char> chutes_errados; // O vector é um vetor dinâmico

bool letra_existe(char chute) {
    for (char letra : palavra_secreta) { // A partir do C++11 é possível iterar uma palavra da maneira como está no código ao invés de for (int i=0; i < palavra_secreta.size(); i++)
        if (chute == letra){
            return true;
        }
    }

    return false;
}

bool acertou_palavra() {
    for (char letra : palavra_secreta) {
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
    for (char letra : palavra_secreta) {
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

vector<string> le_arquivo() {
    ifstream arquivo;

    arquivo.open("palavras.txt"); // Abertura de aqrquivo

    if (!arquivo.is_open()){ // Verficando se o arquivo está aberto
        cout << "Erro na leitura do arquivo de palavras." << endl;
        exit(1);
    }

    int quantidade_palavras;
    arquivo >> quantidade_palavras;

    vector<string> lista;
    string palavra;

    for (int linha=0; linha<quantidade_palavras; linha++){
        arquivo >> palavra;
        lista.push_back(palavra);
    }

    arquivo.close();

    return lista;
}

void escolhe_palavra() {
    vector <string> palavras = le_arquivo();

    srand(time(NULL));
    int indice = rand() % palavras.size();

    palavra_secreta = palavras[indice];
}

void escrever_arquivo(vector <string> nova_lista) {
    ofstream arquivo;
    arquivo.open("palavras.txt");

    if (arquivo.is_open()){
        arquivo << nova_lista.size() << endl;

        for (string palavra : nova_lista){
            arquivo << palavra << endl;
        }
    } else {
        cout << "Erro na abertura do arquivo de palavras." << endl;
    }

    arquivo.close();
}

void adiciona_palavra(){
    cout << "Digite a nova palavra em letras maiúsculas:";
    string palavra_nova;
    cin >> palavra_nova;

    vector <string> lista_palavras = le_arquivo();

    lista_palavras.push_back(palavra_nova);

    escrever_arquivo(lista_palavras);
}

int main() {
    cout << "Bem-vindo ao jogo da forca!" << endl;

    escolhe_palavra();

    while (!acertou_palavra() && !enforcou()) {
        mostra_erros();

        mostra_palavra();

        chuta();
    }

    cout << "Palavra secreta: " << palavra_secreta << endl << endl;

    if (enforcou()) {
        cout << "Que pena! Tente novamente!" << endl;
    } else {
        cout << "Parabéns! Você acertou a palavra secreta!" << endl;
        cout << "Deseja adicionar uma nova palavra? (S/N)" << endl;
        char resposta;
        cin >> resposta;

        if (resposta == 'S') {
            adiciona_palavra();
        }
    }

    cout<< "Fim de Jogo!" << endl;
}

// Para facilitar a compilação do arquivo, foi criado um arquivo Makefile.
/* Explicação para o arquivo Makefile:
    1ª parte: definição de variáveis.
    2ª parte: criação do comando de compilação (all:). Nessa parte é escrito que o arquivo a ser compilado é o TARGET, e que ele necessita do arquivo SRC para isso acontecer, além do próprio comando para compilar.
    3ª parte: criação do comando de rodar o arquivo (run:). Nessa parte é escrito como o arquivo TARGET é rodado.
    4ª parte: craição do comando para apagar o arquivo (clean:). Está escrito o comando para apagar o arquivo TARGET.
*/