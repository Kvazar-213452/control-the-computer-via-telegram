#include <windows.h>
#include <iostream>
#include <string>

bool SetWallpaper(const std::string& imagePath) {
    if (GetFileAttributes(imagePath.c_str()) == INVALID_FILE_ATTRIBUTES) {
        return false;
    }

    int result = SystemParametersInfo(SPI_SETDESKWALLPAPER, 0, (PVOID)imagePath.c_str(), SPIF_UPDATEINIFILE | SPIF_SENDCHANGE);

    if (result == 0) {
        return false;
    }

    return true;
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        return 1;
    }

    std::string imagePath = argv[1];

    if (!SetWallpaper(imagePath)) {
        return 1;
    }

    return 0;
}
