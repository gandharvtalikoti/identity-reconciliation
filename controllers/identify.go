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
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
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

	// Step 1: Build map of all related contacts (email or phone matches)
	var allContacts []models.Contact
	email := deref(req.Email)
	phone := deref(req.PhoneNumber)

	database.DB.
		Where("email = ? OR phone_number = ?", email, phone).
		Order("created_at asc").
		Find(&allContacts)

	// Step 2: Determine primary contact
	primary := allContacts[0]
	if primary.LinkPrecedence == "secondary" && primary.LinkedID != nil {
		var actualPrimary models.Contact
		database.DB.First(&actualPrimary, *primary.LinkedID)
		primary = actualPrimary
	}

	// Step 2b: If multiple primaries exist, merge them
	primaryContacts := []models.Contact{}
	for _, c := range allContacts {
		if c.LinkPrecedence == "primary" {
			primaryContacts = append(primaryContacts, c)
		}
	}

	// If multiple primaries exist, merge under the oldest one
	if len(primaryContacts) > 1 {
		mainPrimary := primaryContacts[0]
		for _, other := range primaryContacts[1:] {
			// Downgrade to secondary
			database.DB.Model(&other).Updates(models.Contact{
				LinkedID:       &mainPrimary.ID,
				LinkPrecedence: "secondary",
				UpdatedAt:      time.Now().UTC(),
			})
		}
		primary = mainPrimary
	}

	// Step 3: See if new info exists (not already stored)
	alreadyExists := false
	for _, c := range allContacts {
		if c.Email == email && c.PhoneNumber == phone {
			alreadyExists = true
			break
		}
	}

	// Step 4: If new combo, create secondary contact
	if !alreadyExists {
		newContact := models.Contact{
			Email:          email,
			PhoneNumber:    phone,
			LinkedID:       &primary.ID,
			LinkPrecedence: "secondary",
			CreatedAt:      time.Now().UTC(),
			UpdatedAt:      time.Now().UTC(),
		}
		database.DB.Create(&newContact)
		allContacts = append(allContacts, newContact)
	}

	// Step 5: Format response
	emails := map[string]bool{}
	phones := map[string]bool{}
	secondaryIDs := []uint{}

	for _, c := range allContacts {
		if c.ID == primary.ID {
			emails[c.Email] = true
			phones[c.PhoneNumber] = true
			continue
		}
		if c.Email != "" {
			emails[c.Email] = true
		}
		if c.PhoneNumber != "" {
			phones[c.PhoneNumber] = true
		}
		if c.LinkPrecedence == "secondary" {
			secondaryIDs = append(secondaryIDs, c.ID)
		}
	}

	resp := IdentifyResponse{}
	resp.Contact.PrimaryContactId = primary.ID
	for e := range emails {
		resp.Contact.Emails = append(resp.Contact.Emails, e)
	}
	for p := range phones {
		resp.Contact.PhoneNumbers = append(resp.Contact.PhoneNumbers, p)
	}
	resp.Contact.SecondaryContactIds = secondaryIDs

	c.JSON(http.StatusOK, resp)

}

func deref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
