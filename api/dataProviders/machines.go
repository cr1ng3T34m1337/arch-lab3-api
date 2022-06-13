package dataProviders

import (
	"encoding/json"
	"fmt"
)

type Machine struct {
	Id            int    `json:"id"`
	Os            string `json:"os"`
	Processor     string `json:"processor"`
	WorkingStatus string `json:"workingStatus"`
	ConnectedTo   int    `json:"connectedTo"`
}

func (s *Store) SetWorkingStatus(id int, status string) error {
	str := fmt.Sprintf("update machines set workingStatus=\"%s\" where id=%v;", status, id)
	sqlRes, err := s.Db.Exec(str)
	if err != nil {
		return err
	}
	num, err := sqlRes.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return nil
	}

	str = fmt.Sprintf("select connectedTo from machines where id=%v;", id)
	row := s.Db.QueryRow(str)

	var connectedTo int
	if err = row.Scan(&connectedTo); err != nil {
		return err
	}

	str = fmt.Sprintf("select usedMachines from balancers where id=%d;", connectedTo)
	row = s.Db.QueryRow(str)
	var usedMachinesStr string
	if err = row.Scan(&usedMachinesStr); err != nil {
		return err
	}
	var usedMachines []int
	if err := json.Unmarshal([]byte(usedMachinesStr), &usedMachines); err != nil {
		return err
	}
	index := findInSlice(id, usedMachines)
	if index == -1 {
		usedMachines = append(usedMachines, id)
	} else {
		usedMachines[index] = usedMachines[len(usedMachines)-1]
		usedMachines = usedMachines[0 : len(usedMachines)-1]
	}
	res, err := json.Marshal(usedMachines)
	if err != nil {
		return err
	}
	str = fmt.Sprintf("update balancers set usedMachines=\"%s\" where id=%d;", string(res), connectedTo)
	_, err = s.Db.Query(str)
	if err != nil {
		return err
	}

	return nil
}

func findInSlice(searched int, slice []int) int {
	for index, val := range slice {
		if val == searched {
			return index
		}
	}
	return -1
}
