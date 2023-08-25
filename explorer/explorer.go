package explorer

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Explorer() {
	dbexplorer, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/eventLog")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*") // 设置模板路径

	r.GET("/", func(c *gin.Context) {
		rows, err := dbexplorer.Query("SELECT owner, tokenid FROM owners")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var data []struct {
			Owner   string
			TokenID string
		}

		for rows.Next() {
			var d struct {
				Owner   string
				TokenID string
			}
			err := rows.Scan(&d.Owner, &d.TokenID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			data = append(data, d)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Data": data,
		})
	})

	r.Run(":8080")
}
