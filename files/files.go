package files

import (
	"os"
	"password/app-password/output"
)

type VaultDb struct {
	filename string
}

func NewVaultDb(name string) *VaultDb {
	return &VaultDb{
		filename: name,
	}
}

func (db *VaultDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *VaultDb) Write(content []byte) {
	file, err := os.Create(db.filename)

	output.PrintError(err)

	_, err = file.Write(content)
	defer file.Close()

	output.PrintError(err)
}
