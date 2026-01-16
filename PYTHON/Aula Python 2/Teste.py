def criar_conta(numero, titular, saldo, limite) -> dict:
    conta = {'numero': numero, 'titular': titular, 'saldo': saldo, 'limite': limite}
    return conta

def depositar(conta, valor):
    conta['saldo'] += valor

def sacar(conta, valor):
    conta['saldo'] -= valor

def extrato(conta):
    print('O saldo da sua conta é {}'.format(conta['saldo']))