package models

import "database/sql"

type Project struct {
  Id int
  Name string
}

type ProjectModel struct {
  db *sql.DB
}

func NewProjectModel(db *sql.DB) *ProjectModel {
  return &ProjectModel{db: db}
}

func (pm *ProjectModel) Create(name string) error {
  _, err := pm.db.Exec("INSERT INTO PROJECT (NAME) VALUES (?)", name)
  return err
}

func (pm *ProjectModel) ReadById(id int) (*Project, error) {
  row := pm.db.QueryRow("SELECT NAME FROM PROJECT WHERE ID = ?", id)

  if err := row.Err(); err != nil {
    return nil, err
  }

  p := &Project{Id: id}

  if err := row.Scan(&p.Name); err != nil {
    return nil, err
  }

  return p, nil
}

func (pm *ProjectModel) ReadAll() ([]*Project, error) {
  rows, err := pm.db.Query("SELECT NAME FROM PROJECT")
  defer rows.Close()

  if err != nil {
    return nil, err
  }

  pl := []*Project{}

  for rows.Next() {
    p := &Project{}

    if err := rows.Scan(&p.Id, &p.Name); err != nil {
      return nil, err
    }

    pl = append(pl, p)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }

  return pl, nil
}

func (pm *ProjectModel) Update(id int, name string) error {
  _, err := pm.db.Exec("UPDATE PROJECT SET NAME = ? WHERE ID = ?", name, id)
  return err
}

func (pm *ProjectModel) Delete(id int) error {
  _, err := pm.db.Exec("DELETE FROM PROJECT WHERE ID = ?", id)
  return err
}
