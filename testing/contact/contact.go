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
func (c *Contact) SetTrackingData() {
// SOLUTION: Notice that we've changed the type of the functions receivers to *Contact. Prior
// to that, the function receiver was passed in by value and therefore changes to the fields within
// the struct of the receiver did not take affect since we did not specifically referenced it as a pointer
	
	//if .InsertedAt is a zero-value time...
	if c.InsertedAt.IsZero() {
		//set it to the current time
		c.InsertedAt = time.Now()
	}
	//set .UpdatedAt to the current time
	c.UpdatedAt = time.Now()
}