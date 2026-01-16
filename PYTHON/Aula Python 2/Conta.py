class Conta:
    def __init__(self, numero, titular, saldo, limite) -> None:
        self.numero = numero
        self.titular = titular
        self.saldo = saldo
        self.limite = limite
    
    def extrato(self):
        print(f'Saldo:{self.saldo} do titular {self.titular}')

    def depositar(self, valor):
        self.saldo = valor + self.saldo

    def sacar(self, valor):
        self.saldo = self.saldo - valor
        
    pass