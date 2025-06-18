package controllers

import (
	"bitespeed-identity/database"
	"bitespeed-identity/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IdentifyRequest struct {
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phoneNumber"`
}

type IdentifyResponse struct {
	Contact struct {
		PrimaryContactId    uint     `json:"primaryContatctId"`
		Emails              []string `json:"emails"`
		PhoneNumbers        []string `json:"phoneNumbers"`
		SecondaryContactIds []uint   `json:"secondaryContactIds"`
	} `json:"contact"`
}

func IdentifyContact(c *gin.Context) {
	var req IdentifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var existingContacts []models.Contact

	// Query for matching email or phone
	database.DB.Where("email = ? OR phone_number = ?", req.Email, req.PhoneNumber).Find(&existingContacts)

	if len(existingContacts) == 0 {
		// Create new primary contact
		newContact := models.Contact{
			Email:          deref(req.Email),
			PhoneNumber:    deref(req.PhoneNumber),
			LinkPrecedence: "primary",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		database.DB.Create(&newContact)

		resp := IdentifyResponse{}
		resp.Contact.PrimaryContactId = newContact.ID
		resp.Contact.Emails = []string{newContact.Email}
		resp.Contact.PhoneNumbers = []string{newContact.PhoneNumber}
		resp.Contact.SecondaryContactIds = []uint{}

		c.JSON(http.StatusOK, resp)
		return
	}

	// We'll handle this logic later in step 6
	c.JSON(http.StatusOK, gin.H{"msg": "Found existing contacts, merge logic coming next"})
}

func deref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
