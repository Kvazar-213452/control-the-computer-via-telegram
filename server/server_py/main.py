from flask import Flask, request, jsonify
from head_com.func import (
    set_sound_volume,
    set_mouse_speed,
    change_screen_brightness
)

app = Flask(__name__)

@app.route('/set_volume', methods=['POST'])
def index_1():
    data = request.get_json()
    val = data.get('volume', None)
    if val is None:
        return jsonify({'error': 'No volume value provided'}), 400

    set_sound_volume(val)

    return jsonify({'message': 'ok'})

@app.route('/set_mouse_speed', methods=['POST'])
def index_2():
    data = request.get_json()
    val = data.get('volume', None)
    if val is None:
        return jsonify({'error': 'No volume value provided'}), 400

    set_mouse_speed(val)

    return jsonify({'message': 'ok'})


@app.route('/change_screen_brightness', methods=['POST'])
def index_0():
    data = request.get_json()
    val = data.get('volume', None)
    if val is None:
        return jsonify({'error': 'No volume value provided'}), 400

    change_screen_brightness(val)

    return jsonify({'message': 'ok'})

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=4444)
