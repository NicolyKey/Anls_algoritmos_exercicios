using exercicio_01;
using Xunit;
using Xunit.Abstractions;

namespace exercicio_01.Tests;

public class Questao02Tests
{
    private readonly ITestOutputHelper _output;

    public Questao02Tests(ITestOutputHelper output)
    {
        _output = output;
    }

    [Fact]
    public void ExecutarComparacao()
    {
        var questao = new Questao02();
        
        var originalOut = Console.Out;
        using var writer = new StringWriter();
        Console.SetOut(writer);
        
        questao.ExecutarComparacao(20);
        
        Console.SetOut(originalOut);
        var output = writer.ToString();
        
        _output.WriteLine(output);
        Assert.Contains("Tempo inserção Lista posição aleatória:", output);
        Assert.Contains("repeticao =", output);
    }
}
