package main

import (
  "testing"
)

func TestFlattenEmpty(t *testing.T) {
  empty := ``

  want := ``

  got, _ := Flatten(empty)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJsonWithEmptyValue1(t *testing.T) {
  json := `{}`

  want := ``

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJsonWithEmptyValue2(t *testing.T) {
  json := `{ "test": {} }`

  want := `test: {}
`

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJsonWithEmptyArray1(t *testing.T) {
  json := `[]`

  want := ``

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJsonWithEmptyArray2(t *testing.T) {
  json := `{ "test": [] }`

  want := `test: []
`

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJson1(t *testing.T) {
  json := `{ "test": "coincoin" }
`

  want := `test: "coincoin"
`

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJson2(t *testing.T) {
  json := `[ "coincoin", "cuicui", "coucou" ]
`

  want := `[0]: "coincoin"
[1]: "cuicui"
[2]: "coucou"
`

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf(`Flatten(): want "%v", got "%v"`, want, got)
  }
}

func TestFlattenJson3(t *testing.T) {
  json := ` 
{ "test": "coincoin" }

`

  want := `test: "coincoin"
`

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenJson4(t *testing.T) {
  json := `{
  "a": "Easy!",
  "b": {
    "c": 2,
    "d": true,
    "my.super/key.test": [
      "une\\vilaine` + "`string`" + `",
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
}
`

  want := `a: "Easy!"
b."my.super/key.test"[0]: "une\\vilaine` + "`string`" + `"
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
`

  got, _ := Flatten(json)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  } 
}

func TestFlattenYamlEmpty1(t *testing.T) {
  yaml := ""

  want := ""

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlEmpty2(t *testing.T) {
  yaml := `---
`

  want := ""

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlEmpty3(t *testing.T) {
  yaml := `---
---
---
`

  want := ""

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlMap(t *testing.T) {
  yaml := `coincoin: test
`

  want := `---
coincoin: "test"
`

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlArray(t *testing.T) {
  yaml := `- coincoin
- cuicui
- coucou
`

  want := `---
[0]: "coincoin"
[1]: "cuicui"
[2]: "coucou"
`

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlDocument(t *testing.T) {
  yaml := `---
a: Easy!
b:
  c: 2
  d: true
  my.super/key.test:
  - une\vilaine` + "`string`" + `
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
  i: "true"
`

  want := `---
a: "Easy!"
b."my.super/key.test"[0]: "une\\vilaine` + "`string`" + `"
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
`

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlDocumentWithEmptyValues(t *testing.T) {
  yaml := `---
# Source: avatar/charts/mongodb/templates/standalone/dep-sts.yaml
test1: {}
test2: []
test3: 
`

  want := `---
test1: {}
test2: []
test3: <nil>
`

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}

func TestFlattenYamlDocuments(t *testing.T) {
  yaml := `---
a: Easy!
b:
  c: 2
  d: true
  my.super/key.test:
  - une\vilaine` + "`string`" + `
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
  i: "true"
---
a: Easy!
b:
  c: 2
  d: true
  my.super/key.test:
  - une\vilaine` + "`string`" + `
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
  i: "true"
`

  want := `---
a: "Easy!"
b."my.super/key.test"[0]: "une\\vilaine` + "`string`" + `"
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
---
a: "Easy!"
b."my.super/key.test"[0]: "une\\vilaine` + "`string`" + `"
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
`

  got, _ := Flatten(yaml)

  if want != got {
    t.Fatalf("Flatten(): want:\n%v\n, got:\n%v", want, got)
  }
}
