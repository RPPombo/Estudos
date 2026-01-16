#Importando bibliotecas
from openai import OpenAI
from dotenv import load_dotenv
import os

load_dotenv() #Carregando variáveis de ambiente

cliente = OpenAI(api_key = os.getenv('OPENAI_API_KEY')) #Criando a conexão com o OpenAI

resposta = cliente.chat.completions.create(
    messages = [
        {"role":"system", "content": "Listar apenas os nomes dos produtos, sem considerar descrição."}, #role = Papel da mensagem e content = conteúdo
        { "role": "user", "content": "Liste 3 produtos sustentáveis"} 
    ],
    model = "gpt-4"
)

print(resposta.choices[0].message.content) #Recebimento da mensagem da IA