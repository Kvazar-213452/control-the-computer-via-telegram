#include "webview.h"
#include <windows.h>
#include <stdio.h>
#include <stdlib.h>

// ⣿⠟⣽⣿⣿⣿⣿⣿⢣⠟⠋⡜⠄⢸⣿⣿⡟⣬⢁⠠⠁⣤⠄⢰⠄⠇⢻⢸
// ⢏⣾⣿⣿⣿⠿⣟⢁⡴⡀⡜⣠⣶⢸⣿⣿⢃⡇⠂⢁⣶⣦⣅⠈⠇⠄⢸⢸
// ⣹⣿⣿⣿⡗⣾⡟⡜⣵⠃⣴⣿⣿⢸⣿⣿⢸⠘⢰⣿⣿⣿⣿⡀⢱⠄⠨⢸
// ⣿⣿⣿⣿⡇⣿⢁⣾⣿⣾⣿⣿⣿⣿⣸⣿⡎⠐⠒⠚⠛⠛⠿⢧⠄⠄⢠⣼
// ⣿⣿⣿⣿⠃⠿⢸⡿⠭⠭⢽⣿⣿⣿⢂⣿⠃⣤⠄⠄⠄⠄⠄⠄⠄⠄⣿⡾
// ⣼⠏⣿⡏⠄⠄⢠⣤⣶⣶⣾⣿⣿⣟⣾⣾⣼⣿⠒⠄⠄⠄⡠⣴⡄⢠⣿⣵
// ⣳⠄⣿⠄⠄⢣⠸⣹⣿⡟⣻⣿⣿⣿⣿⣿⣿⡿⡻⡖⠦⢤⣔⣯⡅⣼⡿⣹
// ⡿⣼⢸⠄⠄⣷⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣕⡜⡌⡝⡸⠙⣼⠟⢱⠏
// ⡇⣿⣧⡰⡄⣿⣿⣿⣿⡿⠿⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣋⣪⣥⢠⠏⠄
// ⣧⢻⣿⣷⣧⢻⣿⣿⣿⡇⠄⢀⣀⣀⡙⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠂⠄⠄
// ⢹⣼⣿⣿⣿⣧⡻⣿⣿⣇⣴⣿⣿⣿⣷⢸⣿⣿⣿⣿⣿⣿⣿⣿⣰⠄⠄⠄
// ⣼⡟⡟⣿⢸⣿⣿⣝⢿⣿⣾⣿⣿⣿⢟⣾⣿⣿⣿⣿⣿⣿⣿⣿⠟⠄⡀⡀
// ⣿⢰⣿⢹⢸⣿⣿⣿⣷⣝⢿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠛⠉⠄⠄⣸⢰⡇
// ⣿⣾⣹⣏⢸⣿⣿⣿⣿⣿⣷⣍⡻⣛⣛⣛⡉⠁⠄⠄⠄⠄⠄⠄⢀⢇⡏⠄

void SetWindowIcon(HWND hwnd, LPCWSTR iconPath) {
    HICON hIcon = (HICON)LoadImageW(NULL, iconPath, IMAGE_ICON, 0, 0, LR_LOADFROMFILE | LR_DEFAULTSIZE);
    if (hIcon) {
        SendMessage(hwnd, WM_SETICON, ICON_BIG, (LPARAM)hIcon);
        SendMessage(hwnd, WM_SETICON, ICON_SMALL, (LPARAM)hIcon);
    }
}

char* read_base64_from_file(const char* filepath) {
    FILE* file = fopen(filepath, "rb");
    if (!file) return NULL;

    fseek(file, 0, SEEK_END);
    long length = ftell(file);
    fseek(file, 0, SEEK_SET);

    char* buffer = (char*)malloc(length + 1);
    if (!buffer) {
        fclose(file);
        return NULL;
    }

    fread(buffer, 1, length, file);
    buffer[length] = '\0';
    fclose(file);
    return buffer;
}

void SetupWebview(webview_t w, const char* title, int height, int width, const char* base64_video) {
    webview_set_title(w, title);
    webview_set_size(w, width, height, WEBVIEW_HINT_NONE);

    size_t base64_len = strlen(base64_video);
    if (base64_len == 0) {
        printf("Помилка: base64_video порожній!\n");
        return;
    }

    size_t html_size = 8192 + base64_len;
    char* html_template = (char*)malloc(html_size);
    if (!html_template) {
        printf("Помилка: Не вдалося виділити пам'ять для HTML шаблону.\n");
        return;
    }

    snprintf(html_template, html_size,
        "<html><body style='margin:0;background:#000;'>"
        "<video width='100%%' height='100%%' loop autoplay muted>"
        "<source src='data:video/mp4;base64,%s' type='video/mp4'>"
        "</video></body></html>", base64_video);

    webview_set_html(w, html_template);

    free(html_template);
}


int main(int argc, char *argv[]) {
    if (argc < 4) {
        printf("Використання: %s <title> <height> <width>\n", argv[0]);
        return 1;
    }

    const char* title = argv[1];
    int height = atoi(argv[2]);
    int width = atoi(argv[3]);

    char* base64_video = read_base64_from_file("./data.temp");
    if (!base64_video) {
        printf("Не вдалося зчитати файл data.temp\n");
        return 1;
    }

    webview_t w = webview_create(0, NULL);
    if (!w) {
        free(base64_video);
        return -1;
    }

    SetupWebview(w, title, height, width, base64_video);
    free(base64_video);

    HWND hwnd = (HWND)webview_get_window(w);
    SetWindowIcon(hwnd, L"icon.ico");

    webview_run(w);
    webview_destroy(w);
    return 0;
}
