import telebot
import json
import pygame
import os
import pyautogui
import ctypes
import wmi
import sys
from ctypes import cast, POINTER
from comtypes import CLSCTX_ALL
from pycaw.pycaw import AudioUtilities, IAudioEndpointVolume

with open('unix.json', 'r') as f:
    data = json.load(f)

pygame.init()

if data['article']['bot'] != "unix":
    sys.exit()
if data['article']['dicord'] != "https://discord.gg/article":
    sys.exit()
if data['article']['version'] != "2.20":
    sys.exit()
if data['article']['paradise'] != "1.00":
    sys.exit()


def play_music(file_path):
    pygame.mixer.music.load(file_path)
    pygame.mixer.music.play()

def close_window():
    pyautogui.hotkey('alt', 'f4')

def sleep_windows():
    ctypes.windll.Powrprof.SetSuspendState(0, 1, 0)

def shutdown_windows():
    os.system("shutdown /s /t 1")

def reboot_windows():
    os.system("shutdown /r /t 1")

def change_screen_brightness(percentage):
    brightness = int(percentage * 255 / 100)
    c = wmi.WMI(namespace='wmi')
    methods = c.WmiMonitorBrightnessMethods()[0]
    methods.WmiSetBrightness(brightness, 0)

def set_sound_volume(volume_percentage):
    devices = AudioUtilities.GetSpeakers()
    interface = devices.Activate(
        IAudioEndpointVolume._iid_, CLSCTX_ALL, None)
    volume = cast(interface, POINTER(IAudioEndpointVolume))
    volume.SetMasterVolumeLevelScalar(volume_percentage / 100, None)

def set_mouse_speed(speed):
    SPI_SETMOUSESPEED = 113
    ctypes.windll.user32.SystemParametersInfoW(SPI_SETMOUSESPEED, 0, speed, 0)



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
