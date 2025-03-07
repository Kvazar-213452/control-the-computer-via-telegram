// Program.cs
using System;
using System.IO;
using System.Threading.Tasks;
using Newtonsoft.Json;
using Telegram.Bot;
using MyTelegramBot;

class Program {
    static async Task Main(string[] args) {
        string json = File.ReadAllText("token.json");
        dynamic tokenData = JsonConvert.DeserializeObject(json);
        string token = tokenData.token;

        var botClient = new TelegramBotClient(token);

        var me = await botClient.GetMeAsync();
        Console.WriteLine($"Hello! I am {me.Username}!");

        BotHandler handler = new BotHandler(botClient);

        botClient.OnMessage += handler.HandleMessages;

        botClient.StartReceiving();

        Console.WriteLine("Press any key to exit...");
        Console.ReadKey();

        botClient.StopReceiving();
    }
}
