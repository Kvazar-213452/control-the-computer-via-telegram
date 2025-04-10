g++ -o main.exe sound_volume.cpp -static -static-libgcc -static-libstdc++ -lole32 -loleaut32 -lwindowscodecs -mwindows
gcc -o mouse_speed mouse_speed.c -luser32 -mwindows
