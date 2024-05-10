# Ejemplo de API Rest con Go y SQLite3

Estructura simple e inicial de una API Rest donde el servidor es desarrollado en Go (Golang) utilizando como base de datos SQLite3 (local).

> NOTA:  
Importante! Para ejecutar el binario generado (*.exe) es necesario tener instalado en el directorio %PATH% un compilador de C. Esto se debe al paquete utilizado como driver para la conexión con la base de datos.

Para más información leer el siguiente issue: [Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stu](https://github.com/go-gorm/gorm/issues/6468)

El ejemplo fue extraído de la siguiente web: [API Rest con Go (Golang) y SQLite3](https://medium.com/@orlmonteverde/api-rest-con-go-golang-y-sqlite3-e378af30719c)

## Compilar y generar archivo ejecutable (.exe)

Para generar el archivo binario ejecutable se debe realizar el siguiente comando:

```
go build
```

## Ejecutar el servidor

Para poner en funcionamiento el servidor, :

```console
./go-web-server.exe -migrate
```

Es importante incluir el flag `-migrate` para que se creen las tablas necesarias en la base de datos si es que no existen.


## Repositorio del paquete SQLite3

### go-sqlite3

URL del respositorio: [https://github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)