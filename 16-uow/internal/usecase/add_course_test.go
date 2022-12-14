package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/virb30/curso-go/16-uow/internal/repository"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories`(id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses`(id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY(category_id) REFERENCES categories(id));")

	input := InputUseCase{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryId: 2,
	}

	ctx := context.Background()
	useCase := NewAddCourseUseCase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
