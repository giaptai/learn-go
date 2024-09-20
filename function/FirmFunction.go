package function

import (
	"firm.com/connectDB"
	"firm.com/models"
	"fmt"
)

// select f.id, f.name, f.url, f.location, b.type, r.name as region, f2.name as holding from firm as f
//
//	left join region as r on f.region_id = r.id
//	left join biz_size as b on b.id = f.biz_size_id
//	left join firm as f2 on f.holding_id = f2.id;

func GetFirm(page int) ([]models.VWFirm, error) {
	var firms []models.VWFirm
	query := "SELECT id, logo_img, name, region FROM vw_firm LIMIT 10 OFFSET ?" //string
	rows, err := connectdb.Db.Query(query, (page-1)*10)
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

func InsertFirm(firm *models.Firm) (*models.Firm, error) {
	query := "INSERT INTO firm (name, region_id, biz_size_id, url, location) VALUES (?,?,?,?,?)"
	_, err := connectdb.Db.Exec(query, &firm.Firm, &firm.Region, &firm.BizSize, &firm.Website, &firm.Address)
	if err != nil {
		return nil, err
	}
	return firm, nil
}

func GetSingleFirm(id int) (models.Firm, error) {
	fmt.Printf("Company id: %d\n", id)
	var firm models.Firm
	query := "SELECT id, logo_img, firm, url, location, biz_size, region, holding FROM vw_firmdetail WHERE id = ?"
	row := connectdb.Db.QueryRow(query, id)
	if err := row.Scan(&firm.ID, &firm.LogoImg, &firm.Firm,
		&firm.Website, &firm.Address, &firm.BizSize,
		&firm.Region, &firm.Holding); err != nil {
		return firm, err
	}
	return firm, nil
}

func UpdateFirm(firm *models.VWFirm) (*models.VWFirm, error) {
	fmt.Println(&firm.Name, "--", firm.Name)
	query := "UPDATE firm SET name = ? WHERE id = ?"
	_, err := connectdb.Db.Exec(query, firm.Name, firm.ID)
	if err != nil {
		return nil, err
	}
	return firm, nil
}

func DeleteFirm(id int) error {
	query := "DELETE FROM firm WHERE id = ?"
	_, err := connectdb.Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
