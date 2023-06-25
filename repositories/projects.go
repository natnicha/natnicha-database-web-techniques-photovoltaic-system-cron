package repositories

import (
	"photovoltaic-system-cron/db"
)

type Project struct {
	Id     int `json:"id"`
	UserId int `json:"userId"`
}

func GetProjects() ([]Project, error) {
	query := `SELECT pj.id, pj.user_id, pj.start_at at time zone 'utc' as start_at
		FROM projects pj
		INNER join products p on p.project_id = pj.id 
		WHERE CAST((pj.start_at at time zone 'utc' + INTERVAL '30 day') as date) <= CAST((now()) as date)
		AND pj.is_printed = false
	`
	var project []Project
	tx := db.Database.Raw(query).Find(&project)
	if tx.Error != nil {
		return []Project{}, tx.Error
	}
	return project, nil
}
