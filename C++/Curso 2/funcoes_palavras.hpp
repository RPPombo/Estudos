#pragma once
#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <ctime>
#include <cstdlib>
#include "funcoes_arquivo.hpp"

std::string escolhe_palavra();

bool acertou_palavra(const std::string& palavra_secreta, const std::map<char, bool>& chutou);

void mostra_palavra(const std::string& palavra_secreta, const std::map <char,bool>& chutou);

bool letra_existe(const std::string& palavra_secreta ,char chute);