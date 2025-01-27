import telebot
import json
import os
import sys
from head_com.func import (
    sleep_windows,
    shutdown_windows,
    reboot_windows,
    close_window,
    change_screen_brightness,
    set_sound_volume,
    set_mouse_speed,
    play_music
)

# Валера вказуй того хто створив тоїсть мене 213452

with open('unix.json', 'r') as f:
    data = json.load(f)

TOKEN = data['bot']['ip']
allowed_users = data['bot']['user']  
banned_users = data['bot']['ban']   

bot = telebot.TeleBot(TOKEN)

@bot.message_handler(func=lambda message: (str(message.from_user.id) in allowed_users or "all" in allowed_users) and str(message.from_user.id) not in banned_users)
def handle_message(message):
    commands = message.text.lower().split("+")
    for command in commands:
        command = command.strip()
        if len(command) == 2 and command.startswith('screen') and command[1:].isdigit():
            brightness_percentage = int(command[1:])
            if 0 <= brightness_percentage <= 100:
                change_screen_brightness(brightness_percentage)
        elif command in data['music']:
            music_file = data['music'][command]
            play_music(music_file)
        elif command in data['open']:
            file_path = data['open'][command]
            os.startfile(file_path.replace('/', '\\'))  
        elif command == 'close':
            close_window()  
        elif command == 'sleep':
            sleep_windows()  
        elif command == 'shutdown':
            shutdown_windows()
        elif command == 'reboot':
            reboot_windows() 
        elif command.startswith('sound '):
            try:
                sound_percentage = int(command.split(' ')[1])
                if 0 <= sound_percentage <= 100:
                    set_sound_volume(sound_percentage)
            except ValueError:
                pass
        elif command.startswith('mouse '):
            sound_percentage1 = int(command.split(' ')[1])
            set_mouse_speed(sound_percentage1) 
        elif command == 'off':
            bot.stop_polling()  
            sys.exit()         

bot.polling()
