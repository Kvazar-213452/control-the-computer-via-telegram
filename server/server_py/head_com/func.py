import os
import pythoncom
import ctypes
import wmi
import sounddevice as sd
import soundfile as sf

from ctypes import cast, POINTER
from pycaw.pycaw import AudioUtilities, IAudioEndpointVolume
from comtypes import CLSCTX_ALL, CoInitialize, CoUninitialize

def set_mouse_speed(speed):
    SPI_SETMOUSESPEED = 113
    ctypes.windll.user32.SystemParametersInfoW(SPI_SETMOUSESPEED, 0, speed, 0)

def set_sound_volume(volume_percentage):
    CoInitialize()

    try:
        devices = AudioUtilities.GetSpeakers()
        interface = devices.Activate(
            IAudioEndpointVolume._iid_, CLSCTX_ALL, None)
        volume = cast(interface, POINTER(IAudioEndpointVolume))
        volume.SetMasterVolumeLevelScalar(volume_percentage / 100, None)
    finally:
        CoUninitialize()

def change_screen_brightness(percentage):
    pythoncom.CoInitialize()

    brightness = int(percentage * 255 / 100)
    c = wmi.WMI(namespace='wmi')

    methods = c.WmiMonitorBrightnessMethods()[0]
    methods.WmiSetBrightness(brightness, 0)

    pythoncom.CoUninitialize()
