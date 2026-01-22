package domain

import "errors"

var (
	ErrEventNotFound     = errors.New("event not found")
	ErrNoAvailableSlots  = errors.New("no available slots for this event")
	ErrBookingNotFound   = errors.New("booking not found")
	ErrBookingExpired    = errors.New("booking has already expired")
	ErrAlreadyConfirmed  = errors.New("booking is already confirmed")
)
