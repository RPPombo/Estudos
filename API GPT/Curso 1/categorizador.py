#Importando bibliotecas
from openai import OpenAI
from dotenv import load_dotenv
import os

load_dotenv() #Carregando variáveis de ambiente

cliente = OpenAI(api_key = os.getenv('OPENAI_API_KEY')) #Criando a conexão com o OpenAI

modelo ="gpt-4"

def categoriza_produtos(nome_produto, lista_categorias):
    prompt_sistema = f"""
        Você é um categorizador de produtos.
        Você deve assumir as categorias presentes na lista abaixo.

        # Lista de Categorias Válidas
        {lista_categorias.split(',')}

        # Formato da Saída
        Produto: Nome do Produto
        Categoria: apresente a categoria do produto

        # Exemplo de Saída
        Produto: Escova elétrica com recarga solar
        Categoria: Eletrônicos Verdes
    """

    resposta = cliente.chat.completions.create(
        messages = [
            {"role":"system", "content": prompt_sistema}, #role = Papel da mensagem e content = conteúdo
            { "role": "user", "content": nome_produto} 
        ], #Mensagens
        model = modelo, #Escolha do modelo
        temperature = 1, #Definição da temperatura (quão criativo/aleatório o modelo será)
        max_tokens = 256, #Definição do número máximo de tokens que o modelo pode gerar
        n = 1 #Definição do número de respostas
    )

    return resposta.choices[0].message.content

categorias = input("Informe as categorias válidas, separadas por vírgula: ")

while True:
    nome_produto = input("Informe o nome do produto: ")
    if nome_produto == 'sair':
        break
    else:
        print(categoriza_produtos(nome_produto,categorias))
