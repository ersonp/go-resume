package handlers

import (
	"github.com/ersonp/go-resume/common"
	"github.com/ersonp/go-resume/models"
	"github.com/gin-gonic/gin"
)

func ShowHomePage(c *gin.Context) {
	// Call the render function with the name of the template to render
	myinfo := models.GetMyInfoData()
	profession := models.GetProfessionList()
	education := models.GetEducationList()
	experience := models.GetExperienceList()
	project := models.GetProjectList()
	skill := models.GetMySkillList()
	smedia := models.GetSocialMediaList()
	common.Render(c, gin.H{
		"title":       "Home",
		"MyInfo":      myinfo,
		"Profession":  profession,
		"Education":   education,
		"Experience":  experience,
		"Project":     project,
		"Skill":       skill,
		"SocialMedia": smedia,
	}, "home.html")
}
