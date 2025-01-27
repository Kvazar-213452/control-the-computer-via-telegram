from gtts import gTTS
import pythoncom
import requests
import ctypes
import wmi
from ctypes import cast, POINTER
from pycaw.pycaw import AudioUtilities, IAudioEndpointVolume
from comtypes import CLSCTX_ALL, CoInitialize, CoUninitialize
from head_com.config import server

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

def text_to_speech_ukraine(text, output_file='output.mp3'):
    try:
        tts = gTTS(text=text, lang='uk', slow=False)
        tts.save(output_file)
        upload_file_to_go_server(output_file)
    except Exception as e:
        print(f"error {e}")

def upload_file_to_go_server(file_path):
    url = server + "upload_mp3"
    files = {'file': open(file_path, 'rb')}
    response = requests.post(url, files=files)
    
    if response.status_code == 200:
        print("godd")
    else:
        print("not good")
