package function

import (
	"firm.com/connectDB"
	"firm.com/models"
)

func GetBizSizes() ([]models.BizSize, error) {
	// var ss string
	// ss= "select f.id, f.name, f.url, f2.name as holding, b.type, r.name FROM firm AS f" +
	// 	"left join biz_size AS b On f.biz_size_id=b.id" +
	// 	"left join region AS r ON f.region_id=r.id " +
	// 	"left join firm AS f2 ON f2.id=f.holding_id"
	var bizsizes []models.BizSize
	query := "SELECT id, type, brief FROM biz_size" //string
	rows, err := connectdb.ConnMySql().Query(query)
	if err != nil {
		return nil, err
	}
	//
	defer rows.Close()
	for rows.Next() {
		var bizsize models.BizSize
		if err := rows.Scan(&bizsize.ID, &bizsize.TYPE, &bizsize.BRIEF); err != nil {
			return nil, err
		}
		bizsizes = append(bizsizes, bizsize)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bizsizes, nil
}

func InsertBizSize(bizsize *models.BizSize) (*models.BizSize, error) {
	query := "INSERT INTO biz_size (type, brief) VALUES (?,?)"
	_, err := connectdb.ConnMySql().Exec(query, &bizsize.TYPE, &bizsize.BRIEF)
	if err != nil {
		return nil, err
	}
	return bizsize, nil
}

func GetSingleBizSize(id int) (models.BizSize, error) {
	var bizsize models.BizSize
	query := "SELECT type, brief FROM biz_size WHERE id = ?"
	row := connectdb.ConnMySql().QueryRow(query, id)
	if err := row.Scan(&bizsize.ID, &bizsize.TYPE, &bizsize.BRIEF); err != nil {
		return bizsize, err
	}
	return bizsize, nil
}

func UpdateBizSize(bizsize *models.BizSize) (*models.BizSize, error) {
	query := "UPDATE biz_size SET type = ?, brief=? WHERE id = ?"
	_, err := connectdb.ConnMySql().Exec(query, bizsize.TYPE, bizsize.BRIEF, bizsize.ID)
	if err != nil {
		return nil, err
	}
	return bizsize, nil
}

func DeleteBizSize(id int) error {
	query := "DELETE FROM biz_size WHERE id = ?"
	_, err := connectdb.ConnMySql().Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
