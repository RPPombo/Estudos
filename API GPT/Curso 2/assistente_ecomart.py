from openai import OpenAI
from dotenv import load_dotenv
import os
from helpers import *
from selecionar_persona import *
import json
from tools_ecomart import *

load_dotenv()

cliente = OpenAI(api_key=os.getenv('API_KEY'))
modelo = "gpt-4.1-nano"
contexto = carrega("dados/ecomart.txt")

def criar_lista_ids():
    lista_ids_arquivos = list()
    file_dados = cliente.files.create(
        file=open("dados/dados_ecomart.txt","rb"),
        purpose="assistants"
    )
    lista_ids_arquivos.append(file_dados.id)

    file_politicas = cliente.files.create(
        file=open("dados/politicas_ecomart.txt","rb"),
        purpose="assistants"
    )
    lista_ids_arquivos.append(file_politicas.id)

    file_produtos = cliente.files.create(
        file=open("dados/produtos_ecomart.txt","rb"),
        purpose="assistants"
    )
    lista_ids_arquivos.append(file_produtos.id)

    return lista_ids_arquivos

def pegar_json():
    file_name = "assistentes.json"

    if not os.path.exists(file_name):
        thread_id = criar_thread()
        file_id_list = criar_lista_ids()
        assistant_id = criar_assistente(file_id_list)
        data = {
            "assistant_id": assistant_id.id,
            "thread_id": thread_id.id,
            "file_ids": file_id_list
        }

        with open(file_name , "w", encoding="utf-8") as file:
            json.dump(data, file, ensure_ascii= False, indent=4)
        print(f"arquivo {file_name} criado")

    try:
        with open(file_name, "r",encoding="utf-8") as file:
            data = json.load(file)
            return data
    except FileNotFoundError:
        print("arquivo não encontrado")

def criar_thread():
    return cliente.beta.threads.create()

def criar_assistente(file_ids: list):
    assistente = cliente.beta.assistants.create(
        name = "Atendente Ecomart",
        instructions = f"""
        Você é um chatbot de atendimento a clientes de um e-commerce. 
        Você não deve responder perguntas que não sejam dados do ecommerce informado!
        Além disso, acesse os arquivos associados a você e a thread para responder as perguntas.
        """,
        model = modelo,
        tools= minhas_tools,
    )

    return assistente

"""
thread = cliente.beta.threads.create(
    messages=[
        {
            "role": "user",
            "content": "Liste os produtos"
        }
    ]
) #Criando uma thread

cliente.beta.threads.messages.create(
    thread_id=thread.id,
    role = "user",
    content =  " da categoria moda sustentável"
) #Acessando a thread anterior e adicionando conteúdo

run = cliente.beta.threads.runs.create(
    thread_id=thread.id,
    assistant_id=assistente.id
) #Rodando a thread

while run.status !="completed":
    run = cliente.beta.threads.runs.retrieve(
        thread_id=thread.id,
        run_id=run.id
    ) #Verificando o status da thread

historico = cliente.beta.threads.messages.list(thread_id=thread.id).data

for mensagem in reversed(historico):
    print(f"role: {mensagem.role}\nConteúdo: {mensagem.content[0].text.value}")
"""