import os
import pyautogui
import ctypes
import wmi
from ctypes import cast, POINTER
from comtypes import CLSCTX_ALL
from pycaw.pycaw import AudioUtilities, IAudioEndpointVolume
import sounddevice as sd
import soundfile as sf

def sleep_windows():
    os.system("rundll32.exe powrprof.dll,SetSuspendState Sleep")

def shutdown_windows():
    os.system("shutdown /s /t 1")

def reboot_windows():
    os.system("shutdown /r /t 1")

def close_window():
    pyautogui.hotkey('alt', 'f4')

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

def play_music(file_path):
    data, samplerate = sf.read(file_path)
    sd.play(data, samplerate)
    sd.wait()