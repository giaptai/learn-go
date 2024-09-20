package function

import (
	"firm.com/connectDB"
	"firm.com/models"
)

func GetRegions() ([]models.Region, error) {
	var regions []models.Region
	query := "SELECT id, name, description FROM region"
	rows, err := connectdb.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var region models.Region
		if err := rows.Scan(&region.ID, &region.NAME, &region.DESC); err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return regions, nil
}

func InsertRegion(region *models.Region) (*models.Region, error) {
	query := "INSERT INTO region (name, description) VALUES (?,?)"
	_, err := connectdb.Db.Exec(query, &region.NAME, &region.DESC)
	if err != nil {
		return nil, err
	}
	return region, nil
}

func GetSingleRegion(id int) (models.Region, error) {
	var region models.Region
	query := "SELECT id, name, description FROM region WHERE id =?"
	row := connectdb.Db.QueryRow(query, id)
	if err := row.Scan(&region.ID, &region.NAME, region.DESC); err != nil {
		return region, nil
	}
	return region, nil
}

func UpdateRegion(region *models.Region) (*models.Region, error) {
	query := "UPDATE region SET name =?, description=? WHERE id =?"
	_, err := connectdb.Db.Exec(query, region.NAME, &region.DESC, region.ID)
	if err != nil {
		return nil, err
	}
	return region, nil
}

func DeleteRegion(id int) error {
	query := "DELETE FROM region WHERE id =?"
	_, err := connectdb.Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
