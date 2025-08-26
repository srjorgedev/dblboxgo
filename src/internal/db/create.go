package db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTables(db *sql.DB) {
	tables := []string{
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_general (
			_id TEXT UNIQUE PRIMARY KEY,
			num_id INTEGER UNIQUE,
			chapter INTEGER,
			rarity INTEGER,
			type INTEGER,
			transform BOOLEAN,
			lf BOOLEAN,
			zenkai BOOLEAN,
			kit_type INTEGER,
			zenkai_kit_type INTEGER,
			tag_switch BOOLEAN
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_name (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			name INTEGER
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_affinity (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			affinity INTEGER
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_tag (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			tag INTEGER
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_traits (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			trait INTEGER
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_held_cards (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			card1 TEXT,
			card2 TEXT
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_stats_min (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			health INTEGER,
			base_strike_attack INTEGER,
			base_blast_attack INTEGER,
			base_strike_defense INTEGER,
			base_blast_defense INTEGER,
			critic INTEGER,
			ki_recovery INTEGER
		);
	`,
		`
		CREATE TABLE IF NOT EXISTS unit_parameters_stats_max (
			_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			_unit_id TEXT,
			health INTEGER,
			base_strike_attack INTEGER,
			base_blast_attack INTEGER,
			base_strike_defense INTEGER,
			base_blast_defense INTEGER,
			critic INTEGER,
			ki_recovery INTEGER
		);
	`,
		`CREATE TABLE IF NOT EXISTS data_tag (
		_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
		data_en TEXT,
		data_es TEXT,
		data_fr TEXT,
		data_jp TEXT
	)`,
		`CREATE TABLE IF NOT EXISTS data_chapter (
		_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
		data_en TEXT,
		data_es TEXT,
		data_fr TEXT,
		data_jp TEXT
	)`,
		`CREATE TABLE IF NOT EXISTS data_rarity (
		_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
		data_en TEXT,
		data_es TEXT,
		data_fr TEXT,
		data_jp TEXT
	)`,
		`CREATE TABLE IF NOT EXISTS data_type (
		_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
		data_en TEXT,
		data_es TEXT,
		data_fr TEXT,
		data_jp TEXT
	)`,
		`CREATE TABLE IF NOT EXISTS data_affinity (
		_id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
		data_en TEXT,
		data_es TEXT,
		data_fr TEXT,
		data_jp TEXT
	)`,
	}

	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			fmt.Println("Error creating table:")
			log.Fatal(err)
		}
	}

	log.Println("All tables created successfully.")

}
