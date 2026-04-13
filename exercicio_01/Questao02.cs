using System.Diagnostics;

namespace exercicio_01;
public class Questao02
{
    private List<int> lista = new List<int>();
    private Random random = new Random();

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
            int posicaoAleatoria = lista.Count > 0 ? random.Next(0, lista.Count) : 0;
            
            var tempoLista = MedirTempo(() => lista.Insert(posicaoAleatoria, i));

            Console.WriteLine($"repeticao = {i}");
            Console.WriteLine($"Tempo inserção Lista posição aleatória: {tempoLista} ms");
            Console.WriteLine(new string('-', 30));
        }
    }
}
