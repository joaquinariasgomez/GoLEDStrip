# GoLEDServer :)

## Prerrequisitos

Compilar la librería [rpi_ws281x](https://github.com/jgarff/rpi_ws281x) directamente en la Raspberry Pi, de acuerdo a la [documentación](https://github.com/jgarff/rpi_ws281x#build), para luego copiar los ficheros `*.a` a `/usr/local/lib` y los ficheros `*.h` a `/usr/local/include`.

## Documentación de peticiones

### POST
Tipo ("type"):
- "change-mode"
- "set-brightness"
- "set-color"

Modos ("change-mode"):
Instrucción ("instruction"):
- "office-lights"
- "static-color"
  - [Opcional] "args":
    1. Color, en hexadecimal, que settea

Ajustes de brillo ("set-brightness"):
Instrucción ("instruction"):
- "increase"
  - [Opcional] "args":
    1. Cantidad, en enteros, que sube de brillo
- "decrease"
  - [Opcional] "args":
    1. Cantidad, en enteros, que baja de brillo

Ajustes de color ("set-color"):
Instrucción ("instruction"):
- Color, en hexadecimal, que settea

## TODO
- Quizás crear una extensión de Google Chrome, a modo de frontend. Usar la referencia https://www.youtube.com/playlist?list=PLC3y8-rFHvwg2-q6Kvw3Tl_4xhxtIaNlY