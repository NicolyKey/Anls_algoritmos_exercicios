using exercicio_01;
using Xunit;
using Xunit.Abstractions;

namespace exercicio_01.Tests;

public class Questao01Tests
{
    private readonly ITestOutputHelper _output;

    public Questao01Tests(ITestOutputHelper output)
    {
        _output = output;
    }

    [Fact]
    public void ExecutarComparacao_ComN20_DeveExecutarSemErros()
    {
        var questao = new Questao01();
        
        var originalOut = Console.Out;
        using var writer = new StringWriter();
        Console.SetOut(writer);
        
        questao.ExecutarComparacao(20);
        
        Console.SetOut(originalOut);
        var output = writer.ToString();
        
        _output.WriteLine(output);
        Assert.Contains("Tempo inserção Lista:", output);
        Assert.Contains("Tempo inserção Lista Encadeada:", output);
    }

    [Theory]
    [InlineData(10, 10)]
    [InlineData(1000, 10)]
    [InlineData(100000, 10)]
    public void CompararCapacidadesIniciais_ComDiferentesCapacidades_DeveExecutarSemErros(int capacidade, int repeticoes)
    {
        var questao = new Questao01();
        
        var originalOut = Console.Out;
        using var writer = new StringWriter();
        Console.SetOut(writer);
        
        questao.CompararCapacidadesIniciais(capacidade, repeticoes);
        
        Console.SetOut(originalOut);
        var output = writer.ToString();
        
        _output.WriteLine(output);
        Assert.Contains($"Capacidade inicial: {capacidade}", output);
        Assert.Contains("Tempo inserção Lista:", output);
    }

    [Fact]
    public void ExecutarTodosOsTestes()
    {
        var questao = new Questao01();
        
        var originalOut = Console.Out;
        using var writer = new StringWriter();
        Console.SetOut(writer);
        
        _output.WriteLine("=== Executando Comparação ===");
        questao.ExecutarComparacao(20);
        
        _output.WriteLine("\n=== Comparando Capacidades Iniciais ===");
        questao.CompararCapacidadesIniciais(10, 10);
        questao.CompararCapacidadesIniciais(1000, 10);
        questao.CompararCapacidadesIniciais(100000, 10);
        
        Console.SetOut(originalOut);
        var output = writer.ToString();
        
        _output.WriteLine(output);
    }
}
