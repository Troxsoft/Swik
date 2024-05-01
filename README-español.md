# SwikDB

## ¿Que es swikDB?

- Una base de datos Key/Value caracteristicas:
  - **Facilidad**
  - **Utiliza** `Javascript`
  - **Todo en un solo ejecutable**
  - **Simplicidad(esto no significa que no sea potente)**

## API

- ### `db.get(key:string)`
  Obtiene la clave que se le especifica si no devuelve null
- ### `db.getAll()`
  Obtiene todas las claves y valores como un array de objetos
- ### `db.getIf(logic:(e)=>bool)`
  Obtiene la primera clave y valores que cumpla con la logica establecida
- ### `db.getAllIf(logic:(e)=>bool)`
  Obtiene todas las claves y valores que cumpla con la logica establecida
- ### `db.set(key:string)`
  Cambia el valor de la clave,siempre devuelve undefined
- ### `db.remove(key:string)`
  Elimina la clave especificada,si se elimino devuelve true ,pero el el caso
  contrario devuelve false
- ### Y mas(proximamente se actualiza el readme,en el codigo fuente se puede mirar)

## Ejemplo:

```js
db.set("isGood", true);
db.get("isGood");
db.Remove("isGood");
db.set("a", 1);
db.set("b", 2);
db.removeIf((e) => e.key.key == "b");
```

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
