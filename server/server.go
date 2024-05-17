package server

import (
	"github.com/Theshawa/syntax-highlighter/highlighter"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	port        string
	engine      *gin.Engine
	highlighter *highlighter.Highlighter
}

func (s *Server) Run() error {
	return s.engine.Run(":" + s.port)
}

func (s *Server) GetLanguageOptions() []struct {
	Name  string
	Value string
} {
	languages := make([]struct {
		Name  string
		Value string
	}, 0)
	for k, l := range s.highlighter.Languages {
		languages = append(languages, struct {
			Name  string
			Value string
		}{Name: l.Name, Value: k})
	}
	return languages
}

func (s *Server) init() {
	s.port = os.Getenv("PORT")
	s.engine.StaticFS("/public", gin.Dir("server/public", false))
	s.engine.LoadHTMLGlob("server/pages/*")

	s.engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"availableLanguages": s.GetLanguageOptions(),
		})
	})
	s.engine.POST("/output", func(c *gin.Context) {
		code := strings.TrimSpace(c.PostForm("code"))
		language := c.PostForm("language")
		tokens, err := s.highlighter.Highlight(code, language)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"isOutput":           true,
			"code":               code,
			"tokens":             tokens,
			"selectedLanguage":   s.highlighter.Languages[language],
			"availableLanguages": s.GetLanguageOptions(),
		})
	})
}

func NewServer() *Server {
	s := &Server{
		engine:      gin.Default(),
		highlighter: highlighter.New(),
	}
	s.init()
	return s
}
