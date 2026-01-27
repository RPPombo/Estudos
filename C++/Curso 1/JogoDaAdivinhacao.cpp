#include <iostream> // Biblioteca padrão para entrada e saída em C++
using namespace std; // utilizado para indicar que está usando varias funções do std, tirando assim a necessidade de ficar colocando toda hora std::

int main(){
    cout << "Bem-Vindo ao jogo de Adivinhação!" << endl; // É dessa forma que é impresso strings no terminal. O std::endl serve como um \n

    int numero_secreto = 42;

    cout << "O número secreto é "<< numero_secreto << ". Não conte para ninguém!" << endl;
}