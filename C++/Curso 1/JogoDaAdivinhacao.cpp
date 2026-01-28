#include <iostream> // Biblioteca padrão para entrada e saída em C++
using namespace std; // utilizado para indicar que está usando varias funções do std, tirando assim a necessidade de ficar colocando toda hora std::

int main(){
    cout << "Bem-Vindo ao jogo de Adivinhação!" << endl; // É dessa forma que é impresso strings no terminal. O std::endl serve como um \n

    const int NUMERO_SECRETO = 42; // manter a váriavel sempre a mesma com const

    int tentativas = 0;
    double pontos = 1000.0;

    bool acertou = false;

    while (!acertou){
        tentativas ++;
        int chute;
        cout << "Tentativa " << tentativas << endl;

        cout << "Qual o valor do chute? ";
        cin >> chute; // Maneira de receber um input

        cout << "Seu chute foi: " << chute << endl;

        double pontos_perdidos = abs(chute - NUMERO_SECRETO)/2.0; // Quando uma conta é feita apenas com valores inteiros o valor de saída sempre será inteiro, por isso o 2.0

        pontos -= pontos_perdidos;

        acertou = chute == NUMERO_SECRETO;
        bool menor = chute < NUMERO_SECRETO;

        if (acertou) {
            cout << "Parabéns você acertou!" << endl;
            cout << "Foram " << tentativas << " tentativas utilizadas!" << endl;
        } else if (menor) {
            cout << "Seu chute foi menor que o número secreto!" << endl;
        } else {
            cout << "Seu chute foi maior que o número secreto!" << endl;
        }
    }

    cout.precision(2); // Responsável por colocar precisão nas váriaveis numéricas do terminal 
    cout << fixed; // Responsável por fixar as casas decimais no mesmo lugar
    cout << "Sua pontuação foi de " << pontos << " pontos." << endl;
    cout << "Fim de jogo!" << endl;
}