# Secret Detector

Secret Detector es un proyecto nuevo que busca ayudar a detectar secretos o credenciales expuestas en archivos y código.

El proyecto acaba de comenzar, por lo que sus funciones todavía están en desarrollo. Con el tiempo se añadirán mejoras y nuevas capacidades.

## Estado del proyecto

En desarrollo. La estructura y las funciones pueden cambiar a medida que el proyecto avance.

## Ejecutar el escáner de líneas

El programa lee desde la entrada estándar y muestra únicamente las líneas que contienen IDs de claves de acceso de AWS, con su número de línea. Se recomienda ejecutarlo desde WSL:

```bash
go run ./cmd/secret-detector < testdata/sample.txt
```

Los ejemplos positivos y negativos están en `testdata/aws-access-keys.txt` y `testdata/aws-access-keys-negative.txt`. Las reglas basadas en patrones se definen con `rules.Rule` (`ID`, `Description`, `Pattern` y `MinEntropy`, opcional) y se agregan individualmente al `rules.Registry` mediante `Add`. Para sumar una detección por expresión regular, crea su definición en `internal/services` y registra el resultado en `cmd/secret-detector/main.go`; no hace falta ampliar un constructor con más argumentos. La detección AWS identifica IDs de acceso (por ejemplo, prefijos `AKIA` y `ASIA`); no valida si la clave es activa ni detecta la clave secreta de AWS. La carga de reglas definidas por usuarios desde YAML queda para una etapa futura.

El tamaño máximo predeterminado por línea es 1 MiB. Consulta `.env.example` para ver las variables disponibles. Go no carga archivos `.env` por sí mismo; en WSL puedes exportar la variable antes de ejecutar:

```bash
export SECRET_DETECTOR_MAX_LINE_BYTES=2097152
go run ./cmd/secret-detector < testdata/sample.txt
```

## Próximos pasos

Algunas ideas para el desarrollo futuro:

- Detectar distintos tipos de secretos, como claves de API y tokens.
- Indicar en qué archivo y línea se encontró una posible exposición.
- Reducir falsos positivos y permitir ajustar las reglas de detección.
- Añadir instrucciones de instalación y uso cuando estén disponibles.

## Contribuciones

Las sugerencias y contribuciones serán bienvenidas.
