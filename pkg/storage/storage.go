package storage

import (
	"encoding/json"
	"os"
)

func Save(path string, record map[string]any) error {
	records, err := Load(path)
	if err != nil {
		return err
	}

	records = append(records, record)

	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Load(path string) ([]map[string]any, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []map[string]any{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var records []map[string]any
	err = json.Unmarshal(data, &records)
	return records, err
}

func DeleteByHash(path string, targetHash string) (bool, error) {
	records, err := Load(path)
	if err != nil {
		return false, err
	}

	newRecords := make([]map[string]any, 0)
	found := false

	for _, r := range records {
		if val, ok := r["hash"]; ok {
			if strVal, ok := val.(string); ok && strVal == targetHash {
				found = true
				continue
			}
		}
		newRecords = append(newRecords, r)
	}

	if !found {
		return false, nil
	}

	if len(newRecords) == 0 {
		err := os.Remove(path)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	data, err := json.MarshalIndent(newRecords, "", "  ")
	if err != nil {
		return false, err
	}

	return true, os.WriteFile(path, data, 0644)
}
