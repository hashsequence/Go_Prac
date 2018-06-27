package main

import (
    "bufio"
    "fmt"
    "os"
  //  "io"
    "encoding/json"
  //  "strings"
    //"strconv"
    "reflect"
)

type Queue struct {
  data []string
  size int
  back int
  front int
}

func NewQueue() *Queue {
  return &Queue{ []string{},0,  0,  0 }
}

func (q* Queue) Add(value string) {
  q.data = append(q.data, value)
  q.size = len(q.data)
  q.back = q.size - 1
}

func (q* Queue) Pop() (string, bool) {
  temp := q.data[0]
  if q.size > 0 {
    q.data = q.data[1:]
    q.front = 0
    q.size = len(q.data)
    return temp,true
  }
  return "", false
}

func (q* Queue) IsEmpty() bool {
  return q.size == 0
}

func NewQueries() *Queue {
  queries := NewQueue()


  scanner := bufio.NewScanner(os.Stdin)
   for scanner.Scan() {
       queries.Add(scanner.Text())
   }
/*
  for {
      reader := bufio.NewReader(os.Stdin)
     //fmt.Print("Enter text: ")
     text, err := reader.ReadString('\n')
     if err != nil {
        if err == io.EOF {
            break
        }
    }

    // fmt.Println(text)
     queries.Add(text)
  }
  */
  return queries
}

type Store struct{
  storage []string
}

func (s *Store) Exec(query string, results *[]string) {
  //parse query
  defer  func() { if p := recover(); p != nil {
        return
    }
  }()

  pos := Pos(query,' ')
  command  := query[0: pos]
  doc := query[pos + 1:]
  //fmt.Println("document: " + doc)

  switch command {
  case  "add":
      s.Add(doc)
  case "get":
//    fmt.Println("getting: " + doc)
   s.Get(doc, results)
  case "delete":
    s.Delete(doc)
  default:
    fmt.Println("command does not exist")
  }
}

func (s *Store) Add(document string) {
  fmt.Println("-------------------------------------------")
  fmt.Println("ADD | Document: ", document, "\n")
  s.storage = append (s.storage, document)
  fmt.Println("-------------------------------------------")
}

func (s *Store) Delete(document string) {
  for i, page := range s.storage {
    if CheckIfPageContainsDoc(page, document) {
      s.storage = append(s.storage[:i], s.storage[i+1:]...)
    }
  }
}

func (s *Store) Get(document string, results *[]string) {
  defer  func() { if p := recover(); p != nil {
        return
    }
  }()
    fmt.Println("-------------------------------------------")
  for _, page := range s.storage {
   fmt.Println("GET | PAGE: ", page)
    if CheckIfPageContainsDoc(page, document) {
      // *results = append (*results, document)
        fmt.Println(document, " is in ", page)
    } else {
        fmt.Println(document, " is not in ", page)
    }

  }
    fmt.Println("-------------------------------------------")

}

func (s *Store) Process(queries *Queue) []string {
  res := []string{}
//  fmt.Println("queries.data: ", queries.data)
  for _,query := range queries.data {
    s.Exec(query, &res)
  }
  return res
}
/*******************************************
helper functions
*********************************************/

/*
  return index of first occurence of <value> in the string
*/
func Pos(s string, value rune) int {
    for k, v := range s {
        if (v == value) {
            return k
        }
    }
    return -1
}

/*
converts bool to string
*/
/*****************************************************************
CheckIfPageContainsDoc : check if the document is within the page of the storage


******************************************************************/









func CheckIfPageContainsDoc(page, document string) (flag bool) {
  /*
  bool, for JSON booleans
  float64, for JSON numbers
  string, for JSON strings
  []interface{}, for JSON arrays
  map[string]interface{}, for JSON objects
  nil for JSON null
*/
defer  func() { if p := recover(); p != nil {
      fmt.Errorf("Get paniced!!")
      flag = false
      return
  }
}()
  //pg_byte, _ := json.Marshal(page)
//  doc_byte, _ := json.Marshal(document)
  var jsonObject interface{}
  var jsonObject2 interface{}
  //json.Unmarshal(pg_byte, &jsonObject)
  //json.Unmarshal(doc_byte, &jsonObject2)
  json.Unmarshal([]byte(page), &jsonObject)
  json.Unmarshal([]byte(document), &jsonObject2)
  pg := jsonObject.(map[string]interface{})
  doc := jsonObject2.(map[string]interface{})
  //pg_str :=  fmt.Sprintf("%s",pg_byte)
  //doc_str := fmt.Sprintf("%s", doc_byte)
//  pg_str, _ = strconv.Unquote(pg_str)
  //doc_str, _ = strconv.Unquote(doc_str)
fmt.Println("GET | page: ", pg, "\n")
fmt.Println("GET | document: ", doc, "\n")

  flag = false
  for doc_key, doc_value := range doc {
   if _, ok:= pg[doc_key]; ok {
     flag = true
     if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Bool || reflect.TypeOf(pg[doc_key]).Kind() == reflect.Float64 || reflect.TypeOf(pg[doc_key]).Kind() == reflect.String {
       fmt.Println("comparing ", pg[doc_key], " and ", pg[doc_key])
       if reflect.TypeOf(pg[doc_key]).Kind() == reflect.TypeOf(doc_value).Kind() {
         if !reflect.DeepEqual(pg[doc_key], doc_value) {
           flag = false
           return
         }
       } else {
         flag = false
         return
       }
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Map {
       if reflect.TypeOf(doc_value).Kind() == reflect.Map {
         fmt.Println("comparing ", pg[doc_key].(map[string]interface{}), " and ", pg[doc_key].(map[string]interface{}))
         /*
            if reflect.DeepEqual(pg[doc_key].(map[string]interface{}),doc_value.(map[string]interface{})) {
              flag = false
              return
            }
            */

         if !CompareObj(pg[doc_key].(map[string]interface{}), doc_value.(map[string]interface{})) {
           flag = false
           return
         }

       }
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Slice {
       if reflect.TypeOf(doc_value).Kind() == reflect.Slice {
          fmt.Println("comparing ", pg[doc_key].([]interface{}), " and ", pg[doc_key].([]interface{}))
          /*
         if !reflect.DeepEqual(pg[doc_key].([]interface{}),doc_value.([]interface{})) {
           flag = false
           return
         }
         */
         if !CompareArr(pg[doc_key].([]interface{}), doc_value.([]interface{})) {
           flag = false
           return
         }

       }
     }
   } /*else {
     for pg_key, pg_value := range pg {
       if reflect.TypeOf(pg_value).Kind() == reflect.Map {

       } else if reflect.TypeOf(pg_value).Kind() == reflect.Slice {

       }
     }
   }*/
  }

  return
}

/***************************************************
Compare Functions
*****************************************************/
func CompareObj (pg, doc map[string]interface{}) (flag bool) {
  defer  func() { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
        flag = false
        return
    }
  }()
  flag = true
  for doc_key, doc_value := range doc {
   if _, ok:= pg[doc_key]; ok && reflect.TypeOf(pg[doc_key]).Kind() == reflect.TypeOf(doc_value).Kind() {
     if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Bool || reflect.TypeOf(pg[doc_key]).Kind() == reflect.Float64 || reflect.TypeOf(pg[doc_key]).Kind() == reflect.String {
       if reflect.TypeOf(pg[doc_key]).Kind() == reflect.TypeOf(doc_value).Kind() {
         if pg[doc_key] != doc_value {
           flag = false
           return
         }
       } else {
         flag = false
         return
       }
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Map {
       if reflect.TypeOf(doc_value).Kind() == reflect.Map {
         if !CompareObj(pg[doc_key].(map[string]interface{}), doc_value.(map[string]interface{})) {
           flag = false
           return
         }
       }
     } else if reflect.TypeOf(pg[doc_key]).Kind() == reflect.Slice {
       if reflect.TypeOf(doc_value).Kind() == reflect.Slice {
         if !CompareArr(pg[doc_key].([]interface{}), doc_value.([]interface{})) {
           flag = false
           return
         }
       }
     } else {
       flag = false
     }
   } else {
     flag =  false
   }
 }
 return
}

func CompareArr(pg, doc []interface{}) (flag bool) {
  defer  func() { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
        flag = false
        return
    }
  }()
  for _, doc_value := range doc {
    if !Contains(pg, doc_value) {
      return false
    }
  }
  return true
}

func Contains(s []interface{}, e interface{}) bool {
  defer  func() bool { if p := recover(); p != nil {
        fmt.Errorf("Get paniced!!")
        return false
    }
    return true
  }()
    for _, a := range s {
      switch reflect.ValueOf(e).Kind() {
      case reflect.Map:
        if reflect.TypeOf(a).Kind() == reflect.Map {
          if CompareObj(a.(map[string]interface{}), e.(map[string]interface{})) {
            return true
          }
        }
      case reflect.Slice:
        if reflect.TypeOf(a).Kind() == reflect.Slice {
          if CompareArr(a.([]interface{}), e.([]interface{})) {
            return true
          }
        }
      default:
        if reflect.DeepEqual(a,e) {
          return true
        }
      }
    }
    return false
}


func main() {
  datastore := &Store{[]string{}}
  queries := NewQueries()
  datastore.Process(queries)
  fmt.Println("\n")
  for _, value := range datastore.storage {
      fmt.Println("STORAGE: ", value)
  }
  fmt.Println("\n")
  for _, value := range queries.data {
      fmt.Println("QUERIES: ", value)
  }



}
