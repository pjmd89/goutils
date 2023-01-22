# goutils

## Herramientas varias para golang

### Lista de herramientas

- dbutils:
    - db
    - tags
- jsonutils:
    - json
- systemutils:
    - argon2i
    - debugmode
    - fileexists
    - fileinfo
    - filesystem
    - goroutine_id
    - log
- videoutils:
    - ffmpeg
    - m3u8
    - mediainfo

## db

Interfaces y estructura para la estandarizacion de consultas a bases de dato.

## tags

Genera una estructura y consulta el tag `gql` establecidos en una estructura.

## json

Establece un `struct` definido partiendo de los datos dados en un archivo json.

## argon2i

Establece, compara y devuelve cadenas de cifrado en formato argon2i.

## debugmode

Define una variable booleana en dependencia del tag `debugmode`.

## fileexist

Chequea un fun archivo existe.

## fileinfo

Establece un objeto con toda la informacion necesario de un archivo dado.

## filesystem

Define una interface y  estructura compatible con `embed.FS` con la finalidad de poder utilizar archivos embebidos o rutas de carpetas en la misma estructura de código.

## goroutine_id

Devuelve el id de la rutina de go en ejecucion.

## log

Genera archivos de logs.

## ffmpeg

Genera archivos de videos usando el codec `H.264` partiendo de cualquier video dado. Requiere de `ffmpeg` instalado en su sistema.

## m3u8

Genera archivos `m3u8`.

## mediainfo

Devuelve la información de un archivo de video o audio dado. Requiere de `mediainfo` instalado en su sistema.
