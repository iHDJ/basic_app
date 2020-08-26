package seed

import "github.com/jinzhu/gorm"

var (
	authorities = map[string][]string{
		"manager.enumeration": []string{
			"query",
			"show",
			"create",
			"update",
			"destroy",
			"batch_destroy",
		},

		"manager.group": []string{
			"query",
			"create",
			"update",
			"destroy",
		},
	}
)

func SeedAuthority(db *gorm.DB) (err error) {

	return
}
