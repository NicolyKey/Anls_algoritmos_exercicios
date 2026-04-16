Resultado da run dos testes

//---------------//

Numero de elementos por teste: 100000

Q1 - Insercao no final (20 execucoes) \
Slice (vetor):     917.56µs \
LinkedList:        2.962085ms \
-> Slice foi 3.23x mais rapida

Q1.2 - Influencia da capacidade inicial (Slice) \
Capacidade inicial 10: media = 593.92µs \
Capacidade inicial 1000: media = 744.71µs \
Capacidade inicial 100000: media = 99.64µs

Q2 - Insercao em posicao aleatoria (20 execucoes)  \
Slice (vetor):     275.63353ms \
LinkedList:        12.3664551s \
-> Slice foi 44.87x mais rapida

Q3.1 - Remocao do primeiro elemento (20 execucoes) \
Slice (vetor):     50.005µs \
LinkedList:        251.295µs \
-> Slice foi 5.03x mais rapida

Q3.2 - Remocao do ultimo elemento (20 execucoes) \
Slice (vetor):     0s \
LinkedList:        7.205663995s \
-> Slice foi +Infx mais rapida

Q4 - Remocao em indice aleatorio (20 execucoes) \
Slice (vetor):     596.210095ms \
LinkedList:        4.92371621s \
-> Slice foi 8.26x mais rapida

Q5 - Acesso aleatorio (10.000 acessos) (20 execucoes) \
Slice (vetor):     50.3µs \
LinkedList:        689.081265ms \
-> Slice foi 13699.43x mais rapida
