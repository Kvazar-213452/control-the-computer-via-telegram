import argparse
from gtts import gTTS

def text_to_speech_ukraine(text, output_file='output.mp3'):
    try:
        tts = gTTS(text=text, lang='uk', slow=False)
        tts.save(output_file)
        print(f"Файл збережено: {output_file}")
    except Exception as e:
        print(f"Сталася помилка: {e}")

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description="Перетворення тексту в мовлення.")
    parser.add_argument('text', help="Текст для перетворення в мовлення")
    parser.add_argument('--output', default='output.mp3', help="Шлях до файлу для збереження (за замовчуванням output.mp3)")
    
    args = parser.parse_args()

    text_to_speech_ukraine(args.text, args.output)
