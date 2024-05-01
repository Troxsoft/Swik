function _class_call_check(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
        throw new TypeError("Cannot call a class as a function");
    }
}
function _defineProperties(target, props) {
    for(var i = 0; i < props.length; i++){
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor) descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
    }
}
function _create_class(Constructor, protoProps, staticProps) {
    if (protoProps) _defineProperties(Constructor.prototype, protoProps);
    if (staticProps) _defineProperties(Constructor, staticProps);
    return Constructor;
}
function _define_property(obj, key, value) {
    if (key in obj) {
        Object.defineProperty(obj, key, {
            value: value,
            enumerable: true,
            configurable: true,
            writable: true
        });
    } else {
        obj[key] = value;
    }
    return obj;
}
var DB = /*#__PURE__*/ function() {
    "use strict";
    function DB() {
        _class_call_check(this, DB);
    }
    _create_class(DB, [
        {
            key: "get",
            value: function get(key) {
                var _data = _get_(key);
                var data = JSON.parse(_data);
                if (typeof data.sucess !== "undefined") {
                    return '{"sucess":false}';
                }
                var kv = new KeyValue(_data);
                return kv;
            }
        },
        {
            key: "set",
            value: function set(key, value) {
                var kv = new KeyValue('{\n            "'.concat(key, '":"').concat(value, '"\n\n        }\n        ')).set(value);
                return kv;
            }
        },
        {
            key: "getAll",
            value: function getAll() {
                return new KeyValueList();
            }
        },
        {
            key: "getAllKeys",
            value: function getAllKeys() {
                return new KeyList();
            }
        },
        {
            key: "getAllValues",
            value: function getAllValues() {
                return new ValueList();
            }
        },
        {
            key: "getKey",
            value: function getKey(key) {
                var dd = JSON.parse(_getKey_(key));
                if (typeof dd.sucess !== "undefined") {
                    return '{"sucess":false}';
                }
                return new Key(dd.key);
            }
        },
        {
            key: "getValue",
            value: function getValue(key) {
                var dd = JSON.parse(_getValue_(key));
                if (typeof dd.sucess !== "undefined") {
                    return '{"sucess":false}';
                }
                return new Value(new Key(key));
            }
        },
        {
            key: "remove",
            value: function remove(key) {
                var h = new KeyValueList();
                return h.remove(key);
            }
        },
        {
            key: "clear",
            value: function clear(key) {
                var h = new KeyValueList();
                return h.clear();
            }
        }
    ]);
    return DB;
}();
var Key = /*#__PURE__*/ function() {
    "use strict";
    function Key(key) {
        _class_call_check(this, Key);
        _define_property(this, "_key", void 0);
        this._key = key;
    }
    _create_class(Key, [
        {
            key: "key",
            value: function key() {
                return this._key;
            }
        },
        {
            key: "val",
            value: function val() {
                return new Value(this);
            }
        },
        {
            key: "value",
            value: function value() {
                return _getKey_(this._key);
            }
        }
    ]);
    return Key;
}();
var Value = /*#__PURE__*/ function() {
    "use strict";
    function Value(key) {
        _class_call_check(this, Value);
        _define_property(this, "_key", void 0);
        this._key = key;
    }
    _create_class(Value, [
        {
            key: "key",
            value: function key() {
                return this._key;
            }
        },
        {
            key: "val",
            value: function val() {
                var j = _getValue_(this._key.key());
                return JSON.parse(j).value;
            }
        },
        {
            key: "value",
            value: function value() {
                return _getValue_(this._key.key());
            }
        }
    ]);
    return Value;
}();
var KeyValueList = /*#__PURE__*/ function() {
    "use strict";
    function KeyValueList() {
        _class_call_check(this, KeyValueList);
    }
    _create_class(KeyValueList, [
        {
            key: "value",
            value: function value() {
                return _getAll_();
            }
        },
        {
            key: "remove",
            value: function remove(key) {
                var f = JSON.parse(_remove_(key));
                if (!f.sucess) {
                    return f;
                }
                return this;
            }
        },
        {
            key: "clear",
            value: function clear() {
                _clear_();
                return this;
            }
        }
    ]);
    return KeyValueList;
}();
var KeyList = /*#__PURE__*/ function() {
    "use strict";
    function KeyList() {
        _class_call_check(this, KeyList);
    }
    _create_class(KeyList, [
        {
            key: "value",
            value: function value() {
                return _getAllKeys_();
            }
        }
    ]);
    return KeyList;
}();
var ValueList = /*#__PURE__*/ function() {
    "use strict";
    function ValueList() {
        _class_call_check(this, ValueList);
    }
    _create_class(ValueList, [
        {
            key: "value",
            value: function value() {
                return _getAllValues_();
            }
        }
    ]);
    return ValueList;
}();
// npx swc .\api.ts -o api.js
var KeyValue = /*#__PURE__*/ function() {
    "use strict";
    function KeyValue(json) {
        _class_call_check(this, KeyValue);
        _define_property(this, "_key", void 0);
        _define_property(this, "_value", void 0);
        var h = JSON.parse(json);
        this._key = Object.keys(h)[0];
        this._value = Object.values(h)[0];
    }
    _create_class(KeyValue, [
        {
            key: "key",
            value: function key() {
                return new Key(this._key);
            }
        },
        {
            key: "val",
            value: function val() {
                return new Value(this.key());
            }
        },
        {
            key: "value",
            value: function value() {
                return '{"'.concat(this._key, '":"').concat(this._value, '"}');
            }
        },
        {
            key: "set",
            value: function set(value) {
                var h = _set_(this._key, value);
                this._value = value;
                if (!JSON.parse(h).sucess) {
                    return '{"sucess":false}';
                }
                return this;
            }
        }
    ]);
    return KeyValue;
}();
var db = new DB() // hhe
;

