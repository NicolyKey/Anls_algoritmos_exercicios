using System.Diagnostics;

namespace exercicio_01;
public class Questao01
{
    private List<int> lista = new List<int>();
    private LinkedList<int> listaEncadeada = new LinkedList<int>();

    static double MedirTempo(Action action, int repeticoes = 10)
    {
        long ticks = 0;

        for(int i = 0; i < repeticoes; i++)
        {
            GC.Collect();
            GC.WaitForPendingFinalizers();
            GC.Collect();

            Stopwatch stopwatch = Stopwatch.StartNew();
            action();
            stopwatch.Stop();

            ticks += stopwatch.ElapsedTicks;
        }

        return ticks / (double)repeticoes;
    }

    public void ExecutarComparacao(int repeticoes)
    {
        for(int i = 0; i < repeticoes; i++)
        {
            var tempoLista = MedirTempo(() => lista.Add(i));
            var tempoListaEncadeada = MedirTempo(() => listaEncadeada.AddLast(i));

            Console.WriteLine($"repeticao = {i}");
            Console.WriteLine($"Tempo inserção Lista: {tempoLista} ms");
            Console.WriteLine($"Tempo inserção Lista Encadeada: {tempoListaEncadeada} ms");
            Console.WriteLine(new string('-', 30));
        }
    }

    public void  CompararCapacidadesIniciais(int capacidade, int repeticoes)  
    {
        for(int i = 0; i < repeticoes; i++)
        {
            lista = new List<int>(capacidade);
          
            var tempoLista = MedirTempo(() => lista.Add(1));

            Console.WriteLine($"capacidade = {capacidade}");
            Console.WriteLine($"Tempo inserção Lista: {tempoLista} ms");
            Console.WriteLine(new string('-', 30));
        }
    }

}
