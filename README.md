# Servidor Raspberry + Tira de Leds WS281X

Servidor escrito en Golang que permite una comunicación con la tira de Leds WS281X, vía API, de forma que puedas controlar la tira de LEDs con una serie de peticiones http desde cualquier cliente.

## Prerrequisitos

Compilar la librería [rpi_ws281x](https://github.com/jgarff/rpi_ws281x) directamente en la Raspberry Pi, de acuerdo a la [documentación](https://github.com/jgarff/rpi_ws281x#build), para luego copiar los ficheros `*.a` a `/usr/local/lib` y los ficheros `*.h` a `/usr/local/include`.

## Documentación de peticiones

### Encender las luces

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Restaura la configuración de la tira de leds y enciende las luces con la animación por defecto)</code></summary>

##### Body

> | type      |  command                            |
> |-----------|-------------------------------------|
> | "startup" |  `{"instruction":"any","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                            |
> |---------------|---------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"startup","command":{"instruction":"any", "args":[]}}`     |
> | `400`         | `application/json`        | `Error message`                                                     |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Apagar las luces

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Apaga las luces y las deja en modo ahorro de energía)</code></summary>

##### Body

> | type       |  command                            |
> |------------|-------------------------------------|
> | "shutdown" |  `{"instruction":"any","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                            |
> |---------------|---------------------------|---------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"shutdown","command":{"instruction":"any", "args":[]}}`    |
> | `400`         | `application/json`        | `Error message`                                                     |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Subir brillo

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Sube el brillo de la tira de LEDs un 25%)</code></summary>

##### Body

> | type             |  command                                 |
> |------------------|------------------------------------------|
> | "set-brightness" |  `{"instruction":"increase","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                                       |
> |---------------|---------------------------|--------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"set-brightness","command":{"instruction":"increase", "args":[]}}`    |
> | `400`         | `application/json`        | `Error message`                                                                |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Subir brillo cierta cantidad

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Sube el brillo de la tira de LEDs una cantidad especificada)</code></summary>

##### Body

> | type             |  command                                     |
> |------------------|----------------------------------------------|
> | "set-brightness" |  `{"instruction":"increase","args":["10"]}`  |


##### Responses

> | http code     | content-type              | response                                                                           |
> |---------------|---------------------------|------------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"set-brightness","command":{"instruction":"increase", "args":["10"]}}`    |
> | `400`         | `application/json`        | `Error message`                                                                    |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Bajar brillo

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Baja el brillo de la tira de LEDs un 25%)</code></summary>

##### Body

> | type             |  command                                 |
> |------------------|------------------------------------------|
> | "set-brightness" |  `{"instruction":"decrease","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                                       |
> |---------------|---------------------------|--------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"set-brightness","command":{"instruction":"decrease", "args":[]}}`    |
> | `400`         | `application/json`        | `Error message`                                                                |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Bajar brillo cierta cantidad

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Baja el brillo de la tira de LEDs una cantidad especificada)</code></summary>

##### Body

> | type             |  command                                     |
> |------------------|----------------------------------------------|
> | "set-brightness" |  `{"instruction":"decrease","args":["10"]}`  |


##### Responses

> | http code     | content-type              | response                                                                           |
> |---------------|---------------------------|------------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"set-brightness","command":{"instruction":"decrease", "args":["10"]}}`    |
> | `400`         | `application/json`        | `Error message`                                                                    |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Modo oficina

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Pone las luces en modo oficina)</code></summary>

##### Body

> | type          |  command                                      |
> |---------------|-----------------------------------------------|
> | "change-mode" |  `{"instruction":"office-lights","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                                          |
> |---------------|---------------------------|-----------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"change-mode","command":{"instruction":"office-lights", "args":[]}}`     |
> | `400`         | `application/json`        | `Error message`                                                                   |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Modo Arcoiris Bolas

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Pone las luces en modo arcoiris bolas)</code></summary>

##### Body

> | type          |  command                                      |
> |---------------|-----------------------------------------------|
> | "change-mode" |  `{"instruction":"rainbow balls","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                                          |
> |---------------|---------------------------|-----------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"change-mode","command":{"instruction":"rainbow balls", "args":[]}}`     |
> | `400`         | `application/json`        | `Error message`                                                                   |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Modo Arcoiris Continuo

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Pone las luces en modo arcoiris continuo)</code></summary>

##### Body

> | type          |  command                                           |
> |---------------|----------------------------------------------------|
> | "change-mode" |  `{"instruction":"rainbow continuous","args":[]}`  |


##### Responses

> | http code     | content-type              | response                                                                               |
> |---------------|---------------------------|----------------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"change-mode","command":{"instruction":"rainbow continuous", "args":[]}}`     |
> | `400`         | `application/json`        | `Error message`                                                                        |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------

### Modo Color Estático

<details>
 <summary><code>POST</code> <code><b>/action</b></code> <code>(Pone las luces en modo color estático)</code></summary>

##### Body

> | type          |  command                                               |
> |---------------|--------------------------------------------------------|
> | "change-mode" |  `{"instruction":"static-color","args":["0xc1e1c1"]}`  |


##### Responses

> | http code     | content-type              | response                                                                                   |
> |---------------|---------------------------|--------------------------------------------------------------------------------------------|
> | `201`         | `application/json`        | `{"type":"change-mode","command":{"instruction":"static-color", "args":["0xc1e1c1"]}}`     |
> | `400`         | `application/json`        | `Error message`                                                                            |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8888/action
> ```

</details>

------------------------------------------------------------------------------------------