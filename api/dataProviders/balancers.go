package dataProviders

import "encoding/json"

type Balancer struct {
	Id                 int   `json:"id"`
	UsedMachines       []int `json:"usedMachines"`
	TotalMachinesCount int   `json:"totalMachinesCount"`
}

func (s *Store) ListBalancers() ([]*Balancer, error) {
	rows, err := s.Db.Query("select id, usedMachines, totalMachinesCount from balancers limit 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Balancer
	for rows.Next() {
		var b Balancer
		var usedMachines string
		if err := rows.Scan(&b.Id, &usedMachines, &b.TotalMachinesCount); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(usedMachines), &b.UsedMachines); err != nil {
			return nil, err
		}
		res = append(res, &b)
	}
	if res == nil {
		res = make([]*Balancer, 0)
	}
	return res, nil
}
