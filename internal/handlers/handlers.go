package handlers

import (
	"BlogSite/internal/storage/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"strconv"
)

func HomeHandler(c *gin.Context) {
	allPosts := database.AllPosts()
	c.HTML(http.StatusOK, "index.html", allPosts)

}

func AddPostHandler(c *gin.Context) {
	uploadPath := "internal/assets/"
	imageName := ""

	postTitle := c.PostForm("postTitle")
	postDescription := c.PostForm("postDescription")
	postImage, err := c.FormFile("postImage")
	if err == nil {
		imageName = uuid.New().String() + filepath.Ext(postImage.Filename)
		imagePath := uploadPath + imageName
		err = c.SaveUploadedFile(postImage, imagePath)
	} else {
		fmt.Println(err)
	}
	database.AddPost(postTitle, postDescription, imageName)
}

func AddPostPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "add-post.html", nil)
}

func PostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedPost := database.PostByID(id)
	c.HTML(http.StatusOK, "post-page.html", selectedPost)

}
