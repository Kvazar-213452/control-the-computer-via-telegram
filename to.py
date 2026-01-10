import base64

input_path = 'shell.zip'
output_path = 'data.datat'

with open(input_path, 'rb') as f:
    file_data = f.read()
    encoded_data = base64.b64encode(file_data).decode('utf-8')

with open(output_path, 'w') as f:
    f.write(encoded_data)

print(f"✅ Base64 записано в {output_path}")
