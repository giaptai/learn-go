package function

import (
	"firm.com/connectDB"
	"firm.com/models"
	"fmt"
)

func GetFirm() ([]models.VWFirm, error) {
	fmt.Println("s")
	var firms []models.VWFirm
	query := "SELECT id, logo_img, name, region FROM vw_firm" //string
	rows, err := connectdb.ConnMySql().Query(query)
	if err != nil {
		return nil, err
	}
	//
	defer rows.Close()
	for rows.Next() {
		var firm models.VWFirm
		if err := rows.Scan(&firm.ID, &firm.LogoImg, &firm.Name, &firm.Region); err != nil {
			return nil, err
		}
		firms = append(firms, firm)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return firms, nil
}

func InsertFirm(firm *models.VWFirm) (*models.VWFirm, error) {
	query := "INSERT INTO firm (name, region_id) VALUES (?,?)"
	_, err := connectdb.ConnMySql().Exec(query, &firm.Name, 192)
	if err != nil {
		return nil, err
	}
	return firm, nil
}

func GetSingleFirm(id int) (models.VWFirm, error) {
	fmt.Println(id)
	var firm models.VWFirm
	query := "SELECT id, logo_img, name, region FROM vw_firm WHERE id = ?"
	row := connectdb.ConnMySql().QueryRow(query, id)
	if err := row.Scan(&firm.ID, &firm.LogoImg, &firm.Name, &firm.Region); err != nil {
		return firm, err
	}
	return firm, nil
}

func UpdateFirm(firm *models.VWFirm) (*models.VWFirm, error) {
	fmt.Println(&firm.Name, "--", firm.Name)
	query := "UPDATE firm SET name = ? WHERE id = ?"
	_, err := connectdb.ConnMySql().Exec(query, firm.Name, firm.ID)
	if err != nil {
		return nil, err
	}
	return firm, nil
}

func DeleteFirm(id int) error {
	query := "DELETE FROM firm WHERE id = ?"
	_, err := connectdb.ConnMySql().Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
