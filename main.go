package main

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"path/filepath"
)

func getImages(dir string) ([]string, error) {
	var images []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			ext := filepath.Ext(path)
			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif" || ext == ".PNG" || ext == ".JPG" || ext == ".JPEG" || ext == ".GIF" {
				images = append(images, "/"+filepath.ToSlash(path))
			}
		}
		return nil
	})
	return images, err
}

func main() {
	router := gin.Default()

	router.Static("/Img", "./Img")
	router.LoadHTMLGlob("HTML/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/search", func(c *gin.Context) {
		keywords := c.Query("q")
		searchPath := "Img/" + keywords
		images, err := getImages(searchPath)
		if err != nil {
			c.HTML(http.StatusOK, "search.html", gin.H{
				"error": err,
			})
		}
		c.HTML(http.StatusOK, "search.html", gin.H{
			"Images": images,
		})
	})

	_ = router.Run(":8080")
}
