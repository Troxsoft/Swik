#SwikDB

## What is swikDB?

- A database Key/value characteristics:
  - **Facility**
  - **Uses** `Javascript`
  - **All in one executable**
  - **Simplicity (this does not mean it is not powerful)**

##API

- ### `db.get(key:string)`
  Gets the specified key if it does not return null
- ### `db.getAll()`
  Gets all keys and values ​​as an array of objects.
- ### `db.getIf(logic:(e)=>bool)`
  Gets the first key and values ​​that meet the established logic.
- ### `db.getAllIf(logic:(e)=>bool)`
  Gets all keys and values ​​that comply with the established logic.
- ### `db.set(key:string)`
  Change the value of the key, always returns undefined
- ### `db.remove(key:string)`
  Deletes the specified key, if deleted it returns true, but the case otherwise
  returns false
- ### And more (the readme file will be updated soon, you can see it in the source code)

## Example:

```js
db.set("isGood", true);
db.get("isGood");
db.Remove("isGood");
db.set("a", 1);
db.set("b", 2);
db.removeIf((e) => e.key.key == "b");
```

#### Readme original en español.

[![License: MIT](https://img.shields.io/badge/License-MIT-Yellow.svg)](https://opensource.org/licenses/MIT)
