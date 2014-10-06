// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// xpjson is experimenting json datas
package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	type User struct {
		ID       int
		Login    string
		Nom      string
		Prenom   string
		URole    int
		PDVLabel string
		Supprime bool
	}
	u := User{123, "superlogin", "Martin", "Jacques", 3, "Epi d'or", false}
	b, e := json.Marshal(u)
	if e != nil {
		log.Fatal(e)
	}
	os.Stdout.Write(b)
}
