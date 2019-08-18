// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// bdd connect to a mysql database and retrieve some datas and print them
package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// dbConnect tries to connect to a database whth given dsn and driver
// every second before timeout reached.
// as seen on https://alex.dzyoba.com/blog/go-connect-loop
func dbConnect(driver, dsn string, timeout time.Duration) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutReached := time.After(timeout)

	for {
		select {
		case <-timeoutReached:
			return nil, fmt.Errorf("Failed to connect to database after %s timeout", timeout)
		case <-ticker.C:
			err = db.Ping()
			if err == nil {
				return db, nil
			}
			log.Printf("Impossible de se connecter. Erreur = %v", err)
		}
	}

}

func main() {

	db, err := dbConnect("mysql", "admin:admin@/prod?parseTime=true&multiStatements=true", 3*time.Second)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir la db. Erreur = %v", err)
	}

	r := db.QueryRow("SET @T=1000;SELECT @T;")

	var v int64
	r.Scan(&v)
	println(v)

	rows, err := db.Query("SELECT DATE_FORMAT(v_date,'%y%m'),n_libelle,floor(sum(vd_qte*IF(p_vendupoids,1.0,p_poids))) FROM famille,produit,vente,vente_d,fournisseur,nutrition WHERE v_pdv=536 AND v_date>='2013-01-01' AND v_date<'2014-01-01' AND vd_produit=p_id AND vd_vid=v_id AND p_famille=fa_id AND p_fourn=fo_id AND fo_pdv=v_pdv AND fa_stats=n_id AND v_supprime=0 GROUP BY fa_stats,MONTH(v_date)")

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var (
		date     string
		cat      string
		quantite int
		res      []string
	)
	for rows.Next() {
		err := rows.Scan(&date, &cat, &quantite)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, date, cat, strconv.Itoa(quantite))
	}
	println(len(res))
	for i := 0; i < len(res)/3; i++ {
		println(res[3*i], ";", res[3*i+1], ";", res[3*i+2])
	}
}
