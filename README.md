# Anls_algoritmos_exercicios

equipe: Ana Carolina Hausmann, Artur Baldo, Nicoly K. Lima Araujo

### Questão 1
Compare a velocidade de inserção de novos elementos ao final das duas estruturas.
1. Após 20 testes, qual foi o tempo médio que cada lista demorou para ser preenchida? Qual
delas foi mais rápida? Quantas vezes ela foi mais rápida que a outra?
  R: Para testar foi necessário rodar o comando `dotnet test --logger "console;verbosity=detailed"`
  
  ```bash
     N = 0
    Tempo inserção Lista: 152,2 ms
    Tempo inserção Lista Encadeada: 354,7 ms
    ------------------------------
    N = 1
    Tempo inserção Lista: 5,6 ms
    Tempo inserção Lista Encadeada: 7,3 ms
    ------------------------------
    N = 2
    Tempo inserção Lista: 2,6 ms
    Tempo inserção Lista Encadeada: 9,4 ms
    ------------------------------
    N = 3
    Tempo inserção Lista: 6,8 ms
    Tempo inserção Lista Encadeada: 10,4 ms
    ------------------------------
    N = 4
    Tempo inserção Lista: 4,5 ms
    Tempo inserção Lista Encadeada: 8,4 ms
    ------------------------------
    N = 5
    Tempo inserção Lista: 2,6 ms
    Tempo inserção Lista Encadeada: 7,3 ms
    ------------------------------
    N = 6
    Tempo inserção Lista: 4,9 ms
    Tempo inserção Lista Encadeada: 7,2 ms
    ------------------------------
    N = 7
    Tempo inserção Lista: 2,3 ms
    Tempo inserção Lista Encadeada: 8,7 ms
    ------------------------------
    N = 8
    Tempo inserção Lista: 2,6 ms
    Tempo inserção Lista Encadeada: 7,7 ms
    ------------------------------
    N = 9
    Tempo inserção Lista: 3,1 ms
    Tempo inserção Lista Encadeada: 6,8 ms
    ------------------------------
    N = 10
    Tempo inserção Lista: 6,1 ms
    Tempo inserção Lista Encadeada: 11,4 ms
    ------------------------------
    N = 11
    Tempo inserção Lista: 5,7 ms
    Tempo inserção Lista Encadeada: 10,1 ms
    ------------------------------
    N = 12
    Tempo inserção Lista: 10,5 ms
    Tempo inserção Lista Encadeada: 12,3 ms
    ------------------------------
    N = 13
    Tempo inserção Lista: 9,6 ms
    Tempo inserção Lista Encadeada: 8,2 ms
    ------------------------------
    N = 14
    Tempo inserção Lista: 5,7 ms
    Tempo inserção Lista Encadeada: 9 ms
    ------------------------------
    N = 15
    Tempo inserção Lista: 7,6 ms
    Tempo inserção Lista Encadeada: 9,3 ms
    ------------------------------
    N = 16
    Tempo inserção Lista: 6 ms
    Tempo inserção Lista Encadeada: 15,3 ms
    ------------------------------
    N = 17
    Tempo inserção Lista: 9 ms
    Tempo inserção Lista Encadeada: 6,2 ms
    ------------------------------
    N = 18
    Tempo inserção Lista: 4,5 ms
    Tempo inserção Lista Encadeada: 6,5 ms
    ------------------------------
    N = 19
    Tempo inserção Lista: 3,2 ms
    Tempo inserção Lista Encadeada: 5,7 ms
    ------------------------------
  ```
  R: a linkedList é mais lenta por possuir mais operações e também não ter uma linearidade ao acessar os elementos. Quando eu faço um addLast, eu precsiso atualizar os ponteitos de previous e next daquele node em específico, por essas razões a List é linearmente mais rápida que a LinkedList.
  
3. No caso do ArrayList, avalie se alterar a capacidade inicial influencia no desempenho.
Realize 10 execuções para cada uma das seguintes capacidades iniciais:
• 10
• 1000
• 100.000
Calcule a média dos tempos e responda:
• Qual configuração foi mais rápida?
• Por que isso acontece?

R: Ao definir uma capacidade inicial no array nos três casos possuí uma variação no tempo de inserção entre 4ms até 3,2ms, porém o de 10 elementos foi o mais rápido. Devido a quantidade de espaço que eu inicialmente aloco na memória ao instanciar o array.

### Questão 2
Repita o experimento anterior, mas agora cada inserção deve ocorrer em uma posição aleatória
válida da lista.
1. Após 20 testes, qual foi o tempo médio que cada lista demorou para ser preenchida? Qual
delas foi mais rápida? Quantas vezes ela foi mais rápida?