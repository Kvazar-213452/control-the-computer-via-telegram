from flask import Flask, request, jsonify
from ctypes import cast, POINTER
from pycaw.pycaw import AudioUtilities, IAudioEndpointVolume
from comtypes import CLSCTX_ALL, CoInitialize, CoUninitialize

app = Flask(__name__)

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

@app.route('/set_volume', methods=['POST'])
def set_volume():
    data = request.get_json()
    volume_percentage = data.get('volume', None)
    if volume_percentage is None:
        return jsonify({'error': 'No volume value provided'}), 400

    print(volume_percentage)
    set_sound_volume(volume_percentage)

    return jsonify({'message': f'Рівень звуку змінено на {volume_percentage}%'})

# Запуск Flask серверу
if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
