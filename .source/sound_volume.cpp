#include <windows.h>
#include <mmdeviceapi.h>
#include <endpointvolume.h>
#include <comdef.h>
#include <stdio.h>
#include <stdlib.h>

#pragma comment(lib, "Ole32.lib")

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <volume_percentage>\n", argv[0]);
        return -1;
    }

    float volume_percentage = atof(argv[1]);

    if (volume_percentage < 0 || volume_percentage > 100) {
        printf("Please enter a volume percentage between 0 and 100.\n");
        return -1;
    }

    CoInitialize(NULL);

    IMMDeviceEnumerator* deviceEnumerator = NULL;
    IMMDevice* defaultDevice = NULL;
    IAudioEndpointVolume* volume = NULL;

    HRESULT hr = CoCreateInstance(__uuidof(MMDeviceEnumerator), NULL, CLSCTX_ALL, __uuidof(IMMDeviceEnumerator), (void**)&deviceEnumerator);
    if (FAILED(hr)) {
        printf("CoCreateInstance failed\n");
        return -1;
    }

    hr = deviceEnumerator->GetDefaultAudioEndpoint(eRender, eConsole, &defaultDevice);
    if (FAILED(hr)) {
        printf("GetDefaultAudioEndpoint failed\n");
        return -1;
    }

    hr = defaultDevice->Activate(__uuidof(IAudioEndpointVolume), CLSCTX_ALL, NULL, (void**)&volume);
    if (FAILED(hr)) {
        printf("Activate failed\n");
        return -1;
    }

    float volume_scalar = volume_percentage / 100.0f;

    hr = volume->SetMasterVolumeLevelScalar(volume_scalar, NULL);
    if (FAILED(hr)) {
        printf("Failed to set volume level\n");
        return -1;
    }

    printf("Volume set to %.2f%%\n", volume_percentage);

    volume->Release();
    defaultDevice->Release();
    deviceEnumerator->Release();

    CoUninitialize();
    return 0;
}
