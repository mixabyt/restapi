package sqlstore

import "mdl/internal/app/model"

type MeasureUnitsRepository struct {
	store *Store
}

func (m *MeasureUnitsRepository) GetAll() ([]*model.MeasureUnits, error) {
	rows, err := m.store.db.Query("SELECT * FROM measure_units")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measureUnits []*model.MeasureUnits
	for rows.Next() {
		var measureUnit model.MeasureUnits
		if err := rows.Scan(&measureUnit.ID, &measureUnit.Name, &measureUnit.AllowFractional); err != nil {
			return nil, err
		}
		measureUnits = append(measureUnits, &measureUnit)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return measureUnits, nil

}
