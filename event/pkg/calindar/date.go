package calindar

import "errors"

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetYear(year int) error {
	if year < 1 || year > 9999 {
		return errors.New("year out of range")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("month out of range")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("day out of range")
	}
	d.day = day
	return nil
}

func (e *Event) SetTitle(event string) error {
	if len(event) > 30 {
		return errors.New("event too long")
	}
	e.title = event
	return nil
}

func (d *Date) GetYear() int {
	return d.year
}
func (d *Date) GetMonth() int {
	return d.month
}
func (d *Date) GetDay() int {
	return d.day
}
func (e *Event) GetTitle() string {
	return e.title
}
