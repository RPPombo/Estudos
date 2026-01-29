#include <iostream> // Biblioteca padrão para entrada e saída em C++
#include <cctype> // Biblioteca para transformações em chars
#include <cstdlib> // Biblioteca padrão de C
#include <ctime> // Bilioteca para o receber o tempo
using namespace std; // utilizado para indicar que está usando varias funções do std, tirando assim a necessidade de ficar colocando toda hora std::

int main(){
    cout << "Bem-Vindo ao jogo de Adivinhação!" << endl;// É dessa forma que é impresso strings no terminal. O std::endl serve como um \n

    cout << "Escolha o nível de dificuldade:" << endl;
    cout << "(F)Fácil | (M) Médio | (D) Difícil" << endl;

    char dificuldade;
    cin >> dificuldade;
    dificuldade = toupper(dificuldade); // deixando a entrada sempre em uppercase

    int total_tentativas;

    if (dificuldade == 'F' ){
        total_tentativas = 15; 
    } else if (dificuldade == 'M') {
        total_tentativas = 10;
    } else {
        total_tentativas = 5;
    }

    srand(time(NULL));
    const int NUMERO_SECRETO = rand() % 100; // manter a váriavel sempre a mesma com const

    int tentativas = 1;
    double pontos = 1000.0;

    bool acertou = false;

    while (!acertou && tentativas <= total_tentativas){
        
        int chute;
        cout << "Tentativa " << tentativas << " de " << total_tentativas << endl;

        cout << "Qual o valor do chute? ";
        cin >> chute; // Maneira de receber um input

        cout << "Seu chute foi: " << chute << endl;

        double pontos_perdidos = abs(chute - NUMERO_SECRETO)/2.0; // Quando uma conta é feita apenas com valores inteiros o valor de saída sempre será inteiro, por isso o 2.0

        pontos -= pontos_perdidos;

        acertou = chute == NUMERO_SECRETO;
        bool menor = chute < NUMERO_SECRETO;

        if (acertou) {
            cout << "\nParabéns você acertou!" << endl;
            cout << "\nForam " << tentativas << " tentativas utilizadas!" << endl;
            cout.precision(2); // Responsável por colocar precisão nas váriaveis numéricas do terminal 
            cout << fixed; // Responsável por fixar as casas decimais no mesmo lugar
            cout << "Sua pontuação foi de " << pontos << " pontos." << endl;
        } else if (menor) {
            cout << "Seu chute foi menor que o número secreto!" << endl;
            tentativas ++;
        } else {
            cout << "Seu chute foi maior que o número secreto!" << endl;
            tentativas ++;
        }
    }

    if (!acertou) {
        cout << "\nVocê não acertou! Tente novamente!" << endl;
    }
    
    cout << "\nFim de jogo!" << endl;
}