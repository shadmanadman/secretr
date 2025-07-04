package internal

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"filippo.io/age"
	_ "github.com/mattn/go-sqlite3"
)

func SecretsDir() string {
	dir := filepath.Join(os.Getenv("Home"), ".secrert")
	os.MkdirAll(dir, 0700)
	return dir
}

func dbPath() string {
	return filepath.Join(SecretsDir(), "secrets.db")
}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS secrets(
	name TEXT PRIMARY KEY,
	value BLOB
	);`)

	return db, err
}

func StoreSecret(name, plaintext, passphrase string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	defer db.Close()

	r, err := age.NewScryptRecipient(strings.TrimSpace(passphrase))

	if err != nil {
		return err
	}

	var out bytes.Buffer
	w, err := age.Encrypt(&out, r)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, plaintext)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, plaintext)
	if err != nil {
		return err
	}
	w.Close()

	_, err = db.Exec(`
	INSERT OR REPLACE INTO secrets(name,value) 	VALUES(?,?)`, name, out.Bytes())

	return err
}

func RetrieveSecret(name, passphrase string) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}

	defer db.Close()

	var data []byte
	err = db.QueryRow(`SELECT value FROM secrets WHERE name =?`, name).Scan(&data)
	if err != nil {
		return "", errors.New("secret not found")
	}

	id, err := age.NewScryptIdentity(strings.TrimSpace(passphrase))
	if err != nil {
		return "", err
	}

	r, err := age.Decrypt(bytes.NewReader(data), id)
	if err != nil {
		return "", err
	}

	decrypted, err := io.ReadAll(r)
	return string(decrypted), err
}

func DeleteSecret(name string) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	defer db.Close()

	res, err := db.Exec(`DELETE FROM secrets WHERE name=?`, name)
	if err != nil {
		return err
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("secret not found")
	}

	return nil
}

func ListSecrets() ([]string, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(`SELECT name FROM secrets ORDER BY name`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []string
	for rows.Next() {
		var name string
		rows.Scan(&name)
		results = append(results, name)
	}

	return results, nil
}
