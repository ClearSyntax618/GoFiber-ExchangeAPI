# Simple Exchange API with Fiber

Cree un servicio que proporcione la tasa de cambio más reciente e histórica para la moneda.
Por favor, utilice una arquitectura de red de tres niveles en este desafío.
También puede considerar el uso de pub / sub, modo de trabajador para procesar la solicitud y hacerla más escalable.

## ¿Cómo funciona la API?
Es una API demasiado simple. Solamente contiene dos endpoints `GET`.

| Method | Resource |
| ------- | --------- |
| `GET` | / | 
| `GET` | /historical|

Cada endpoint retorna una respuesta `JSON`.

## ¿Qué retorna?

### Index
El primer `GET` retorna los valores actuales (en USD) de las monedas.

#### Respuesta
```json
{
  "rates": {
    "AED": 3.67304,
    "AFN": 75.202857,
    "ALL": 100.169585,
    "AMD": 401.67,
    "ANG": 1.802404,
    "AOA": 825.9294,
    "ARS": 350.0651,
    "AUD": 1.578187,
    "AWG": 1.8,
    "AZN": 1.7,
    "BAM": 1.852882,
    "BBD": 2,
    "BDT": 110.255837,
    "BGN": 1.8563,
    "BHD": 0.376986,
    "BIF": 2837.881393,
    "BMD": 1,
    "BND": 1.370003,
    ...
  }
}
```

### Historical
En cambio, al segundo `GET` se le debe pasar un campo `date` previamente desde el Frontend y con ese campo retorna los valores históricos de las monedas (en USD).

#### Parámetros
```json
{
  "date": "2011-12-01"
}
```

#### Respuesta
```json
{
  "rates": {
    "AED": 3.67125,
    "AFN": 48.322483,
    "ALL": 103.393684,
    "AMD": 381.533882,
    "ANG": 1.789736,
    "AOA": 94.934567,
    "ARS": 4.281433,
    "AUD": 0.97754,
    "AWG": 1.789736,
    "AZN": 0.785944,
    "BAM": 1.449622,
    "BBD": 2,
    "BDT": 76.903915,
    "BGN": 1.450858,
    "BHD": 0.376544,
    "BIF": 1229.910393,
    "BMD": 1,
    "BND": 1.282019,
    ...
  }
}
```
