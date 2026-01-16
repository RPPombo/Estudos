// Esse arquivo serve para as funções serem declaradas antes de serem própriamente escritas, para que não haja a preocupação da ordem delas no código
void abertura();
void fazer_chute(char chutes[26], int* chutes_dados);
int ja_chutou(char chutes[26], int chutes_dados, char letra);
void desenha_forca(char palavrasecreta[20], char chutes[26], int chutes_dados);
int ja_enforcou(int chutes_dados, char palavrasecreta[20], char chutes[26]);
int ja_ganhou(char palavrasecreta[20], int chutes_dados, char chutes[26]);
void escolhe_palavra(char palavrasecreta[20]);
int chutes_errados(int chutes_dados, char palavrasecreta[20], char chutes[26]);