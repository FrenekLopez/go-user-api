# Semana 1: Fundamentos de Go

Se detalla los aprendizajes técnicos de la Semana 1 del proyecto Go User API, enfocándose en los conceptos básicos de programación en Go practicados a través de programas pequeños.

## Programas Completados
1. **hello.go**:
   - **Descripción**: Imprime "Hola, Go!" en la consola.
   - **Aprendizajes**:
     - Declaración de `package main` para programas ejecutables.
     - Uso de `fmt.Println` para salida básica.
     - Estructura mínima de un programa en Go.
2. **circle.go**:
   - **Descripción**: Calcula el área de un círculo con radio 5.0.
   - **Aprendizajes**:
     - Importación del paquete `math` para usar constantes como `math.Pi`.
     - Declaración de constantes con `const` y variables de punto flotante con `float64`.
     - Formateo de salida con `fmt.Printf` y precisión (e.g., `%.2f`).
3. **factorial.go**:
   - **Descripción**: Calcula el factorial de 5 usando una función personalizada.
   - **Aprendizajes**:
     - Definición de funciones con parámetros y tipos de retorno (`func factorial(n int) int`).
     - Uso de bucles `for` y condicionales `if`.
     - Manejo de casos base (e.g., 0! = 1).
4. **palindrome.go**:
   - **Descripción**: Verifica si "radar" es un palíndromo, ignorando mayúsculas y espacios.
   - **Aprendizajes**:
     - Manipulación de cadenas con el paquete `strings` (`strings.ToLower`, `strings.ReplaceAll`).
     - Implementación de manejo de errores con tipos `error`.
     - Iteración sobre cadenas con índices para comparar caracteres.
5. **vowels.go**:
   - **Descripción**: Cuenta las vocales en "Hola Mundo".
   - **Aprendizajes**:
     - Iteración sobre cadenas usando `for` con `range` para acceder a runas.
     - Uso de condicionales para verificar vocales (`a`, `e`, `i`, `o`, `u`).
     - Aplicación de `strings.ToLower` para procesamiento insensible a mayúsculas.

## Reflexiones
- **Desafíos**: Al principio enfrenté el error `main redeclared` al ejecutar `go run .` en `cmd/week1/`. Aprendí a ejecutar programas individualmente (`go run <archivo>`) para evitar conflictos con múltiples funciones `main()`.
- **Progreso**: Gané confianza en la sintaxis de Go, el uso de paquetes, y estructuras de control. Los commits diarios reforzaron el hábito de control de versiones.
- **Herramientas**: Usar VS Code con la extensión de Go facilitó el formateo (`gofmt`) y la depuración.

## Próximos Pasos
- Semana 2: Explorar concurrencia en Go con goroutines y canales, preparando el terreno para el desarrollo de la API.