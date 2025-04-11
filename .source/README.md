g++ -o main.exe sound_volume.cpp -static -static-libgcc -static-libstdc++ -lole32 -loleaut32 -lwindowscodecs -mwindows
gcc -o mouse_speed mouse_speed.c -static -static-libgcc -luser32 -mwindows
gcc -o set_key set_key.c -static -static-libgcc -luser32 -mwindows
gcc -o close close.c -static -static-libgcc -luser32 -mwindows
g++ -o set_bg set_bg.cpp -static -static-libgcc -luser32 -mwindows

gcc main.c -o shell_web.exe -L. -lwebview -mwindows
