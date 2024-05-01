declare function _get_(key: string)
declare function _getKey_(key: string)
declare function _remove_(key: string)
declare function _getValue_(key: string)
declare function _getAllKeys_()
declare function _getAllValues_()
declare function _getAll_()
declare function _clear_()
declare function _set_(key: string, value: string)
class DB {
    constructor() {

    }
    get(key: string): KeyValue | string {

        const _data = _get_(key)
        const data = JSON.parse(_data)
        if (typeof data.sucess !== "undefined") {
            return `{"sucess":false}`

        }
        const kv = new KeyValue(_data)
        return kv
    }
    set(key: string, value: string): KeyValue | string {



        const kv = new KeyValue(`{
            "${key}":"${value}"

        }
        `).set(value)

        return kv
    }

    getAll(): KeyValueList {
        return new KeyValueList()
    }
    getAllKeys(): KeyList {
        return new KeyList()
    }
    getAllValues(): ValueList {
        return new ValueList()
    }
    getKey(key: string): Key | string {
        const dd = JSON.parse(_getKey_(key))
        if (typeof dd.sucess !== "undefined") {
            return `{"sucess":false}`
        }
        return new Key(dd.key)
    }

    getValue(key: string): Value | string {
        const dd = JSON.parse(_getValue_(key))
        if (typeof dd.sucess !== "undefined") {
            return `{"sucess":false}`
        }
        return new Value(new Key(key))
    }
    remove(key: string): string | KeyValueList {
        const h = new KeyValueList()
        return h.remove(key)
    }
    clear(key: string): KeyValueList {
        const h = new KeyValueList()
        return h.clear()
    }
}
class Key {
    private _key: string
    constructor(key: string) {
        this._key = key
    }
    key(): string {
        return this._key
    }
    val(): Value {
        return new Value(this)
    }
    value(): string {
        return _getKey_(this._key)
    }
}
class Value {
    private _key: Key
    constructor(key: Key) {
        this._key = key
    }
    key(): Key {
        return this._key
    }

    val(): string {
        const j = _getValue_(this._key.key())
        return JSON.parse(j).value
    }

    value(): string {
        return _getValue_(this._key.key())
    }
}
class KeyValueList {

    constructor() {
    }
    value(): string {
        return _getAll_()
    }
    remove(key: string): KeyValueList | string {
        const f = JSON.parse(_remove_(key))
        if (!f.sucess) {
            return f
        }
        return this
    }
    clear(): KeyValueList {
        _clear_()
        return this
    }
}
class KeyList {

    constructor() {
    }
    value(): string {
        return _getAllKeys_()
    }
}
class ValueList {

    constructor() {
    }


    value(): string {
        return _getAllValues_()
    }
}
// npx swc .\api.ts -o api.js
class KeyValue {
    private _key: string
    private _value: string
    constructor(json: string) {
        let h = JSON.parse(json)

        this._key = Object.keys(h)[0]
        this._value = Object.values(h)[0] as string
    }
    public key(): Key {
        return new Key(this._key)
    }
    public val(): Value {
        return new Value(this.key())
    }
    public value(): string {
        return `{"${this._key}":"${this._value}"}`
    }

    public set(value: string): KeyValue | string {
        const h = _set_(this._key, value)
        this._value = value
        if (!JSON.parse(h).sucess) {
            return `{"sucess":false}`
        }
        return this

    }

}
const db = new DB()


// hhe