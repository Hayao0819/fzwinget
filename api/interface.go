package api

import "fmt"

func (p Package) Label() string {
	return fmt.Sprintf("[%v] %v", p.ID, p.Name)
}

func (p Package) Detail() string {
	return fmt.Sprintf(
		"[%v]\n%v\n\n───────────────────\n\npublished by %v\n\nlicenced under the %v\n\ntags: %v\nversions: %v\n\nURL: %v\nLicense URL: %v",
		p.ID,
		p.Name,
		p.Publisher,
		p.License,
		p.Tags,
		p.Versions,
		//humanize.Time(p.CreatedAt),
		//humanize.Time(p.UpdatedAt),
		p.Homepage,
		p.LicenseURL,
	)
}
