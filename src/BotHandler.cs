// src/BotHandler.cs
using System;
using Telegram.Bot;
using Telegram.Bot.Args;

namespace MyTelegramBot {
    public class BotHandler {
        private TelegramBotClient botClient;

        public BotHandler(TelegramBotClient client) {
            botClient = client;
        }

        public async void HandleMessages(object sender, MessageEventArgs e) {
            if (e.Message.Text != null) {
                Console.WriteLine($"Received a message from {e.Message.From.Username}: {e.Message.Text}");

                if (e.Message.Text.StartsWith("#help")) {
                    await botClient.SendTextMessageAsync(e.Message.Chat.Id, "Here is the list of commands:\n#help - Shows this help message\n...");
                } else {
                    await botClient.SendTextMessageAsync(e.Message.Chat.Id, "You said: " + e.Message.Text);
                }
            }
        }
    }
}
