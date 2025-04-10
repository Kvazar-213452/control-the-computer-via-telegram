#include <windows.h>
#include <stdio.h>
#include <string.h>

// Функція для перетворення символу на код клавіші
WORD GetKeyCode(const char* key) {
    if (strcmp(key, "a") == 0) return 0x41; // 'A'
    if (strcmp(key, "b") == 0) return 0x42; // 'B'
    if (strcmp(key, "c") == 0) return 0x43; // 'C'
    if (strcmp(key, "w") == 0) return 0x57; // 'W'
    if (strcmp(key, "s") == 0) return 0x53; // 'S'
    if (strcmp(key, "f1") == 0) return 0x70; // F1
    if (strcmp(key, "enter") == 0) return VK_RETURN; // Enter
    if (strcmp(key, "space") == 0) return VK_SPACE; // Space
    if (strcmp(key, "ctrl") == 0) return VK_CONTROL; // Ctrl
    if (strcmp(key, "alt") == 0) return VK_MENU; // Alt
    return 0;
}

void PressKey(WORD key) {
    INPUT input[2] = {0};

    input[0].type = INPUT_KEYBOARD;
    input[0].ki.wVk = key;

    input[1].type = INPUT_KEYBOARD;
    input[1].ki.wVk = key;
    input[1].ki.dwFlags = KEYEVENTF_KEYUP;

    SendInput(2, input, sizeof(INPUT));
}

int main(int argc, char* argv[]) {
    if (argc < 2) {
        printf("Please provide the key to press (e.g., 'a', 'w', 'f1', 'space')\n");
        return -1;
    }

    const char* keyName = argv[1];

    WORD keyCode = GetKeyCode(keyName);

    if (keyCode == 0) {
        printf("Invalid key name: %s\n", keyName);
        return -1;
    }

    PressKey(keyCode);
    printf("Key '%s' has been pressed and released.\n", keyName);

    return 0;
}
