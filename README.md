# flatten

Transform yaml or json input stream to a flat hierarchic key/value collection, sorted by alphabetic order.

Useful to provide diff usage between 2 documents.

Support yaml with multi-documents input stream.

* Example with yaml content:

```bash
$ echo '---
a: Easy!
b:
  c: 2
  d: true
  my.super/key.test:
  - une\vilaine `string`
  - 4
  - another.key.too:
      x: 1
      y: 5
      z: 6.666666666666666666666666
  e:
  - true
  - false
  f: |
    coincoin
      cuicui\npouetpouet
    je\mets\des\antislash
  g: 3.1415
  h: 3.141592654
  i: "true"' | flatten
```

... will send following output:

```bash
---
a: "Easy!"
b."my.super/key.test"[0]: "une\\vilaine `string`"
b."my.super/key.test"[1]: 4
b."my.super/key.test"[2]."another.key.too".x: 1
b."my.super/key.test"[2]."another.key.too".y: 5
b."my.super/key.test"[2]."another.key.too".z: 6.666666666666667
b.c: 2
b.d: true
b.e[0]: true
b.e[1]: false
b.f: "coincoin\n  cuicui\\npouetpouet\nje\\mets\\des\\antislash\n"
b.g: 3.1415
b.h: 3.141592654
b.i: "true"
```

* Example with json content:

```bash
$ echo '{
  "a": "Easy!",
  "b": {
    "c": 2,
    "d": true,
    "my.super/key.test": [
      "une\\vilaine `string`",
      4, {
        "another.key.too": {
          "x": 1,
          "y": 5,
          "z": 6.666666666666666666666666
        }
      }
    ],
    "e": [
      true,
      false
    ],
    "f": "coincoin\n  cuicui\\npouetpouet\nje\\mets\\des\\antislash\n",
    "g": 3.1415,
    "h": 3.141592654,
    "i": "true"
  }
}' | flatten
```

... will send following output:

```bash
a: "Easy!"
b."my.super/key.test"[0]: "une\\vilaine `string`"
b."my.super/key.test"[1]: 4
b."my.super/key.test"[2]."another.key.too".x: 1
b."my.super/key.test"[2]."another.key.too".y: 5
b."my.super/key.test"[2]."another.key.too".z: 6.666666666666667
b.c: 2
b.d: true
b.e[0]: true
b.e[1]: false
b.f: "coincoin\n  cuicui\\npouetpouet\nje\\mets\\des\\antislash\n"
b.g: 3.1415
b.h: 3.141592654
b.i: "true"
```

## Build

Requires go >= v1.20

```bash
go build
```

## Install

```bash
sudo ln -s ${PWD}/flatten /usr/local/bin/flatten
```

## Usage

* yaml file:

```bash
cat myfile.yaml | flatten
```

* json file:

```bash
cat myfile.yaml | yaml2json
```
