## Busca binária
A ideia princiapal desse algoritmo é agilizar a busca por um item dentro de uma lista ordenada.
Ao inves de passarmos item a item da lista, dividimos ela ao meio e comparamos se o item que buscamos
está na direita ou a esqueda da lista (se é maior ou menor).

EX: [1, 2, 3, 4, 5, 6] -> busque o numero "6"

Falando de complexidade, utilizando BigO, esse é um algoritmo "Log n" (log de n)
onde "n" é a quantidade de itens da lista.

Ao resolvendo a equação temos o número máximo de interações que iremos fazer na lista para
encontrar o alvo.

Ex: Imagine uma lista com 16 item. "Log 16" (Log de 16 na base 2)
    Log 16 = 2^x  // Qual numero que elevado a 2 mais se aproxima de 16?
    2^4 = 2x2x2x2 = 16 
    logo Log 16 = 4 interações.


Qual a quantidade máxima de interações para encontrarmos um item em uma lista de tamanho 128?
Log 128 = 2^x
2ˆ7 = 2x2x2x2x2x2x2 = 128
Logo Log 128 = 7


Tempo de Execução Linerar "O(n)":
    Busca simples, a quantidade de interações é igual ao tamanho da lista.
    100 items -> 100 tentativas

Tempo de Execução Logaritmica "O(Log n)":
    Busca binaria, a quantidade de interações é sempre feita na metada a lista a cada interação.
    100 items -> 7 tentativas
