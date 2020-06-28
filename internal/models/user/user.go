package user

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/irvanherz/aldebaran/internal/config"
)

type User struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Dob       string `json:"dob,omitempty"`
	Role      int    `json:"role,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

func ReadByEmail(email string) (*User, string) {
	user := User{}

	err := config.Db.QueryRow("SELECT * FROM user WHERE email=?", email).Scan(&user.Id, &user.Email, &user.Phone, &user.Name, &user.Gender, &user.Dob, &user.Role, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == nil {
		return &user, ""
	} else if err == sql.ErrNoRows {
		return nil, ""
	} else {
		return nil, "DATABASE_ERROR"
	}
}

func ReadById(userId int64) (*User, string) {
	user := User{}

	err := config.Db.QueryRow("SELECT * FROM user WHERE id=?", userId).Scan(&user.Id, &user.Email, &user.Phone, &user.Name, &user.Gender, &user.Dob, &user.Role, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == nil {
		return &user, ""
	} else if err == sql.ErrNoRows {
		return nil, ""
	} else {
		return nil, "DATABASE_ERROR"
	}
}

func ReadMany(page int, itemPerPage int, filters string) (*[]User, string) {
	limit := itemPerPage
	offset := (page - 1) * itemPerPage
	var whereClause = ""
	if filters != "" {
		whereClause = "WHERE " + filters
	}

	rows, err := config.Db.Query("SELECT * FROM user "+whereClause+" LIMIT ? OFFSET ?", limit, offset)
	if err == nil {
		var users []User
		for rows.Next() {
			var user User
			rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Dob, &user.Role, &user.Password, &user.CreatedAt, &user.UpdatedAt)
			users = append(users, user)
		}
		return &users, ""
	} else if err == sql.ErrNoRows {
		return nil, ""
	} else {
		return nil, "DATABASE_ERROR"
	}
}

func Write(user User) (*User, string) {
	stmt, prep_err := config.Db.Prepare("INSERT INTO user SET email=?, phone=?, name=?, gender=?, dob=?, role=?, password=?")
	if prep_err != nil {
		return nil, prep_err.Error()
	}
	res, exec_err := stmt.Exec(user.Email, user.Phone, user.Name, user.Gender, user.Dob, user.Role, user.Password)
	if exec_err != nil {
		me, _ := exec_err.(*mysql.MySQLError)
		if me.Number == 1062 {
			return nil, "Email has been registered"
		} else {
			return nil, me.Message
		}
	} else {
		insert_id, _ := res.LastInsertId()
		user.Id = insert_id
		return &user, ""
	}
}
