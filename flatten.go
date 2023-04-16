package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "github.com/clevandowski/yamltool"
  "gopkg.in/yaml.v3"
  "io"
  "os"
  "sort"
  "strconv"
  "strings"
)

func isJsonArray(s string) bool {
  firstChar := s[0:1]
  if firstChar == "[" {
    return true
  } else {
    return false
  }
}

func isYamlArray(s string) bool {
  firstChars := s[0:2]
  if firstChars == "- " {
    return true
  } else {
    return false
  }
}

func getDocumentFormat(document string) string {
  document = strings.TrimSpace(document)
  if document == "" {
    return "empty"
  }
  firstChar := document[0:1]
  if firstChar == "{" || firstChar == "[" {
    return "json"
  } else {
    return "yaml"
  }
}

func walkMap(s string, nodemap map[string]interface{}) string {
  output := ""
  if len(nodemap) == 0 {
    if s == "" {
      output = ""
    } else {
      output = s + ": {}\n" 
    }
  } else {
    for k, v := range nodemap {
      separator := "."
      if len(s) == 0 {
        separator = ""
      }
      key := k
      if strings.Contains(k, ".") {
        key = "\"" + k + "\""
      }
      output = output + walk(s + separator + key, v)
    }
  }
  return output
}

func walkArray(s string, nodearray []interface{}) string {
  output := ""
  if len(nodearray) == 0 {
    if s == "" {
      output = ""
    } else {
      output = s + ": []\n"
    } 
  } else {
    for n, v := range nodearray {
      output = output + walk(s + "[" + strconv.Itoa(n) + "]", v)
    }
  }
  return output
}


func walk(s string, node interface{}) string {
  output := ""
  if nodemap, ok := node.(map[string]interface{}); ok {
    output = walkMap(s, nodemap)
  } else if nodearray, ok := node.([]interface{}); ok {
    output = walkArray(s, nodearray)
  } else if value, ok := node.(string); ok {
    value = strings.Replace(value, "\\", "\\\\", -1)
    value = strings.Replace(value, "\n", "\\n", -1)
    output = fmt.Sprintf("%v: \"%v\"\n", s, value)
  } else if _, ok := node.(int); ok {
    output = fmt.Sprintf("%v: %v\n", s, node)
  } else if _, ok := node.(float64); ok {
    output = fmt.Sprintf("%v: %v\n", s, node)
  } else if _, ok := node.(bool); ok {
    output = fmt.Sprintf("%v: %v\n", s, node)
  } else if node == nil {
    output = fmt.Sprintf("%v: %v\n", s, node)
  } else {
    output = s
    fmt.Printf("[ERROR] %v - Unknown type %T for %v\n", output, node, node)
  }
  return output
}

func sortLines(input string) string {
  if strings.TrimSpace(input) == "" {
    return ""
  }
  var sorted sort.StringSlice
  sorted = strings.Split(input, "\n")
  sorted.Sort()
  sorted = sorted[1:]
  output := strings.Join(sorted, "\n")
  output = output + "\n"
  return output
}

func flattenJson(document string) (string, error) {
  var m any
  if isJsonArray(document) {
    m = make([]interface{}, 1)
  } else {
    m = make(map[string]interface{})
  }
  err := json.Unmarshal([]byte(document), &m)
  if err != nil {
    return "", err
  }
  unsorted := walk("", m)
  sorted := sortLines(unsorted)
  return sorted, nil
}


func flattenYaml(data string) (string, error) {
  documents := yamltool.SplitDocuments(data)
  output := ""
  for _, document := range documents {
    output += "---\n"
    var m any
    if isYamlArray(document) {
      m = make([]interface{}, 1)
    } else {
      m = make(map[string]interface{})
    }
    err := yaml.Unmarshal([]byte(document), &m)
    if err != nil {
      return "", err
    }
    unsorted := walk("", m)
    sorted := sortLines(unsorted)
    output += sorted
  }
  return output, nil
}

func Flatten(data string) (string, error) {
  switch getDocumentFormat(data) {
  case "json":
    return flattenJson(data)
  case "yaml":
    return flattenYaml(data)
  case "empty":
    return "", nil
  default:
    return "", errors.New("unknown document type")
  }
}

func main() {
  stdin, err := io.ReadAll(os.Stdin)
  if err != nil {
    panic(err)
  }

  data := string(stdin)
  flatten, err := Flatten(data)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%v", flatten)
}
