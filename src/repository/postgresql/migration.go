package postgresql

var migrations = []string{
	"CREATE TABLE IF NOT EXISTS $1.cars_an2 ( " +
		"id serial NOT NULL, " +
		"link varchar(255) NOT NULL UNIQUE, " +
		"title varchar(255), " +
		"short_description varchar(50000), " +
		"price int," +
		"post_date varchar(25), " +
		"views_count int, " +
		"average_price int, " +
		"city varchar(255), " +
		"car_mark varchar(255), " +
		"car_model varchar(255), " +
		"car_year int, " +
		"car_type varchar(255), " +
		"engine_volume float, " +
		"transmission varchar(255), " +
		"wheel_location varchar(255), " +
		"car_color varchar(255), " +
		"wheel_drive varchar(255), " +
		"rastamozhen bool, " +
		"full_description varchar(500000), " +
		"promotions varchar(255), " +
		"promotion_price int, " +
		"created_at timestamp without time zone, " +
		"updated_at timestamp without time zone, " +
		"PRIMARY KEY (id)); ",
}
