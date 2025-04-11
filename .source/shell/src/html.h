#ifndef HTML_H
#define HTML_H

#include <stdio.h>

const char* html_(const char* port) {
    static char text[256];
    snprintf(text, sizeof(text),
             "<style>iframe{position: fixed;height: 100%%;width: 100%%;top: 0%%;left: 0%%;}</style>"
             "<iframe src='http://127.0.0.1:%s' frameborder='0'></iframe>", port);
    return text;
}

#endif