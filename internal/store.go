package internal

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"filippo.io/age"
)

func secretsDir() string {
	dir := filepath.Join(os.Getenv("Home"), ".secrert")
	os.MkdirAll(dir, 0700)
	return dir
}

func StoreSecret(name, plaintext, passphrase string) error {
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

	w.Close()

	path := filepath.Join(secretsDir(), name+".secret")
	return os.WriteFile(path, out.Bytes(), 0600)
}

func RetrieveSecret(name, passphrase string) (string, error) {
	path := filepath.Join(secretsDir(), name+".secret")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	d, err := age.NewScryptIdentity(strings.TrimSpace(passphrase))

	if err != nil {
		return "", err
	}

	r, err := age.Decrypt(bytes.NewReader(data), d)
	if err != nil {
		return "", err
	}

	out, err := io.ReadAll(r)
	return string(out), err
}
