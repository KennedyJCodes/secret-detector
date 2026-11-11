# Secret Detector

Secret Detector es un proyecto nuevo que busca ayudar a detectar secretos o credenciales expuestas en archivos y código.

El proyecto acaba de comenzar, por lo que sus funciones todavía están en desarrollo. Con el tiempo se añadirán mejoras y nuevas capacidades.

## Estado del proyecto

En desarrollo. La estructura y las funciones pueden cambiar a medida que el proyecto avance.

## Ejecutar el escáner de líneas

El programa lee desde la entrada estándar y muestra cada línea con su número. Se recomienda ejecutarlo desde WSL:

```bash
go run ./cmd/secret-detector < testdata/sample.txt
```

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
