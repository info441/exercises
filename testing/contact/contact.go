package contact

import (
	"time"
)

//Contact represents a contact record
type Contact struct {
	FirstName  string
	LastName   string
	InsertedAt time.Time
	UpdatedAt  time.Time
}

//SetTrackingData sets the tracking fields of the Contact.
//If the InsertedAt field is a zero-value, this will set
//that field to the current time to track when the record
//was first inserted. It will also update the UpdatedAt
//field to the current time.
func (c Contact) SetTrackingData() {
	// TODO:
	//BUG: there's a subtle bug here that you need
	//to discover and fix. Run the automated tests
	//that call this method on a Contact instance
	//and check to see if these tracking fields
	//were set correctly.
	
	//if .InsertedAt is a zero-value time...
	if c.InsertedAt.IsZero() {
		//set it to the current time
		c.InsertedAt = time.Now()
	}
	//set .UpdatedAt to the current time
	c.UpdatedAt = time.Now()
}