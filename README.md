# GoLEDServer :)

## Prerrequisitos

Compilar la librería [rpi_ws281x](https://github.com/jgarff/rpi_ws281x) directamente en la Raspberry Pi, de acuerdo a la [documentación](https://github.com/jgarff/rpi_ws281x#build), para luego copiar los ficheros `*.a` a `/usr/local/lib` y los ficheros `*.h` a `/usr/local/include`.

## Documentación de peticiones
Peticiones:
Action types: "set-brightness", "set-mode", "set-config" (para configurar qué parte quiero iluminar, por ejemplo)
Modes: "office-lights", ..

- POST a /action con body:
  {
  "type": "set-mode",
  "mode": "office-lights"
  }